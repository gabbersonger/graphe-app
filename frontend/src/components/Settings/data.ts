import SettingsGeneral from "@/components/Settings/built/General/index.svelte";
import SettingsAppearence from "@/components/Settings/built/Appearence/index.svelte";
import SettingsShortcuts from "@/components/Settings/built/Shortcuts/index.svelte";
import SettingsFormatting from "@/components/Settings/built/Formatting/index.svelte";
import SettingsSearch from "@/components/Settings/built/Search/index.svelte";
import SettingsInstantDetails from "@/components/Settings/built/InstantDetails/index.svelte";
import SettingsVersion from "@/components/Settings/built/Version/index.svelte";

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
