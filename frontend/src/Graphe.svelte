<script lang="ts">
    import App from "@/routes/App.svelte";
    import Settings from "@/routes/Settings.svelte";
    import Loading from "@/routes/Loading.svelte";

    import {
        graphe_mode,
        graphe_theme,
        grapheManager,
    } from "@/lib/managers/grapheManager";
    import { createThemeStyles } from "@/static/themes";

    $: if ($graphe_theme) {
        const theme_values = createThemeStyles($graphe_theme);
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
        isolation: isolate;
        z-index: 0;
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
