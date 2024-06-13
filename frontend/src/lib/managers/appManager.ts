import { get, writable, type Writable } from "svelte/store";
import { EventsOn, EventsOff, EventsEmit } from "!wails/runtime/runtime";

import type { database } from "!wails/go/models";
import type { BibleRef, BibleVersion } from "@/lib/Scripture/types";
import { instantDetails, updateBaseData } from "@/lib/Scripture/manager";
import type { ModalName } from "@/components/Modals/data";
import { sidebarData, type SidebarSection } from "@/components/Sidebar/data";

// Types

export type AppMode = "passage" | "search";
export type SearchQuery = {};
export type SearchResult = {};

// Data stores

export const app_mode: Writable<AppMode> = writable("passage");
export const app_modal: Writable<ModalName | ""> = writable("");
export const app_sidebar: Writable<boolean> = writable(false);
export const app_sidebarSection: Writable<SidebarSection> = writable(
  sidebarData[0].name,
);

export const app_version: Writable<BibleVersion> = writable();
export const app_data: Writable<database.ScriptureSection[]> = writable([]);
export const app_instantDetails: Writable<database.ScriptureWordData> =
  writable();
export const app_currentRef: Writable<BibleRef> = writable(40_001_001);

// Functions to handle events

function handleMode(mode: AppMode) {
  app_mode.set(mode);
  if (mode == "search") EventsEmit("app:modal", "search");
  else if (get(app_modal) != "") EventsEmit("app:modal:close");
}

function handleModal(data: ModalName) {
  if (data == "text" && get(app_mode) == "search") return;
  app_modal.update((val) => (val == data ? "" : data));
}

function handleModalClose() {
  app_modal.set("");
}

function handleSidebar(data: SidebarSection) {
  app_sidebar.set(true);
  app_sidebarSection.set(data);
}

function handleSidebarToggle() {
  app_sidebar.update((x) => !x);
}

function handleVersion(version: BibleVersion) {
  if (get(app_version) != version) {
    app_currentRef.set(null);
    app_data.set([]);
    app_version.set(version);
    updateBaseData();
  }
}

function handleGoTo(ref: BibleRef) {
  EventsEmit("visualiser:goto", ref);
}

function handleInstantDetails(ref: BibleRef, word_num: number) {
  let current = get(app_instantDetails);
  if (!(current && current.ref == ref && current.word_number == word_num))
    instantDetails(ref, word_num);
}

function handleInstantDetailsHide() {
  app_instantDetails.set(null);
}

// Register events

export function appManager(_: HTMLElement) {
  const handlers: {
    [key: string]: (...data: any) => void;
  } = {
    "app:mode": handleMode,
    "app:modal": handleModal,
    "app:modal:close": handleModalClose,
    "app:sidebar": handleSidebar,
    "app:sidebar:toggle": handleSidebarToggle,
    "app:version": handleVersion,
    "app:goto": handleGoTo,
    "app:instantdetails": handleInstantDetails,
    "app:instantdetails:hide": handleInstantDetailsHide,
  };

  for (const [event, callback] of Object.entries(handlers)) {
    EventsOn(event, callback);
  }

  // Start the app
  EventsEmit("app:version", "esv");

  return {
    destroy() {
      for (const [event, _] of Object.entries(handlers)) {
        EventsOff(event);
      }
    },
  };
}
