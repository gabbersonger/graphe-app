import { Logger } from "!/graphe/internal/logger";
import { Events } from "@wailsio/runtime";

export const assertUnreachable = (_: never): never => {
  throw new Error("Cannot reach here");
};

export function GrapheLog(severity: "info" | "error", message: string) {
  const prefix = "Javascript";
  switch (severity) {
    case "info":
      Logger.Log(prefix, message);
      return;
    case "error":
      Logger.Log(prefix, message);
      throw Error(message);
  }
}

export function GrapheEvent(name: string, data: any = null) {
  const formatted_data = data == null ? null : [data];
  Events.Emit({
    name: name,
    data: formatted_data,
  });
}

export function throttle<R, A extends any[]>(
  fn: (...args: A) => R,
  delay: number,
): (...args: A) => R | undefined {
  let timeout: ReturnType<typeof setTimeout> | null = null;
  return (...args: A) => {
    if (!timeout) {
      const val = fn(...args);
      timeout = setTimeout(() => {
        timeout = null;
      }, delay);
      return val;
    }
  };
}
