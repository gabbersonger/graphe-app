<script lang="ts">
    import Virtualiser from "@/components/MainWindow/Virtualiser.svelte";
    import { app_data, app_currentRef } from "@/lib/appManager";
    import {
        bibleRefToString,
        bibleRefToVersionBookTitle,
        isRefBookStart,
    } from "@/lib/Scripture/ref";
    import InstantDetails from "@/components/MainWindow/InstantDetails.svelte";
    import type { BibleRef } from "@/lib/Scripture/types";
    import { EventsEmit } from "!wails/runtime/runtime";

    let cp = { section: 0, block: 0 };
    $: if ($app_data.length > 0 && cp)
        $app_currentRef = $app_data[cp.section].blocks[cp.block].range.start;

    function callInstantDetails(ref: BibleRef, word_num: number) {
        EventsEmit("app:instantdetails", ref, word_num);
    }

    let instant_details_timer: ReturnType<typeof setTimeout>;
    const INSTANT_DETAILS_DELAY = 50; // unit: ms
    function handleWordMouseEnter(ref: BibleRef, word_num: number) {
        clearTimeout(instant_details_timer);
        instant_details_timer = setTimeout(
            () => callInstantDetails(ref, word_num),
            INSTANT_DETAILS_DELAY,
        );
    }

    function handleWordMouseLeave(e: MouseEvent) {
        clearTimeout(instant_details_timer);
        EventsEmit("app:instantdetails:hide");
    }
</script>

<div id="content">
    {#if $app_data}
        <Virtualiser items={$app_data} bind:current_position={cp} let:row>
            <div class="block">
                {#if isRefBookStart(row.range.start)}
                    <div class="heading">
                        {bibleRefToVersionBookTitle(row.range.start, "gnt")}
                    </div>
                {/if}

                <span class="ref">
                    {bibleRefToString(row.range.start, "short")}
                </span>
                {#each row.verses as verse, verse_index}
                    <div class="verse" style="display: inline">
                        {#if verse_index > 0}
                            <sup>{verse.ref % 1000}</sup>
                        {/if}
                        {#each verse.words as word, word_index}
                            {word.pre}<span
                                class="word"
                                on:mouseenter={(e) =>
                                    handleWordMouseEnter(
                                        verse.ref,
                                        word_index + 1,
                                    )}
                                on:mouseleave={handleWordMouseLeave}
                                >{word.text}</span
                            >{word.post}{" "}
                        {/each}
                    </div>
                {/each}
            </div>
        </Virtualiser>
    {/if}

    <InstantDetails />
</div>

<style>
    #content {
        position: relative;
        width: 100%;
        height: 100%;
        overflow: hidden;
    }

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
        /* font-family: "Neuton"; */
        font-family: "SBL Greek";
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
        font-family: "SBL Greek";
        font-size: 1.2rem;
        cursor: pointer;
    }

    .word:hover {
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
