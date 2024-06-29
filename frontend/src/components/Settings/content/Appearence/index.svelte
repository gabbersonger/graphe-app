<script lang="ts">
    import SettingsHeading from "@/components/Settings/ui/SettingsHeading.svelte";
    import SettingsSubHeading from "@/components/Settings/ui/SettingsSubHeading.svelte";
    import SettingsParagraph from "@/components/Settings/ui/SettingsParagraph.svelte";
    import Command from "@/components/Settings/ui/Command.svelte";
    import ZoomSelector from "@/components/Settings/content/Appearence/ZoomSelector.svelte";
    import FontSelector from "@/components/Settings/content/Appearence/FontSelector.svelte";
    import ThemeSelector from "@/components/Settings/content/Appearence/ThemeSelector.svelte";
</script>

<SettingsHeading>Appearence</SettingsHeading>

<SettingsSubHeading>Scaling / Zoom</SettingsSubHeading>
<SettingsParagraph>
    You can change the zoom level with <Command text="⌘+" />/<Command
        text="⌘-"
    /> and reset to the default zoom with <Command text="⌘0" />.
</SettingsParagraph>
<ZoomSelector />

<SettingsSubHeading>Fonts</SettingsSubHeading>
<SettingsParagraph>
    Choose which font is used for the different parts of the application.
</SettingsParagraph>
<FontSelector />

<SettingsSubHeading>Theme</SettingsSubHeading>
<SettingsParagraph>Choose a theme for the application.</SettingsParagraph>
<ThemeSelector />
