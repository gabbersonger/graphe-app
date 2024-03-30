<script lang="ts">
    import type { app } from "!wails/go/models";
    import { createVirtualizer } from "@tanstack/svelte-virtual";
    import { bibleRefToString } from "@/lib/Scripture/ref";

    export let text: app.ScriptureSection;

    let virtualListEl: HTMLDivElement;
    let virtualItemEls: HTMLDivElement[] = [];

    let count = text.blocks.length;
    $: virtualizer = createVirtualizer<HTMLDivElement, HTMLDivElement>({
        count,
        getScrollElement: () => virtualListEl,
        estimateSize: () => 80,
    });

    $: items = $virtualizer.getVirtualItems();

    $: {
        if (virtualItemEls.length)
            virtualItemEls.forEach((el) => $virtualizer.measureElement(el));
    }
</script>

<div
    bind:this={virtualListEl}
    style="position: relative; height: {$virtualizer.getTotalSize()}px; width: 100%;"
>
    {#each text.blocks as block, idx}
        <div class="block" bind:this={virtualItemEls[idx]} data-index={idx}>
            <span class="ref">
                {bibleRefToString(block.range.start, "short")}
            </span>
            {#each block.verses as verse, index}
                <div class="verse" style="display: inline">
                    {#if index > 0}
                        <sup>{verse.ref % 1000}</sup>
                    {/if}
                    {#each verse.words as word}
                        <span class="word">{word}</span>{" "}
                    {/each}
                </div>
            {/each}
        </div>
    {/each}
</div>

<style>
    .block {
        padding-bottom: 1rem;
        color: var(--clr-text);
        font-size: 1rem;
        line-height: 1.7;
    }

    .ref {
        font-weight: bold;
        font-size: 0.8rem;
        color: var(--clr-main);
        background: var(--clr-background-sub);
        padding: 0.3rem;
        border-radius: 0.1rem;
    }

    sup {
        vertical-align: super;
        font-weight: bold;
        font-size: 0.7rem;
        color: var(--clr-text-sub);
    }

    .word {
        font-family: "Accordance";
        font-size: 1.2rem;
    }

    .word:hover {
        cursor: pointer;
        background: var(--clr-background-dark);
        outline: 0.2rem solid var(--clr-background-dark);
    }
</style>
