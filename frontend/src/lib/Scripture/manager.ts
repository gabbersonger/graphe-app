import {
  app_data,
  app_instantDetails,
  app_version,
} from "@/lib/managers/appManager";
import { GetScriptureSection, GetScriptureWord } from "!wails/go/app/App";
import { get } from "svelte/store";
import type { BibleRef } from "@/lib/Scripture/types";
import { createVersionRange } from "@/lib/Scripture/version";

export async function updateBaseData() {
  const range = createVersionRange(get(app_version));
  const data = await GetScriptureSection(range);
  app_data.set(data);
}

export async function instantDetails(ref: BibleRef, word_number: number) {
  let data = await GetScriptureWord(get(app_version), ref, word_number);
  app_instantDetails.set(data);
}
