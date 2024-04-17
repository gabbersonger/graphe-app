import { app_data, app_version } from "@/lib/appManager";
import { GetScriptureSections } from "!wails/go/app/App";
import { createBibleRef } from "./ref";
import { get } from "svelte/store";

export async function updateBaseData() {
  let ranges = [
    {
      start: createBibleRef("Matthew", 1),
      end: createBibleRef("Matthew", 28, "end"),
    },
    {
      start: createBibleRef("Mark", 1),
      end: createBibleRef("Mark", 16, "end"),
    },
    {
      start: createBibleRef("Luke", 1),
      end: createBibleRef("Revelation", 22, "end"),
    },
  ];
  let data = await GetScriptureSections(get(app_version), ranges);
  app_data.set(data);
}
