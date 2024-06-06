<script lang="ts">
    import type { database } from "!wails/go/models";
    import { app_version } from "@/lib/appManager";
    import type { BibleRef } from "@/lib/Scripture/types";
    import { getChapter, getVerse, refToString } from "@/lib/Scripture/ref";
    import { EventsEmit } from "!wails/runtime/runtime";

    export let block: database.ScriptureBlock;

    let instant_details_timeout: ReturnType<typeof setTimeout> = null;

    const INSTANT_DETAILS_DELAY = 50;
    function handleMouseEnter(ref: BibleRef, word_num: number) {
        clearTimeout(instant_details_timeout);
        instant_details_timeout = setTimeout(() => {
            EventsEmit("app:instantdetails", ref, word_num);
        }, INSTANT_DETAILS_DELAY);
    }

    function handleMouseLeave() {
        clearTimeout(instant_details_timeout);
        EventsEmit("app:instantdetails:hide");
    }
</script>

<div class="block">
    {#if "details" in block.verses[0]}
        {#each block.verses[0].details as detail}
            <div class={detail.type == 0 ? "title" : "heading"}>
                {detail.data}
            </div>
        {/each}
    {/if}

    {#if !("continuation" in block.verses[0])}
        <span class="ref">
            {refToString($app_version, block.range.start, "short")}
        </span>
    {/if}

    {#each block.verses as verse, index}
        <div class="verse">
            {#if index > 0}
                <sup>
                    {getVerse(verse.ref) == 1
                        ? `${getChapter(verse.ref)}:${getVerse(verse.ref)}`
                        : getVerse(verse.ref)}
                </sup>
            {/if}
            {#each verse.words as word}
                {word.pre}<span
                    class="word"
                    class:hoverable={!("no_instant_details" in word)}
                    on:mouseenter={"no_instant_details" in word
                        ? null
                        : (e) => handleMouseEnter(verse.ref, word.word_num)}
                    on:mouseleave={"no_instant_details" in word
                        ? null
                        : handleMouseLeave}>{word.text}</span
                >{word.post}{word.post != "-" ? " " : ""}
                {#if "details" in word}
                    {#each word.details as detail}
                        {#if detail.type == 0}
                            <br />
                        {:else if detail.type == 1}
                            <!-- TODO: handle indent -->
                        {:else if detail.type == 2}
                            <!-- TODO: handle footnote -->
                        {:else if detail.type == 3}
                            <!-- TODO: handle crossref -->
                        {/if}
                    {/each}
                {/if}
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
    }

    .title {
        display: block;
        text-align: center;
        font-family: var(--font-title);
        font-size: 3em;
        padding-block: 3rem;
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
        font-family: var(--font-content);
        font-size: 1.2rem;
    }

    .hoverable:hover {
        cursor: pointer;
        background: var(--clr-main);
        color: var(--clr-background);
        outline: 0.5ch solid var(--clr-main);
        border-radius: 0.15ch;
    }

    *::selection {
        background: var(--clr-selection);
        color: var(--clr-background);
    }
</style>
