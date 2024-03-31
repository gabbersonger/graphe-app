import { Cog, SquareSigma } from "lucide-svelte";
import SidebarSettings from "@/components/sidebar/SidebarSettings.svelte";
import SidebarFormatting from "@/components/sidebar/SidebarFormatting.svelte";

export const sidebarData = [
  {
    name: "formatting",
    icon: SquareSigma,
    window: SidebarFormatting,
  },
  {
    name: "settings",
    icon: Cog,
    window: SidebarSettings,
  },
] as const;

export type SidebarSection = (typeof sidebarData)[number]["name"];
