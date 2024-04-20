<script lang="ts">
    import ModalWrapper from "@/components/Modals/ModalWrapper.svelte";
    import ModalButtons from "@/components/Modals/ModalButtons.svelte";
    import Input from "@/components/ui/Input.svelte";
    import { BookText, LibraryBig } from "lucide-svelte";

    import { EventsEmit } from "!wails/runtime/runtime";
    import { versionData } from "@/lib/Scripture/data";
    import { app_version } from "@/lib/appManager";
    import type { BibleVersion } from "@/lib/Scripture/types";

    let value: string = "";
    let isInputFocused = true;

    function normaliseString(string: string): string {
        return string.toLowerCase().replaceAll(" ", "");
    }

    // TODO: update searching to use abbreviations
    const searchStrings = Object.keys(versionData).map((x) => ({
        name: x as BibleVersion,
        string:
            normaliseString(x) +
            "|" +
            normaliseString(versionData[x].full_name) +
            "|" +
            normaliseString(versionData[x].language),
    }));

    type AvailableVersion = (typeof versionData)[BibleVersion] & {
        name: keyof typeof versionData;
    };
    function filterVersions(query: string): AvailableVersion[] {
        query = normaliseString(query);
        return searchStrings
            .map((x) => ({ ...x, index: x.string.indexOf(query) }))
            .filter((x) => x.index >= 0)
            .sort((a, b) => a.index - b.index)
            .map((x) => ({ name: x.name, ...versionData[x.name] }));
    }

    $: available_versions = filterVersions(value);

    function clickVersion(index: number) {
        let version = available_versions[index];
        EventsEmit("app:version", version.name);
        EventsEmit("ui:modal:close");
    }

    function checkForInputEnter(e: KeyboardEvent) {
        if (e.code == "Enter" && available_versions.length > 0) clickVersion(0);
    }
</script>

<ModalWrapper>
    <div slot="header">
        <Input
            bind:value
            bind:focus={isInputFocused}
            on:keypress={checkForInputEnter}
            icon={LibraryBig}
            placeholder="Change the Version"
        />
    </div>

    <div slot="content">
        <ModalButtons
            items={available_versions}
            rowData={{ number: 4, maxwidth: 10 }}
            onItemClick={clickVersion}
            icon={BookText}
            subheading={(index) => available_versions[index].language}
            heading={(index) => available_versions[index].full_name}
            isActive={(index) => available_versions[index].name == $app_version}
            isFocused={(index) =>
                value.length > 0 && index == 0 && isInputFocused}
        />
    </div>
</ModalWrapper>
