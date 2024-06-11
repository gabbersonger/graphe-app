import { EventsEmit, EventsOff, EventsOn } from "!wails/runtime/runtime";
import type { SettingSection } from "@/components/Settings/data";
import { writable, type Writable } from "svelte/store";

// Data stores

export const settings_section: Writable<SettingSection> = writable("general");

// Functions to handle events

function handleSectionMode(mode: SettingSection) {
  console.log("HEARD");
  settings_section.set(mode);
}

export function settingsManager(_: HTMLElement) {
  const handlers: {
    [key: string]: (...data: any) => void;
  } = {
    "settings:section": handleSectionMode,
  };

  for (const [event, callback] of Object.entries(handlers)) {
    EventsOn(event, callback);
  }

  EventsEmit("settings:section", "appearence");

  return {
    destroy() {
      for (const [event, _] of Object.entries(handlers)) {
        EventsOff(event);
      }
    },
  };
}
