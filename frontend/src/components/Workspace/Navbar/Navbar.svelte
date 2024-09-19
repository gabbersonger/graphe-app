<script lang="ts">
    import NavbarItem from "@/components/Workspace/Navbar/NavbarItem.svelte";
    import {
        BookOpenText,
        Search,
        Sigma,
        TextSelect,
        LibraryBig,
        NotepadText,
    } from "lucide-svelte";
    import {
        workspace_mode,
        workspace_version,
        workspace_ref,
    } from "@/lib/stores";
    import { Events } from "@wailsio/runtime";
    import { graphe_settings, workspace_sidebar } from "@/lib/stores";
    import type { SettingsValues } from "!/graphe/internal/settings";
    import {
        ScriptureRefStringType,
        ScriptureService,
    } from "!/graphe/internal/scripture";

    export let nav_floating = true;
    let width: number;

    const BREAKPOINTS = [
        {
            value: 50,
            text: 400,
            nav: 350,
        },
        {
            value: 70,
            text: 500,
            nav: 385,
        },
        {
            value: 90,
            text: 600,
            nav: 450,
        },
        {
            value: 110,
            text: 680,
            nav: 520,
        },
        {
            value: 130,
            text: 760,
            nav: 590,
        },
        {
            value: 150,
            text: 850,
            nav: 640,
        },
        {
            value: 170,
            text: 910,
            nav: 700,
        },
        {
            value: 190,
            text: 1030,
            nav: 790,
        },
    ] as const;

    function getBreakpoint(
        settings: SettingsValues | undefined,
        type: "text" | "nav",
    ): number {
        if (settings == undefined) return 0;

        const zoom = settings.appearence.zoom;
        let curr = 0;
        for (let i = 0; i < BREAKPOINTS.length; i++) {
            if (zoom >= BREAKPOINTS[i].value) {
                curr = BREAKPOINTS[i][type];
            } else break;
        }
        return curr;
    }

    $: nav_breakpoint = getBreakpoint($graphe_settings, "nav");
    $: text_breakpoint = getBreakpoint($graphe_settings, "text");
    $: nav_floating = width > nav_breakpoint;

    let current_ref_string = "";
    async function getCurrentRefString() {
        ScriptureService.RefToString(
            $workspace_ref as number,
            $workspace_version as string,
            ScriptureRefStringType.StringChapter,
        ).then((value) => {
            current_ref_string = value;
        });

        let value = await ScriptureService.RefToString(
            $workspace_ref as number,
            $workspace_version as string,
            ScriptureRefStringType.StringChapter,
        );
        console.log("asd");
        return value;
    }
    $: if ($workspace_version != undefined && $workspace_ref != null) {
        getCurrentRefString();
    }
</script>

<div
    id="navbar"
    class:topbar={width <= nav_breakpoint}
    bind:clientWidth={width}
>
    <div class="container">
        <div class="wrapper wrapper-nav">
            <NavbarItem
                icon={TextSelect}
                text={width > text_breakpoint ? "passage" : ""}
                on:click={() =>
                    Events.Emit({
                        name: "window:workspace:mode",
                        data: "passage",
                    })}
                tooltip="Passage Mode"
                command={$graphe_settings?.shortcuts.passageMode}
                selected={$workspace_mode == "passage"}
            />

            <NavbarItem
                icon={Search}
                text={width > text_breakpoint ? "search" : ""}
                on:click={() =>
                    Events.Emit({
                        name: "window:workspace:mode",
                        data: "search",
                    })}
                tooltip="Search Mode"
                command={$graphe_settings?.shortcuts.searchMode}
                selected={$workspace_mode == "search"}
            />

            <div class="separator"></div>

            <NavbarItem
                icon={LibraryBig}
                text={$workspace_version}
                on:click={() =>
                    Events.Emit({
                        name: "window:workspace:modal",
                        data: "version",
                    })}
                tooltip="Choose Version"
                command={$graphe_settings?.shortcuts.chooseVersion}
            />

            <NavbarItem
                icon={BookOpenText}
                text={current_ref_string}
                on:click={() =>
                    Events.Emit({
                        name: "window:workspace:modal",
                        data: "text",
                    })}
                tooltip="Choose Text"
                command={$graphe_settings?.shortcuts.chooseText}
                disabled={$workspace_mode == "search"}
            />

            <div class="separator"></div>

            <NavbarItem
                icon={Sigma}
                on:click={() =>
                    Events.Emit({
                        name: "window:workspace:modal",
                        data: "functions",
                    })}
                tooltip="Functions"
                command={$graphe_settings?.shortcuts.openFunctions}
            />

            <NavbarItem
                icon={NotepadText}
                on:click={() =>
                    Events.Emit({
                        name: "window:workspace:sidebar",
                        data: null,
                    })}
                tooltip="Analytics"
                command={$graphe_settings?.shortcuts.openAnalytics}
                selected={$workspace_sidebar}
            />
        </div>
    </div>

    <div class="navbar-fade"></div>
</div>

<style>
    #navbar {
        position: relative;
        width: 100%;
        height: 100%;
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: flex-start;
        user-select: none;
        -webkit-user-select: none;
    }

    .container {
        width: 100%;
        height: 100%;
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: center;
    }

    .wrapper {
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: flex-start;
        gap: 1.2rem;
        background: red;
    }

    .wrapper-nav {
        height: min(2rem, 85%);
        background: var(--clr-background-sub);
        border-radius: 0.4rem;
        padding: 0 1.8em;
    }

    .separator {
        height: 1rem;
        width: 0.1ch;
        background: var(--clr-text-sub);
    }

    .navbar-fade {
        --size-fade-height: 20px;
        position: absolute;
        width: 90vw;
        height: var(--size-fade-height);
        right: 5vw;
        bottom: calc(-1 * var(--size-fade-height));
        background: linear-gradient(
            180deg,
            var(--clr-background) 30%,
            rgba(0, 0, 0, 0) 100%
        );
        z-index: 1;
        pointer-events: none;
    }

    #navbar.topbar {
        background: red;
        background: var(--clr-background-sub);
        padding-block: 0;
    }

    #navbar.topbar .container {
        justify-content: flex-start;
        padding-left: 1vw;
    }

    #navbar.topbar .wrapper-nav {
        padding: 0;
        padding-inline: 0.2rem;
        gap: 0.5rem;
    }
</style>
