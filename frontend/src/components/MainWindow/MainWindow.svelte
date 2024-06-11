<script lang="ts">
    import TextVirtualiser from "@/components/MainWindow/TextVirtualiser.svelte";
    import InstantDetails from "@/components/MainWindow/InstantDetails.svelte";
    import VersionInfo from "@/components/MainWindow/VersionInfo.svelte";
    import type { BibleRef, BibleVersion } from "@/lib/Scripture/types";
    import {
        app_data,
        app_version,
        app_currentRef,
    } from "@/lib/managers/appManager";
    import { versionData } from "@/lib/Scripture/data";

    let current_verse: BibleRef;
    $: if ($app_data.length > 0 && current_verse)
        $app_currentRef = current_verse;

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
    <TextVirtualiser data={$app_data} bind:current_verse />
    <VersionInfo />
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

    #content[data-language-heading="English"] {
        --font-title: "Neuton";
    }

    #content[data-language-heading="Ancient Greek"] {
        --font-title: "SBL Greek";
    }
</style>
