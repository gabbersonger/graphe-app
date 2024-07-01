<script lang="ts">
    import { EventsEmit } from "!wails/runtime/runtime";

    import {
        type ThemeName,
        themeData,
        createThemeStylesString,
    } from "@/static/themes";
    import { graphe_settings } from "@/lib/stores";

    function selectTheme(theme: ThemeName) {
        EventsEmit("graphe:setting", ["appearence", "theme"], theme);
    }
</script>

<div id="themes">
    {#each themeData as theme}
        <button on:click={() => selectTheme(theme.name)}>
            <div
                class="theme-display"
                class:selected={$graphe_settings.appearence.theme == theme.name}
            >
                <div
                    class="theme-wrapper"
                    style={createThemeStylesString(theme.name)}
                >
                    <div class="block">
                        <div class="bubble dark">
                            <div class="text muted"></div>
                            <div class="text main"></div>
                            <div class="separator"></div>
                            <div class="text muted short"></div>
                            <div class="text muted short"></div>
                            <div class="separator"></div>
                            <div class="text muted short"></div>
                            <div class="text muted short"></div>
                        </div>
                    </div>
                    <div class="block">
                        <div class="row">
                            <div class="bubble small dark">
                                <div class="text short main"></div>
                            </div>
                            <div class="bubble small long">
                                <div class="text long"></div>
                            </div>
                        </div>
                        <div class="row">
                            <div class="bubble small long">
                                <div class="text long"></div>
                                <div class="sub"></div>
                                <div class="text long"></div>
                            </div>
                        </div>
                        <div class="row">
                            <div class="bubble small long">
                                <div class="text long"></div>
                                <div class="text blank"></div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

            <p>{theme.name}</p>
        </button>
    {/each}
</div>

<style>
    #themes {
        display: grid;
        grid-template-columns: repeat(auto-fit, minmax(190px, 1fr));
        gap: 1vw;
    }

    #themes button {
        border: none;
        cursor: pointer;
        background: none;
    }

    #themes button p {
        padding: 0;
        padding-top: 0.7rem;
        margin: 0;
        font-family: var(--font-system);
        font-size: 0.8rem;
        color: var(--clr-text);
        text-transform: capitalize;
    }

    .theme-display {
        position: relative;
        aspect-ratio: 2 / 1;
        background: var(--clr-background-sub);
        border: 5px solid var(--clr-background-sub);
        border-radius: 5px;
    }

    .theme-display.selected {
        outline: 2px solid var(--clr-main);
    }

    .theme-display .theme-wrapper {
        position: relative;
        width: 100%;
        height: 100%;
        isolation: isolate;
        overflow: hidden;
        display: flex;
        flex-direction: column;
        gap: 5%;
        padding: 5%;
        border-radius: inherit;
        background: var(--clr-background);
    }

    .theme-display .block {
        display: flex;
        flex-direction: row;
        width: 100%;
        display: flex;
        flex-direction: column;
    }

    .theme-display .bubble {
        height: 20px;
        border-radius: 4px;
        padding-inline: 6px;
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: flex-start;
        gap: 6px;
    }

    .theme-display .bubble.small {
        height: 20px;
        padding-inline: 6px;
    }

    .theme-display .bubble.long {
        width: 100%;
    }

    .theme-display .row {
        display: flex;
        flex-direction: row;
        align-items: center;
    }

    .theme-display .text {
        width: 20%;
        height: 12px;
        background: var(--clr-text);
        border-radius: 4px;
    }

    .theme-display .text.short {
        width: 10%;
        min-width: 14px;
    }

    .theme-display .text.long {
        width: 100%;
    }

    .theme-display .separator {
        width: 2px;
        height: 12px;
        background: var(--clr-text-sub);
    }

    .theme-display .sub {
        width: 10%;
        height: 10px;
        border-radius: 3px;
        margin-bottom: 10px;
        background: var(--clr-text-sub);
    }

    .theme-display .dark {
        background: var(--clr-background-sub);
    }

    .theme-display .main {
        background: var(--clr-main);
    }

    .theme-display .muted {
        background: var(--clr-text-sub);
    }

    .theme-display .blank {
        background: unset;
    }
</style>
