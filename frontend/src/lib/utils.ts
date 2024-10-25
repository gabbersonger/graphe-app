import { Logger } from "!/graphe/internal/logger";

export const assertUnreachable = (_: never): never => {
  throw new Error("Cannot reach here");
};

export function GrapheLog(severity: "info" | "error", message: string) {
  const prefix = "Javascript: ";
  switch (severity) {
    case "info":
      Logger.Log(prefix, message);
      return;
    case "error":
      Logger.Log(prefix, message);
      throw Error(message);
  }
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

// export function throttle<F extends Function>(
//   fn: (args: T) => any,
//   delay: number,
// ) {
//   let timeout: ReturnType<typeof setTimeout> | null = null;
//   return (...args: Parameters<T>) => {
//     if (!timeout) {
//       fn(...args);
//       timeout = setTimeout(() => {
//         timeout = null;
//       }, delay);
//     }
//   };
// }
// export function throttle(fn: Function, delay: number) {
//   let timeout: ReturnType<typeof setTimeout> | null = null;
//   return (...args: any[]) => {
//     if (!timeout) {
//       fn(...args);
//       timeout = setTimeout(() => {
//         timeout = null;
//       }, delay);
//     }
//   };
// }
