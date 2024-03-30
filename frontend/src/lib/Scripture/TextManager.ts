import type { app } from "!wails/go/models";
import type { BibleRef } from "@/lib/Scripture/types";
import { isValidBibleRef } from "@/lib/Scripture/ref";
import { isValidBibleRange } from "@/lib/Scripture/range";
import { GrapheError } from "@/lib/utils";

// const MAX_DIVISIONS = 5;
// const MIN_CHAPTER_PER_DIVISION = 10;

// const countChapterSpan = (start: BibleRef, end: BibleRef): number => {
//   // TODO
//   return 260;
// };

// const getSectionOffsets = (start: BibleRef, end: BibleRef): number[] => {
//   const chapter_span = countChapterSpan(start, end);
//   const amount_sections = Math.max(
//     Math.min(
//       Math.floor(chapter_span / MIN_CHAPTER_PER_DIVISION),
//       MAX_DIVISIONS,
//     ),
//     1,
//   );
//   const chapters_per_section = Math.floor(chapter_span / amount_sections);
//   const sections = new Array(amount_sections).fill(chapters_per_section);
//   sections[0] = chapter_span - chapters_per_section * (amount_sections - 1);
//   return sections;
// };

// const getFetchingRanges = (
//   start: BibleRef,
//   section_offsets: number[],
// ): app.ScriptureRange[] => {
//   let ranges: app.ScriptureRange[] = [];
//   ranges.push({
//     start: start,
//     end: createOffsetBibleRef(start, section_offsets[0], "chapter", "end"),
//   });
//   for (let i = 1; i < section_offsets.length; i++) {
//     const last = ranges[ranges.length - 1].end;
//     ranges.push({
//       start: createOffsetBibleRef(last, 1, "chapter", "start"),
//       end: createOffsetBibleRef(last, 1 + section_offsets[i], "chapter", "end"),
//     });
//   }
//   return ranges;
// };

export class TextManager {
  static GetScriptureSection(
    version: "esv" | "gnt" | "lxx" | "hot",
    start: BibleRef,
    end: BibleRef,
  ): app.ScriptureSection[] {
    const range = { version: version, start: start, end: end };
    if (!isValidBibleRef(start))
      GrapheError(`Invalid start (${start}) passed to \`GetScriptureSection\``);
    if (!isValidBibleRef(end))
      GrapheError(`Invalid end (${end}) passed to \`GetScriptureSection\``);
    if (!isValidBibleRange(range))
      GrapheError(
        `Invalid range (${start} to ${end}) passed to \`GetScriptureSection\``,
      );

    // const section_offsets = getSectionOffsets(start, end);
    // const bible_ranges = getFetchingRanges(start, section_offsets);
    // console.log(bible_ranges);
    return [];
  }
}
