import { LogError, LogInfo } from "!wails/runtime/runtime";

export const assertUnreachable = (_: never): never => {
  throw new Error("Cannot reach here");
};

export function GrapheLog(severity: "info" | "error", message: string) {
  const prefix = "[Javascript] ";
  switch (severity) {
    case "info":
      LogInfo(prefix + message);
      return;
    case "error":
      LogError(prefix + message);
      throw Error(message);
  }
}
