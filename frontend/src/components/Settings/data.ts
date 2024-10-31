import SettingsGeneral from "@/components/Settings/content/General/index.svelte";
import SettingsAppearence from "@/components/Settings/content/Appearence/index.svelte";
import SettingsShortcuts from "@/components/Settings/content/Shortcuts/index.svelte";
import SettingsFormatting from "@/components/Settings/content/Formatting/index.svelte";
import SettingsSearch from "@/components/Settings/content/Search/index.svelte";
import SettingsInstantDetails from "@/components/Settings/content/InstantDetails/index.svelte";
import SettingsVersion from "@/components/Settings/content/Version/index.svelte";

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
export function getSettingCategories() {
  let categories: string[] = [];
  for (let i = 0; i < settingsData.length; i++) {
    if (!categories.includes(settingsData[i].category)) {
      categories.push(settingsData[i].category);
    }
  }
  return categories as (typeof settingsData)[number]["category"][];
}
