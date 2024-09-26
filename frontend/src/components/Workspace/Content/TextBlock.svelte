<script lang="ts">
    import type { ScriptureBlock } from "!/graphe/internal/data";

    export let block: ScriptureBlock;
</script>

<div class="block">
    {#each block.verses as verse, index}
        <div class="verse">
            {#if index == 0}
                <span class="ref">{verse.ref_string}</span>
            {:else}
                <sup>{verse.ref_string}</sup>
            {/if}

            {#each verse.words as word}
                {word.pre}<span class="word">{word.text}</span>{word.post}{" "}
            {/each}
        </div>
    {/each}
</div>

<style>
    .block {
        position: relative;
        display: block;
        padding-bottom: 1rem;
        color: var(--clr-text);
        font-size: 1rem;
        line-height: 1.7;
    }

    .verse {
        display: inline;
        font-family: var(--font-content);
        font-size: 1.2rem;
    }

    .ref {
        font-family: var(--font-system);
        font-weight: bold;
        font-size: 0.8rem;
        color: var(--clr-main);
        background: var(--clr-background-sub);
        padding: 0.3rem;
        border-radius: 0.1rem;
    }

    sup {
        vertical-align: super;
        font-family: var(--font-system);
        font-weight: bold;
        font-size: 0.7rem;
        color: var(--clr-text-sub);
    }
</style>
