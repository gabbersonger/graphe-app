import { Sigma, NotepadText } from "lucide-svelte";
import SidebarAnalytics from "@/components/sidebar/built/SidebarAnalytics.svelte";
import SidebarFunctions from "@/components/sidebar/built/SidebarFunctions.svelte";

export const sidebarData = [
  {
    name: "functions",
    icon: Sigma,
    window: SidebarFunctions,
  },
  {
    name: "analytics",
    icon: NotepadText,
    window: SidebarAnalytics,
  },
] as const;

export type SidebarSection = (typeof sidebarData)[number]["name"];
