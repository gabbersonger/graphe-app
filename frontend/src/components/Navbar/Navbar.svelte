<script lang="ts">
    import NavbarItem from "@/components/Navbar/NavbarItem.svelte";
    import Button from "@/components/ui/Button.svelte";
    import {
        BookOpenText,
        PanelRight,
        PanelRightClose,
        Search,
        Sigma,
        TextSelect,
        LibraryBig,
        ClipboardType,
    } from "lucide-svelte";
    import { app_mode, app_version, app_currentRef } from "@/lib/appManager";
    import { ui_showSidebar } from "@/lib/uiManager";
    import { refToString } from "@/lib/Scripture/ref";
    import { EventsEmit } from "!wails/runtime/runtime";

    let width: number;
</script>

<div id="navbar" bind:clientWidth={width}>
    <div class="container">
        {#if width > 450}
            <div class="wrapper"></div>
        {/if}

        <div class="wrapper wrapper-nav">
            <NavbarItem
                icon={TextSelect}
                text={width > 650 ? "passage" : ""}
                on:click={() => EventsEmit("app:mode", "passage")}
                tooltip="Passage Mode"
                command="⌘P"
                selected={$app_mode == "passage"}
            />

            <NavbarItem
                icon={Search}
                text={width > 650 ? "search" : ""}
                on:click={() => EventsEmit("app:mode", "search")}
                tooltip="Search Mode"
                command="⌘S"
                selected={$app_mode == "search"}
            />

            <div class="separator"></div>

            <NavbarItem
                icon={LibraryBig}
                text={$app_version}
                on:click={() => EventsEmit("ui:modal", "version")}
                tooltip="Choose Version"
                command="⌘D"
            />

            <NavbarItem
                icon={BookOpenText}
                text={$app_version && $app_currentRef
                    ? refToString($app_version, $app_currentRef, "chapter")
                    : ""}
                on:click={() => EventsEmit("ui:modal", "text")}
                tooltip="Choose Text"
                command="⌘F"
                disabled={$app_mode == "search"}
            />

            <div class="separator"></div>

            <NavbarItem
                icon={Sigma}
                on:click={() => EventsEmit("ui:modal", "functions")}
                tooltip="Functions"
                command="⌘E"
            />

            <NavbarItem
                icon={ClipboardType}
                on:click={() => EventsEmit("ui:modal", "appearence")}
                tooltip="Appearence"
                command="⌘R"
            />
        </div>

        <div class="wrapper">
            <Button
                icon={$ui_showSidebar ? PanelRightClose : PanelRight}
                on:click={() => ($ui_showSidebar = !$ui_showSidebar)}
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
        justify-content: space-between;
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
