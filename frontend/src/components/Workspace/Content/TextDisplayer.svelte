<script lang="ts">
    import TextBlock from "@/components/Workspace/Content/TextBlock.svelte";

    import type { ScriptureSection } from "!/graphe/internal/data";
    import {
        ScriptureService,
        type ScriptureRef,
    } from "!/graphe/internal/scripture";
    import { onMount } from "svelte";
    import { Events } from "@wailsio/runtime";

    const NUM_BLOCKS_DISPLAY = 15;
    const MAX_HEIGHT = 100_000;

    export let data: Array<ScriptureSection>;
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
    let min_height = MAX_HEIGHT;

    function reset() {
        current_verse = undefined;
        visible_above.length = 0;
        visible_below.length = 0;
        visible_above_elements.length = 0;
        visible_below_elements.length = 0;
        pivot_height = 0;
        min_height = MAX_HEIGHT;
    }

    // Called when new data is loaded
    function load() {
        reset();
        getVirtualBlocks(visible_below, 0);
        current_verse = data[0].blocks[0].verses[0].ref;
        if (visible_below.length == n_blocks) {
            min_height = 0;
        }
    }

    function getVirtualBlocks(
        output: VirtualBlock[],
        start: number,
        length: number | undefined = undefined,
    ) {
        if (length == undefined) length = NUM_BLOCKS_DISPLAY;
        if (start < 0 || start >= n_blocks || length < 0) return [];
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

    let viewport: HTMLDivElement;
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
        visible_above = getVirtualBlocksAbove(visible_above, new_pivot);
        visible_below = getVirtualBlocks(visible_below, new_pivot);
    }

    function resize(width: number) {
        // TODO
    }

    function goto(block: number) {
        // TODO
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

    let content: HTMLDivElement;
    const resize_observer = new ResizeObserver((e) =>
        resize(e[0].contentRect.width),
    );
    onMount(() => {
        resize_observer.observe(content);
        Events.On("window:workspace:visualiser:goto", gotoRef);
        return () => {
            resize_observer.unobserve(content);
            Events.Off("window:workspace:visualiser:goto");
        };
    });
</script>

<div class="container">
    <div class="viewport" bind:this={viewport} on:scroll={refresh}>
        <div class="content" bind:this={content}>
            <div
                class="rows"
                style="--pivot-height: {pivot_height}px; min-height: {min_height}px"
            >
                <div
                    class="collection"
                    style:bottom={`calc(100% - ${pivot_height}px)`}
                >
                    {#each visible_above as vb, index (vb.index)}
                        <div bind:this={visible_above_elements[index]}>
                            <TextBlock
                                block={data[vb.section].blocks[vb.block]}
                            />
                        </div>
                    {/each}
                </div>

                <div class="collection" style:top={`${pivot_height}px`}>
                    {#each visible_below as vb, index (vb.index)}
                        <div bind:this={visible_below_elements[index]}>
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

    .viewport::-webkit-scrollbar {
        display: none;
    }

    .content {
        position: relative;
        width: 90%;
        max-width: 80ch;
        margin: 0 auto;
    }

    .rows {
        position: relative;
    }

    .collection {
        position: absolute;
        width: 100%;
    }
</style>
