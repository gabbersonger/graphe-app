import { app_data, app_instantDetails, app_version } from "@/lib/appManager";
import { GetScriptureSections, GetScriptureWord } from "!wails/go/app/App";
import { get } from "svelte/store";
import { createBibleRef } from "@/lib/Scripture/ref";
import type { BibleRef } from "@/lib/Scripture/types";
import type { app } from "!wails/go/models";

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

export async function instantDetails(ref: BibleRef, word_number: number) {
  let data = await GetScriptureWord(ref, word_number);
  app_instantDetails.set(data);
}
