import { settings } from "!wails/go/models";
import { graphe_settings } from "@/lib/stores";
import { GetSettings, UpdateSetting } from "!wails/go/app/App";
import { GrapheLog } from "@/lib/utils";

import { defaultTheme } from "@/static/themes";

export function defaultSettings(): settings.SettingsValues {
  return new settings.SettingsValues({
    appearence: {
      theme: defaultTheme,
      font: {
        system: "",
        greek: "",
        hebrew: "",
        english: "",
      },
    },
  });
}

export async function getSavedSettings() {
  const settings = await GetSettings();
  // TODO: Set the specific settings that are not default
  graphe_settings.set(settings);
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

export async function updateSetting(setting: string[], value: any) {
  const setting_updated = await UpdateSetting(setting, value);
  if (setting_updated) {
    GrapheLog("info", `Setting updated: ${setting.join("/")} -> ${value}`);
    updateSettingStore(setting, value);
  }
}
