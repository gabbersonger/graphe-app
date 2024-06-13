<script lang="ts">
    import { EventsEmit } from "!wails/runtime/runtime";

    import {
        themeData,
        createThemeStyles,
        type ThemeName,
    } from "@/static/themes";
    import { ui_theme } from "@/lib/managers/uiManager";

    function selectTheme(theme: ThemeName) {
        EventsEmit("ui:theme", theme);
    }
</script>

<div id="themes">
    {#each themeData as theme}
        <button on:click={() => selectTheme(theme.name)}>
            <div class="theme-display" class:selected={$ui_theme == theme.name}>
                <div
                    class="theme-wrapper"
                    style={createThemeStyles(theme.name)}
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
        grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
        gap: 1.2rem;
    }

    #themes button {
        border: none;
        cursor: pointer;
        background: none;
    }

    #themes button p {
        padding: 0;
        margin: 0;
        padding-top: 0.7rem;
        font-size: 0.8rem;
        text-transform: capitalize;
        color: var(--clr-text);
    }

    .theme-display {
        position: relative;
        aspect-ratio: 2 / 1;
        background: var(--clr-background-sub);
        border: 0.3rem solid var(--clr-background-sub);
        border-radius: 0.3rem;
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
        gap: 0.5rem;
        padding: 1rem;
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
        height: 2rem;
        border-radius: 0.2rem;
        padding-inline: 0.5rem;
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: flex-start;
        gap: 0.5rem;
    }

    .theme-display .bubble.small {
        height: 1.4rem;
        padding-inline: 0.3rem;
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
        width: 10ch;
        height: 0.8rem;
        background: var(--clr-text);
        border-radius: 0.3rem;
    }

    .theme-display .text.short {
        width: 5ch;
    }

    .theme-display .text.long {
        width: 100%;
    }

    .theme-display .separator {
        width: 2px;
        height: 1rem;
        background: var(--clr-text-sub);
    }

    .theme-display .sub {
        width: 3ch;
        height: 0.7rem;
        border-radius: 0.2rem;
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
