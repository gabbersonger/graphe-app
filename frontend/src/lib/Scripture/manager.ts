import {
  app_data,
  app_instantDetails,
  app_range,
  app_version,
} from "@/lib/managers/appManager";
import { GetScriptureSection, GetScriptureWord } from "!wails/go/app/App";
import { get } from "svelte/store";
import type { BibleRef } from "@/lib/Scripture/types";

export async function updateBaseData() {
  const data = await GetScriptureSection(get(app_range));
  app_data.set(data);
}

export async function instantDetails(ref: BibleRef, word_number: number) {
  let data = await GetScriptureWord(get(app_version), ref, word_number);
  app_instantDetails.set(data);
}
