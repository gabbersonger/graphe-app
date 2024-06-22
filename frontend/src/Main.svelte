<script lang="ts">
    import Workspace from "@/components/Workspace/index.svelte";
    import Settings from "@/components/Settings/index.svelte";
    import Loading from "@/components/Loading/index.svelte";

    import { graphe_mode, graphe_settings } from "@/lib/stores";
    import { grapheManager } from "@/lib/managers/graphe";
    import { createThemeStyles } from "@/static/themes";

    $: if ($graphe_settings) {
        const theme_values = createThemeStyles(
            $graphe_settings.appearence.theme,
        );
        for (let i = 0; i < theme_values.length; i++) {
            document.body.style.setProperty(
                theme_values[i].variable,
                theme_values[i].value,
            );
        }
    }
</script>

<div id="window" use:grapheManager>
    {#if $graphe_mode != "loading"}
        <div class="app">
            <Workspace />
        </div>
    {/if}

    {#if $graphe_mode == "settings"}
        <div class="settings">
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
</style>
