<script lang="ts">
    import ModalWrapper from "@/components/Modals/ModalWrapper.svelte";
    import Input from "@/components/ui/Input.svelte";
    import { BookText, LibraryBig } from "lucide-svelte";

    import { EventsEmit } from "!wails/runtime/runtime";
    import { versionData } from "@/lib/Scripture/data";
    import type { BibleVersion } from "@/lib/Scripture/types";
    import { app_version } from "@/lib/appManager";

    let value: string = "";
    $: available_versions = filterVersions(value);

    function normaliseString(string: string): string {
        return string.toLowerCase().replaceAll(" ", "");
    }

    const searchStrings = Object.keys(versionData).map((x) => ({
        name: x as BibleVersion,
        string:
            normaliseString(x) +
            "|" +
            normaliseString(versionData[x].fullname) +
            "|" +
            normaliseString(versionData[x].language),
    }));
    function filterVersions(query: string): BibleVersion[] {
        query = normaliseString(query);
        return searchStrings
            .map((x) => ({ ...x, index: x.string.indexOf(query) }))
            .filter((x) => x.index >= 0)
            .sort((a, b) => a.index - b.index)
            .map((x) => x.name);
    }

    function clickVersion(version: string) {
        EventsEmit("app:version", version as BibleVersion);
        EventsEmit("ui:modal:close");
    }

    let isInputFocused = true;
    function checkForInputEnter(e: KeyboardEvent) {
        if (e.code == "Enter" && available_versions.length > 0) {
            clickVersion(available_versions[0]);
        }
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
        <div class="version-buttons">
            {#each available_versions as version, index}
                <button
                    on:click={() => clickVersion(version)}
                    class:active={version == $app_version}
                    class:hover={true}
                    class:focus={value.length > 0 &&
                        index == 0 &&
                        isInputFocused}
                    tabindex="0"
                >
                    <span class="language">
                        <BookText />
                        {versionData[version].language}
                    </span>
                    <span class="title">{versionData[version].fullname}</span>
                </button>
            {/each}
        </div>
    </div>
</ModalWrapper>

<style>
    .version-buttons {
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        justify-content: flex-start;
        gap: 1rem;
    }

    button {
        position: relative;
        width: calc((100% - 3rem) / 4);
        min-width: 10rem;
        aspect-ratio: 2 / 1;
        background: none;
        border: 2px solid var(--clr-background-sub);
        border-radius: 0.2rem;
        color: var(--clr-text);
        font-size: 1rem;
        display: flex;
        flex-direction: column;
        align-items: flex-start;
        padding: 1em;
        text-align: left;
    }

    button .language {
        font-size: 0.8em;
        color: var(--clr-text-sub);
        padding-bottom: 0.3em;
        display: flex;
        align-items: center;
        gap: 0.5em;
    }

    button .language > :global(svg) {
        height: 1em;
        width: 1em;
    }

    button .title {
        font-size: 1em;
        text-align: left;
        line-height: 1.3;
    }

    button:not(.active):hover {
        background: var(--clr-background-sub);
        border-color: var(--clr-background-dark);
        cursor: pointer;
    }

    button.active {
        color: var(--clr-text-highlight);
        color: var(--clr-text);
        background: var(--clr-background-sub);
    }

    button.active .language {
        color: var(--clr-text);
        color: var(--clr-main);
    }

    button:focus,
    button.focus {
        border-color: var(--clr-main);
        outline: none;
    }
</style>
