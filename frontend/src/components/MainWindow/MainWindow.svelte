<script lang="ts">
    import type { app } from "!wails/go/models";
    import { GetScriptureSections } from "!wails/go/app/App";
    import { createBibleRef } from "@/lib/Scripture/ref";
    import WindowText from "./WindowText.svelte";

    let text: app.ScriptureSection;
    async function onStartup() {
        const texts = await GetScriptureSections("gnt", [
            {
                start: createBibleRef("Matthew", 1),
                end: createBibleRef("Matthew", 28, "end"),
            },
            {
                start: createBibleRef("Mark", 1),
                end: createBibleRef("Revelation", 22, "end"),
            },
        ]);
        if (texts.length > 0) text = texts[0];
    }
    onStartup();
</script>

<div id="content">
    {#if text}
        <WindowText {text} />
    {/if}
</div>

<style>
    #content {
        position: relative;
        width: 100%;
        height: 100%;
        overflow: hidden;
    }
</style>
