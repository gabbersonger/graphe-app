import { writable, type Writable } from "svelte/store";
import { sidebarData, type SidebarSection } from "@/components/Sidebar/data";
import type { BibleRef } from "@/lib/Scripture/types";

export const ui_showSidebar = writable(false);
export const ui_sidebarSection: Writable<SidebarSection> = writable(
  sidebarData[0].name,
);

export const ui_currentRef: Writable<BibleRef> = writable(40_001_001);
