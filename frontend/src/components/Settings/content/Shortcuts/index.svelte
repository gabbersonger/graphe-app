<script lang="ts">
    import SettingsHeading from "@/components/Settings/ui/SettingsHeading.svelte";
    import SettingsSubHeading from "@/components/Settings/ui/SettingsSubHeading.svelte";
    import SettingsParagraph from "@/components/Settings/ui/SettingsParagraph.svelte";
    import ShortcutTable from "@/components/Settings/content/Shortcuts/ShortcutTable.svelte";
    import { Keyboard } from "lucide-svelte";
</script>

<SettingsHeading>Shortcuts</SettingsHeading>

<div class="warning">
    <Keyboard />
    <span>All shortcuts are disabled while on this page</span>
</div>

<SettingsSubHeading>Manage Shortcuts</SettingsSubHeading>
<SettingsParagraph>
    Keyboard shortcuts allow you to do certain tasks more easily by clicking a
    combination of keys. You can set or remove keybinds below.
</SettingsParagraph>
<ShortcutTable />

<style>
    .warning {
        display: flex;
        flex-direction: row;
        align-items: center;
        gap: 0.5rem;
        margin-top: 1rem;
        padding: 1rem;
        background: var(--clr-background-sub);
        border-radius: 0.2rem;
        color: var(--clr-main);
        font-family: var(--font-system);
        font-weight: 500;
        font-size: 0.8rem;
        line-height: 1.5;
    }

    .warning > :global(svg) {
        height: 1rem;
        width: 1rem;
    }
</style>
