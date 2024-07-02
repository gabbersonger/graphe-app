<script lang="ts">
    import SettingsHeading from "@/components/Settings/ui/SettingsHeading.svelte";
    import SettingsSubHeading from "@/components/Settings/ui/SettingsSubHeading.svelte";
    import SettingsParagraph from "@/components/Settings/ui/SettingsParagraph.svelte";
    import ShortcutTable from "@/components/Settings/content/Shortcuts/ShortcutTable.svelte";
</script>

<SettingsHeading>Shortcuts</SettingsHeading>

<SettingsSubHeading>Manage Shortcuts</SettingsSubHeading>
<SettingsParagraph>
    Keyboard shortcuts allow you to do certain tasks more easily by clicking a
    combination of keys. You can set or remove keybinds below.
</SettingsParagraph>
<ShortcutTable />
