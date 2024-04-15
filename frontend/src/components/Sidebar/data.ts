import { CircleHelp } from "lucide-svelte";
import SidebarTodo from "@/components/sidebar/SidebarTodo.svelte";

export const sidebarData = [
  {
    name: "todo",
    icon: CircleHelp,
    window: SidebarTodo,
  },
] as const;

export type SidebarSection = (typeof sidebarData)[number]["name"];
