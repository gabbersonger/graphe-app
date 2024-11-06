import { EventHandler } from "@/lib/event_handler";
import { graphe_mode, settings_section } from "@/lib/stores";
import type { SettingSection } from "@/components/Settings/data";
import { get } from "svelte/store";
import { z } from "zod";

export const z_settings_section = z.union([
  z.literal("general"),
  z.literal("appearence"),
  z.literal("shortcuts"),
  z.literal("version"),
  z.literal("formatting"),
  z.literal("search"),
  z.literal("instantdetails"),
]);
function handleSection(section: SettingSection) {
  if (get(settings_section) == "shortcuts" && section != "shortcuts") {
    // TODO: enable shortcuts
  } else if (section == "shortcuts") {
    // TODO: disable shortcuts
  }

  settings_section.set(section);
  if (get(graphe_mode) != "settings") {
    graphe_mode.set("settings");
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
