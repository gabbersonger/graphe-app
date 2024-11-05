<script lang="ts">
    import TextBlock from "@/components/Workspace/Content/TextBlock.svelte";
    import type { ScriptureSection } from "!/graphe/internal/data";
    import {
        ScriptureService,
        type ScriptureRef,
    } from "!/graphe/internal/scripture";
    import { throttle } from "@/lib/utils";
    import { EventHandler } from "@/lib/event_handler";
    import { onMount } from "svelte";
    import { z } from "zod";

    export let data: ScriptureSection[];
    $: n_blocks = data.reduce((a, c) => a + c.blocks.length, 0) ?? 0;
    $: if (n_blocks == 0) reset();
    $: if (n_blocks > 0) init();

    export let current_verse: ScriptureRef | undefined = undefined;

    const NUM_BLOCK_BUFFER = 20;
    const SCROLL_MAX_HEIGHT = 100_000;
    const SCROLL_BUFFER = 10_000;

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

    let viewport: HTMLDivElement;

    function reset() {
        current_verse = undefined;
        visible_above.length = 0;
        visible_below.length = 0;
        visible_above_elements.length = 0;
        visible_below_elements.length = 0;
        pivot_height = 0;
    }

    function init() {
        reset();
        update({ index: 0, section: 0, block: 0 });
    }

    function refresh() {
        if (n_blocks == 0) return;

        const scroll = viewport.scrollTop;

        let pivot: VirtualBlock;
        let offset = 0;
        let remaining = scroll - pivot_height;

        if (remaining >= 0) {
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

        // Deal with edge-case: fast scrolling towards the top
        if (remaining < 0) {
            const last_above = visible_above_elements.findLast(
                (e) => e != null,
            );
            if (
                (last_above == undefined ||
                    remaining < -last_above.scrollHeight) &&
                pivot.index < NUM_BLOCK_BUFFER
            ) {
                update({ index: 0, section: 0, block: 0 }, 0, 0);
                return;
            }
        }

        update(pivot, offset, remaining);
    }
    const throttled_refresh = throttle(refresh, 100);

    function getVirtualBlocks(
        output: VirtualBlock[],
        start: number,
        length: number = NUM_BLOCK_BUFFER,
    ) {
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

    function update(
        vb: VirtualBlock,
        offset: number = 0,
        remaining: number = 0,
    ) {
        // Update current verse
        const pivot_block = data[vb.section].blocks[vb.block];
        current_verse = pivot_block.verses[0].ref;

        // Load the correct virtual blocks
        const pivot = vb.index;
        const start = Math.max(pivot - NUM_BLOCK_BUFFER, 0);
        visible_above = getVirtualBlocks(visible_above, start, pivot - start);
        visible_below = getVirtualBlocks(visible_below, pivot);

        // Update the pivot height
        if (pivot == 0) {
            pivot_height = 0;
        } else if (pivot == n_blocks - 1) {
            pivot_height = SCROLL_MAX_HEIGHT - viewport.clientHeight;
        } else if (pivot_height + offset < SCROLL_BUFFER) {
            pivot_height = SCROLL_MAX_HEIGHT / 2;
        } else if (pivot_height + offset > SCROLL_MAX_HEIGHT - SCROLL_BUFFER) {
            pivot_height = SCROLL_MAX_HEIGHT / 2;
        } else {
            pivot_height += offset;
        }

        viewport.scrollTop = pivot_height + remaining;
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
                        update({ index: nth_item, section: i, block: j });
                        return;
                    }
                    nth_item += 1;
                }
            } else nth_item += section.blocks.length;
        }
    }

    onMount(() => {
        const events = new EventHandler();
        events.subscribe("window:workspace:text:goto", gotoRef, z.number());
        return () => {
            events.shutdown();
        };
    });
</script>

<div class="container">
    <div class="viewport" bind:this={viewport} on:scroll={throttled_refresh}>
        <div class="content">
            <div
                class="rows"
                style:min-height={`${SCROLL_MAX_HEIGHT}px`}
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
