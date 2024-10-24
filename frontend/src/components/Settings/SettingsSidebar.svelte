<script lang="ts">
    import { Events } from "@wailsio/runtime";

    import { settings_section } from "@/lib/stores";
    import {
        settingsData,
        getSettingCategories,
        type SettingSection,
    } from "@/components/Settings/data";

    function selectSection(section: SettingSection) {
        Events.Emit({ name: "window:settings:section", data: section });
    }
</script>

<div class="buttons">
    {#each getSettingCategories() as category}
        <div class="title">{category}</div>
        {#each settingsData.filter((x) => x.category == category) as s}
            <button
                on:click={() => selectSection(s.name)}
                class:selected={$settings_section == s.name}
            >
                {s.display}
            </button>
        {/each}
    {/each}
</div>

<style>
    .buttons {
        display: flex;
        flex-direction: column;
        gap: 0.5rem;

        --size-button-padding: 0.6rem;
    }

    .title {
        padding-inline: var(--size-button-padding);
        font-family: var(--font-system);
        font-size: 0.7rem;
        font-weight: bold;
        color: var(--clr-text-sub);
        text-transform: uppercase;
    }

    .title:not(:first-child) {
        padding-top: 1rem;
    }

    button {
        background: none;
        border: none;
        padding: 0;
        padding: 0.3em var(--size-button-padding);
        border-radius: 0.2em;
        cursor: pointer;
        font-family: var(--font-system);
        font-size: 0.9rem;
        font-weight: 500;
        color: var(--clr-text);
        text-align: left;
    }

    button.selected {
        background: var(--clr-background-sub);
        color: var(--clr-main);
    }
</style>
