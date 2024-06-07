<script lang="ts">
    import App from "@/routes/App.svelte";
    import Settings from "@/routes/Settings.svelte";
    import Loading from "@/routes/Loading.svelte";

    import { ui_theme } from "@/lib/uiManager";
    import { createThemeStyles } from "@/static/themes";
    import { graphe_mode, grapheManager } from "@/lib/grapheManager";
</script>

<div id="window" style={createThemeStyles($ui_theme)} use:grapheManager>
    {#if $graphe_mode != "loading"}
        <div class="app">
            <App />
        </div>
    {/if}

    {#if $graphe_mode == "settings"}
        <div class="settings">
            <Settings />
        </div>
    {/if}

    {#if $graphe_mode == "loading"}
        <div class="loding">
            <Loading />
        </div>
    {/if}
</div>

<style>
    #window {
        position: relative;
        width: 100%;
        height: 100%;
    }

    .app,
    .settings {
        position: absolute;
        inset: 0;
        isolation: isolate;
    }

    .settings {
        z-index: 1;
    }
</style>
