import Main from "./Main.svelte";
import "./main.css";

const app = new Main({
  target: document.getElementById("graphe") as HTMLElement,
});

export default app;
