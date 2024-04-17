<script lang="ts">
    import Virtualiser from "@/components/MainWindow/Virtualiser.svelte";

    import { app_data, app_currentRef } from "@/lib/appManager";
    import {
        bibleRefToString,
        bibleRefToVersionBookTitle,
        isRefBookStart,
    } from "@/lib/Scripture/ref";

    let cp = { section: 0, block: 0 };
    $: if ($app_data.length > 0 && cp)
        $app_currentRef = $app_data[cp.section].blocks[cp.block].range.start;

    let scrollVirtualiser: (_: number) => void;
</script>

<div id="content">
    {#if $app_data}
        <Virtualiser
            items={$app_data}
            bind:current_position={cp}
            bind:scrollToItem={scrollVirtualiser}
            let:row
        >
            <div class="block">
                {#if isRefBookStart(row.range.start)}
                    <div class="heading">
                        {bibleRefToVersionBookTitle(row.range.start, "gnt")}
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
    {/if}
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
    }
</style>
