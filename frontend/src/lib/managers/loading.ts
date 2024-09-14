import { Events } from "@wailsio/runtime";
import { EventHandler } from "@/lib/event_handler";
import { getSavedSettings } from "@/lib/settings_handler";

export function loadingManager(_: HTMLElement) {
  const events = new EventHandler();

  getSavedSettings().then((x) => {
    Events.Emit({ name: "graphe:mode", data: "settings" });
  });

  return {
    destroy() {
      events.shutdown();
    },
  };
}
