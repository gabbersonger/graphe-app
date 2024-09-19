<script lang="ts">
    import TextVirtualiser from "@/components/Workspace/Content/TextVirtualiser.svelte";
    import InstantDetails from "@/components/Workspace/Content/InstantDetails.svelte";

    import {
        workspace_data,
        workspace_version,
        workspace_ref,
    } from "@/lib/stores";
    import {
        ScriptureService,
        type ScriptureRef,
    } from "!/graphe/internal/scripture";

    let current_verse: ScriptureRef;
    $: if ($workspace_data.length > 0 && current_verse)
        $workspace_ref = current_verse;

    let language: string;
    let languageHeadings: string;
    $: updateLanguages($workspace_version);
    async function updateLanguages(version: string | undefined) {
        if (!version || version == undefined) {
            return;
        }
        language = await ScriptureService.GetVersionLanguage(version);
        languageHeadings =
            await ScriptureService.GetVersionLanguageHeadings(version);
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
