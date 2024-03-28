import { Warn } from "!wails/app/App";

export const assertUnreachable = (_: never): never => {
  throw new Error("Cannot reach here");
};

export const GrapheError = (message: string): never => {
  Warn("Javascript: " + message);
  throw Error(message);
};
