<script lang="ts">
    import type { app } from "!wails/go/models";
    import { GetScriptureSections } from "!wails/go/app/App";
    import {
        biblePointToRef,
        createBiblePoint,
        bibleRefToPoint,
        biblePointToString,
    } from "@/lib/data/bibleReference";

    let text: app.ScriptureSection;
    async function onStartup() {
        const start = performance.now();
        const texts = await GetScriptureSections("gnt", [
            {
                start: biblePointToRef(createBiblePoint("Matthew", 1)),
                end: biblePointToRef(createBiblePoint("Mark", 16, "end")),
            },
            {
                start: biblePointToRef(createBiblePoint("Luke", 1)),
                end: biblePointToRef(createBiblePoint("John", 21, "end")),
            },
            {
                start: biblePointToRef(createBiblePoint("Acts", 1)),
                end: biblePointToRef(createBiblePoint("Romans", 16, "end")),
            },
            {
                start: biblePointToRef(createBiblePoint("1 Corinthians", 1)),
                end: biblePointToRef(
                    createBiblePoint("2 Thessalonians", 3, "end"),
                ),
            },
            {
                start: biblePointToRef(createBiblePoint("1 Timothy", 1)),
                end: biblePointToRef(createBiblePoint("Revelation", 22, "end")),
            },
        ]);
        const end = performance.now();
        console.log(`Took: ${end - start}ms`);

        if (texts.length > 0) text = texts[0];
        console.log(text);
    }
    onStartup();
</script>

<div id="content">
    <div class="container">
        <div class="wrapper">
            {#if text}
                {#each text.blocks as block}
                    <div class="block">
                        <span class="ref">
                            {biblePointToString(
                                bibleRefToPoint(block.range.start),
                                "short",
                            )}
                        </span>
                        {#each block.verses as verse, index}
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
                {/each}
            {/if}
        </div>
    </div>
</div>

<style>
    #content {
        --size-max-width: 80ch;

        position: relative;
        width: 100%;
        height: 100%;
        display: flex;
        flex-direction: row;
        justify-content: center;
        overflow: hidden;

        box-sizing: border-box;
    }

    .container {
        position: relative;
        height: 100%;
        width: 100%;
        max-width: var(--size-max-width);
        overflow: scroll;
    }

    .wrapper {
        position: absolute;
        inset: 0;
        overflow: scroll;
    }

    .wrapper::-webkit-scrollbar {
        display: none;
    }

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
        color: var(--clr-main);
    }

    .word {
        font-family: "Accordance";
        font-size: 1.2rem;
    }

    .word:hover {
        cursor: pointer;
        background: var(--clr-background-dark);
        outline: 0.2rem solid var(--clr-background-dark);
    }
</style>
