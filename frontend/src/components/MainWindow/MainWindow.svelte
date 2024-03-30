<script lang="ts">
    import type { app } from "!wails/go/models";
    import { GetScriptureSections } from "!wails/go/app/App";
    import { createBibleRef } from "@/lib/Scripture/ref";
    import BibleDisplay from "./BibleDisplay.svelte";

    let text: app.ScriptureSection;
    async function onStartup() {
        const texts = await GetScriptureSections("gnt", [
            {
                start: createBibleRef("Matthew", 1),
                end: createBibleRef("Revelation", 22, "end"),
            },
        ]);
        if (texts.length > 0) text = texts[0];
    }
    onStartup();
</script>

<div id="content">
    <div class="container">
        <div class="wrapper">
            {#if text}
                <BibleDisplay {text} />
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
        width: 95%;
        max-width: var(--size-max-width);
        overflow: scroll;
    }

    .wrapper {
        position: absolute;
        inset: 0;
        overflow: scroll;
    }

    /* .wrapper::-webkit-scrollbar {
        display: none;
    } */
</style>
