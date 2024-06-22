import { bibleData, versionData } from "@/lib/Scripture/data";
import type { BibleRef, BibleVersion } from "@/lib/Scripture/types";
import { GrapheLog, assertUnreachable } from "@/lib/utils";
import { getVersionBookIndex, isValidVersion } from "@/lib/Scripture/version";

export function getBook(ref: BibleRef): number {
  return (ref - (ref % 1_000_000)) / 1_000_000;
}

export function getChapter(ref: BibleRef): number {
  return ((ref % 1_000_000) - (ref % 1_000)) / 1_000;
}

export function getVerse(ref: BibleRef): number {
  return ref % 1_000;
}

/**
 * Determines if given chapter has a superscript (verse 0) in a specific version
 * @throws Will throw if `version` is not valid
 * @param {number} version - The version in which to check
 * @param {number} book - The book number to check
 * @param {number} chapter - The chapter to check
 */
function isSuperscriptChapter(
  version: BibleVersion,
  book: number,
  chapter: number,
): boolean {
  if (!isValidVersion(version))
    GrapheLog(
      "error",
      `Invalid version (${version}) passed to \`isSuperscriptChapter\``,
    );

  const bookDataIndex = getVersionBookIndex(version, book);
  if (bookDataIndex < 0) return false;
  const bookData = versionData[version].books[bookDataIndex];

  if ("superscripts" in bookData) {
    const superscripts = bookData.superscripts;
    return superscripts.some((c) => c == chapter);
  }
  return false;
}

/**
 * Determines if a BibleRef is strongly valid. That is, the chapter and verse
 * exist in given bible book in that specific version. (Missing verses in versions
 * will be an invalid BibleRef.)
 * @throws Will throw if `version` is not valid
 * @param {BibleVersion} version - The version in which to check
 * @param {BibleRef} ref - The reference to be validated
 */
export function isValidRef(version: BibleVersion, ref: BibleRef): boolean {
  if (!isValidVersion(version))
    GrapheLog(
      "error",
      `Invalid version (${version}) passed to \`isSuperscriptChapter\``,
    );
  if (ref == null) return false;

  const verse = getVerse(ref);
  const chapter = getChapter(ref);
  const book = getBook(ref);

  // Handle book invalid
  const bookDataIndex = getVersionBookIndex(version, book);
  if (bookDataIndex < 0) return false;
  const bookData = versionData[version].books[bookDataIndex];

  // Handle prologue case
  if (chapter == 0 && "prologue" in bookData) {
    return verse > 0 && verse <= bookData.prologue;
    // Note: This is because no prologue has superscript
  }

  // Handle chapter invalid
  if (chapter <= 0 || chapter > bookData.num_chapters) return false;

  // Handle verse invalid
  if (verse == 0) return isSuperscriptChapter(version, book, chapter);
  if (verse < 0 || verse > bookData.num_verses[chapter - 1]) return false;

  // Handle missing sections
  for (let i = 0; i < bookData.missing_sections.length; i++) {
    const section = bookData.missing_sections[i];
    if (getChapter(section.start) == chapter) {
      const sectionStartVerse = getVerse(section.start);
      const sectionEndVerse = getVerse(section.end);
      if (verse >= sectionStartVerse && verse <= sectionEndVerse) return false;
    }
  }

  return true;
}

/**
 * Create a valid BibleRef
 * @throws Will throw if `version` is invalid or the resulting BibleRef is not valid.
 * @param {BibleVersion} version - The version for the new ref
 * @param {number} book - The book for the new ref
 * @param {number} chapter - The chapter for the new ref
 * @param {number|"start"|"end"} verse - The verse for the new ref.
 * The options are: (1) **any valid number**, (2) **"start"** which
 * puts the resulting ref at the start of the chapter, and (3) **"end"**
 * which puts the resulting ref at the end of the chapter.
 */
export function createRef(
  version: BibleVersion,
  book: number,
  chapter: number,
  verse: number | "start" | "end" = "start",
) {
  if (!isValidVersion(version))
    GrapheLog(
      "error",
      `Invalid version (${version}) passed to \`createBibleRef\``,
    );

  let ref: BibleRef = 0;

  const bookDataIndex = getVersionBookIndex(version, book);
  if (bookDataIndex >= 0) {
    let verse_num: number;
    const bookData = versionData[version].books[bookDataIndex];

    if (chapter == 0 && "prologue" in bookData) {
      if (verse == "start") verse_num = 1;
      else if (verse == "end") verse_num = bookData.prologue;
      else verse_num = verse;
    } else if (chapter > 0 && chapter <= bookData.num_chapters) {
      if (verse == "start")
        verse_num = isSuperscriptChapter(version, book, chapter) ? 0 : 1;
      else if (verse == "end") verse_num = bookData.num_verses[chapter - 1];
      else verse_num = verse;
    }

    ref = parseInt(
      String(book) +
        String(chapter).padStart(3, "0") +
        String(verse_num).padStart(3, "0"),
    );
  }

  if (!isValidRef(version, ref))
    GrapheLog(
      "error",
      `Invalid combination of version (${version}), book (${book}), chapter (${chapter}) and verse (${verse}) passed to \`createRef\``,
    );
  return ref;
}

/**
 * Convert a BibleRef into a scripture reference string
 * @throws Will throw if `version` is not valid
 * @returns blank if invalid BibleRef in version
 * @param {BibleVersion} version - The version to check against
 * @param {BibleRef} ref - The reference to be converted
 * @param {"short"|"long"|"chapter"} format - Determines the format of the scripture reference string produced...
 * **short** – A shorter reference (e.g. Gen 1:1, John 15:12, 3Jo 14).
 * **long** – A longer reference (e.g. Genesis 1:1, John 15:12, 3 John 14).
 * **chapter** – A chapter reference (e.g. Genesis 1, John 15, 3 John).
 * **book** – A book reference (e.g. Genesis, John, 3 John).
 */
export function refToString(
  version: BibleVersion,
  ref: BibleRef,
  format: "short" | "long" | "chapter" | "book",
) {
  if (!isValidVersion(version))
    GrapheLog(
      "error",
      `Invalid version (${version}) passed to \`refToString\``,
    );
  if (!isValidRef(version, ref)) return "";

  const verse = getVerse(ref);
  const chapter = getChapter(ref);
  const book = getBook(ref);

  const bookDataIndex = getVersionBookIndex(version, book);
  const versionBookData = versionData[version].books[bookDataIndex];
  const bookData = bibleData[versionBookData.book_number - 1];

  const isSingleChapterBook = versionBookData.num_chapters == 1;
  switch (format) {
    case "short":
      return isSingleChapterBook
        ? `${bookData.short} ${verse}`
        : `${bookData.short} ${chapter}:${verse}`;
    case "long":
      return isSingleChapterBook
        ? `${bookData.name} ${verse}`
        : `${bookData.name} ${chapter}:${verse}`;
    case "chapter":
      return isSingleChapterBook
        ? `${bookData.short}`
        : `${bookData.short} ${chapter}`;
    case "book":
      return bookData.name;
  }
  assertUnreachable(format);
}
