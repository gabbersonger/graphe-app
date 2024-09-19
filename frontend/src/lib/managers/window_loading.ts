import { SettingsDB } from "!/graphe/internal/settings";
import { graphe_settings, type GrapheMode } from "@/lib/stores";
import { EventHandler } from "@/lib/event_handler";
import { Events } from "@wailsio/runtime";

async function getSavedSettings() {
  const settings = await SettingsDB.GetSettings();
  graphe_settings.set(settings);
  return true;
}

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
