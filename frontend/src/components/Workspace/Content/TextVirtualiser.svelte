<script lang="ts">
    import ScriptureBlock from "@/components/Workspace/Content/ScriptureBlock.svelte";
    import type { data } from "!wails/go/models";
    import type { BibleRef } from "@/lib/Scripture/types";
    import { onMount, tick } from "svelte";
    import { EventsOff, EventsOn } from "!wails/runtime/runtime";
    import { isRefInRange } from "@/lib/Scripture/range";

    const DEFAULT_BLOCKS_DISPLAY = 15;
    const PRELOAD_INTERVAL = 200;
    const PRELOAD_BATCH = 50;
    const RESIZE_DELAY = 50;
    const GOTO_SCROLL_OFFSET = -20;

    export let data: Array<data.ScriptureSection>;
    $: n_blocks = data.reduce((a, c) => a + c.blocks.length, 0) ?? 0;
    $: if (n_blocks == 0) reset();
    $: if (n_blocks > 0) load();

    export let current_verse: BibleRef = undefined;
    let current_block: number = 0;
    let current_offset: number = 0;
    $: if (n_blocks > 0 && visible.length > 0 && visible_elements[0]) {
        const first_block = data[visible[0].section].blocks[visible[0].block];
        const verse_elements = visible_elements[0].querySelectorAll(".verse");
        for (let i = verse_elements.length - 1; i >= 0; i--) {
            const verse_element = verse_elements[i] as HTMLElement;
            if (verse_element && verse_element.offsetTop < current_offset) {
                if (i < first_block.verses.length)
                    current_verse = first_block.verses[i].ref;
                break;
            }
        }
    }

    let viewport: HTMLDivElement;
    let content: HTMLDivElement;

    type VirtualBlock = {
        index: number;
        section: number;
        block: number;
    };

    let mode: "precise" | "locked" = "precise";
    let visible: VirtualBlock[] = [];
    let visible_elements: HTMLDivElement[] = [];
    let content_height: number = 0;
    let locked_block: number = 0;
    let locked_block_offset: number = 0;
    let locked_content_height: number = 0;
    let locked_heights: number[] = [];
    let locked_pre: VirtualBlock[] = [];
    let locked_pre_elements: HTMLDivElement[] = [];

    let registered_width: number = undefined;
    let resize_timeout: ReturnType<typeof setTimeout> = null;
    let preloaded_data: Map<number, number[]> = new Map();
    let preloading: VirtualBlock[] = [];
    let preloading_elements: HTMLDivElement[] = [];
    let preloaded_count: number = 0;
    let preload_interval: ReturnType<typeof setInterval> = null;

    function reset() {
        current_verse = undefined;
        current_block = 0;
        current_offset = 0;
        mode = "precise";
        visible.length = 0;
        visible_elements.length = 0;
        content_height = 0;
        locked_block = 0;
        locked_block_offset = 0;
        locked_content_height = 0;
        locked_heights.length = 0;
        locked_pre.length = 0;
        locked_pre_elements.length = 0;
        registered_width = undefined;
        clearTimeout(resize_timeout);
        preloaded_data.clear();
        preloading.length = 0;
        preloading_elements.length = 0;
        clearInterval(preload_interval);
    }

    function load() {
        if (mode == "precise") {
            getVirtualBlocks(visible, current_block);
        } else if (mode == "locked") {
            getVirtualBlocks(visible, locked_block);
        }

        // Begin preloading if it is not already fully done for the registered_width
        if (preloaded_data.has(registered_width)) {
            const width_data = preloaded_data.get(registered_width);
            for (let i = 1; i < width_data.length; i++) {
                if (width_data[i] == 0) {
                    preloaded_count = i - 1;
                    break;
                }
                if (i == width_data.length - 1) preloaded_count = n_blocks;
            }
            content_height = width_data[preloaded_count];
            if (preloaded_count >= locked_block) unlock();
            if (preloaded_count == n_blocks) return;
        } else {
            preloaded_data.set(registered_width, Array(n_blocks + 1).fill(0));
            preloaded_count = 0;
        }
        preload_interval = setInterval(preloadBatch, PRELOAD_INTERVAL);
    }

    async function preloadBatch() {
        if (!preloaded_data.has(registered_width)) {
            preloaded_data.set(registered_width, Array(n_blocks + 1).fill(0));
            return;
        }
        const width_data = preloaded_data.get(registered_width);

        // Get height of currently preloading
        for (let i = 0; i < preloading.length; i++) {
            if (!preloading_elements[i]) return;
            const block = preloading[i].index;
            content_height =
                width_data[block] + preloading_elements[i].offsetHeight;
            width_data[block + 1] = content_height;
            if (mode == "locked" && current_block <= block) unlock();
        }

        // Finish or set up next batch
        preloaded_count += preloading.length;
        if (preloaded_count == n_blocks) {
            clearInterval(preload_interval);
            preloading.length = 0;
        } else {
            const batch = Math.min(PRELOAD_BATCH, n_blocks - preloaded_count);
            preloading = getVirtualBlocks(preloading, preloaded_count, batch);
            await tick();
        }
    }

    async function refreshPrecise(scroll: number) {
        const width_data = preloaded_data.get(registered_width);
        for (let i = 0; i < preloaded_count; i++) {
            if (width_data[i] > scroll) {
                current_block = Math.max(i - 1, 0);
                current_offset = scroll - width_data[current_block];
                visible = getVirtualBlocks(visible, current_block);
                await tick();
                return;
            }
        }
    }

    async function updateLockedVisible(
        block: number,
        length: number = DEFAULT_BLOCKS_DISPLAY,
    ) {
        visible = getVirtualBlocks(visible, block, length);
        await tick();
        for (let i = 0; i < visible.length; i++) {
            if (!visible_elements[i]) continue;
            locked_heights[locked_block + i] = visible_elements[i].clientHeight;
        }

        const start_locked_pre = Math.max(0, block - DEFAULT_BLOCKS_DISPLAY);
        const diff = block - start_locked_pre;
        locked_pre = getVirtualBlocks(locked_pre, start_locked_pre, diff);
        await tick();
        for (let i = 0; i < locked_pre.length; i++) {
            if (!locked_pre_elements[i]) continue;
            locked_heights[start_locked_pre + i] =
                locked_pre_elements[i].clientHeight;
        }
    }

    async function refreshLocked(scroll: number) {
        const amount_scrolled = Math.abs(scroll - locked_block_offset);
        if (amount_scrolled < 1) {
            return updateLockedVisible(locked_block);
        }

        if (amount_scrolled > viewport.clientHeight * 2) {
            // Estimate
            const average_block = locked_content_height / n_blocks;
            const new_block = Math.floor(scroll / average_block);

            // If estimate falls within the preloaded data -> unlock and load from there
            if (new_block <= preloaded_count) {
                current_block = new_block;
                current_offset =
                    scroll -
                    preloaded_data.get(registered_width)[current_block];
                visible = getVirtualBlocks(visible, current_block);
                unlock();
                await tick();
                return;
            }

            // Otherwise, loading new elements -> update locked block
            locked_block = new_block;
            locked_block_offset = scroll;
            current_block = locked_block;
            current_offset = 0;
            return updateLockedVisible(locked_block);
        }
    }

    function refresh() {
        const scroll = viewport.scrollTop;
        if (mode == "precise") refreshPrecise(scroll);
        else if (mode == "locked") refreshLocked(scroll);
    }

    function resize(width: number) {
        if (width == registered_width) return;
        if (registered_width == undefined) {
            registered_width = width;
            return;
        }

        if (mode == "precise") lock();

        registered_width = width;
        preloaded_count = 0;
        content_height = 0;
        clearInterval(preload_interval);
        clearTimeout(resize_timeout);
        preloading.length = 0;
        resize_timeout = setTimeout(load, RESIZE_DELAY);
    }

    async function goto(block: number) {
        if (block < 0 || block >= n_blocks) return;

        if (block <= preloaded_count) {
            const blockPosition = preloaded_data.get(registered_width)[block];
            const scrollPosition = Math.max(
                blockPosition + GOTO_SCROLL_OFFSET,
                0,
            );
            viewport.scrollTop = scrollPosition + 1;
        } else {
            current_block = block;
            current_offset = 0;
            locked_block = block;
            locked_content_height =
                (n_blocks / preloaded_count) * content_height;
            locked_block_offset = (locked_content_height / n_blocks) * block;
            locked_heights = Array(n_blocks).fill(0);
            mode = "locked";
            await tick();
            viewport.scrollTop = locked_block_offset;
        }
    }

    function gotoRef(ref: BibleRef) {
        let nth_item = 0;
        for (let i = 0; i < data.length; i++) {
            const section = data[i];
            if (isRefInRange(ref, section.range)) {
                for (let j = 0; j < section.blocks.length; j++) {
                    const block = section.blocks[j];
                    if (isRefInRange(ref, block.range)) {
                        return goto(nth_item);
                    }
                    nth_item += 1;
                }
            } else nth_item += section.blocks.length;
        }
    }

    async function lock() {
        if (mode == "locked") return;

        locked_block = current_block;
        locked_block_offset =
            preloaded_data.get(registered_width)[locked_block];
        locked_content_height = (n_blocks / preloaded_count) * content_height;
        locked_heights = Array(n_blocks).fill(0);
        mode = "locked";
        await tick();
        updateLockedVisible(locked_block);
    }

    async function unlock() {
        if (mode != "locked") return;

        mode = "precise";
        await tick();
        viewport.scrollTop =
            preloaded_data.get(registered_width)[current_block] +
            current_offset;

        locked_block = 0;
        locked_block_offset = 0;
        locked_content_height = 0;
        locked_heights.length = 0;
        locked_pre.length = 0;
        locked_pre_elements.length = 0;
    }

    function getVirtualBlocks(
        output: VirtualBlock[],
        start: number,
        length: number = -1,
    ): VirtualBlock[] {
        if (length == -1) length = DEFAULT_BLOCKS_DISPLAY;
        if (start < 0 || start >= n_blocks || length < 0) return;
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

    const resize_observer = new ResizeObserver((e) =>
        resize(e[0].contentRect.width),
    );
    onMount(() => {
        resize_observer.observe(content);
        EventsOn("window:workspace:visualiser:goto", gotoRef);
        return () => {
            resize_observer.unobserve(content);
            EventsOff("window:workspace:visualiser:goto");
        };
    });
</script>

<div class="container">
    <div class="viewport" bind:this={viewport} on:scroll={refresh}>
        <div class="content" bind:this={content}>
            <div
                class="rows"
                style:height={`${mode == "precise" ? content_height : locked_content_height}px`}
            >
                <div
                    class="offseter-pre"
                    style:bottom={`calc(100% - ${locked_block_offset}px`}
                >
                    {#each locked_pre as vb, index (vb.index)}
                        <div bind:this={locked_pre_elements[index]}>
                            <ScriptureBlock
                                block={data[vb.section].blocks[vb.block]}
                            />
                        </div>
                    {/each}
                </div>

                <div
                    class="offseter"
                    style:top={mode == "precise"
                        ? preloaded_data.has(registered_width)
                            ? `${preloaded_data.get(registered_width)[visible[0].index ?? 0]}px`
                            : `0px`
                        : `${locked_block_offset}px`}
                >
                    {#each visible as vb, index (vb.index)}
                        <div bind:this={visible_elements[index]}>
                            <ScriptureBlock
                                block={data[vb.section].blocks[vb.block]}
                            />
                        </div>
                    {/each}
                </div>
            </div>
        </div>
    </div>
</div>

<div class="container" style="z-index:-1">
    <div class="viewport">
        <div class="preloader">
            <div class="rows">
                <div class="offseter">
                    {#each preloading as vb, index (vb.index)}
                        <div bind:this={preloading_elements[index]}>
                            <ScriptureBlock
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

    .offseter,
    .offseter-pre {
        position: absolute;
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
