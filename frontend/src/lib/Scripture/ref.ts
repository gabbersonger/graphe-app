import { bibleData } from "@/lib/Scripture/data";
import type { BibleBook, BibleRef } from "@/lib/Scripture/types";
import { GrapheError, assertUnreachable } from "@/lib/utils";

/**
 * Determines if given chapter is a superscript psalm
 * @param {number} chapter - The psalm chapter to check.
 */
const isSuperscriptPsalmChapter = (chapter: number): boolean => {
  return bibleData[18].superscripts.some((c) => c == chapter);
};

/**
 * Determines if a BibleRef is valid (that is, chapter and verse exist in given bible book)
 * @param {BibleRef} ref - The reference to be validated.
 */
export const isValidBibleRef = (ref: BibleRef): boolean => {
  if (ref == null) return false;
  const verse = ref % 1000;
  const chapter = ((ref % 1000000) - verse) / 1000;
  const book = (ref - chapter * 1000 - verse) / 1000000;

  return (
    book > 0 &&
    book <= bibleData.length &&
    chapter > 0 &&
    chapter <= bibleData[book - 1].num_chapters &&
    (verse > 0 ||
      (verse == 0 && book == 19 && isSuperscriptPsalmChapter(chapter)))
  );
};

export const isRefBookStart = (ref: BibleRef): boolean => {
  if (!isValidBibleRef) return false;
  const verse = ref % 1000;
  const chapter = ((ref % 1000000) - verse) / 1000;
  return chapter == 1 && verse == 1;
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
  if (
    book_index >= 0 &&
    chapter > 0 &&
    chapter <= bibleData[book_index].num_chapters
  ) {
    let verse_num: number;
    if (verse == "start") {
      verse_num =
        book == "Psalms" && isSuperscriptPsalmChapter(chapter) ? 0 : 1;
    } else if (verse == "end") {
      verse_num = bibleData[book_index].num_verses[chapter - 1];
    } else {
      verse_num = verse;
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

  const verse = ref % 1000;
  const chapter = ((ref % 1000000) - verse) / 1000;
  const book = (ref - chapter * 1000 - verse) / 1000000;
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

// /**
//  * Create a valid BibleRef at the specified offset away (will return valid BibleRef even if only partial offset away).
//  * @throws Will throw if `ref` is not valid.
//  * @param {BibleRef} ref - The original BibleRef from which offset is measured.
//  * @param {number} offset - The offset magnitude.
//  * @param {"book"|"chapter"|"verse"} offset_unit - The units for the offset.
//  * @param {"start"|"end"} mode - Determines what mode of offset is taken.
//  *  **start** – unit immediately below {@link offset_unit} is given lowest possible value (e.g. 1 John 2:28 + 1 chapter => 1 John 3:1, or + 1 book => 2 John 1).
//  *  **end** – unit immediately below  {@link offset_unit} is given highest possible value (e.g. 1 John 2:28 + 1 chapter => 1 John 3:24, or + 1 book => 2 John 13).
//  */
// export const createOffsetBibleRef = (
//   ref: BibleRef,
//   offset: number,
//   offset_unit: "book" | "chapter" | "verse",
//   mode: "start" | "end",
// ): BibleRef => {
//   if (!isValidBibleRef(ref))
//     GrapheError(`Invalid ref (${ref}) passed to \`createOffsetBibleRef\``);
//   if (offset == 0) return ref;

//   const verse = ref % 1000;
//   const chapter = ((ref % 1000000) - verse) / 1000;
//   const book = (ref - chapter * 1000 - verse) / 1000000;

//   const step = -offset / Math.abs(offset);
//   switch (offset_unit) {
//     case "book":
//       for (let i = offset; i != step; i += step) {
//         const index = book + i - 1;
//         if (index < 0 && index >= bibleData.length) continue;
//         const new_chapter = mode == "start" ? 1 : bibleData[index].num_chapters;
//         return createBibleRef(bibleData[index].name, new_chapter, mode);
//       }
//     case "chapter":
//       let chapters_left = offset;
//       let new_book = book;
//       let new_chapter = chapter;

//       // If moving backwards
//       if (offset < 0) {
//         // position at chapter 1 (or as far back as allowed)
//         let move_amount = Math.min(Math.abs(chapters_left), new_chapter - 1);
//         new_chapter = new_chapter - move_amount;
//         chapters_left += move_amount;

//         while (chapters_left < 0) {
//           if (new_book == 1) break;
//           new_book -= 1;
//           move_amount = Math.min(
//             Math.abs(chapters_left),
//             bibleData[new_book - 1].num_chapters,
//           );
//           new_chapter = bibleData[new_book - 1].num_chapters - move_amount + 1;
//           chapters_left += move_amount;
//         }
//       }

//       // If moving forwards
//       if (offset > 0) {
//         // position at final chapter of book (or as far forward as allowed)
//         let move_amount = Math.min(
//           chapters_left,
//           bibleData[new_book - 1].num_chapters - new_chapter,
//         );
//         new_chapter = new_chapter + move_amount;
//         chapters_left -= move_amount;

//         while (chapters_left > 0) {
//           if (new_book == bibleData.length) break;
//           new_book += 1;
//           move_amount = Math.min(
//             chapters_left,
//             bibleData[new_book - 1].num_chapters,
//           );
//           new_chapter = move_amount;
//           chapters_left -= move_amount;
//         }
//       }

//       return createBibleRef(bibleData[new_book - 1].name, new_chapter, mode);
//     case "verse":
//       // TODO
//       break;
//   }

//   return 40001001;
// };
