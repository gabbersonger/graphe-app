<script lang="ts">
    import TextBlock from "@/components/Workspace/Content/TextBlock.svelte";

    import type { ScriptureSection } from "!/graphe/internal/data";
    import {
        ScriptureService,
        type ScriptureRef,
    } from "!/graphe/internal/scripture";
    import { onMount, tick } from "svelte";
    import { Events } from "@wailsio/runtime";

    const NUM_BLOCKS_DISPLAY = 20;
    const FULL_HEIGHT = 100_000; // This should be large
    const BUFFER_ZONE = 10_000;

    let viewport: HTMLDivElement;
    let visible_above_container: HTMLDivElement;
    let visible_below_container: HTMLDivElement;

    export let data: ScriptureSection[];
    $: n_blocks = data.reduce((a, c) => a + c.blocks.length, 0) ?? 0;
    $: if (n_blocks == 0) reset();
    $: if (n_blocks > 0) load();

    export let current_verse: ScriptureRef | undefined = undefined;

    type VirtualBlock = {
        index: number;
        section: number;
        block: number;
    };
    let visible_above: VirtualBlock[] = [];
    let visible_below: VirtualBlock[] = [];
    let visible_above_elements: HTMLDivElement[] = [];
    let visible_below_elements: HTMLDivElement[] = [];
    let pivot_height = 0;
    let mode: "exact" | "middle" | "top" | "bottom" = "top";

    function reset() {
        current_verse = undefined;
        visible_above.length = 0;
        visible_below.length = 0;
        visible_above_elements.length = 0;
        visible_below_elements.length = 0;
        pivot_height = 0;
        mode = "top";
    }

    // Called when new data is loaded
    function load() {
        reset();
        current_verse = data[0].blocks[0].verses[0].ref;
        displayAroundPivot(0);
    }

    function getVirtualBlocks(
        output: VirtualBlock[],
        start: number,
        length: number | undefined = undefined,
    ) {
        if (length == undefined) length = NUM_BLOCKS_DISPLAY;
        if (start < 0 || start >= n_blocks || length <= 0) {
            output.length = 0;
            return output;
        }
        if (start + length > n_blocks) length = n_blocks - start;

        // Ensure the output array is of right length
        const old_output_length = output.length;
        output.length = length;
        for (let i = old_output_length; i < length; i++)
            output[i] = { index: 0, section: 0, block: 0 };

        // Fill the output array
        let virtual_rows_filled = 0;
        let curr_section = 0;
        let n_to_skip = start;
        while (virtual_rows_filled < length) {
            if (data[curr_section].blocks.length <= n_to_skip) {
                n_to_skip -= data[curr_section].blocks.length;
                curr_section += 1;
                continue;
            }

            const n_possible = data[curr_section].blocks.length - n_to_skip;
            const n_to_take = Math.min(
                length - virtual_rows_filled,
                n_possible,
            );
            for (let i = 0; i < n_to_take; i++) {
                const n = virtual_rows_filled + i;
                output[n].index = start + n;
                output[n].section = curr_section;
                output[n].block = n_to_skip + i;
            }
            curr_section++;
            n_to_skip = 0;
            virtual_rows_filled += n_to_take;
        }
        return output;
    }

    function getVirtualBlocksAbove(output: VirtualBlock[], pivot: number) {
        const start = Math.max(pivot - NUM_BLOCKS_DISPLAY, 0);
        const length = pivot - start;
        return getVirtualBlocks(output, start, length);
    }

    async function refresh() {
        if (n_blocks == 0) return;
        let scroll = viewport.scrollTop;

        let vae_length = 0;
        for (; vae_length < visible_above_elements.length; vae_length++)
            if (visible_above_elements[vae_length] == null) break;

        // Get the new pivot and the offset within that new pivot
        let new_pivot_offset = 0;
        let offset_remaining = 0;
        if (scroll > pivot_height) {
            offset_remaining = scroll - pivot_height;
            for (let i = 0; i < visible_below_elements.length; i++) {
                const height = visible_below_elements[i].clientHeight;
                if (offset_remaining - height > 0) {
                    offset_remaining -= height;
                } else {
                    new_pivot_offset = i;
                    break;
                }
            }
        } else {
            offset_remaining *= -1;
            for (let i = 1; i <= vae_length; i++) {
                offset_remaining -=
                    visible_above_elements[vae_length - i].clientHeight;
                if (offset_remaining <= 0) {
                    new_pivot_offset = -i;
                    break;
                }
            }
        }

        if (new_pivot_offset == 0) return;

        // Update current verse
        let new_virtual: VirtualBlock;
        if (new_pivot_offset > 0) {
            new_virtual = visible_below[new_pivot_offset];
        } else {
            new_virtual =
                visible_above[visible_above.length + new_pivot_offset];
        }
        const new_block = data[new_virtual.section].blocks[new_virtual.block];
        current_verse = new_block.verses[0].ref;
        // FIX: currently this is a very approximate current verse
        // (does not take into account offset_remaining)

        // Update pivot height
        let moved_height = 0;
        if (new_pivot_offset > 0) {
            for (let i = 0; i < new_pivot_offset; i++)
                moved_height += visible_below_elements[i].clientHeight;
        } else {
            for (let i = 1; i <= -new_pivot_offset; i++)
                moved_height -=
                    visible_above_elements[vae_length - i].clientHeight;
        }
        pivot_height += moved_height;

        // Update data being displayed
        const new_pivot = visible_below[0].index + new_pivot_offset;
        displayAroundPivot(new_pivot);
    }

    async function displayAroundPivot(pivot: number) {
        visible_above = getVirtualBlocksAbove(visible_above, pivot);
        visible_below = getVirtualBlocks(visible_below, pivot);

        const top_reached =
            visible_above.length > 0 && visible_above[0].index == 0;
        const bottom_reached =
            visible_below.length > 0 &&
            visible_below[visible_below.length - 1].index == n_blocks - 1;

        if (top_reached && bottom_reached) {
            mode = "exact";
        } else if (top_reached) {
            mode = "top";
        } else if (bottom_reached) {
            mode = "bottom";
        } else {
            mode = "middle";
        }
    }

    function goto(block: number) {
        pivot_height = FULL_HEIGHT / 2;
        displayAroundPivot(block);
        viewport.scrollTop = pivot_height;
    }

    async function gotoRef(ref: ScriptureRef) {
        let nth_item = 0;
        for (let i = 0; i < data.length; i++) {
            const section = data[i];
            const section_contains = await ScriptureService.RangeContains(
                section.range,
                ref,
            );
            if (section_contains) {
                for (let j = 0; j < section.blocks.length; j++) {
                    const block_contains = await ScriptureService.RangeContains(
                        section.blocks[j].range,
                        ref,
                    );
                    if (block_contains) {
                        return goto(nth_item);
                    }
                    nth_item += 1;
                }
            } else nth_item += section.blocks.length;
        }
    }

    function handleEvent(event_data: any) {
        gotoRef(event_data.data as ScriptureRef);
    }

    onMount(() => {
        Events.On("window:workspace:text:goto", handleEvent);
        return () => {
            Events.Off("window:workspace:text:goto");
        };
    });
</script>

<div class="container">
    <div class="viewport" bind:this={viewport} on:scroll={refresh}>
        <div class="content">
            <div
                class="rows"
                data-mode={mode}
                style:--full-height="{FULL_HEIGHT}px"
                style:--pivot-height="{pivot_height}px"
            >
                <div
                    class="collection"
                    bind:this={visible_above_container}
                    style:bottom={`calc(100% - ${pivot_height}px)`}
                >
                    {#each visible_above as vb, index (vb.index)}
                        <div
                            bind:this={visible_above_elements[index]}
                            style="outline: 2px solid red;"
                        >
                            <TextBlock
                                block={data[vb.section].blocks[vb.block]}
                            />
                        </div>
                    {/each}
                </div>

                <div
                    class="collection"
                    bind:this={visible_below_container}
                    style:top={`${pivot_height}px`}
                >
                    {#each visible_below as vb, index (vb.index)}
                        <div
                            bind:this={visible_below_elements[index]}
                            style="outline: 2px solid blue;"
                        >
                            <TextBlock
                                block={data[vb.section].blocks[vb.block]}
                            />
                        </div>
                    {/each}
                </div>
            </div>
        </div>
    </div>
</div>

<style>
    .container {
        position: absolute;
        inset: 0;
    }

    .viewport {
        position: relative;
        width: 100%;
        height: 100%;
        overflow-y: scroll;
    }

    /* .viewport::-webkit-scrollbar {
        display: none;
    } */

    .content {
        position: relative;
        width: 90%;
        max-width: 80ch;
        margin: 0 auto;
    }

    .rows {
        position: relative;
        height: var(--full-height);
    }

    .collection {
        width: 100%;
    }

    .rows[data-mode="exact"] {
        height: unset;
    }

    .rows[data-mode="exact"],
    .rows[data-mode="top"] {
        display: flex;
        flex-direction: column;
        justify-content: flex-start;
    }

    .rows[data-mode="middle"] .collection {
        position: absolute;
    }

    .rows[data-mode="bottom"] {
        display: flex;
        flex-direction: column;
        justify-content: flex-end;
    }
</style>
