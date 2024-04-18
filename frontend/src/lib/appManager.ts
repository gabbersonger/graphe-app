import { get, writable, type Writable } from "svelte/store";
import { EventsOn, EventsOff, EventsEmit } from "!wails/runtime/runtime";
import { ui_modal } from "@/lib/uiManager";

import type { app } from "!wails/go/models";
import type { BibleRef, BibleVersion } from "@/lib/Scripture/types";
import { updateBaseData } from "@/lib/Scripture/manager";

// Types

export type AppMode = "passage" | "search";
export type SearchQuery = {};
export type SearchResult = {};

// Data stores

export const app_mode: Writable<AppMode> = writable("passage"); // passage/search mode
export const app_version: Writable<BibleVersion> = writable("gnt"); // esv/hot/lxx/gnt
export const app_range: Writable<app.ScriptureRange> = writable(); // passage mode: what's visible
export const app_search_query: Writable<SearchQuery> = writable();
export const app_search_result: Writable<SearchResult> = writable(); // search mode: what's visible
export const app_data: Writable<app.ScriptureSection[]> = writable([]); // all the data for version
export const app_currentRef: Writable<BibleRef> = writable(40_001_001);

// Functions to handle events

function handleAppMode(mode: AppMode) {
  app_mode.set(mode);
  if (mode == "search") EventsEmit("ui:modal", "search");
  else if (get(ui_modal) != "") EventsEmit("ui:modal:close");
}

function handleAppVersion(version: BibleVersion) {
  if (get(app_version) != version) {
    app_version.set(version);
    updateBaseData();
  }
}

function handleAppSearch(query: SearchQuery) {
  console.log(`TODO: handleAppSearch for ${query}`);
}

function handleAppRange(range: app.ScriptureRange) {
  console.log(`TODO: handleAppPassageRange for ${range}`);
}

function handleAppGoTo(ref: BibleRef) {
  EventsEmit("visualiser:goto", ref);
}

// Register events

export function appManager(_: HTMLElement) {
  updateBaseData();

  const handlers: {
    [key: string]: (...data: any) => void;
  } = {
    "app:mode": handleAppMode,
    "app:version": handleAppVersion,
    "app:search": handleAppSearch,
    "app:range": handleAppRange,
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
