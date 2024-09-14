import { Logger } from "!/log/slog";

export const assertUnreachable = (_: never): never => {
  throw new Error("Cannot reach here");
};

export function GrapheLog(severity: "info" | "error", message: string) {
  const prefix = "Javascript: ";
  switch (severity) {
    case "info":
      Logger.Info(prefix + message);
      return;
    case "error":
      Logger.Error(prefix + message);
      throw Error(message);
  }
}
