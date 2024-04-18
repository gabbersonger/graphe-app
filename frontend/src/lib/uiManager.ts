import { writable, type Writable } from "svelte/store";
import { sidebarData, type SidebarSection } from "@/components/Sidebar/data";
import { defaultTheme, themeData, type ThemeName } from "@/static/themes";
import { EventsOff, EventsOn } from "!wails/runtime/runtime";
import type { ModalName } from "@/components/Modals/data";
import { app_mode } from "@/lib/appManager";
import { get } from "svelte/store";

// Data stores

export const ui_theme: Writable<ThemeName> = writable(defaultTheme);
export const ui_showSidebar = writable(false);
export const ui_sidebarSection: Writable<SidebarSection> = writable(
  sidebarData[0].name,
);
export const ui_modal: Writable<ModalName | ""> = writable("");

// Functions to handle events

function handleUIModal(data: ModalName) {
  if (data == "text" && get(app_mode) == "search") return;
  ui_modal.update((val) => (val == data ? "" : data));
}

function handleUIModalClose() {
  ui_modal.set("");
}

function handleUISidebarToggle() {
  ui_showSidebar.update((x) => !x);
}

function handleUITheme(data: ThemeName) {
  ui_theme.set(data);
}

// TODO: remove
function handleUIThemeToggle() {
  ui_theme.update(
    (theme) =>
      themeData[
        (themeData.findIndex((t) => t.name == theme) + 1) % themeData.length
      ].name,
  );
}

// Register events

export function uiManager(_: HTMLElement) {
  const handlers: {
    [key: string]: (...data: any) => void;
  } = {
    "ui:modal": handleUIModal,
    "ui:modal:close": handleUIModalClose,
    "ui:sidebar:toggle": handleUISidebarToggle,
    "ui:theme": handleUITheme,

    "ui:theme:toggle": handleUIThemeToggle,
  };

  for (const [event, callback] of Object.entries(handlers)) {
    EventsOn(event, callback);
  }

  return {
    destroy() {
      for (const [event, _] of Object.entries(handlers)) {
        EventsOff(event);
      }
    },
  };
}
