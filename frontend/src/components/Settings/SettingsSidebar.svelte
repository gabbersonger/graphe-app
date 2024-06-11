<script lang="ts">
    import { EventsEmit } from "!wails/runtime/runtime";

    import { settings_section } from "@/lib/managers/settingsManager";
    import {
        settingsData,
        type SettingSection,
    } from "@/components/Settings/data";

    function selectSection(section: SettingSection) {
        EventsEmit("settings:section", section);
    }
</script>

<div class="buttons">
    <div class="title">App Settings</div>
    {#each settingsData as s}
        <button
            on:click={() => selectSection(s.name)}
            class:selected={$settings_section == s.name}
        >
            {s.display}
        </button>
    {/each}
</div>

<style>
    .buttons {
        display: flex;
        flex-direction: column;
        gap: 0.5em;

        --size-button-padding: 0.6rem;
    }

    .title {
        font-size: 0.7rem;
        font-weight: bold;
        color: var(--clr-text);
        text-transform: uppercase;
        padding-inline: var(--size-button-padding);
    }

    button {
        background: none;
        border: none;
        font-size: 0.9rem;
        color: var(--clr-text);
        cursor: pointer;
        padding: 0;
        text-align: left;
        padding: 0.3em var(--size-button-padding);
        border-radius: 0.2em;
        font-weight: 500;
    }

    button.selected {
        background: var(--clr-background-sub);
        color: var(--clr-main);
    }
</style>
