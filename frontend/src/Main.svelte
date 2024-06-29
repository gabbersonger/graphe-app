<script lang="ts">
    import Workspace from "@/components/Workspace/index.svelte";
    import Settings from "@/components/Settings/index.svelte";
    import Loading from "@/components/Loading/index.svelte";

    import { graphe_mode } from "@/lib/stores";
    import { grapheManager } from "@/lib/managers/graphe";
    import { uiManager } from "@/lib/managers/ui";
</script>

<svelte:body use:uiManager />

<div id="window" use:grapheManager>
    {#if $graphe_mode != "loading"}
        <div class="app">
            <Workspace />
        </div>

        <div class="settings" class:shown={$graphe_mode == "settings"}>
            <Settings />
        </div>
    {/if}

    {#if $graphe_mode == "loading"}
        <div class="loading">
            <Loading />
        </div>
    {/if}
</div>

<style>
    #window {
        position: relative;
        width: 100%;
        height: 100%;
        isolation: isolate;
        z-index: 0;
    }

    .app,
    .settings,
    .loading {
        position: absolute;
        inset: 0;
        isolation: isolate;
    }

    .settings {
        z-index: 1;
    }

    .settings:not(.shown) {
        display: none;
    }
</style>
