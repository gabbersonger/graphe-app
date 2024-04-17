import { get, writable, type Writable } from "svelte/store";
import { EventsOn, EventsOff, EventsEmit } from "!wails/runtime/runtime";
import { ui_modal } from "@/lib/uiManager";

import type { app } from "!wails/go/models";
import type { BibleRange, BibleRef, BibleVersion } from "@/lib/Scripture/types";
import { updateBaseData } from "@/lib/Scripture/manager";

// Types

export type AppMode = "passage" | "search";
export type SearchQuery = {};

// Data stores

export const app_mode: Writable<AppMode> = writable("passage");
export const app_version: Writable<BibleVersion> = writable("gnt");
export const app_range: Writable<BibleRange> = writable();
export const app_search: Writable<SearchQuery> = writable(true);
export const app_data: Writable<app.ScriptureSection[]> = writable([]);
export const app_currentRef: Writable<BibleRef> = writable(40_001_001);

// Functions to handle events

function handleAppMode(mode: AppMode) {
  app_mode.set(mode);
  if (mode == "search") EventsEmit("ui:modal", "search");
  else if (get(ui_modal) != "") EventsEmit("ui:modal:closeAll");
}

function handleAppText(version: BibleVersion) {
  console.log(`TODO: handleAppText for ${version}`);
}

function handleAppSearch(query: SearchQuery) {
  console.log(`TODO: handleAppSearch for ${query}`);
}

function handleAppPassageRange(range: BibleRange) {
  console.log(`TODO: handleAppPassageRange for ${range}`);
}

function handleAppPassageWhole() {
  console.log(`TODO: handleAppPassageWhole`);
}

function handleAppGoTo(ref: BibleRef) {
  console.log(`TODO: handleAppGoTo for ${ref}`);
}

// Register events

export function appManager(_: HTMLElement) {
  updateBaseData();

  const handlers: {
    [key: string]: (...data: any) => void;
  } = {
    "app:mode": handleAppMode,
    "app:text": handleAppText,
    "app:search": handleAppSearch,
    "app:passage:range": handleAppPassageRange,
    "app:passage:whole": handleAppPassageWhole,
    "app:goto": handleAppGoTo,
  };

  for (const [event, callback] of Object.entries(handlers)) {
    EventsOn(event, callback);
  }

  return {
    destroy() {
      for (const [event, _] of Object.entries(handlers)) {
        EventsOff(event);
      }
    },
  };
}
