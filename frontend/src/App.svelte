<script lang="ts">
    import Navbar from "@/components/Navbar/Navbar.svelte";
    import MainWindow from "@/components/MainWindow/MainWindow.svelte";
    import Sidebar from "@/components/Sidebar/Sidebar.svelte";
    import Modals from "@/components/Modals/Modals.svelte";

    import { ui_showSidebar, ui_theme } from "@/lib/stores";
    import { themeData } from "@/lib/theme-data";
    import { WindowIsFullscreen } from "!wails/runtime/runtime";
    import { eventListener } from "@/lib/eventListener";

    let isFullscreen = false;
    async function checkIfFullscreen(_: number) {
        let temp = await WindowIsFullscreen();
        isFullscreen = temp;
    }
    let innerWidth: number;
    $: checkIfFullscreen(innerWidth);

    $: theme = themeData.find((t) => t.name == $ui_theme);
</script>

<svelte:window bind:innerWidth />

<div
    id="app"
    class:sidebar={$ui_showSidebar}
    class:fullscreen={isFullscreen}
    style:--clr-background={theme.colors.background}
    style:--clr-background-sub={theme.colors.backgroundSub}
    style:--clr-background-dark={theme.colors.backgroundDark}
    style:--clr-main={theme.colors.main}
    style:--clr-text={theme.colors.text}
    style:--clr-text-sub={theme.colors.textSub}
    style:--clr-text-highlight={theme.colors.textHighlight}
    use:eventListener
>
    <nav><Navbar /></nav>
    <main><MainWindow /></main>
    <aside><Sidebar /></aside>
    <Modals />
</div>

<style>
    #app {
        position: relative;
        width: 100%;
        height: 100%;
        display: grid;
        grid-template-columns: 1fr;
        grid-template-rows: var(--size-navbar-height) 1fr;
        grid-template-areas: "navbar" "content";
        background: var(--clr-background);
    }

    #app.sidebar {
        grid-template-columns: 1fr var(--size-sidebar-width);
        grid-template-areas: "navbar sidebar" "content sidebar";
    }

    #app:not(.sidebar) aside {
        display: none;
    }

    #app nav {
        grid-area: navbar;
        --wails-draggable: drag;
    }

    #app:not(.fullscreen) nav {
        padding-left: var(--size-navbar-clear-left);
    }

    #app main {
        grid-area: content;
    }

    #app aside {
        grid-area: sidebar;
    }
</style>
