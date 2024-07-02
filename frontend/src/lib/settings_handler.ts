import { graphe_settings } from "@/lib/stores";
import { GetSettings, UpdateSetting } from "!wails/go/app/App";
import { GrapheLog } from "@/lib/utils";
import { get } from "svelte/store";

export async function getSavedSettings() {
  const settings = await GetSettings();
  graphe_settings.set(settings);
  return true;
}

function updateSettingStore(setting: string[], value: any) {
  graphe_settings.update((s) => {
    try {
      let item = s;
      for (let i = 0; i < setting.length - 1; i++) item = item[setting[i]];
      item[setting[setting.length - 1]] = value;
    } catch (e) {
      GrapheLog(
        "error",
        `Error updating setting store for ${setting.join("/")} to ${value}`,
      );
    }
    return s;
  });
}

function parseSettingValue(setting: string[], value: any): any {
  if (
    setting.length == 2 &&
    setting[0] == "appearence" &&
    setting[1] == "zoom"
  ) {
    const current_zoom = get(graphe_settings).appearence.zoom;
    if (value == "in") {
      return Math.min(current_zoom + 10, 200);
    } else if (value == "out") {
      return Math.max(current_zoom - 10, 50);
    } else if (value == "reset") {
      return 100;
    }
  }
  return value;
}

export async function updateSetting(setting: string[], value: any) {
  const parsed_value = parseSettingValue(setting, value);
  const setting_updated = await UpdateSetting(setting, parsed_value);
  if (setting_updated) {
    GrapheLog(
      "info",
      `Setting updated: ${setting.join("/")} -> ${parsed_value}`,
    );
    updateSettingStore(setting, parsed_value);
  }
}
