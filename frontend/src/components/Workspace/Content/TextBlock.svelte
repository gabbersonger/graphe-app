<script lang="ts">
    import type { ScriptureBlock } from "!/graphe/internal/data";
    import { GrapheEvent, throttle } from "@/lib/utils";

    export let block: ScriptureBlock;

    let instant_details_timeout: ReturnType<typeof setTimeout>;

    const INSTANT_DETAILS_DELAY = 50;
    function handleMouseEnter(ref: number, word_number: number) {
        clearTimeout(instant_details_timeout);
        instant_details_timeout = setTimeout(() => {
            GrapheEvent("window:workspace:instantdetails", {
                ref: ref,
                word_number: word_number,
            });
        }, INSTANT_DETAILS_DELAY);
    }

    const throttled_handleMouseEnter = throttle(handleMouseEnter, 50);

    function handleMouseLeave() {
        clearTimeout(instant_details_timeout);
        GrapheEvent("window:workspace:instantdetails:hide");
    }
</script>

<div class="block">
    {#if block.verses.length > 0 && "details" in block.verses[0] && block.verses[0].details != undefined}
        {#each block.verses[0]?.details as detail}
            <div class={detail.type == 0 ? "title" : "heading"}>
                {detail.data}
            </div>
        {/each}
    {/if}

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
                        ? (e) =>
                              throttled_handleMouseEnter(
                                  verse.ref,
                                  word.word_num,
                              )
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

    .title {
        display: block;
        text-align: center;
        font-family: var(--font-title);
        font-size: 3em;
        padding-block: 3rem;
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
