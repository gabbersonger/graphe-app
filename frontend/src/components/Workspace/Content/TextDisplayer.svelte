<script lang="ts">
    import TextBlock from "@/components/Workspace/Content/TextBlock.svelte";
    import type { ScriptureSection } from "!/graphe/internal/data";
    import {
        ScriptureService,
        type ScriptureRef,
    } from "!/graphe/internal/scripture";
    import { onMount, tick } from "svelte";
    import { Events } from "@wailsio/runtime";
    import { throttle } from "@/lib/utils";

    const NUM_BLOCK_BUFFER = 20;
    const SCROLLABLE_HEIGHT = 100_000;
    const SCROLLABLE_BUFFER = 10_000;

    let viewport: HTMLDivElement;

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
    let pivot_height: number = 0;
    let scroll_locked = false;

    function reset() {
        current_verse = undefined;
        visible_above.length = 0;
        visible_below.length = 0;
        visible_above_elements.length = 0;
        visible_below_elements.length = 0;
        pivot_height = 0;
        scroll_locked = false;
    }

    function load() {
        reset();
        loadVirtualBlock({ index: 0, section: 0, block: 0 });
    }

    function refresh() {
        if (scroll_locked) return;
        if (n_blocks == 0) return;

        const scroll = viewport.scrollTop;
        const { pivot, offset, remaining } = calculatePivot(scroll);

        if (pivot.index == visible_below[0].index) return;
        throttled_update(pivot, offset, remaining);
    }
    const throttled_refresh = throttle(refresh, 10);

    async function update(vb: VirtualBlock, offset: number, remaining: number) {
        // Update current verse
        const new_block = data[vb.section].blocks[vb.block];
        current_verse = new_block.verses[0].ref;

        // Handle reaching the start of blocks
        if (vb.index == 0) return loadVirtualBlock(vb, remaining);

        // Handle reaching the end of blocks
        if (vb.index + NUM_BLOCK_BUFFER >= n_blocks)
            return loadVirtualBlock(vb, remaining);

        // Handle getting too close to edge while in the middle
        if (
            pivot_height + offset < SCROLLABLE_BUFFER ||
            pivot_height + offset > SCROLLABLE_HEIGHT - SCROLLABLE_BUFFER
        ) {
            return loadVirtualBlock(vb, remaining);
        }

        pivot_height += offset;
        updateViritualBlocks(vb.index);
    }
    const throttled_update = throttle(update, 10);

    function calculatePivot(scroll: number): {
        pivot: VirtualBlock;
        offset: number;
        remaining: number;
    } {
        let pivot: VirtualBlock;
        let offset = 0;
        let remaining = scroll - pivot_height;

        if (remaining > 0) {
            const last = visible_below_elements.findLastIndex((e) => e != null);
            let i = 0;
            while (remaining > 0 && i < last) {
                const height = visible_below_elements[i].clientHeight;
                if (remaining - height > 0) {
                    remaining -= height;
                    offset += height;
                } else break;
                i++;
            }
            pivot = visible_below[i];
        } else {
            let i = visible_above_elements.findLastIndex((e) => e != null);
            while (remaining < 0 && i >= 0) {
                const height = visible_above_elements[i].clientHeight;
                remaining += height;
                offset -= height;
                if (remaining >= 0) break;
                i--;
            }
            pivot = i < 0 ? visible_below[0] : visible_above[i];
        }

        return {
            pivot: pivot,
            offset: offset,
            remaining: remaining,
        };
    }

    function getVirtualBlocks(
        output: VirtualBlock[],
        start: number,
        length: number | undefined = undefined,
    ) {
        if (length == undefined) length = NUM_BLOCK_BUFFER;
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

    function updateViritualBlocks(pivot: number) {
        const start = Math.max(pivot - NUM_BLOCK_BUFFER, 0);
        visible_above = getVirtualBlocks(visible_above, start, pivot - start);
        visible_below = getVirtualBlocks(visible_below, pivot);
    }

    async function loadVirtualBlock(vb: VirtualBlock, offset: number = 0) {
        // Stop re-loading if it is the current block already loaded
        if (visible_below.length > 0 && visible_below[0].index == vb.index)
            return;

        // Stop re-loading bottom if bottom already loaded
        if (
            visible_below.length > 0 &&
            visible_below[0].index < vb.index &&
            visible_below[visible_below.length - 1].index == n_blocks - 1
        )
            return;

        updateViritualBlocks(vb.index);
        current_verse = data[vb.section].blocks[vb.block].verses[0].ref;

        if (visible_below.length > 0 && visible_below[0].index == 0) {
            // At the top
            pivot_height = 0;
        } else if (
            visible_below.length > 0 &&
            visible_below[visible_below.length - 1].index == n_blocks - 1
        ) {
            // At the bottom
            pivot_height = SCROLLABLE_HEIGHT;
            scroll_locked = true;
            await tick();
        } else {
            // In the middle
            pivot_height = SCROLLABLE_HEIGHT / 2;
        }

        // Scroll to the new content
        const current_scroll = viewport ? viewport.scrollTop : 0;
        if (pivot_height == current_scroll) return;

        scroll_locked = true;
        viewport.scrollTop = pivot_height + offset;
        setTimeout(() => {
            scroll_locked = false;
        }, 10);
        // FIX: this relies on 10ms being enough time for scrollTop to finish
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
                        loadVirtualBlock({
                            index: nth_item,
                            section: i,
                            block: j,
                        });
                    }
                    nth_item += 1;
                }
            } else nth_item += section.blocks.length;
        }
    }

    onMount(() => {
        Events.On("window:workspace:text:goto", (event_data: any) => {
            gotoRef(event_data.data as ScriptureRef);
        });
        return () => {
            Events.Off("window:workspace:text:goto");
        };
    });
</script>

<div class="container">
    <div class="viewport" bind:this={viewport} on:scroll={throttled_refresh}>
        <div class="content">
            <div
                class="rows"
                style:min-height={`${SCROLLABLE_HEIGHT}px`}
                style:--pivot-height={`${pivot_height}px`}
            >
                <div class="collection above">
                    {#each visible_above as vb, index (vb.index)}
                        <div bind:this={visible_above_elements[index]}>
                            <TextBlock
                                block={data[vb.section].blocks[vb.block]}
                            />
                        </div>
                    {/each}
                </div>

                <div class="collection below">
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
        height: var(--full-height);
    }

    .collection {
        position: absolute;
        width: 100%;
    }

    .collection.above {
        bottom: calc(100% - var(--pivot-height));
    }

    .collection.below {
        top: var(--pivot-height);
    }
</style>
