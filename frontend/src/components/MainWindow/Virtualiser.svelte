<script lang="ts">
    import type { app } from "!wails/go/models";
    import { tick } from "svelte";

    export let items: Array<app.ScriptureSection>;
    let num_blocks = items.reduce((acc, cur) => acc + cur.blocks.length, 0);

    let visible: Array<{
        index: number;
        section: number;
        block: number;
    }> = [];
    let row_elements: HTMLElement[] = [];
    let row_offsets: number[] = [];
    let rows_height: number = 0;

    let viewport: HTMLElement;

    let current_item = 0;
    let current_offset = 0;

    export let current_position: { section: number; block: number };
    let last_registered_current_item = 0;
    $: if (current_item != last_registered_current_item) {
        last_registered_current_item = current_item;
        let remaining = current_item;
        for (let i = 0; i < items.length; i++) {
            if (remaining < items[i].blocks.length) {
                current_position.section = i;
                current_position.block = remaining;
                break;
            }
            remaining -= items[i].blocks.length;
        }
    }

    let historic_row_offsets = {};
    let content_width: number = undefined;
    let last_calculated_width: number = undefined;
    let last_registered_width: number = undefined;

    function setVisible(start: number, length: number = -1) {
        if (length == -1) length = Math.min(20, num_blocks);
        if (start < 0 || start >= num_blocks) return;
        if (length < 0 || start + length > num_blocks) return;

        visible.length = length;
        let curr_visible = 0;
        let curr_section = 0;
        let how_many_to_skip = start;
        while (curr_visible < length) {
            if (items[curr_section].blocks.length <= how_many_to_skip) {
                how_many_to_skip -= items[curr_section].blocks.length;
                curr_section += 1;
                continue;
            }

            const how_many_possible =
                items[curr_section].blocks.length - how_many_to_skip;
            const how_many_to_take = Math.min(
                length - curr_visible,
                how_many_possible,
            );
            for (let i = 0; i < how_many_to_take; i++) {
                visible[curr_visible + i] = {
                    index: start + curr_visible + i,
                    section: curr_section,
                    block: how_many_to_skip + i,
                };
            }
            curr_section++;
            how_many_to_skip = 0;
            curr_visible += how_many_to_take;
        }
    }

    async function firstLoad() {
        if (content_width in historic_row_offsets) {
            row_offsets = historic_row_offsets[content_width];
            refresh();
            return;
        }

        setVisible(0, num_blocks);
        await tick();

        if (row_offsets.length == 0) row_offsets = Array(num_blocks).fill(0);

        last_calculated_width = content_width;
        rows_height = 0;
        for (let i = 0; i < row_elements.length; i++) {
            row_offsets[i] = rows_height;
            rows_height += row_elements[i].offsetHeight;
        }

        viewport.scrollTop = row_offsets[current_item] + current_offset;
        refresh();
    }

    function refresh() {
        const scroll = viewport.scrollTop;
        for (let i = 0; i < row_offsets.length; i++) {
            if (row_offsets[i] > scroll) {
                current_item = i - 1;
                current_offset = scroll - row_offsets[current_item];
                setVisible(current_item);
                return;
            }
        }
    }

    export function scrollToItem(n: number) {
        if (n < 0 || n >= num_blocks) return;
        let scrollPosition = Math.max(row_offsets[n] - 20, 0);
        viewport.scrollTop = scrollPosition;
    }

    let resize_timer: ReturnType<typeof setTimeout>;
    $: if (content_width && content_width != last_registered_width) {
        last_registered_width = content_width;
        clearTimeout(resize_timer);
        if (
            last_calculated_width &&
            !(last_calculated_width in historic_row_offsets)
        )
            historic_row_offsets[last_calculated_width] = row_offsets;
        resize_timer = setTimeout(firstLoad, 50);
    }
</script>

<div class="container">
    <div class="viewport" bind:this={viewport} on:scroll={refresh}>
        <div class="content" bind:clientWidth={content_width}>
            <div class="rows" style={`min-height: ${rows_height}px`}>
                <div
                    class="offseter"
                    style={row_offsets.length > 0 && visible.length > 0
                        ? `top: ${row_offsets[visible[0].index]}px`
                        : ""}
                >
                    {#each visible as row (row.index)}
                        <div bind:this={row_elements[row.index]}>
                            <slot row={items[row.section].blocks[row.block]} />
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

    .content {
        position: relative;
        width: 90%;
        max-width: 80ch;
        margin: 0 auto;
    }

    .rows {
        position: relative;
        /* top: 0;
        left: 0;
        width: 100%; */
    }

    .offseter {
        position: absolute;
        top: 0;
        width: 100%;
    }
</style>
