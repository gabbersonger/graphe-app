<script lang="ts">
    import ModalResultView from "@/components/Workspace/Modals/ModalResultView.svelte";
    import { LibraryBig } from "lucide-svelte";

    import { onMount } from "svelte";
    import { workspace_version } from "@/lib/stores";
    import {
        ScriptureService,
        ScriptureVersionBasicInfo,
    } from "!/graphe/internal/scripture";
    import { GrapheEvent } from "@/lib/utils";

    function normaliseString(string: string): string {
        return string.toLowerCase().replaceAll(" ", "");
    }

    let search_data: ScriptureVersionBasicInfo[] = [];
    let search_strings: Array<{
        name: string;
        full_name: string;
        search_string: string;
    }> = [];
    onMount(async () => {
        search_data = await ScriptureService.GetVersionsBasicData();
        search_strings = search_data.map((v) => ({
            name: v.name,
            full_name: v.full_name,
            search_string:
                normaliseString(v.full_name) +
                "|" +
                normaliseString(v.name) +
                "|" +
                normaliseString(v.language),
        }));
    });

    function filterVersions(query: string, _: any[]) {
        query = normaliseString(query);
        return search_strings
            .map((x) => ({ ...x, index: x.search_string.indexOf(query) }))
            .filter((x) => x.index >= 0)
            .filter((x) => x.name != $workspace_version)
            .sort((a, b) => a.index - b.index)
            .map((x) => ({
                value: x.name,
                display: x.full_name,
            }));
    }

    let value = "";
    let available_versions: { value: string; display: string }[] = [];
    $: available_versions = filterVersions(value, search_strings);

    function chooseVersion(index: number) {
        const version = available_versions[index].value;
        GrapheEvent("window:workspace:version", version);
        GrapheEvent("window:workspace:modal:close");
    }
</script>

<ModalResultView
    icon={LibraryBig}
    placeholder="Choose a Version"
    bind:value
    results={available_versions}
    chooseResult={chooseVersion}
    noResults="There are no versions that match your search"
/>
