import { EventHandler } from "@/lib/event_handler";
import { graphe_mode, type GrapheMode } from "@/lib/stores";
import { getSavedSettings, updateSetting } from "@/lib/settings_handler";

function handleMode(mode: GrapheMode) {
  graphe_mode.set(mode);
}

export function grapheManager(_: HTMLElement) {
  const events = new EventHandler();
  events.subscribe("graphe:mode", handleMode);
  events.subscribe("graphe:setting", updateSetting);

  getSavedSettings();

  return {
    destroy() {
      events.shutdown();
    },
  };
}
