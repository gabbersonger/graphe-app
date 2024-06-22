import { EventHandler } from "@/lib/event_handler";
import { graphe_mode, type GrapheMode } from "@/lib/stores";

function handleMode(mode: GrapheMode) {
  graphe_mode.set(mode);
}

export function grapheManager(_: HTMLElement) {
  const events = new EventHandler();
  events.subscribe("graphe:mode", handleMode);

  return {
    destroy() {
      events.shutdown();
    },
  };
}
