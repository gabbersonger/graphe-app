<script lang="ts">
    import type { app } from "!wails/go/models";
    import type { BibleRef } from "@/lib/Scripture/types";

    import { onMount, tick } from "svelte";
    import { EventsOff, EventsOn } from "!wails/runtime/runtime";
    import { isRefInRange } from "@/lib/Scripture/range";

    const DEFAULT_BLOCKS_DISPLAY = 15;
    const PRELOAD_INTERVAL = 200; // TODO: change to 200
    const PRELOAD_BATCH_SIZE = 50;
    const RESIZE_DELAY = 50;
    const GOTO_SCROLL_OFFSET = -20;

    export let data: Array<app.ScriptureSection>;
    $: n_blocks =
        data.length > 0 ? data.reduce((a, c) => a + c.blocks.length, 0) : 0;
    $: if (n_blocks != 0) load();
    $: if (n_blocks == 0) clear();

    let mode: "precise" | "locked" = "precise";
    let current_block: number = 0;
    let current_offset: number = 0;

    export let current_position: { section: number; block: number } = undefined;
    $: if (current_block >= 0 && n_blocks > 0) {
        let remaining = current_block;
        for (let i = 0; i < data.length; i++) {
            if (remaining < data[i].blocks.length) {
                if (current_position == undefined)
                    current_position = { section: i, block: remaining };
                else {
                    current_position.section = i;
                    current_position.block = remaining;
                }
                break;
            }
            remaining -= data[i].blocks.length;
        }
    }

    type VirtualRow = {
        index: number;
        section: number;
        block: number;
    };
    let precise_visible: VirtualRow[] = [];
    let locked_visible: VirtualRow[] = [];
    let preloading: VirtualRow[] = [];

    let precise_computed_offsets: Map<number, number[]> = new Map();
    let precise_content_height: number = 0;
    let precise_content_offsets: number[] = [];

    let locked_content_height: number = 0;
    let locked_block: number = 0;
    let locked_block_offset: number = 0;
    let locked_elements: HTMLDivElement[] = [];
    let locked_backwards_offsets: number[] = [];
    let locked_forwards_offsets: number[] = [];
    let locked_offset_gap: number = 0;

    let preloading_elements: HTMLDivElement[] = [];
    let preloading_done: number = 0;
    let preloading_interval: ReturnType<typeof setInterval> = null;

    let registered_width: number = undefined;
    let resize_timeout: ReturnType<typeof setTimeout> = null;

    let viewport: HTMLDivElement;
    let content: HTMLDivElement;

    async function load() {
        if (n_blocks == 0) return;

        // Check if we already know the precise information
        if (precise_computed_offsets.has(registered_width)) {
            precise_content_offsets =
                precise_computed_offsets.get(registered_width);
            precise_content_height = precise_content_offsets.pop();
            return unlock();
        }

        // Load the rows in and start preloading all rows
        if (mode == "precise") {
            getVirtualRows(precise_visible, current_block);
        } else if (mode == "locked") {
            getVirtualRows(
                locked_visible,
                Math.max(0, locked_block - DEFAULT_BLOCKS_DISPLAY),
                DEFAULT_BLOCKS_DISPLAY * 2,
            );
            locked_visible = locked_visible;
            await tick();
        }

        clearInterval(preloading_interval);
        preloading_done = 0;
        preloading_interval = setInterval(preloadBatch, PRELOAD_INTERVAL);
    }

    function clear() {
        mode = "precise";
        current_block = 0;
        current_offset = 0;

        current_position = undefined;

        precise_visible = [];
        locked_visible = [];
        preloading = [];

        precise_computed_offsets.clear();
        precise_content_height = 0;
        precise_content_offsets = [];

        locked_content_height = 0;
        locked_block = 0;
        locked_block_offset = 0;
        locked_elements = [];
        locked_backwards_offsets = [];
        locked_forwards_offsets = [];
        locked_offset_gap = 0;

        preloading_elements = [];
        preloading_done = 0;
        clearInterval(preloading_interval);
        preloading_interval = null;

        registered_width = undefined;
        clearTimeout(resize_timeout);
        resize_timeout = null;
    }

    function resize() {
        // Prevent unnecessary firing - on init or height change
        const content_width = content.offsetWidth;
        if (content_width == registered_width) return;
        if (registered_width == undefined) {
            registered_width = content_width;
            return;
        }
        registered_width = content_width;

        if (mode == "precise") {
            locked_content_height = precise_content_height;
            locked_block = current_block;
            locked_block_offset = precise_content_offsets[current_block];
            locked_elements.length = 0;
            locked_backwards_offsets = [];
            locked_forwards_offsets = [];
            locked_offset_gap = 0;
            mode = "locked";
        }

        precise_content_height = 0;
        clearTimeout(resize_timeout);
        resize_timeout = setTimeout(load, RESIZE_DELAY);

        clearInterval(preloading_interval);
        preloading_done = 0;
    }

    async function refresh() {
        const scroll = viewport.scrollTop;

        if (mode == "precise") {
            for (let i = 0; i < precise_content_offsets.length; i++) {
                if (precise_content_offsets[i] > scroll) {
                    current_block = i - 1;
                    current_offset =
                        scroll - precise_content_offsets[current_block];
                    getVirtualRows(precise_visible, current_block);

                    precise_visible = precise_visible;
                    await tick();
                    return;
                }
            }
        } else if (mode == "locked") {
            // Save the heights of locked elements loaded in
            for (let i = 0; i < locked_elements.length; i++) {
                const block_index = locked_visible[i].index;
                const offset_from_locked_block = block_index - locked_block;
                if (offset_from_locked_block >= 0) {
                    locked_forwards_offsets[offset_from_locked_block] =
                        locked_elements[i].offsetHeight;
                } else if (offset_from_locked_block < 0) {
                    locked_backwards_offsets[-offset_from_locked_block - 1] =
                        locked_elements[i].offsetHeight;
                }
            }

            const amount_scrolled = scroll - locked_block_offset;
            if (amount_scrolled == 0) return;

            const direction_offsets =
                amount_scrolled > 0
                    ? locked_forwards_offsets
                    : locked_backwards_offsets;
            const direction = amount_scrolled > 0 ? 1 : -1;

            let new_center = undefined;
            let amount_left = Math.abs(amount_scrolled);
            for (let i = 0; i < direction_offsets.length; i++) {
                if (amount_left - direction_offsets[i] <= 0) {
                    new_center = locked_block + direction * i;
                    if (direction == -1) new_center--;
                    break;
                }
                amount_left -= direction_offsets[i];
            }
            current_block = new_center;
            current_offset = amount_left;

            const new_start = Math.max(new_center - DEFAULT_BLOCKS_DISPLAY, 0);
            const length = DEFAULT_BLOCKS_DISPLAY * 2;
            const new_end = new_start + length - 1;
            getVirtualRows(locked_visible, new_start, length);

            if (direction > 0 && new_start > locked_block) {
                locked_offset_gap = locked_forwards_offsets
                    .slice(0, new_start - locked_block)
                    .reduce((a, b) => a + b, 0);
            } else if (direction < 0 && new_end < locked_block - 1) {
                locked_offset_gap =
                    -1 *
                    locked_backwards_offsets
                        .slice(0, locked_block - 1 - new_end)
                        .reduce((a, b) => a + b, 0);
            }

            locked_visible = locked_visible;
            locked_offset_gap = locked_offset_gap;
            await tick();
        }
    }

    async function scrollTo(index: number) {
        if (index < 0 || index >= n_blocks) return;

        if (index <= preloading_done) {
            const scrollPosition = Math.max(
                precise_content_offsets[index] + GOTO_SCROLL_OFFSET,
                0,
            );
            viewport.scrollTop = scrollPosition + 1;
        } else {
            const estimate_content_height =
                (n_blocks / preloading_done) * precise_content_height;
            const estimate_block_offset =
                (index / n_blocks) * estimate_content_height;
            locked_content_height = estimate_content_height;
            locked_block = index;
            locked_block_offset = estimate_block_offset;
            locked_elements.length = 0;
            locked_backwards_offsets = [];
            locked_forwards_offsets = [];
            locked_offset_gap = 0;
            mode = "locked";
            await tick();
            viewport.scrollTop = estimate_block_offset + 1;

            precise_content_height = 0;
            load();
        }
    }

    // Helper functions

    async function unlock() {
        mode = "precise";
        await tick();
        viewport.scrollTop =
            precise_content_offsets[current_block] + current_offset;

        locked_visible.length = 0;
        locked_content_height = 0;
        locked_block = 0;
        locked_block_offset = 0;
        locked_elements.length = 0;
        locked_backwards_offsets.length = 0;
        locked_forwards_offsets.length = 0;
        locked_offset_gap = 0;
    }

    function getVirtualRows(
        output: VirtualRow[],
        start: number,
        length: number = -1,
    ) {
        if (length == -1) length = DEFAULT_BLOCKS_DISPLAY;
        if (start < 0 || start >= n_blocks || length < 0) return;
        if (start + length > n_blocks) length = n_blocks - start;

        const old_output_length = output.length;
        output.length = length;
        for (let i = old_output_length; i < length; i++)
            output[i] = { index: 0, section: 0, block: 0 };

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
    }

    async function preloadBatch() {
        preloading = preloading;
        await tick();

        for (let i = 0; i < preloading.length; i++) {
            if (!preloading_elements[i]) return;
            const nth_preload = preloading_done + i;
            precise_content_offsets[nth_preload] = precise_content_height;
            precise_content_height += preloading_elements[i].offsetHeight;

            if (mode == "locked" && locked_block == nth_preload) unlock();
        }

        preloading_done += preloading.length;
        if (preloading_done == n_blocks) {
            clearInterval(preloading_interval);
            preloading.length = 0;
            precise_content_offsets.push(precise_content_height);
            precise_computed_offsets.set(
                registered_width,
                precise_content_offsets,
            );
            return;
        }

        const batch = Math.min(PRELOAD_BATCH_SIZE, n_blocks - preloading_done);
        getVirtualRows(preloading, preloading_done, batch);
    }

    function gotoRef(ref: BibleRef) {
        let nth_item = 0;
        for (let i = 0; i < data.length; i++) {
            const section = data[i];
            if (isRefInRange(ref, section.range)) {
                for (let j = 0; j < section.blocks.length; j++) {
                    const block = section.blocks[j];
                    if (isRefInRange(ref, block.range)) {
                        return scrollTo(nth_item);
                    }
                    nth_item += 1;
                }
            } else nth_item += section.blocks.length;
        }
    }

    // Initialise listeners
    const resizeObserver = new ResizeObserver((_) => resize());
    onMount(() => {
        resizeObserver.observe(content);
        EventsOn("visualiser:goto", gotoRef);
        return () => {
            resizeObserver.unobserve(content);
            EventsOff("visualiser:goto");
        };
    });
</script>

<div class="container">
    <div class="viewport" bind:this={viewport} on:scroll={refresh}>
        <div class="content" bind:this={content}>
            {#if mode == "locked"}
                <div class="rows" style:height={`${locked_content_height}px`}>
                    <div
                        class="offseter"
                        style:bottom={`${locked_content_height - locked_block_offset}px`}
                    >
                        {#each locked_visible as row, index (row.index)}
                            {#if row.index < locked_block}
                                <div bind:this={locked_elements[index]}>
                                    <slot
                                        name="row"
                                        row={data[row.section].blocks[
                                            row.block
                                        ]}
                                    />
                                </div>
                            {/if}
                        {/each}

                        {#if locked_offset_gap < 0}
                            <div
                                class="offseter-gap"
                                style:height={`${Math.abs(locked_offset_gap)}px`}
                            />
                        {/if}
                    </div>
                    <div
                        class="offseter"
                        style:top={`${locked_block_offset}px`}
                    >
                        {#if locked_offset_gap > 0}
                            <div
                                class="offseter-gap"
                                style:height={`${locked_offset_gap}px`}
                            />
                        {/if}

                        {#each locked_visible as row, index (row.index)}
                            {#if row.index >= locked_block}
                                <div bind:this={locked_elements[index]}>
                                    <slot
                                        name="row"
                                        row={data[row.section].blocks[
                                            row.block
                                        ]}
                                    />
                                </div>
                            {/if}
                        {/each}
                    </div>
                </div>
            {:else if mode == "precise"}
                <div
                    class="rows"
                    style:min-height={`${precise_content_height}px`}
                >
                    <div
                        class="offseter"
                        style:top={precise_content_offsets.length > 0 &&
                        precise_visible.length > 0
                            ? `${precise_content_offsets[precise_visible[0].index]}px`
                            : `0px`}
                    >
                        {#each precise_visible as row (row.index)}
                            <slot
                                name="row"
                                row={data[row.section].blocks[row.block]}
                            />
                        {/each}
                    </div>
                </div>
            {/if}
        </div>
    </div>
</div>

<div class="container" style="z-index: -1;">
    <div class="viewport">
        <div class="preloader">
            <div class="rows">
                <div class="offseter">
                    {#each preloading as row, index (row.index)}
                        <div bind:this={preloading_elements[index]}>
                            <slot
                                name="row"
                                row={data[row.section].blocks[row.block]}
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
        width: 10px;
    }

    .viewport::-webkit-scrollbar-corner {
        background: transparent;
    }

    .viewport::-webkit-scrollbar-thumb {
        background: var(--clr-text-sub);
        -webkit-transition: 0.125s;
        transition: 0.125s;
        border-radius: 4px !important;
    }

    .viewport::-webkit-scrollbar-track {
        background: 0 0;
    }

    .content,
    .preloader {
        position: relative;
        width: 90%;
        max-width: 80ch;
        margin: 0 auto;
    }

    .rows {
        position: relative;
    }

    .offseter {
        position: absolute;
        width: 100%;
    }

    .offseter-gap {
        position: relative;
        width: 100%;
    }

    .preloader .offseter {
        position: absolute;
        top: 0;
        width: 100%;
        pointer-events: none;
        opacity: 0;
    }
</style>
