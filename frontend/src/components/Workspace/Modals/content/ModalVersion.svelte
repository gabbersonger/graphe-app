<script lang="ts">
    import ModalResultView from "../ModalResultView.svelte";
    import { LibraryBig } from "lucide-svelte";

    import { versionData } from "@/lib/Scripture/data";
    import type { BibleVersion } from "@/lib/Scripture/types";
    import { EventsEmit } from "!wails/runtime/runtime";
    import { workspace_version } from "@/lib/stores";

    function normaliseString(string: string): string {
        return string.toLowerCase().replaceAll(" ", "");
    }

    const search_strings = Object.keys(versionData).map((x) => ({
        name: x as BibleVersion,
        string:
            normaliseString(versionData[x].full_name) +
            "|" +
            normaliseString(x) +
            "|" +
            normaliseString(versionData[x].language),
    }));

    function filterVersions(query: string) {
        query = normaliseString(query);
        return search_strings
            .map((x) => ({ ...x, index: x.string.indexOf(query) }))
            .filter((x) => x.index >= 0)
            .filter((x) => x.name != $workspace_version)
            .sort((a, b) => a.index - b.index)
            .map((x) => ({
                value: x.name,
                display: versionData[x.name].full_name,
            }));
    }

    let value = "";
    let available_versions: { value: string; display: string }[] = [];
    $: available_versions = filterVersions(value);

    function chooseVersion(index: number) {
        const version = available_versions[index].value;
        EventsEmit("window:workspace:version", version);
        EventsEmit("window:workspace:modal:close");
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
