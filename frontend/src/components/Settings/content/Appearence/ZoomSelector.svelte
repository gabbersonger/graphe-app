<script lang="ts">
    import { Slider } from "bits-ui";
    import { graphe_settings } from "@/lib/stores";
    import { GrapheEvent } from "@/lib/utils";

    $: value = [$graphe_settings?.appearence.zoom ?? 100];

    function onZoomChange(zoom: number) {
        GrapheEvent("graphe:setting", {
            setting: ["appearence", "zoom"],
            value: zoom,
        });
    }
</script>

<div class="wrapper">
    <Slider.Root
        min={50}
        max={200}
        step={10}
        bind:value
        onValueChange={(v) => onZoomChange(v[0])}
        let:thumbs
        let:ticks
        class="root"
    >
        {#each ticks as tick}
            <Slider.Tick {tick} asChild let:builder>
                <div use:builder.action {...builder} class="tick">
                    {#if ![60, 70, 90, 110, 130, 140, 160, 170, 180, 190].includes(tick["data-value"])}
                        <div class="text">
                            {tick["data-value"]}%
                        </div>
                    {/if}
                </div>
            </Slider.Tick>
        {/each}
        <span class="bar">
            <Slider.Range class="range" />
        </span>
        {#each thumbs as thumb}
            <Slider.Thumb {thumb} class="thumb" />
        {/each}
    </Slider.Root>
</div>

<style>
    .wrapper {
        position: relative;
        padding-top: 10px;
        padding-bottom: 20px;
    }

    :global(.root) {
        position: relative;
        width: 100%;
        display: flex;
        align-items: center;
    }

    .bar {
        position: relative;
        height: 5px;
        width: 100%;
        overflow: hidden;
        border-radius: 5px;
        background: var(--clr-background-sub);
    }

    :global(.range) {
        position: absolute;
        height: 100%;
        background: var(--clr-main);
    }

    :global(.thumb) {
        display: block;
        width: 10px;
        height: 18px;
        cursor: pointer;
        border-radius: 5px;
        background: var(--clr-main);
        outline: none;
    }

    :global(.tick) {
        display: block;
        width: 2px;
        height: 15px;
        cursor: pointer;
        background: var(--clr-background-sub);
    }

    :global(.tick) :global(.text) {
        position: absolute;
        top: 20px;
        transform: translateX(calc(-50% + 1.5px));
        font-family: var(--font-system);
        font-size: 0.6rem;
        color: var(--clr-text);
    }
</style>
