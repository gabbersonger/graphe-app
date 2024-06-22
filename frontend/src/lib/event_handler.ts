import { EventsOn, EventsOff, LogInfo } from "!wails/runtime/runtime";

export class EventHandler {
  events: Set<string>;

  constructor() {
    this.events = new Set();
  }

  subscribe(event: string, handler: (...data: any) => void) {
    if (!this.events.has(event)) {
      EventsOn(event, (...data) => {
        LogInfo(`Event Handled: ${event}`);
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
