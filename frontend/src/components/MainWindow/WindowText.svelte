<script lang="ts">
    import type { app } from "!wails/go/models";
    import { bibleRefToString } from "@/lib/Scripture/ref";
    import { ui_currentRef } from "@/stores/app";
    import { createVirtualizer } from "@tanstack/svelte-virtual";
    import type { ScrollEvents } from "lucide-svelte/dist/icons/scroll.svelte";
    import { onMount } from "svelte";

    export let text: app.ScriptureSection;

    let virtualListEl: HTMLDivElement;
    let virtualItemEls: HTMLDivElement[] = [];

    const count = text.blocks.length;
    $: virtualizer = createVirtualizer<HTMLDivElement, HTMLDivElement>({
        count,
        getScrollElement: () => virtualListEl,
        estimateSize: (index: number) => {
            let num_char = text.blocks[index].verses.reduce(
                (acc, cur) => acc + cur.words.length,
                0,
            );
            return Math.ceil(num_char / 120) * 25 + 16;
        },
        overscan: 2000,
    });

    $: virtualItems = $virtualizer.getVirtualItems();

    $: {
        if (virtualItemEls.length)
            virtualItemEls.forEach((el) => $virtualizer.measureElement(el));
    }

    let scrollTimeout: ReturnType<typeof setTimeout>;
    function onScroll() {
        clearTimeout(scrollTimeout);
        scrollTimeout = setTimeout(() => {
            const block = text.blocks[$virtualizer.range.startIndex];
            $ui_currentRef = block.range.start;
        }, 30);
    }
</script>

<div class="container">
    <div class="wrapper" bind:this={virtualListEl} on:scroll={onScroll}>
        <div class="sizer" style={`height: ${$virtualizer.getTotalSize()}px`}>
            <div
                style={`position: absolute; top: 0; left: 0; width: 100%; transform: translateY(${
                    virtualItems[0]
                        ? virtualItems[0].start -
                          $virtualizer.options.scrollMargin
                        : 0
                }px);`}
            >
                {#each virtualItems as row, idx (row.index)}
                    <div
                        class="block"
                        bind:this={virtualItemEls[idx]}
                        data-index={idx}
                    >
                        <span class="ref">
                            {bibleRefToString(
                                text.blocks[row.index].range.start,
                                "short",
                            )}
                        </span>
                        {#each text.blocks[row.index].verses as verse, index}
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
        </div>
    </div>
</div>

<style>
    .container {
        position: absolute;
        inset: 0;
    }

    .wrapper {
        --size-max-width: 80ch;

        position: relative;
        width: 100%;
        height: 100%;
        overflow-y: scroll;
    }

    .sizer {
        position: relative;
        width: 90%;
        max-width: var(--size-max-width);
        margin: 0 auto;
    }

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
