<script lang="ts">
    import Navbar from "@/components/Navbar.svelte";
    import MainWindow from "@/components/MainWindow.svelte";
    import Sidebar from "@/components/Sidebar.svelte";

    import { ui_showSidebar } from "@/stores/app";
</script>

<div id="app" class:sidebar={$ui_showSidebar}>
    <nav><div class="wrapper"><Navbar /></div></nav>
    <main><MainWindow /></main>
    <aside><Sidebar /></aside>
</div>

<style>
    #app {
        --size-navbar-height: 38px;
        --size-navbar-clear-left: 80px;
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

    #app nav .wrapper {
        position: relative;
        width: 100%;
        height: 100%;
        padding-left: var(--size-navbar-clear-left);
    }

    #app main {
        grid-area: content;
    }

    #app aside {
        grid-area: sidebar;
    }
</style>
