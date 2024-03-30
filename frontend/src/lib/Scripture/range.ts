import type { BibleRange } from "@/lib/Scripture/types";
import { isValidBibleRef } from "@/lib/Scripture/ref";
import { bibleData } from "@/lib/Scripture/data";

export const isValidBibleRange = (range: BibleRange): boolean => {
  if (
    !isValidBibleRef(range.start) ||
    !isValidBibleRef(range.end) ||
    range.start >= range.end
  )
    return false;

  const start_book = (range.start - (range.start % 1000000)) / 1000000;
  const end_book = (range.end - (range.end % 1000000)) / 1000000;
  for (let i = start_book; i <= end_book; i++) {
    if (!bibleData[i - 1].version.some((v) => v == range.version)) return false;
  }
  return true;
};
