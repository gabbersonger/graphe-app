import { graphe_settings } from "@/lib/stores";
import { SettingsDB } from "!/graphe/internal/settings";
import { GrapheLog } from "@/lib/utils";
import { get } from "svelte/store";

export async function getSavedSettings() {
  const settings = await SettingsDB.GetSettings();
  graphe_settings.set(settings);
  return true;
}

function updateSettingStore(setting: string[], value: any) {
  graphe_settings.update((s) => {
    if (s == undefined) {
      GrapheLog(
        "error",
        `Setting store is null, while trying to update \`${setting.join("/")}\` to \`${value}\``,
      );
      return;
    }

    try {
      let item: any = s;
      for (let i = 0; i < setting.length - 1; i++) {
        if (!(setting[i] in item)) {
          GrapheLog(
            "error",
            `Error accessing parameter \`${setting[i]}\` in settings store, while trying to update \`${setting.join("/")}\` to \`${value}\``,
          );
        }
        item = item[setting[i]];
      }
      item[setting[setting.length - 1]] = value;
    } catch (e) {
      GrapheLog(
        "error",
        `Error updating setting store for \`${setting.join("/")}\` to \`${value}\` (error: ${e})`,
      );
    }
    return s;
  });
}

function parseSettingValue(setting: string[], value: any): any {
  // Handle zoom
  if (
    setting.length == 2 &&
    setting[0] == "appearence" &&
    setting[1] == "zoom"
  ) {
    const current_graphe_settings = get(graphe_settings);
    if (current_graphe_settings == undefined) {
      return GrapheLog(
        "error",
        "Trying to parse setting value when settings is null",
      );
    }
    const current_zoom = current_graphe_settings.appearence.zoom;
    if (value == "in") {
      return Math.min(current_zoom + 10, 200);
    } else if (value == "out") {
      return Math.max(current_zoom - 10, 50);
    }
  }
  return value;
}

export async function updateSetting(data: { setting: string[]; value: any }) {
  const parsed_value = parseSettingValue(data.setting, data.value);
  const setting_updated = await SettingsDB.UpdateSetting(
    data.setting,
    parsed_value,
  );

  if (setting_updated) {
    GrapheLog(
      "info",
      `Setting updated: ${data.setting.join("/")} -> ${parsed_value}`,
    );
    updateSettingStore(data.setting, parsed_value);
  }
}

export async function resetSetting(setting: string[]) {
  const setting_updated = await SettingsDB.ResetSetting(setting);
  GrapheLog(
    "info",
    `Setting reset: ${setting.join("/")} -> default value: ${setting_updated}`,
  );
  updateSettingStore(setting, setting_updated);
}
