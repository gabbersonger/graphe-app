import { EventsOn, EventsOff } from "!wails/runtime/runtime";
import { GrapheError } from "@/lib/utils";

import { modalData, type ModalName } from "@/components/Modals/data";
import { app_mode, ui_modal, ui_showSidebar } from "@/lib/stores";
import type { AppMode } from "@/lib/manager";
import { get } from "svelte/store";

const event_list = {
  "app:mode": (data: AppMode) => {
    app_mode.set(data);
    if (data == "search")
      ui_modal.update((val) => (val == data ? "" : "search"));
    else if (get(ui_modal) != "") ui_modal.set("");
  },
  "ui:modal": (data: ModalName) => {
    if (!modalData.some((x) => x.name == data))
      return GrapheError(
        `Invalid modal name (\`${data}\`) passed in event (\`ui:modal\`)`,
      );
    if (data == "choosePassage" && get(app_mode) == "search") return;
    ui_modal.update((val) => (val == data ? "" : data));
  },
  "ui:sidebar:toggle": () => {
    ui_showSidebar.update((x) => !x);
  },
};

export function eventListener(_: HTMLElement) {
  for (const [event, callback] of Object.entries(event_list)) {
    EventsOn(event, callback);
  }

  return {
    destroy() {
      for (const [event, _] of Object.entries(event_list)) {
        EventsOff(event);
      }
    },
  };
}
