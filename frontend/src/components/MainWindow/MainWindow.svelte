<script lang="ts">
    import Virtualiser from "@/components/MainWindow/Virtualiser.svelte";
    import InstantDetails from "@/components/MainWindow/InstantDetails.svelte";

    import type { BibleVersion } from "@/lib/Scripture/types";

    import { app_data, app_version, app_currentRef } from "@/lib/appManager";
    import { GetEnvironmentInfo } from "!wails/go/app/App";
    import { onMount } from "svelte";
    import { GitBranch } from "lucide-svelte";
    import { versionData } from "@/lib/Scripture/data";
    import ScriptureBlock from "./ScriptureBlock.svelte";

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

    let language: (typeof versionData)[BibleVersion]["language"];
    let languageHeadings: (typeof versionData)[BibleVersion]["languageHeadings"];
    $: if ($app_version) {
        language = versionData[$app_version].language;
        languageHeadings = versionData[$app_version].languageHeadings;
    }
</script>

<div
    id="content"
    data-language={language}
    data-language-heading={languageHeadings}
>
    <Virtualiser data={$app_data} bind:current_position>
        <ScriptureBlock slot="row" let:row block={row} />
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

    #content[data-language="Ancient Greek"] {
        --font-content: "SBL Greek";
    }

    #content[data-language="English"] {
        --font-content: "Neuton";
    }

    #content[data-language-heading="English"] :global(.heading) {
        --font-heading: "Neuton";
    }

    #content[data-language-heading="Ancient Greek"] :global(.heading) {
        --font-heading: "SBL Greek";
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
