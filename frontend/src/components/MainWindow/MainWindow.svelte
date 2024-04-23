<script lang="ts">
    import Virtualiser from "@/components/MainWindow/Virtualiser.svelte";
    import InstantDetails from "@/components/MainWindow/InstantDetails.svelte";

    import type { BibleRef } from "@/lib/Scripture/types";

    import { EventsEmit } from "!wails/runtime/runtime";
    import { app_data, app_version, app_currentRef } from "@/lib/appManager";
    import { isRefBookStart, refToString } from "@/lib/Scripture/ref";
    import { GetEnvironmentInfo } from "!wails/go/app/App";
    import { onMount } from "svelte";
    import { GitBranch } from "lucide-svelte";

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

    let current_position: { section: number; block: number };
    $: if ($app_data.length > 0 && current_position) {
        $app_currentRef =
            $app_data[current_position.section].blocks[current_position.block]
                .range.start;
    }

    let version: string;
    onMount(async () => {
        let data = await GetEnvironmentInfo();
        version = data.version;
    });
</script>

<div id="content">
    <Virtualiser data={$app_data} bind:current_position>
        <div slot="row" let:row class="block">
            {#if isRefBookStart($app_version, row.range.start)}
                <div class="heading">
                    {refToString($app_version, row.range.start, "book")}
                </div>
            {/if}

            <span class="ref">
                {refToString($app_version, row.range.start, "short")}
            </span>
            {#each row.verses as verse, index}
                <div class="verse" style="display: inline">
                    {#if index > 0}
                        <sup>{verse.ref % 1000}</sup>
                    {/if}
                    {#each verse.words as word}
                        {word.pre}<span
                            class="word"
                            on:mouseenter={(e) =>
                                handleWordMouseEnter(verse.ref, word.word_num)}
                            on:mouseleave={handleWordMouseLeave}
                            >{word.text}</span
                        >{word.post}{" "}
                    {/each}
                </div>
            {/each}
        </div>
    </Virtualiser>

    <div class="version-info"><GitBranch />{version}</div>

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

    .version-info {
        position: absolute;
        bottom: 1rem;
        right: 1rem;
        display: flex;
        align-items: center;
        gap: 0.1rem;
        font-family: monospace;
        font-size: 0.8rem;
        color: var(--clr-text-sub);
        pointer-events: none;
        background: var(--clr-background);
        border-radius: 0.1rem;
        padding: 0.2rem;
    }

    .version-info > :global(svg) {
        height: 1em;
    }
</style>
