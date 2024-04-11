<script lang="ts">
    import { tick } from "svelte";

    export let items: Array<any>;

    let visible = [];
    let row_elements = [];
    let row_position_data = [];
    let rows_height = 0;
    let viewport: HTMLElement;

    export let current_item = 0;
    let current_offset = 0;

    let last_width = 0;
    let historic_row_position_data = {};

    async function firstLoad() {
        // Save positioning data
        if (row_position_data.length > 0) {
            historic_row_position_data[last_width] = row_position_data;
        }
        if (resize_width in historic_row_position_data) {
            row_position_data = historic_row_position_data[resize_width];
            refresh();
            return;
        }

        visible = items.slice(0, items.length).map((data, i) => {
            return { index: i, data };
        });

        await tick();

        last_width = resize_width;
        rows_height = 0;
        for (let i = 0; i < row_elements.length; i++) {
            row_position_data[i] = {};
            row_position_data[i].height = row_elements[i].offsetHeight;
            row_position_data[i].start = row_elements[i].offsetTop;
            rows_height += row_elements[i].offsetHeight;

            if (i == current_item) {
                viewport.scrollTop =
                    row_position_data[current_item].start + current_offset;
            }
        }

        refresh();
    }

    function refresh() {
        const scroll = viewport.scrollTop;
        for (let i = 0; i < row_position_data.length; i++) {
            if (row_position_data[i].start > scroll) {
                current_item = i - 1;
                current_offset = scroll - row_position_data[current_item].start;
                visible = items
                    .slice(current_item, current_item + 20)
                    .map((data, x) => {
                        return { index: x + current_item, data };
                    });
                return;
            }
        }
    }

    let resize_width: number;
    let resize_timer: ReturnType<typeof setTimeout>;
    $: if (resize_width) {
        clearTimeout(resize_timer);
        resize_timer = setTimeout(firstLoad, 50);
    }
</script>

<svelte:window />

<div class="container">
    <div class="viewport" bind:this={viewport} on:scroll={refresh}>
        <div class="content" bind:clientWidth={resize_width}>
            <div class="rows" style="min-height: {rows_height}px">
                <div
                    class="offseter"
                    style={row_position_data.length > 0 && visible.length > 0
                        ? `top: ${row_position_data[visible[0].index].start}px`
                        : ""}
                >
                    {#each visible as row (row.index)}
                        <div bind:this={row_elements[row.index]}>
                            <slot {row} />
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
