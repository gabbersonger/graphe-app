import { writable, type Writable } from "svelte/store";
import type {
  ScriptureRef,
  ScriptureVersion,
} from "!/graphe/internal/scripture";
import type {
  ScriptureSection,
  ScriptureWordData,
} from "!/graphe/internal/data";
import type { SettingsValues } from "!/graphe/internal/settings";
import type { SettingSection } from "@/components/Settings/data";
import type { ModalName } from "@/components/Workspace/Modals/data";
import {} from "!/graphe/internal/settings";

// WHOLE APP
export type GrapheMode = "workspace" | "settings" | "loading";
export const graphe_mode: Writable<GrapheMode> = writable("loading");
export const graphe_settings: Writable<SettingsValues | undefined> =
  writable(undefined);

// SETTINGS WINDOW
export const settings_section: Writable<SettingSection> = writable("general");

// WORKSPACE WINDOW
export type WorkspaceMode = "passage" | "search";
export const workspace_mode: Writable<WorkspaceMode> = writable("passage");
export const workspace_modal: Writable<ModalName | ""> = writable("");
export const workspace_sidebar: Writable<boolean> = writable(false);

export const workspace_version: Writable<ScriptureVersion> = writable("esv");
export const workspace_ref: Writable<ScriptureRef | undefined> =
  writable(undefined);

export const workspace_data: Writable<ScriptureSection[]> = writable([]);
export const workspace_instantDetailsData: Writable<
  ScriptureWordData | undefined
> = writable();
