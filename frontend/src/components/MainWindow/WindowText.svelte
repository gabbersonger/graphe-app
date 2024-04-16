<script lang="ts">
    import type { app } from "!wails/go/models";
    import { bibleRefToString, isRefBookStart } from "@/lib/Scripture/ref";
    import Virtualiser from "@/components/MainWindow/Virtualiser.svelte";
    import { app_currRefVisible } from "@/lib/stores";
    import type { BibleRef } from "@/lib/Scripture/types";

    export let text: app.ScriptureSection;

    let current_item: number;
    $: if (current_item) {
        $app_currRefVisible = text.blocks[current_item].range.start;
    }

    let scrollVirtualiser: (_: number) => void;
</script>

<Virtualiser
    items={text.blocks}
    bind:current_item
    bind:scrollToItem={scrollVirtualiser}
    let:row
>
    <div class="block">
        {#if isRefBookStart(row.range.start)}
            <div class="heading">
                {bibleRefToString(row.range.start, "book")}
            </div>
        {/if}

        <span class="ref">
            {bibleRefToString(row.range.start, "short")}
        </span>
        {#each row.verses as verse, index}
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
</Virtualiser>

<style>
    .block {
        display: block;
        padding-bottom: 1rem;
        color: var(--clr-text);
        font-size: 1rem;
        line-height: 1.7;
    }

    .heading {
        display: block;
        text-align: center;
        font-family: "Neuton";
        font-size: 2em;
        padding-block: 2rem;
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
</style>
