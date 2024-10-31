import { SettingsDB } from "!/graphe/internal/settings";
import { graphe_settings } from "@/lib/stores";
import { GrapheEvent } from "@/lib/utils";

async function getSavedSettings() {
  const settings = await SettingsDB.GetSettings();
  graphe_settings.set(settings);
  return true;
}

export function loadingManager(_: HTMLElement) {
  getSavedSettings().then((x) => {
    GrapheEvent("graphe:mode", "workspace");
  });
}
