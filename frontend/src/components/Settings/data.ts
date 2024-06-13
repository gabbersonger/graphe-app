import SettingsGeneral from "@/components/Settings/built/SettingsGeneral.svelte";
import SettingsAppearence from "@/components/Settings/built/Appearence/SettingsAppearence.svelte";
import SettingsShortcuts from "@/components/Settings/built/SettingsShortcuts.svelte";
import SettingsFormatting from "@/components/Settings/built/SettingsFormatting.svelte";
import SettingsSearch from "@/components/Settings/built/SettingsSearch.svelte";
import SettingsInstantDetails from "@/components/Settings/built/SettingsInstantDetails.svelte";
import SettingsVersion from "@/components/Settings/built/SettingsVersion.svelte";

export const settingsData = [
  {
    name: "general",
    display: "General",
    category: "App Settings",
    content: SettingsGeneral,
  },
  {
    name: "appearence",
    display: "Appearence",
    category: "App Settings",
    content: SettingsAppearence,
  },
  {
    name: "shortcuts",
    display: "Shortcuts",
    category: "App Settings",
    content: SettingsShortcuts,
  },
  {
    name: "version",
    display: "Version",
    category: "App Settings",
    content: SettingsVersion,
  },
  {
    name: "formatting",
    display: "Formatting",
    category: "Functionality",
    content: SettingsFormatting,
  },
  {
    name: "search",
    display: "Search",
    category: "Functionality",
    content: SettingsSearch,
  },
  {
    name: "instantdetails",
    display: "Instant Details",
    category: "Functionality",
    content: SettingsInstantDetails,
  },
] as const;

export type SettingSection = (typeof settingsData)[number]["name"];
