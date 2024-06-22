import { EventHandler } from "../event_handler";

export function windowWorkspaceManager(_: HTMLElement) {
  const events = new EventHandler();

  return {
    destroy() {
      events.shutdown();
    },
  };
}
