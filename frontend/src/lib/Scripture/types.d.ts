import { bibleData } from "@/lib/Scripture/data";

export type BibleBook = (typeof bibleData)[number]["name"];
export type BibleBookAbbreviation = (typeof bibleData)[number]["abbreviation"];
export type BibleTestament = (typeof bibleData)[number]["testament"];
export type BibleVersion = (typeof bibleData)[number]["version"][number];

export type BibleRef = number;
export type BibleRange = {
  version: BibleVersion;
  start: BibleRef;
  end: BibleRef;
};
