import { LogWarning } from "!wails/runtime/runtime";

export const assertUnreachable = (_: never): never => {
  throw new Error("Cannot reach here");
};

export const GrapheError = (message: string): never => {
  LogWarning("Javascript: " + message);
  throw Error(message);
};
