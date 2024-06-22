import { EventsEmit } from "!wails/runtime/runtime";
import { get } from "svelte/store";
import { EventHandler } from "@/lib/event_handler";
import {
  workspace_currentRef,
  workspace_data,
  workspace_instantDetailsData,
  workspace_modal,
  workspace_mode,
  workspace_sidebar,
  workspace_sidebarSection,
  workspace_version,
  type WorkspaceMode,
} from "@/lib/stores";
import { type ModalName } from "@/components/Workspace/Modals/data";
import { type SidebarSection } from "@/components/Workspace/Sidebar/data";
import type { BibleRef, BibleVersion } from "@/lib/Scripture/types";
import { GetScriptureSection, GetScriptureWord } from "!wails/go/app/App";
import { createVersionRange } from "../Scripture/version";

function handleMode(mode: WorkspaceMode) {
  workspace_mode.set(mode);
  if (mode == "search") EventsEmit("window:workspace:modal", "search");
  else if (get(workspace_modal) != "") EventsEmit("window:workspace:modal", "");
}

function handleModal(modal: ModalName) {
  if (modal == "text" && get(workspace_mode) == "search") return;
  workspace_modal.update((val) => (val == modal ? "" : modal));
}

function handleModalClose() {
  workspace_modal.set("");
}

function handleSidebar(section: SidebarSection) {
  workspace_sidebar.set(true);
  workspace_sidebarSection.set(section);
}

function handleSidebarToggle() {
  workspace_sidebar.update((val) => !val);
}

async function updateBaseData() {
  const range = createVersionRange(get(workspace_version));
  const data = await GetScriptureSection(range);
  workspace_data.set(data);
}

function handleVersion(version: BibleVersion) {
  if (get(workspace_version) != version) {
    workspace_currentRef.set(null);
    workspace_data.set([]);
    workspace_version.set(version);
    updateBaseData();
  }
}

function handleGoTo(ref: BibleRef) {
  EventsEmit("window:workspace:visualiser:goto", ref);
}

async function instantDetails(ref: BibleRef, word_number: number) {
  let data = await GetScriptureWord(get(workspace_version), ref, word_number);
  workspace_instantDetailsData.set(data);
}

function handleInstantDetails(ref: BibleRef, word_number: number) {
  let current = get(workspace_instantDetailsData);
  if (!(current && current.ref == ref && current.word_number == word_number)) {
    instantDetails(ref, word_number);
  }
}

function handleInstantDetailsHide() {
  workspace_instantDetailsData.set(undefined);
}

export function windowWorkspaceManager(_: HTMLElement) {
  const events = new EventHandler();
  events.subscribe("window:workspace:mode", handleMode);
  events.subscribe("window:workspace:modal", handleModal);
  events.subscribe("window:workspace:modal:close", handleModalClose);
  events.subscribe("window:workspace:sidebar", handleSidebar);
  events.subscribe("window:workspace:sidebar:toggle", handleSidebarToggle);

  events.subscribe("window:workspace:version", handleVersion);
  events.subscribe("window:workspace:goto", handleGoTo);
  events.subscribe("window:workspace:instantdetails", handleInstantDetails);
  events.subscribe(
    "window:workspace:instantdetails:hide",
    handleInstantDetailsHide,
  );

  EventsEmit("window:workspace:version", "esv");

  return {
    destroy() {
      events.shutdown();
    },
  };
}
