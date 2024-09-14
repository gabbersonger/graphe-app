<script lang="ts">
    import Navbar from "@/components/Workspace/Navbar/Navbar.svelte";
    import MainWindow from "@/components/Workspace/Content/MainWindow.svelte";
    import Sidebar from "@/components/Workspace/Sidebar/Sidebar.svelte";
    import Modals from "@/components/Workspace/Modals/Modals.svelte";

    import { Window } from "@wailsio/runtime";
    import { workspace_modal, workspace_sidebar } from "@/lib/stores";
    import { windowWorkspaceManager } from "@/lib/managers/window_workspace";

    let isFullscreen = false;
    async function checkIfFullscreen(_: number) {
        let temp = await Window.IsFullscreen();
        isFullscreen = temp;
    }
    let innerWidth: number;
    $: checkIfFullscreen(innerWidth);

    let navFloating: boolean;
</script>

<svelte:window bind:innerWidth />

<div
    id="app"
    class:sidebar={$workspace_sidebar}
    class:fullscreen={isFullscreen}
    class:navfloating={navFloating}
    class:modal={$workspace_modal != ""}
    use:windowWorkspaceManager
>
    <nav><Navbar bind:navFloating /></nav>
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

    #app:not(.navfloating) {
        grid-template-rows: var(--size-navbar-height-small) 1fr;
    }

    #app nav {
        grid-area: navbar;
        --wails-draggable: drag;
    }

    #app:not(.fullscreen) nav > :global(#navbar) {
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
