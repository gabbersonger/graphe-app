import { EventHandler } from "@/lib/event_handler";
import { settings_section } from "@/lib/stores";
import type { SettingSection } from "@/components/Settings/data";

function handleSection(section: SettingSection) {
  settings_section.set(section);
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
