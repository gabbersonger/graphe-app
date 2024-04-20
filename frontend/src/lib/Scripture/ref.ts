import { bibleData, versionData } from "@/lib/Scripture/data";
import type { BibleBook, BibleRef, BibleVersion } from "@/lib/Scripture/types";
import { GrapheError, assertUnreachable } from "@/lib/utils";

export const getBook = (ref: BibleRef): number => {
  return (ref - (ref % 1_000_000)) / 1_000_000;
};

export const getChapter = (ref: BibleRef): number => {
  return ((ref % 1_000_000) - (ref % 1_000)) / 1_000;
};

export const getVerse = (ref: BibleRef): number => {
  return ref % 1_000;
};

/**
 * Determines if given chapter has a superscript (verse 0)
 * @param {number} book - The book to check.
 * @param {number} chapter - The chapter to check.
 */
const isSuperscriptChapter = (book: number, chapter: number) => {
  const bookData = bibleData[book - 1];
  if ("superscripts" in bookData) {
    const superscripts = bookData.superscripts;
    return superscripts.some((c) => c == chapter);
  }
  return false;
};

/**
 * Determines if a BibleRef is valid (that is, chapter and verse exist in given bible book)
 * @param {BibleRef} ref - The reference to be validated.
 */
export const isValidBibleRef = (ref: BibleRef): boolean => {
  if (ref == null) return false;
  const verse = getVerse(ref);
  const chapter = getChapter(ref);
  const book = getBook(ref);

  if (book <= 0 && book > bibleData.length) return false;

  if (
    chapter == 0 &&
    bibleData[book - 1].prologue > 0 &&
    verse > 0 &&
    verse <= bibleData[book - 1].prologue
  )
    return true;

  if (chapter <= 0 || chapter > bibleData[book - 1].num_chapters) return false;
  if (verse == 0 && !isSuperscriptChapter(book, chapter)) return false;
  else if (verse < 0 || verse > bibleData[book - 1].num_verses[chapter - 1])
    return false;

  return true;
};

/**
 * Determines if a BibleRef is the start of a book
 * @param {BibleRef} ref - The reference to be checked.
 */
export const isRefBookStart = (ref: BibleRef): boolean => {
  if (!isValidBibleRef) return false;
  return getChapter(ref) == 1 && getVerse(ref) == 1;
};

/**
 * Create a valid BibleRef
 * @throws Will throw if resulting BibleRef is not valid.
 * @param {BibleBook} book - The book name.
 * @param {number} chapter - The chapter of the book.
 * @param {number|"start"|"end"} verse - The verse of the chapter.
 * **start** – first verse of chapter.
 * **end** – last verse of chapter.
 */
export const createBibleRef = (
  book: BibleBook,
  chapter: number,
  verse: number | "start" | "end" = "start",
) => {
  let ref: BibleRef = 0;
  const book_index = bibleData.findIndex((b) => b.name == book);

  if (book_index >= 0) {
    let verse_num: number;

    if (chapter == 0 && bibleData[book_index].prologue > 0) {
      if (verse == "start") {
        verse_num = 1;
      } else if (verse == "end") {
        verse_num = bibleData[book_index].prologue;
      } else {
        verse_num = verse;
      }
    } else if (chapter > 0 && chapter <= bibleData[book_index].num_chapters) {
      if (verse == "start") {
        verse_num = isSuperscriptChapter(book_index + 1, chapter) ? 0 : 1;
      } else if (verse == "end") {
        verse_num = bibleData[book_index].num_verses[chapter - 1];
      } else {
        verse_num = verse;
      }
    }

    ref = parseInt(
      String(book_index + 1) +
        String(chapter).padStart(3, "0") +
        String(verse_num).padStart(3, "0"),
    );
  }

  if (!isValidBibleRef(ref))
    GrapheError(
      `Invalid book (${book}), chapter (${chapter}) and verse (${verse}) passed to \`createBibleRef\``,
    );
  return ref;
};

/**
 * Convert a BibleRef into a scripture reference string
 * @throws Will throw if BibleRef is not valid.
 * @param {BibleRef} ref - The reference to be converted.
 * @param {"short"|"long"|"chapter"} format - Determines the format of the scripture reference string produced...
 * **short** – A shorter reference (e.g. Gen 1:1, John 15:12, 3Jo 14).
 * **long** – A longer reference (e.g. Genesis 1:1, John 15:12, 3 John 14).
 * **chapter** – A chapter reference (e.g. Genesis 1, John 15, 3 John).
 * **book** – A book reference (e.g. Genesis, John, 3 John).
 */
export const bibleRefToString = (
  ref: BibleRef,
  format: "short" | "long" | "chapter" | "book",
) => {
  if (!isValidBibleRef(ref))
    GrapheError(`Invalid ref (${ref}) passed to \`bibleRefToString\``);

  const verse = getVerse(ref);
  const chapter = getChapter(ref);
  const book = getBook(ref);
  const isSingleChapterBook = bibleData[book - 1].num_chapters == 1;
  switch (format) {
    case "short":
      return isSingleChapterBook
        ? `${bibleData[book - 1].abbreviation} ${verse}`
        : `${bibleData[book - 1].abbreviation} ${chapter}:${verse}`;
    case "long":
      return isSingleChapterBook
        ? `${bibleData[book - 1].name} ${verse}`
        : `${bibleData[book - 1].name} ${chapter}:${verse}`;
    case "chapter":
      return isSingleChapterBook
        ? `${bibleData[book - 1].name}`
        : `${bibleData[book - 1].name} ${chapter}`;
    case "book":
      return bibleData[book - 1].name;
  }
  assertUnreachable(format);
};

/**
 * Convert a BibleRef into it's book title in the given bible version
 * @throws Will throw if BibleRef is not valid or BibleRef not in version.
 * @param {BibleRef} ref - A reference within the book.
 * @param {BibleVersion} ver - The version of the text.
 */
export const bibleRefToVersionBookTitle = (
  ref: BibleRef,
  ver: BibleVersion,
): string => {
  if (!isValidBibleRef(ref))
    GrapheError(
      `Invalid ref (${ref}) passed to \`bibleRefToVersionBookTitle\``,
    );

  const book_number = getBook(ref);
  const book_info = versionData[ver].books.find(
    (b) => b.book_number == book_number,
  );
  if (book_info == undefined)
    GrapheError(
      `Valid ref (${ref}) not in bible version (${ver}), passed to \`bibleRefToVersionBookTitle\``,
    );
  return book_info.name;
};
