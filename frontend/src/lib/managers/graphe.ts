import { EventHandler } from "@/lib/event_handler";
import { graphe_mode, settings_section, type GrapheMode } from "@/lib/stores";
import { updateSetting, resetSetting } from "@/lib/settings_handler";
import { get } from "svelte/store";
import { DisableMenu, LoadMenu } from "!wails/go/app/App";

function handleMode(mode: GrapheMode) {
  if (get(graphe_mode) == "settings" && mode != "settings") {
    LoadMenu();
  } else if (mode != "settings" && get(settings_section) == "shortcuts") {
    DisableMenu();
  }
  graphe_mode.set(mode);
}

export function grapheManager(_: HTMLElement) {
  const events = new EventHandler();
  events.subscribe("graphe:mode", handleMode);
  events.subscribe("graphe:setting", updateSetting);
  events.subscribe("graphe:setting:reset", resetSetting);

  return {
    destroy() {
      events.shutdown();
    },
  };
}
