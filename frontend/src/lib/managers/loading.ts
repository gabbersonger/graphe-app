import { EventsEmit } from "!wails/runtime/runtime";
import { EventHandler } from "@/lib/event_handler";
import { getSavedSettings } from "@/lib/settings_handler";

export function loadingManager(_: HTMLElement) {
  const events = new EventHandler();

  getSavedSettings().then((x) => {
    EventsEmit("graphe:mode", "settings");
  });

  return {
    destroy() {
      events.shutdown();
    },
  };
}
