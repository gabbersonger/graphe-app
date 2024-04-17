import { bibleData, versionData } from "@/lib/Scripture/data";

export type BibleBook = (typeof bibleData)[number]["name"];
export type BibleBookAbbreviation = (typeof bibleData)[number]["abbreviation"];
export type BibleTestament = (typeof bibleData)[number]["testament"];
export type BibleVersion = keyof typeof versionData;

export type BibleRef = number;
