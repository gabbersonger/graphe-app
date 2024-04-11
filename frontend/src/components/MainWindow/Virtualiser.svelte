<script lang="ts">
    import { onMount, tick } from "svelte";

    export let items: Array<any>;

    let visible = items;

    let start = 0;
    let end = items.length;

    let row_elements = [];
    let row_position_data = [];
    let rows_height = 0;
    let positionable = false;

    let viewport: HTMLElement;

    $: visible = items.slice(start, end).map((data, i) => {
        return { index: i + start, data };
    });

    async function firstLoad() {
        positionable = false;
        start = 0;
        end = items.length;
        positionable = false;
        visible = items.slice(start, end).map((data, i) => {
            return { index: i + start, data };
        });

        await tick();

        rows_height = 0;
        for (let i = 0; i < row_elements.length; i++) {
            row_position_data[i] = {};
            row_position_data[i].height = row_elements[i].offsetHeight;
            row_position_data[i].start = row_elements[i].offsetTop;
            rows_height += row_elements[i].offsetHeight;
        }
        positionable = true;

        refresh();
    }

    function refresh() {
        const scroll = viewport.scrollTop;
        for (let i = 0; i < row_position_data.length; i++) {
            if (row_position_data[i].start > scroll) {
                start = i - 1;
                end = i + 20;
                return;
            }
        }
    }

    onMount(() => firstLoad);

    let resize_width: number;
    let resize_timer: ReturnType<typeof setTimeout>;
    $: if (resize_width) {
        clearTimeout(resize_timer);
        setTimeout(firstLoad, 3);
    }
</script>

<svelte:window />

<div class="container">
    <div class="viewport" bind:this={viewport} on:scroll={refresh}>
        <div class="content" bind:clientWidth={resize_width}>
            <div class="rows" style="min-height: {rows_height}px">
                {#each visible as row (row.index)}
                    <div
                        style={positionable
                            ? `position: absolute; top: ${row_position_data[row.index].start}px;`
                            : ""}
                        bind:this={row_elements[row.index]}
                    >
                        <slot {row} />
                    </div>
                {/each}
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
</style>
