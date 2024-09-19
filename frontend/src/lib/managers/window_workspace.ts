import { Events } from "@wailsio/runtime";
import { get } from "svelte/store";
import { EventHandler } from "@/lib/event_handler";
import {
  workspace_ref,
  workspace_data,
  workspace_instantDetailsData,
  workspace_modal,
  workspace_mode,
  workspace_sidebar,
  workspace_version,
  type WorkspaceMode,
} from "@/lib/stores";
import { type ModalName } from "@/components/Workspace/Modals/data";
import { DataDB } from "!/graphe/internal/data";
import {
  ScriptureService,
  type ScriptureRef,
  type ScriptureVersion,
} from "!/graphe/internal/scripture";

function handleMode(mode: WorkspaceMode) {
  workspace_mode.set(mode);
  if (mode == "search")
    Events.Emit({ name: "window:workspace:modal", data: "search" });
  else if (get(workspace_modal) != "")
    Events.Emit({ name: "window:workspace:modal", data: "" });
}

function handleModal(modal: ModalName) {
  if (modal == "text" && get(workspace_mode) == "search") return;
  workspace_modal.update((val) => (val == modal ? "" : modal));
}

function handleModalClose() {
  workspace_modal.set("");
}

function handleSidebar(mode: boolean | "toggle") {
  if (typeof mode == "boolean") {
    workspace_sidebar.set(mode);
  } else workspace_sidebar.set(!get(workspace_sidebar));
}

async function updateBaseData() {
  const range = await ScriptureService.GetVersionRange(get(workspace_version));
  const data = await DataDB.GetScriptureSection(range);
  workspace_data.set(data);
}

function handleVersion(version: ScriptureVersion) {
  if (get(workspace_version) != version) {
    workspace_ref.set(undefined);
    workspace_data.set([]);
    workspace_version.set(version);
    updateBaseData();
  }
}

function handleGoTo(ref: ScriptureRef) {
  Events.Emit({ name: "window:workspace:visualiser:goto", data: ref });
}

async function instantDetails(ref: ScriptureRef, word_number: number) {
  let data = await DataDB.GetScriptureWord(
    get(workspace_version),
    ref,
    word_number,
  );
  workspace_instantDetailsData.set(data);
}

function handleInstantDetails(ref: ScriptureRef, word_number: number) {
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

  events.subscribe("window:workspace:version", handleVersion);
  events.subscribe("window:workspace:goto", handleGoTo);
  events.subscribe("window:workspace:instantdetails", handleInstantDetails);
  events.subscribe(
    "window:workspace:instantdetails:hide",
    handleInstantDetailsHide,
  );

  Events.Emit({ name: "window:workspace:version", data: "esv" });

  return {
    destroy() {
      events.shutdown();
    },
  };
}
