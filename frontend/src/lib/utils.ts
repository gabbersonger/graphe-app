import { Throw } from "!wails/app/App";

export const assertUnreachable = (_: never): never => {
  throw new Error("Cannot reach here");
};

export const GrapheError = (message: string): never => {
  Throw("Javascript: " + message);
  throw Error(message);
};
