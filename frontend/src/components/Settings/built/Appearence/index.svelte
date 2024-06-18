<script lang="ts">
    import SettingsHeading from "@/components/Settings/ui/SettingsHeading.svelte";
    import SettingsSubHeading from "@/components/Settings/ui/SettingsSubHeading.svelte";
    import SettingsParagraph from "@/components/Settings/ui/SettingsParagraph.svelte";
    import FontSelector from "@/components/Settings/built/Appearence/FontSelector.svelte";
    import ThemeSelector from "@/components/Settings/built/Appearence/ThemeSelector.svelte";
</script>

<SettingsHeading>Appearence</SettingsHeading>

<SettingsSubHeading>Scaling / Zoom</SettingsSubHeading>

<SettingsSubHeading>Fonts</SettingsSubHeading>
<SettingsParagraph>
    Choose which font is used for the different parts of the application.
</SettingsParagraph>
<FontSelector />

<SettingsSubHeading>Theme</SettingsSubHeading>
<SettingsParagraph>Choose a theme for the application.</SettingsParagraph>
<ThemeSelector />
