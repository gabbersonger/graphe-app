import { Events } from "@wailsio/runtime";
import { GrapheLog } from "@/lib/utils";

export class EventHandler {
  events: Set<string>;

  constructor() {
    this.events = new Set();
  }

  subscribe(event: string, handler: (...data: any) => void) {
    if (!this.events.has(event)) {
      Events.On(event, function (event_data: any) {
        GrapheLog("info", `Event handled by frontend (event: \`${event}\`)`);
        handler(event_data.data);
      });
      this.events.add(event);
    }
  }

  shutdown() {
    for (const event of this.events.values()) {
      Events.Off(event);
    }
  }
}
