<script lang="ts">
    import TextVirtualiser from "@/components/Workspace/Content/TextVirtualiser.svelte";
    import InstantDetails from "@/components/Workspace/Content/InstantDetails.svelte";
    import type { BibleRef, BibleVersion } from "@/lib/Scripture/types";
    import {
        workspace_data,
        workspace_version,
        workspace_currentRef,
    } from "@/lib/stores";
    import { versionData } from "@/lib/Scripture/data";

    let current_verse: BibleRef;
    $: if ($workspace_data.length > 0 && current_verse)
        $workspace_currentRef = current_verse;

    let language: (typeof versionData)[BibleVersion]["language"];
    let languageHeadings: (typeof versionData)[BibleVersion]["languageHeadings"];
    $: if ($workspace_version) {
        language = versionData[$workspace_version].language;
        languageHeadings = versionData[$workspace_version].languageHeadings;
    }
</script>

<div
    id="content"
    data-language={language}
    data-language-heading={languageHeadings}
>
    <TextVirtualiser data={$workspace_data} bind:current_verse />
    <InstantDetails />
</div>

<style>
    #content {
        position: relative;
        width: 100%;
        height: 100%;
        overflow: hidden;
    }

    #content[data-language="greek"] {
        --font-content: var(--font-greek);
    }

    #content[data-language="english"] {
        --font-content: var(--font-english);
    }

    #content[data-language="hebrew"] {
        --font-content: var(--font-hebrew);
    }

    #content[data-language-heading="greek"] {
        --font-title: var(--font-greek);
    }

    #content[data-language-heading="english"] {
        --font-title: var(--font-english);
    }

    #content[data-language-heading="hebrew"] {
        --font-title: var(--font-hebrew);
    }
</style>
