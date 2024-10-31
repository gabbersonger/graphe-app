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
import { GrapheEvent, GrapheLog } from "@/lib/utils";
import { z } from "zod";

function handleReset() {
  handleVersion("esv");
}

const z_workspace_mode = z.union([z.literal("passage"), z.literal("search")]);
function handleMode(mode: WorkspaceMode) {
  workspace_mode.set(mode);
  if (mode == "search") GrapheEvent("window:workspace:modal", "search");
  else if (get(workspace_modal) != "")
    GrapheEvent("window:workspace:modal", "");
}

const z_workspace_modal = z.union([
  z.literal("search"),
  z.literal("version"),
  z.literal("text"),
  z.literal("functions"),
]);
function handleModal(modal: ModalName) {
  if (modal == "text" && get(workspace_mode) == "search") return;
  workspace_modal.update((val) => (val == modal ? "" : modal));
}

function handleModalClose() {
  workspace_modal.set("");
}

const z_sidebar = z.boolean().or(z.literal("toggle"));
function handleSidebar(mode: boolean | "toggle") {
  if (typeof mode == "boolean") {
    workspace_sidebar.set(mode);
  } else workspace_sidebar.set(!get(workspace_sidebar));
}

async function updateBaseData() {
  const version = get(workspace_version);
  if (version == undefined) {
    return GrapheLog(
      "error",
      `[Workspace Manager] Invalid version when doing updateBaseData (version: \`${version}\`)`,
    );
  }
  const range = await ScriptureService.GetVersionRange(version);
  const data = await DataDB.GetScriptureSection(range);
  workspace_ref.set(undefined);
  workspace_data.set(data);
}

function handleVersion(version: ScriptureVersion) {
  if (version == undefined) {
    return GrapheLog(
      "error",
      `[Workspace Manager] Invalid version passed to handleVersion (version: \`${version}\`)`,
    );
  }
  if (get(workspace_version) != version) {
    workspace_ref.set(undefined);
    workspace_data.set([]);
    workspace_version.set(version);
    updateBaseData();
  }
}

function handleGoTo(ref: ScriptureRef) {
  GrapheEvent("window:workspace:text:goto", ref);
}

async function instantDetails(ref: ScriptureRef, word_number: number) {
  const version = get(workspace_version);
  if (version == undefined) {
    return GrapheLog(
      "error",
      `[Workspace Manager] Invalid version when doing instantDetails (version: \`${version}\`)`,
    );
  }
  let data = await DataDB.GetScriptureWord(version, ref, word_number);
  workspace_instantDetailsData.set(data);
}

const z_instant_details = z.object({
  ref: z.number(),
  word_number: z.number(),
});

function handleInstantDetails(data: {
  ref: ScriptureRef;
  word_number: number;
}) {
  let current = get(workspace_instantDetailsData);
  if (
    !(
      current &&
      current.ref == data.ref &&
      current.word_number == data.word_number
    )
  ) {
    instantDetails(data.ref, data.word_number);
  }
}

function handleInstantDetailsHide() {
  workspace_instantDetailsData.set(null);
}

export function windowWorkspaceManager(_: HTMLElement) {
  const events = new EventHandler();
  events.subscribe("window:workspace:reset", handleReset, undefined);

  events.subscribe("window:workspace:mode", handleMode, z_workspace_mode);
  events.subscribe("window:workspace:modal", handleModal, z_workspace_modal);
  events.subscribe("window:workspace:modal:close", handleModalClose, undefined);
  events.subscribe("window:workspace:sidebar", handleSidebar, z_sidebar);

  events.subscribe("window:workspace:version", handleVersion, z.string());
  events.subscribe("window:workspace:goto", handleGoTo, z.number());
  events.subscribe(
    "window:workspace:instantdetails",
    handleInstantDetails,
    z_instant_details,
  );
  events.subscribe(
    "window:workspace:instantdetails:hide",
    handleInstantDetailsHide,
    undefined,
  );

  GrapheEvent("window:workspace:reset");

  return {
    destroy() {
      events.shutdown();
    },
  };
}
