<script lang="ts">
    import type { app } from "!wails/go/models";
    import { bibleRefToString } from "@/lib/Scripture/ref";
    import Virtualiser from "./Virtualiser.svelte";

    export let text: app.ScriptureSection;
</script>

<Virtualiser items={text.blocks} let:row>
    <div class="block">
        <span class="ref">
            {bibleRefToString(row.data.range.start, "short")}
        </span>
        {#each row.data.verses as verse, index}
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
</style>
