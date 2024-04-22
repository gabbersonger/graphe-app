<script lang="ts">
    import type { app } from "!wails/go/models";
    import { EventsOff, EventsOn } from "!wails/runtime/runtime";
    import { isRefInRange } from "@/lib/Scripture/range";
    import type { BibleRef } from "@/lib/Scripture/types";
    import { onMount, tick } from "svelte";

    const DEFAULT_ITEMS_VISIBLE = 20;
    const PRELOAD_INTERVAL = 200;
    const PRELOAD_BATCH_SIZE = 50;
    const RESIZE_DELAY = 50;

    export let data: Array<app.ScriptureSection>;
    $: n_blocks =
        data.length > 0 ? data.reduce((a, c) => a + c.blocks.length, 0) : 0;
    $: if (n_blocks != 0) load();
    $: if (n_blocks == 0) clear();

    // TODO: handle exporting current ref
    // export let current_ref: BibleRef = null;
    let current_block: number = 0;
    let current_offset: number = 0;

    type VirtualRow = {
        index: number;
        section: number;
        block: number;
    };
    let visible: VirtualRow[] = [];
    let preloading: VirtualRow[] = [];
    let row_elements: HTMLDivElement[] = [];

    let total_rows_height: number = 0;
    let row_offsets: number[] = [];
    let n_preloaded: number = 0;
    let preload_interval: ReturnType<typeof setInterval> = null;

    type VirtualiserResizeLock = {
        height: number;
        scroll: number;
        block: number;
        block_offset: number;
    };
    let resize_lock: VirtualiserResizeLock = undefined;
    let content_width: number = undefined;
    let computed_row_offsets: Map<number, number[]> = new Map();
    let resize_timeout: ReturnType<typeof setTimeout> = null;

    let viewport: HTMLDivElement;
    let content: HTMLDivElement;

    // Called once for each resize or text change
    function load() {
        if (n_blocks == 0) return;
        visible = setVisible(current_block);

        if (computed_row_offsets.has(content_width)) {
            row_offsets = computed_row_offsets.get(content_width);
            total_rows_height = row_offsets.pop();
            if (resize_lock != undefined) {
                viewport.scrollTop =
                    row_offsets[resize_lock.block] +
                    (viewport.scrollTop - resize_lock.block_offset);
                resize_lock = undefined;
            }
        } else {
            n_preloaded = 0;
            preload_interval = setInterval(preloadBatch, PRELOAD_INTERVAL);
        }
    }

    // Called for each resizing of content bounding box
    function resize() {
        const current_width = content.offsetWidth;
        if (current_width == content_width) return;
        content_width = current_width;
        if (content_width == undefined) {
            content_width = current_width;
            return;
        }

        content_width = current_width;
        if (resize_lock == undefined) {
            resize_lock = {
                height: total_rows_height,
                scroll: viewport.scrollTop,
                block: current_block,
                block_offset: row_offsets[current_block],
            };
        }

        clearTimeout(resize_timeout);
        resize_timeout = setTimeout(load, RESIZE_DELAY);

        clearInterval(preload_interval);
        preload_interval = null;
        total_rows_height = 0;
        row_offsets = [];
        n_preloaded = 0;
    }

    // Called on scroll
    async function refresh() {
        const scroll = viewport.scrollTop;

        if (resize_lock != undefined) {
            // TODO: get scrolling while locked working
            return;
        }

        for (let i = 0; i < row_offsets.length; i++) {
            if (row_offsets[i] > scroll) {
                current_block = i - 1;
                current_offset = scroll - row_offsets[current_block];
                visible = setVisible(current_block);
                await tick();
                return;
            }
        }
    }

    // Called whenever data passed in changes
    function clear() {
        let resize_timeout: ReturnType<typeof setTimeout>;

        current_block = 0;
        current_offset = 0;

        visible = [];
        preloading = [];
        row_elements = [];

        total_rows_height = 0;
        row_offsets = [];
        n_preloaded = 0;
        clearInterval(preload_interval);
        preload_interval = null;

        resize_lock = undefined;
        content_width = undefined;
        computed_row_offsets.clear();
        clearTimeout(resize_timeout);
        resize_timeout = null;
    }

    // Helper functions (below)

    function getVirtualRows(out: VirtualRow[], start: number, length: number) {
        if (start < 0 || start >= n_blocks) return;
        if (length < 0) return;
        if (start + length > n_blocks) length = n_blocks - start;

        const old_length = out.length;
        out.length = length;
        for (let i = old_length; i < length; i++)
            out[i] = { index: 0, section: 0, block: 0 };

        let virtual_rows_filled = 0;
        let curr_section = 0;
        let how_many_to_skip = start;
        while (virtual_rows_filled < length) {
            if (data[curr_section].blocks.length <= how_many_to_skip) {
                how_many_to_skip -= data[curr_section].blocks.length;
                curr_section += 1;
                continue;
            }

            const how_many_possible =
                data[curr_section].blocks.length - how_many_to_skip;
            const how_many_to_take = Math.min(
                length - virtual_rows_filled,
                how_many_possible,
            );
            for (let i = 0; i < how_many_to_take; i++) {
                const n = virtual_rows_filled + i;
                out[n].index = start + n;
                out[n].section = curr_section;
                out[n].block = how_many_to_skip + i;
            }
            curr_section++;
            how_many_to_skip = 0;
            virtual_rows_filled += how_many_to_take;
        }
    }

    function setVisible(start: number, length: number = -1) {
        if (length == -1) length = Math.min(DEFAULT_ITEMS_VISIBLE, n_blocks);
        getVirtualRows(visible, start, length);
        return visible;
    }

    async function preloadBatch() {
        preloading = preloading;
        await tick();

        for (let i = 0; i < preloading.length; i++) {
            if (!row_elements[n_preloaded + i]) return;
            row_offsets[n_preloaded + i] = total_rows_height;
            total_rows_height += row_elements[n_preloaded + i].offsetHeight;

            if (resize_lock && n_preloaded + i == resize_lock.block) {
                viewport.scrollTop =
                    row_offsets[resize_lock.block] +
                    (viewport.scrollTop - resize_lock.block_offset);
                resize_lock = null;
            }
        }

        n_preloaded += preloading.length;
        if (n_preloaded == n_blocks) {
            preloading.length = 0;
            clearInterval(preload_interval);
            row_offsets.push(total_rows_height);
            computed_row_offsets.set(content_width, row_offsets);
            return;
        }

        const batch_size = Math.min(PRELOAD_BATCH_SIZE, n_blocks - n_preloaded);
        getVirtualRows(preloading, n_preloaded, batch_size);
    }

    function goto(ref: BibleRef) {
        let n_seen = 0;
        for (let i = 0; i < data.length; i++) {
            const section = data[i];
            if (isRefInRange(ref, section.range)) {
                for (let j = 0; j < section.blocks.length; j++) {
                    const block = section.blocks[j];
                    if (isRefInRange(ref, block.range)) {
                        if (row_offsets.length < n_seen) return;
                        const scrollPosition = Math.max(
                            0,
                            row_offsets[n_seen] - 20,
                        );
                        viewport.scrollTop = scrollPosition;
                        return;
                    }
                    n_seen += 1;
                }
            } else n_seen += section.blocks.length;
        }
    }

    const resizeObserver = new ResizeObserver((_) => resize());
    onMount(() => {
        resizeObserver.observe(content);
        EventsOn("visualiser:goto", goto);
        return () => {
            resizeObserver.unobserve(content);
            EventsOff("visualiser:goto");
        };
    });
</script>

<div class="container">
    <div class="viewport" bind:this={viewport} on:scroll={refresh}>
        <div class="content" bind:this={content}>
            <div
                class="rows"
                style:min-height={`${resize_lock == undefined ? total_rows_height : resize_lock.height}px`}
            >
                <div
                    class="offseter"
                    style:top={resize_lock == undefined
                        ? row_offsets.length > 0 && visible.length > 0
                            ? `${row_offsets[visible[0].index]}px`
                            : `0px`
                        : `${resize_lock.block_offset}px`}
                >
                    {#each visible as row (row.index)}
                        <slot
                            name="row"
                            row={data[row.section].blocks[row.block]}
                        />
                    {/each}
                </div>
            </div>
        </div>
    </div>
</div>

<div class="container" style="z-index: -1;">
    <div class="viewport">
        <div class="preloader">
            <div class="rows">
                <div class="offseter">
                    {#each preloading as row (row.index)}
                        <div bind:this={row_elements[row.index]}>
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

    .preloader .offseter {
        position: absolute;
        top: 0;
        width: 100%;
        pointer-events: none;
        opacity: 0;
    }
</style>
