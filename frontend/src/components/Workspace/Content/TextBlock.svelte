<script lang="ts">
    import type { ScriptureBlock } from "!/graphe/internal/data";
    import { GrapheEvent } from "@/lib/utils";

    export let block: ScriptureBlock;

    function handleMouseEnter(ref: number, word_number: number) {
        GrapheEvent("window:workspace:instantdetails", {
            ref: ref,
            word_number: word_number,
        });
    }

    function handleMouseLeave() {
        GrapheEvent("window:workspace:instantdetails:hide");
    }
</script>

<div class="block">
    {#each block.verses as verse, index}
        <div class="verse">
            {#if !verse.continuation}
                {#if index == 0}
                    <span class="ref">{verse.ref_string}</span>
                {:else}
                    <sup>{verse.ref_string}</sup>
                {/if}
            {/if}

            {#each verse.words as word}
                {word.pre}<span
                    class="word"
                    role="tooltip"
                    class:hoverable={word.has_instant_details}
                    on:mouseenter={word.has_instant_details
                        ? (e) => handleMouseEnter(verse.ref, word.word_num)
                        : null}
                    on:mouseleave={word.has_instant_details
                        ? handleMouseLeave
                        : null}>{word.text}</span
                >{word.post}{word.post != "-" ? " " : ""}
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

    .hoverable:hover {
        cursor: pointer;
        background: var(--clr-main);
        color: var(--clr-background);
        outline: 0.5ch solid var(--clr-main);
        border-radius: 0.15ch;
    }
</style>
