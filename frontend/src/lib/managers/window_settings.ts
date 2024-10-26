import { EventHandler } from "@/lib/event_handler";
import { graphe_mode, settings_section } from "@/lib/stores";
import type { SettingSection } from "@/components/Settings/data";
import { get } from "svelte/store";

function handleSection(section: SettingSection) {
  settings_section.set(section);
  if (get(graphe_mode) != "settings") {
    graphe_mode.set("settings");
  }
  if (section == "shortcuts") {
    // TODO: disable menu
  } else {
    // TODO: load menu
  }
}

export function windowSettingsManager(_: HTMLElement) {
  const events = new EventHandler();
  events.subscribe("window:settings:section", handleSection);

  return {
    destroy() {
      events.shutdown();
    },
  };
}
