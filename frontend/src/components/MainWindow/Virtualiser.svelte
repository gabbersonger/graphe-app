<script lang="ts">
    import { tick } from "svelte";

    type T = $$Generic;
    export let items: Array<T>;

    let visible: Array<{ index: number; data: T }> = [];
    let row_elements: HTMLElement[] = [];
    let row_offsets: number[] = [];
    let rows_height: number = 0;

    let viewport: HTMLElement;

    export let current_item = 0;
    let current_offset = 0;

    let historic_row_offsets = {};
    let content_width: number = undefined;
    let last_calculated_width: number = undefined;
    let last_registered_width: number = undefined;

    function setVisible(start: number, length: number = 20) {
        if (start < 0 || start >= items.length) return;
        if (length < 0 || start + length > items.length) return;
        visible = items.slice(start, start + length).map((data, i) => {
            return { index: i + start, data };
        });
    }

    async function firstLoad() {
        if (content_width in historic_row_offsets) {
            row_offsets = historic_row_offsets[content_width];
            refresh();
            return;
        }

        setVisible(0, items.length);
        await tick();

        if (row_offsets.length == 0) row_offsets = Array(items.length).fill(0);

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
        if (n < 0 || n >= items.length) return;
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
                            <slot row={row.data} />
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
