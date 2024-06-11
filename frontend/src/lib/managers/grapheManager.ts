import { EventsEmit, EventsOff, EventsOn } from "!wails/runtime/runtime";
import { writable, type Writable } from "svelte/store";

type GrapheMode = "app" | "settings" | "loading";

// Data stores

export const graphe_mode: Writable<GrapheMode> = writable("loading");

// Functions to handle events

function handleGrapheMode(mode: GrapheMode) {
  graphe_mode.set(mode);
}

export function grapheManager(_: HTMLElement) {
  const handlers: {
    [key: string]: (...data: any) => void;
  } = {
    "graphe:mode": handleGrapheMode,
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
