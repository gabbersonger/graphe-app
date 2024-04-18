import type { BibleRef, BibleVersion } from "@/lib/Scripture/types";
import { getBook, isValidBibleRef } from "@/lib/Scripture/ref";
import { versionData } from "@/lib/Scripture/data";
import type { app } from "!wails/go/models";

const getVersionBookIndex = (
  version: BibleVersion,
  book_number: number,
): number => {
  return versionData[version].books.findIndex(
    (b: { book_number: number }) => b.book_number == book_number,
  );
};

/**
 * Determines if a range is valid
 * @param {app.ScriptureRange} range - The range to check.
 */
export const isValidRange = (range: app.ScriptureRange): boolean => {
  if (!isValidBibleRef(range.start) || !isValidBibleRef(range.end))
    return false;

  const version = range.version as BibleVersion;
  const version_start = getVersionBookIndex(version, getBook(range.start));
  const version_end = getVersionBookIndex(version, getBook(range.end));

  return (
    version_start != -1 &&
    version_end != -1 &&
    (version_start < version_end ||
      (version_start == version_end && range.start < range.end))
  );
};

/**
 * Determines if a reference is within a range
 * @param {BibleRef} ref - The reference to check.
 * @param {app.ScriptureRange} range - The range to check (note each
 * BibleVersion defines book orders differently).
 */
export const isRefInRange = (
  ref: BibleRef,
  range: app.ScriptureRange,
): boolean => {
  const book_start = getBook(range.start);
  const book_end = getBook(range.end);
  if (book_start == book_end) return ref >= range.start && ref <= range.end;

  const book_ref = getBook(ref);
  if (book_ref == book_start) return ref >= range.start;
  if (book_ref == book_end) return ref <= range.end;

  const version = range.version as BibleVersion;
  const version_start = getVersionBookIndex(version, book_start);
  const version_end = getVersionBookIndex(version, book_end);
  const version_ref = getVersionBookIndex(version, book_ref);

  return version_ref >= version_start && version_ref <= version_end;
};
