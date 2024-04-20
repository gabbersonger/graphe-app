import { app_data, app_instantDetails, app_range } from "@/lib/appManager";
import { GetScriptureSections, GetScriptureWord } from "!wails/go/app/App";
import { get } from "svelte/store";
import type { BibleRef } from "@/lib/Scripture/types";

export async function updateBaseData() {
  const data = await GetScriptureSections(get(app_range));
  app_data.set(data); // TODO: not sure this updates visualiser (check once other texts in)
}

export async function instantDetails(ref: BibleRef, word_number: number) {
  let data = await GetScriptureWord(ref, word_number);
  app_instantDetails.set(data);
}
