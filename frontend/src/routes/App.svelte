<script lang="ts">
    import Navbar from "@/components/Navbar/Navbar.svelte";
    import MainWindow from "@/components/MainWindow/MainWindow.svelte";
    import Sidebar from "@/components/Sidebar/Sidebar.svelte";
    import Modals from "@/components/Modals/Modals.svelte";

    import { WindowIsFullscreen } from "!wails/runtime/runtime";
    import { app_modal, app_sidebar } from "@/lib/managers/appManager";
    import { appManager } from "@/lib/managers/appManager";

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
    class:sidebar={$app_sidebar}
    class:fullscreen={isFullscreen}
    class:modal={$app_modal != ""}
    use:appManager
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

    #app.modal > main,
    #app.modal > nav,
    #app.modal > aside {
        filter: blur(2px);
        -webkit-filter: blur(2px);
    }
</style>
