import { writable, type Writable } from "svelte/store";
import type { BibleRef, BibleVersion } from "@/lib/Scripture/types";
import type { database, settings } from "!wails/go/models";
import type { SettingSection } from "@/components/Settings/data";
import type { ModalName } from "@/components/Workspace/Modals/data";
import {
  sidebarData,
  type SidebarSection,
} from "@/components/Workspace/Sidebar/data";
import { defaultSettings } from "@/lib/settings_handler";

// WHOLE APP
export type GrapheMode = "app" | "settings" | "loading";
export const graphe_mode: Writable<GrapheMode> = writable("loading");
export const graphe_settings: Writable<settings.SettingsValues> =
  writable(defaultSettings());

// SETTINGS WINDOW
export const settings_section: Writable<SettingSection> = writable("general");

// WORKSPACE WINDOW
export type WorkspaceMode = "passage" | "search";
export const workspace_mode: Writable<WorkspaceMode> = writable("passage");
export const workspace_modal: Writable<ModalName | ""> = writable("");
export const workspace_sidebar: Writable<boolean> = writable(false);
export const workspace_sidebarSection: Writable<SidebarSection> = writable(
  sidebarData[0].name,
);
export const workspace_version: Writable<BibleVersion> = writable();
export const workspace_data: Writable<database.ScriptureSection[]> = writable(
  [],
);
export const workspace_instantDetailsData: Writable<database.ScriptureWordData> =
  writable();
export const workspace_currentRef: Writable<BibleRef> = writable();
