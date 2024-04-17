import { app_data, app_version } from "@/lib/appManager";
import { GetScriptureSections } from "!wails/go/app/App";
import { createBibleRef } from "./ref";
import { get } from "svelte/store";

export async function updateBaseData() {
  let ranges = [
    {
      version: get(app_version),
      start: createBibleRef("Matthew", 1),
      end: createBibleRef("Matthew", 28, "end"),
    },
    {
      version: get(app_version),
      start: createBibleRef("Mark", 1),
      end: createBibleRef("Mark", 16, "end"),
    },
    {
      version: get(app_version),
      start: createBibleRef("Luke", 1),
      end: createBibleRef("Revelation", 22, "end"),
    },
  ];
  let data = await GetScriptureSections(ranges);
  app_data.set(data); // TODO: not sure this updates visualiser (check once other texts in)
}
