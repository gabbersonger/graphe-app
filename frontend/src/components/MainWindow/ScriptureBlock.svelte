<script lang="ts">
    import type { app } from "!wails/go/models";
    import { EventsEmit } from "!wails/runtime/runtime";
    import { versionData } from "@/lib/Scripture/data";
    import { getBook, isRefBookStart, refToString } from "@/lib/Scripture/ref";
    import type { BibleRef } from "@/lib/Scripture/types";
    import { getVersionBookIndex } from "@/lib/Scripture/version";
    import { app_version } from "@/lib/appManager";

    export let block: app.ScriptureBlock;

    let instant_details_timer: ReturnType<typeof setTimeout>;
    const INSTANT_DETAILS_DELAY = 50;
    function handleWordMouseEnter(ref: BibleRef, word_num: number) {
        clearTimeout(instant_details_timer);
        instant_details_timer = setTimeout(
            () => EventsEmit("app:instantdetails", ref, word_num),
            INSTANT_DETAILS_DELAY,
        );
    }

    function handleWordMouseLeave(e: MouseEvent) {
        clearTimeout(instant_details_timer);
        EventsEmit("app:instantdetails:hide");
    }
</script>

<div class="block">
    {#if isRefBookStart($app_version, block.range.start)}
        <div class="heading">
            {versionData[$app_version].books[
                getVersionBookIndex($app_version, getBook(block.range.start))
            ].display_name}
        </div>
    {/if}

    <span class="ref">
        {refToString($app_version, block.range.start, "short")}
    </span>
    {#each block.verses as verse, index}
        <div class="verse" style="display: inline">
            {#if index > 0}
                <sup>{verse.ref % 1000}</sup>
            {/if}
            {#each verse.words as word}
                {word.pre}<span
                    class="word"
                    class:hoverable={!("no_instant_details" in word)}
                    on:mouseenter={"no_instant_details" in word
                        ? null
                        : (e) => handleWordMouseEnter(verse.ref, word.word_num)}
                    on:mouseleave={"no_instant_details" in word
                        ? null
                        : handleWordMouseLeave}>{word.text}</span
                >{word.post}{" "}
            {/each}
        </div>
    {/each}
</div>

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
        font-family: var(--font-heading);
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
</style>
