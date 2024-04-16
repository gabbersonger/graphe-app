import { writable, type Writable } from "svelte/store";
import { sidebarData, type SidebarSection } from "@/components/Sidebar/data";
import type { BibleRef } from "@/lib/Scripture/types";
import type { ModalName } from "@/components/Modals/data";
import { defaultTheme, type ThemeName } from "@/static/themes";
import type { AppMode } from "@/lib/manager";

// UI elements
export const ui_theme: Writable<ThemeName> = writable(defaultTheme);
export const ui_showSidebar = writable(false);
export const ui_sidebarSection: Writable<SidebarSection> = writable(
  sidebarData[0].name,
);
export const ui_modal: Writable<ModalName | ""> = writable("");

// App Functionality
export const app_mode: Writable<AppMode> = writable("passage");
export const app_textVersion = writable("");
export const app_searchQuery = writable(true);
export const app_data = writable(true);
export const app_currentRef: Writable<BibleRef> = writable(40_001_001);
