<script lang="ts">
    import { EventsEmit } from "!wails/runtime/runtime";
    import NavbarItem from "@/components/Navbar/NavbarItem.svelte";
    import SettingsHeading from "@/components/Settings/ui/SettingsHeading.svelte";
    import SettingsSubHeading from "@/components/Settings/ui/SettingsSubHeading.svelte";
    import {
        BookOpenText,
        LibraryBig,
        NotepadText,
        Search,
        Sigma,
        TextSelect,
    } from "lucide-svelte";
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

<SettingsHeading>Appearence</SettingsHeading>

<SettingsSubHeading>Scaling / Zoom</SettingsSubHeading>

<SettingsSubHeading>Fonts</SettingsSubHeading>

<SettingsSubHeading>Theme</SettingsSubHeading>
<div id="themes">
    {#each themeData as theme}
        <button
            on:click={() => selectTheme(theme.name)}
            class:selected={$ui_theme == theme.name}
            style={createThemeStyles(theme.name)}
        >
            <div class="offset">
                <div class="navbar-wrapper">
                    <NavbarItem
                        icon={TextSelect}
                        text="passage"
                        selected={true}
                    />
                    <NavbarItem icon={Search} text="search" />
                    <div class="separator"></div>
                    <NavbarItem icon={LibraryBig} text="esv" />
                    <NavbarItem icon={BookOpenText} text="gen 1" />
                    <div class="separator"></div>
                    <NavbarItem icon={Sigma} />
                    <NavbarItem icon={NotepadText} />
                </div>
                <div class="text">
                    <div>
                        <span class="ref">Gen 1:1</span>
                        <span class="reg">
                            In the beginning, God created the heavens and the
                            earth.
                        </span>
                        <sup class="verse">2</sup>
                        <span class="reg">
                            The earth was without form and void, and darkness
                            was over the face of the deep. And the Spirit of God
                            was hovering over the face of the waters.
                        </span>
                    </div>
                    <div>
                        <span class="ref">Gen 1:3</span>And God said, “Let there
                        be light,” and there was light.
                        <sup class="verse">4</sup>And God saw that the light was
                        good. And God separated the light from the darkness.
                        <sup class="verse">5</sup>God called the light Day, and
                        the darkness he called Night. And there was evening and
                        there was morning, the first day.
                    </div>
                    <div>
                        <span class="ref">Gen 1:6</span>And God said, “Let there
                        be an expanse in the midst of the waters, and let it
                        separate the waters from the waters.”
                        <sup class="verse">7</sup>And God made the expanse and
                        separated the waters that were under the expanse from
                        the waters that were above the expanse. And it was so.
                        <sup class="verse">8</sup>And God called the expanse
                        Heaven. And there was evening and there was morning, the
                        second day.
                    </div>
                </div>
            </div>
        </button>
    {/each}
</div>

<style>
    #themes {
        display: grid;
        grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
        gap: 1rem;
    }

    #themes button {
        position: relative;
        aspect-ratio: 2 / 1;
        border: none;
        cursor: pointer;
        border-radius: 0.3rem;
        isolation: isolate;
        overflow: hidden;
        background: var(--clr-background);
        border: 1px solid var(--clr-background-dark);
        padding: 1rem;
    }

    #themes button.selected {
        border: 2px solid var(--clr-main);
    }

    #themes button .offset {
        position: absolute;
        top: 1rem;
        left: 1rem;
        min-height: 100%;
        width: 500px;
        min-width: 100%;
        pointer-events: none;
    }

    #themes button .navbar-wrapper {
        height: 36px;
        background: var(--clr-background-sub);
        border-radius: 0.4em;
        padding: 0 1.8em;
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: flex-start;
        gap: 1rem;
    }

    #themes button .navbar-wrapper .separator {
        height: 1rem;
        width: 0.1rem;
        background: var(--clr-text-sub);
    }

    #themes button .text {
        width: 100%;
        padding-block: 1rem;
        font-size: 1rem;
        text-align: left;
    }

    #themes button .text > div {
        position: relative;
        display: block;
        padding-bottom: 1rem;
        color: var(--clr-text);
        font-size: 1rem;
        line-height: 1.7;
    }

    #themes button .text .ref {
        font-weight: bold;
        font-size: 0.8rem;
        color: var(--clr-main);
        background: var(--clr-background-sub);
        padding: 0.3rem;
        border-radius: 0.1rem;
        margin-right: 1ch;
    }

    #themes button .text sup {
        vertical-align: super;
        font-weight: bold;
        font-size: 0.7rem;
        color: var(--clr-text-sub);
        margin-right: 1ch;
    }

    #themes button .text .reg {
        font-family: "Neuton";
        font-size: 1.15rem;
    }
</style>
