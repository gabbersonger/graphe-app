import { EventHandler } from "@/lib/event_handler";
import { settings_section } from "@/lib/stores";
import type { SettingSection } from "@/components/Settings/data";

function handleSection(mode: SettingSection) {
  settings_section.set(mode);
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
