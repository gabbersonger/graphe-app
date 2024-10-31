import { EventHandler } from "@/lib/event_handler";
import { graphe_mode, settings_section } from "@/lib/stores";
import type { SettingSection } from "@/components/Settings/data";
import { get } from "svelte/store";
import { z } from "zod";

export const z_settings_section = z.union([
  z.literal("shortcuts"),
  z.literal("appearance"),
  z.literal("bible"),
  z.literal("general"),
]);
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
  events.subscribe(
    "window:settings:section",
    handleSection,
    z_settings_section,
  );

  return {
    destroy() {
      events.shutdown();
    },
  };
}
