<script lang="ts">
    import type { app } from "!wails/go/models";
    import { GetScriptureSections } from "!wails/go/app/App";
    import { createBibleRef } from "@/lib/Scripture/ref";
    import WindowText from "./WindowText.svelte";

    let texts: Array<app.ScriptureSection>;
    async function onStartup() {
        let data = await GetScriptureSections("gnt", [
            {
                start: createBibleRef("Matthew", 1),
                end: createBibleRef("Matthew", 28, "end"),
            },
            {
                start: createBibleRef("Mark", 1),
                end: createBibleRef("Mark", 16, "end"),
            },
            {
                start: createBibleRef("Luke", 1),
                end: createBibleRef("Revelation", 22, "end"),
            },
        ]);
        texts = data;
    }
    onStartup();
</script>

<div id="content">
    {#if texts}
        <WindowText {texts} />
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
