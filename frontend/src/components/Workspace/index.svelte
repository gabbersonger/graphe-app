<script lang="ts">
    import Navbar from "@/components/Workspace/Navbar/Navbar.svelte";
    import MainWindow from "@/components/Workspace/Content/MainWindow.svelte";
    import Sidebar from "@/components/Workspace/Sidebar/Sidebar.svelte";
    import Modals from "@/components/Workspace/Modals/Modals.svelte";

    import { Window } from "@wailsio/runtime";
    import { workspace_sidebar } from "@/lib/stores";
    import { windowWorkspaceManager } from "@/lib/managers/window_workspace";

    let is_fullscreen = false;
    async function checkIfFullscreen(_: number) {
        let temp = await Window.IsFullscreen();
        is_fullscreen = temp;
    }
    let inner_width: number;
    $: checkIfFullscreen(inner_width);

    let nav_floating: boolean;
</script>

<svelte:window bind:innerWidth={inner_width} />

<div
    id="app"
    class:sidebar={$workspace_sidebar}
    class:fullscreen={is_fullscreen}
    class:navfloating={nav_floating}
    use:windowWorkspaceManager
>
    <div class="content">
        <nav><Navbar bind:nav_floating /></nav>
        <main><MainWindow /></main>
    </div>
    <aside><Sidebar /></aside>
    <Modals />
</div>

<style>
    #app {
        position: relative;
        width: 100%;
        height: 100%;
        display: flex;
        flex-direction: row;
        background: var(--clr-background);
    }

    #app .content {
        position: relative;
        width: calc(
            100% -
                clamp(
                    var(--size-sidebar-width-min),
                    30%,
                    var(--size-sidebar-width-max)
                )
        );
        height: 100%;
        display: flex;
        flex-direction: column;
        isolation: isolate;
    }

    #app:not(.sidebar) .content {
        width: 100%;
    }

    #app:not(.sidebar) aside {
        display: none;
    }

    #app .content nav {
        position: relative;
        width: 100%;
        height: var(--size-navbar-height);
        --wails-draggable: drag;
    }

    #app:not(.fullscreen) .content nav > :global(#navbar) {
        padding-left: var(--size-navbar-clear-left);
    }

    #app:not(.navfloating) .content nav {
        height: var(--size-navbar-height-small);
    }

    #app .content main {
        position: relative;
        width: 100%;
        height: calc(100% - var(--size-navbar-height));
    }

    #app:not(.navfloating) .content main {
        height: calc(100% - var(--size-navbar-height-small));
    }

    #app aside {
        position: relative;
        width: clamp(
            var(--size-sidebar-width-min),
            30%,
            var(--size-sidebar-width-max)
        );
        height: 100%;
    }

    @media (max-width: 600px) {
        #app .content {
            width: 100%;
        }

        #app aside {
            position: absolute;
            top: 0;
            right: 0;
            width: clamp(
                var(--size-sidebar-width-min),
                90%,
                var(--size-sidebar-width-max)
            );
            box-shadow:
                0 4px 6px -1px rgb(0 0 0 / 0.1),
                0 2px 4px -2px rgb(0 0 0 / 0.1);
        }
    }

    @media (max-width: 300px) {
        #app aside {
            width: 100%;
        }
    }
</style>
