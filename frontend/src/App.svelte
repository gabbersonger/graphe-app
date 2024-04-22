<script lang="ts">
    import Navbar from "@/components/Navbar/Navbar.svelte";
    import MainWindow from "@/components/MainWindow/MainWindow.svelte";
    import MainWindow2 from "@/components/MainWindow/MainWindow2.svelte";
    import Sidebar from "@/components/Sidebar/Sidebar.svelte";
    import Modals from "@/components/Modals/Modals.svelte";

    import { ui_modal, ui_showSidebar, ui_theme } from "@/lib/uiManager";
    import { createThemeStyles } from "@/static/themes";
    import { WindowIsFullscreen } from "!wails/runtime/runtime";
    import { appManager } from "@/lib/appManager";
    import { uiManager } from "@/lib/uiManager";

    let isFullscreen = false;
    async function checkIfFullscreen(_: number) {
        let temp = await WindowIsFullscreen();
        isFullscreen = temp;
    }
    let innerWidth: number;
    $: checkIfFullscreen(innerWidth);
</script>

<svelte:window bind:innerWidth />

<div
    id="app"
    class:sidebar={$ui_showSidebar}
    class:fullscreen={isFullscreen}
    class:modal={$ui_modal != ""}
    style={createThemeStyles($ui_theme)}
    use:uiManager
    use:appManager
>
    <nav><Navbar /></nav>
    <!-- <main><MainWindow /></main> -->
    <main><MainWindow2 /></main>
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

    #app.modal > main,
    #app.modal > nav,
    #app.modal > aside {
        filter: blur(2px);
        -webkit-filter: blur(2px);
    }
</style>
