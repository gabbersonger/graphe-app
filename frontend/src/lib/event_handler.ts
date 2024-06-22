import { EventsOn, EventsOff } from "!wails/runtime/runtime";
import { GrapheLog } from "@/lib/utils";

export class EventHandler {
  events: Set<string>;

  constructor() {
    this.events = new Set();
  }

  subscribe(event: string, handler: (...data: any) => void) {
    if (!this.events.has(event)) {
      EventsOn(event, (...data) => {
        GrapheLog("info", `Event Handled: ${event}`);
        handler(...data);
      });
      this.events.add(event);
    }
  }

  shutdown() {
    for (const event of this.events.values()) {
      EventsOff(event);
    }
  }
}
