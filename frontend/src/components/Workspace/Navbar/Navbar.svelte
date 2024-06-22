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
        workspace_currentRef,
    } from "@/lib/stores";
    import { refToString } from "@/lib/Scripture/ref";
    import { EventsEmit } from "!wails/runtime/runtime";

    let width: number;

    const BREAKPOINT = 550;
</script>

<div id="navbar" bind:clientWidth={width}>
    <div class="container">
        <div class="wrapper wrapper-nav">
            <NavbarItem
                icon={TextSelect}
                text={width > BREAKPOINT ? "passage" : ""}
                on:click={() => EventsEmit("window:workspace:mode", "passage")}
                tooltip="Passage Mode"
                command="⌘P"
                selected={$workspace_mode == "passage"}
            />

            <NavbarItem
                icon={Search}
                text={width > BREAKPOINT ? "search" : ""}
                on:click={() => EventsEmit("window:workspace:mode", "search")}
                tooltip="Search Mode"
                command="⌘F"
                selected={$workspace_mode == "search"}
            />

            <div class="separator"></div>

            <NavbarItem
                icon={LibraryBig}
                text={$workspace_version}
                on:click={() => EventsEmit("window:workspace:modal", "version")}
                tooltip="Choose Version"
                command="⌘D"
            />

            <NavbarItem
                icon={BookOpenText}
                text={$workspace_version && $workspace_currentRef
                    ? refToString(
                          $workspace_version,
                          $workspace_currentRef,
                          "chapter",
                      )
                    : ""}
                on:click={() => EventsEmit("window:workspace:modal", "text")}
                tooltip="Choose Text"
                command="⌘T"
                disabled={$workspace_mode == "search"}
            />

            <div class="separator"></div>

            <NavbarItem
                icon={Sigma}
                on:click={() =>
                    EventsEmit("window:workspace:sidebar", "functions")}
                tooltip="Functions"
                command="⌘E"
            />

            <NavbarItem
                icon={NotepadText}
                on:click={() =>
                    EventsEmit("window:workspace:sidebar", "analytics")}
                tooltip="Analytics"
                command="⌘R"
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
        padding: 0 1em;
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
        gap: 1.2em;
    }

    .wrapper-nav {
        height: 36px;
        background: var(--clr-background-sub);
        border-radius: 0.4em;
        padding: 0 1.8em;
    }

    .separator {
        height: 1rem;
        width: 0.1rem;
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
</style>
