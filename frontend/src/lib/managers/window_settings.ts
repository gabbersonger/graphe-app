import { EventHandler } from "@/lib/event_handler";
import { settings_section } from "@/lib/stores";
import type { SettingSection } from "@/components/Settings/data";
import { DisableMenu, LoadMenu } from "!wails/go/app/App";

function handleSection(section: SettingSection) {
  settings_section.set(section);
  if (section == "shortcuts") {
    DisableMenu();
  } else {
    LoadMenu();
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
