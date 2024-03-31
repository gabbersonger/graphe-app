import { writable, type Writable } from "svelte/store";
import { sidebarData, type SidebarSection } from "@/components/Sidebar/data";

export const ui_showSidebar = writable(false);
export const ui_sidebarSection: Writable<SidebarSection> = writable(
  sidebarData[0].name,
);
