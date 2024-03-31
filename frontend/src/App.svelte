<script lang="ts">
    import Navbar from "@/components/Navbar/Navbar.svelte";
    import MainWindow from "@/components/MainWindow/MainWindow.svelte";
    import Sidebar from "@/components/Sidebar.svelte";

    import { ui_showSidebar } from "@/stores/app";
    import { WindowIsFullscreen } from "!wails/runtime/runtime";

    let isFullscreen = false;
    async function checkIfFullscreen(_: number) {
        let temp = await WindowIsFullscreen();
        isFullscreen = temp;
    }
    let innerWidth: number;
    $: checkIfFullscreen(innerWidth);
</script>

<svelte:window bind:innerWidth />

<div id="app" class:sidebar={$ui_showSidebar} class:fullscreen={isFullscreen}>
    <nav><Navbar /></nav>
    <main><MainWindow /></main>
    <aside><Sidebar /></aside>
</div>

<style>
    #app {
        --size-navbar-height: 55px;
        --size-navbar-clear-left: 68px;
        --size-sidebar-width: 300px;

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
