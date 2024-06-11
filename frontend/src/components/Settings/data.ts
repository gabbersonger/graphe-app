import SettingsGeneral from "@/components/Settings/built/SettingsGeneral.svelte";
import SettingsAppearence from "@/components/Settings/built/SettingsAppearence.svelte";
import SettingsShortcuts from "@/components/Settings/built/SettingsShortcuts.svelte";
import SettingsFormatting from "@/components/Settings/built/SettingsFormatting.svelte";
import SettingsSearch from "@/components/Settings/built/SettingsSearch.svelte";
import SettingsInstantDetails from "@/components/Settings/built/SettingsInstantDetails.svelte";
import SettingsVersion from "@/components/Settings/built/SettingsVersion.svelte";

export const settingsData = [
  {
    name: "general",
    display: "General",
    content: SettingsGeneral,
  },
  {
    name: "appearence",
    display: "Appearence",
    content: SettingsAppearence,
  },
  {
    name: "formatting",
    display: "Formatting",
    content: SettingsFormatting,
  },
  {
    name: "search",
    display: "Search",
    content: SettingsSearch,
  },
  {
    name: "instantdetails",
    display: "Instant Details",
    content: SettingsInstantDetails,
  },
  {
    name: "shortcuts",
    display: "Shortcuts",
    content: SettingsShortcuts,
  },
  {
    name: "version",
    display: "Version",
    content: SettingsVersion,
  },
] as const;

export type SettingSection = (typeof settingsData)[number]["name"];
