import { Events } from "@wailsio/runtime";
import { GrapheLog } from "@/lib/utils";
import { z } from "zod";

export class EventHandler {
  events: Set<string>;

  constructor() {
    this.events = new Set();
  }

  subscribe<
    H extends (...data: any) => void,
    S extends Parameters<H> extends []
      ? undefined
      : Parameters<H> extends [any]
        ? z.ZodSchema
        : never,
  >(event: string, handler: H, schema: S) {
    if (!this.events.has(event)) {
      Events.On(event, function (event_data: any) {
        if (schema == null) return handler();

        if (!(event_data.data instanceof Array) || event_data.data.length < 1) {
          console.log(event_data);
          GrapheLog(
            "error",
            `Event data is not an array (event: \`${event}\`)`,
          );
          return;
        }
        const result = schema.safeParse(event_data.data[0]);
        if (result.success) {
          GrapheLog("info", `Event handled by frontend (event: \`${event}\`)`);
          handler(result.data);
        } else {
          console.log(event_data, result);
          GrapheLog(
            "error",
            `Failed to parse event data (event: \`${event}\`)`,
          );
        }
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
