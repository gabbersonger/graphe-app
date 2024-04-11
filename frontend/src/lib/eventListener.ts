import { EventsOn, EventsOff } from "!wails/runtime/runtime";
import { GrapheError } from "@/lib/utils";

import { modalData, type ModalName } from "@/components/Modals/data";
import { ui_modal, ui_showSidebar } from "@/stores/app";

const event_list = {
  "ui:modal": (data: ModalName) => {
    if (!modalData.some((x) => x.name == data))
      return GrapheError(
        `Invalid modal name (\`${data}\`) passed in event (\`ui:modal\`)`,
      );
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
