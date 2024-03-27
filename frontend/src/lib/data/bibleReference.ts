import { bibleData, type BibleBook } from "@/lib/data/bible";
import { GrapheError } from "@/lib/utils";

export type BiblePoint = {
  book: BibleBook;
  chapter: number;
  verse: number;
};

export type BibleRef = number;

export type BibleRange = {
  start: BiblePoint;
  end: BiblePoint;
};

/**
 * Determines if a BiblePoint is valid (that is, chapter and verse exist in given bible book)
 * @param {BiblePoint} point - The point to be validated.
 */
const isValidBiblePoint = (point: BiblePoint): boolean => {
  if (point == null) return false;
  const book_index = bibleData.findIndex((b) => b.name == point.book);
  return !(
    book_index == -1 ||
    (book_index == 18 && point.chapter < 0) ||
    (book_index != 18 && point.chapter <= 0) ||
    point.chapter > bibleData[book_index].num_chapters ||
    point.verse <= 0 ||
    point.verse > bibleData[book_index].num_verses[point.chapter - 1]
  );
};

/**
 * Convert a BiblePoint to a database reference
 * @throws Will throw if `point` is not valid.
 * @param {BiblePoint} point - The point to be converted.
 */
export const biblePointToRef = (point: BiblePoint): BibleRef => {
  if (!isValidBiblePoint(point))
    GrapheError("Invalid point passed to `biblePointToRef`");
  const book_num = bibleData.findIndex((b) => b.name == point.book) + 1;
  const ref =
    String(book_num) +
    String(point.chapter).padStart(3, "0") +
    String(point.verse).padStart(3, "0");
  return parseInt(ref);
};

/**
 * Convert a database reference to a BiblePoint
 * @throws Will throw if `ref` does not produce valid BiblePoint.
 * @param {BibleRef} ref - The database reference (e.g. 1001001 = "Gen 1:1").
 */
export const bibleRefToPoint = (ref: BibleRef): BiblePoint => {
  const verse = ref % 1000;
  const chapter = ((ref % 1000000) - verse) / 1000;
  const book_num = (ref - chapter * 1000 - verse) / 1000000;

  let point: BiblePoint = null;
  if (book_num > 0 && book_num <= bibleData.length) {
    point = {
      book: bibleData[book_num - 1].name,
      chapter: chapter,
      verse: verse,
    };
  }
  if (!isValidBiblePoint(point))
    GrapheError("Invalid ref passed to `bibleRefToPoint`");
  return point;
};

/**
 * Convert a BiblePoint to a reference string
 * @throws Will throw if `point` is not valid.
 * @param {BiblePoint} point - The point to be converted.
 * @param {"short"|"long"} format - The format – "short" (Rev 1:1) or "long" (Revelation 1:1).
 */
export const biblePointToString = (
  point: BiblePoint,
  format: "short" | "long",
): string => {
  if (!isValidBiblePoint(point))
    GrapheError("Invalid point passed to `biblePointToString`");
  for (let i = 0; i < bibleData.length; i++) {
    const book = bibleData[i];
    if (book.name == point.book) {
      switch (format) {
        case "short":
          return `${book.abbreviation} ${point.chapter}:${point.verse}`;
        case "long":
          return `${book.name} ${point.chapter}:${point.verse}`;
      }
    }
  }
  GrapheError("Invalid point passed to `biblePointToString`");
};

/**
 * Create a valid BiblePoint at the start of a specific chapter
 * @throws Will throw if resulting `point` is not valid.
 * @param {BibleBook} book - The book name.
 * @param {number} chapter - The chapter of the book.
 */
export const createBiblePointAtChapterStart = (
  book: BibleBook,
  chapter: number,
): BiblePoint => {
  let point: BiblePoint = {
    book: book,
    chapter: chapter,
    verse: 1,
  };
  if (!isValidBiblePoint(point))
    GrapheError("Invalid book or chapter passed to `getFirstPoint`");
  return point;
};

/**
 * Create a valid BiblePoint at the end of a specific chapter
 * @throws Will throw if resulting `point` is not valid.
 * @param {BibleBook} book - The book name.
 * @param {number} chapter - The chapter of the book.
 */
export const createBiblePointAtChapterEnd = (
  book: BibleBook,
  chapter: number,
): BiblePoint => {
  let point: BiblePoint = null;
  const book_index = bibleData.findIndex((b) => b.name == book);
  if (
    book_index != -1 &&
    chapter > 0 &&
    chapter <= bibleData[book_index].num_chapters
  ) {
    point = {
      book: book,
      chapter: chapter,
      verse: bibleData[book_index].num_verses[chapter - 1],
    };
  }
  if (!isValidBiblePoint(point))
    GrapheError("Invalid book or chapter passed to `getLastPoint`");
  return point;
};
