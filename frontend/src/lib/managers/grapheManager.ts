import { EventsEmit, EventsOff, EventsOn } from "!wails/runtime/runtime";
import { type ThemeName, defaultTheme, themeData } from "@/static/themes";
import { writable, type Writable } from "svelte/store";

type GrapheMode = "app" | "settings" | "loading";

// Data stores

export const graphe_mode: Writable<GrapheMode> = writable("loading");
export const graphe_theme: Writable<ThemeName> = writable(defaultTheme);

// Functions to handle events

function handleMode(mode: GrapheMode) {
  graphe_mode.set(mode);
}

function handleTheme(theme: ThemeName) {
  graphe_theme.set(theme);
}

// TODO: remove
function handleThemeToggle() {
  graphe_theme.update(
    (theme) =>
      themeData[
        (themeData.findIndex((t) => t.name == theme) + 1) % themeData.length
      ].name,
  );
}

export function grapheManager(_: HTMLElement) {
  const handlers: {
    [key: string]: (...data: any) => void;
  } = {
    "graphe:mode": handleMode,
    "graphe:theme": handleTheme,
    "graphe:theme:toggle": handleThemeToggle, // TODO: remove
  };

  for (const [event, callback] of Object.entries(handlers)) {
    EventsOn(event, callback);
  }

  // Start the app
  EventsEmit("graphe:mode", "settings");

  return {
    destroy() {
      for (const [event, _] of Object.entries(handlers)) {
        EventsOff(event);
      }
    },
  };
}
