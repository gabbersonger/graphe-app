import type { BibleRef, BibleVersion } from "@/lib/Scripture/types";
import { getBook, isValidRef } from "@/lib/Scripture/ref";
import { versionData } from "@/lib/Scripture/data";
import type { app } from "!wails/go/models";
import { GrapheError } from "@/lib/utils";
import { getVersionBookIndex, isValidVersion } from "@/lib/Scripture/version";

/**
 * Determines if a range is valid
 * @param {app.ScriptureRange} range - The range to check.
 */
export function isValidRange(range: app.ScriptureRange): boolean {
  if (!isValidVersion(range.version)) return false;

  const version = range.version as BibleVersion;
  if (!isValidRef(version, range.start) || !isValidRef(version, range.end))
    return false;

  const start_book_index = getVersionBookIndex(version, getBook(range.start));
  const end_book_index = getVersionBookIndex(version, getBook(range.end));

  return (
    start_book_index < end_book_index ||
    (start_book_index == end_book_index && range.start < range.end)
  );
}

/**
 * Determines if a reference is within a range
 * @param {BibleRef} ref - The reference to check.
 * @param {app.ScriptureRange} range - The range to check in (which contains
 * a version so we know book order)
 */
export const isRefInRange = (
  ref: BibleRef,
  range: app.ScriptureRange,
): boolean => {
  if (!isValidRange(range))
    GrapheError(
      `Invalid range ({version: ${range.version}, start: ${range.start}, end: ${range.end}) passed to \`isRefInRange\``,
    );
  const version = range.version as BibleVersion;
  if (!isValidRef(version, ref)) return false;

  const start_book = getBook(range.start);
  const end_book = getBook(range.end);
  if (start_book == end_book) return ref >= range.start && ref <= range.end;

  const ref_book = getBook(ref);
  if (ref_book == start_book) return ref >= range.start;
  if (ref_book == end_book) return ref <= range.end;

  const start_book_index = getVersionBookIndex(version, start_book);
  const end_book_index = getVersionBookIndex(version, end_book);
  const ref_book_index = getVersionBookIndex(version, ref_book);

  return ref_book_index >= start_book_index && ref_book_index <= end_book_index;
};
