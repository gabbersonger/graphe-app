import { bibleData } from "./bibleVersionData/_books";
import { lxx } from "./bibleVersionData/lxx";
import { gnt } from "./bibleVersionData/gnt";
import { hot } from "./bibleVersionData/hot";
import { esv } from "./bibleVersionData/esv";

const fs = require("fs");
const crypto = require("crypto");

const VERSION_DATA = {
  lxx: lxx,
  gnt: gnt,
  hot: hot,
  esv: esv,
};

const HASH_FOLDER = "/bibleVersionData";
const HASH_FILE = "/data.lock";
const FRONTEND_FILE = "/../../frontend/src/lib/Scripture/data.ts";
const BACKEND_FILE = "/../../internal/app/scripture_data.go";

function hashFile(filename) {
  var f = fs.readFileSync(__dirname + filename);
  var md5 = crypto.createHash("md5");
  md5.update(f, "utf-8");
  return md5.digest("hex");
}

function haveFilesChanged() {
  let hashes = [];
  fs.readdirSync(__dirname + HASH_FOLDER).forEach((file) => {
    hashes.push(hashFile(HASH_FOLDER + "/" + file));
  });
  const combined_hash = hashes.join("");
  if (combined_hash == fs.readFileSync(__dirname + HASH_FILE)) return false;
  fs.writeFileSync(__dirname + HASH_FILE, combined_hash);
  return true;
}

function createFrontendFile() {
  let data = "// DO NOT EDIT - THIS FILE IS AUTOGENERATED\n\n";

  // Bible Data
  data += "export const bibleData = [\n";
  for (let i = 0; i < bibleData.length; i++) {
    const bookData = bibleData[i];
    data += `  {\n`;
    data += `    name: "${bookData.name}",\n`;
    data += `    short: "${bookData.short}",\n`;
    data += `    abbreviations: [\n`;
    for (let j = 0; j < bookData.abbreviations.length; j++) {
      data += `      "${bookData.abbreviations[j]}",\n`;
    }
    data += `    ],\n`;
    data += `  },\n`;
  }
  data += "] as const;\n\n";

  // Version Data
  data += "export const versionData = {\n";
  for (let [name, version] of Object.entries(VERSION_DATA)) {
    data += `  ${name}: {\n`;
    data += `    full_name: "${version.full_name}",\n`;
    data += `    language: "${version.language}",\n`;
    data += `    books: [\n`;
    for (let i = 0; i < version.books.length; i++) {
      const book = version.books[i];
      data += `      {\n`;
      data += `        book_number: ${book.book_number},\n`;
      data += `        display_name: "${book.display_name}",\n`;
      data += `        num_chapters: ${book.num_chapters},\n`;
      data += `        num_verses: [${book.num_verses.join(", ")}],\n`;
      if ("superscripts" in book) {
        data += `        superscripts: [${book.superscripts.join(", ")}],\n`;
      }
      if ("prologue" in book) {
        data += `        prologue: ${book.prologue},\n`;
      }
      data += `        parallels: [\n`;
      for (let j = 0; j < book.parallels.length; j++) {
        const parallel = book.parallels[j];
        data += `          {\n`;
        data += `            start: ${parallel.start},\n`;
        data += `            end: ${parallel.end},\n`;
        data += `          },\n`;
      }
      data += `        ],\n`;
      data += `        missing_sections: [\n`;
      for (let j = 0; j < book.missing_sections.length; j++) {
        const section = book.missing_sections[j];
        data += `          {\n`;
        data += `            start: ${section.start},\n`;
        data += `            end: ${section.end},\n`;
        data += `          },\n`;
      }
      data += `        ],\n`;
      data += `      },\n`;
    }
    data += `    ],\n`;
    data += `  },\n`;
  }
  data += "} as const;\n";

  // Write the data
  fs.writeFileSync(__dirname + FRONTEND_FILE, data);
}

function createBackendFile() {
  let data = "package app\n";

  // Types
  data += "// DO NOT EDIT - THIS FILE IS AUTOGENERATED\n\n";
  data += "type BookData struct {\n";
  data += "  name          string\n";
  data += "  short         string\n";
  data += "  abbreviations []string\n";
  data += "}\n\n";
  data += "type VersionBibleRange struct {\n";
  data += "  start int\n";
  data += "  end   int\n";
  data += "}\n\n";
  data += "type VersionBookData struct {\n";
  data += "  book_number      int\n";
  data += "  display_name     string\n";
  data += "  num_chapters     int\n";
  data += "  num_verses       []int\n";
  data += "  superscripts     []int\n";
  data += "  prologue         int\n";
  data += "  parallels        []VersionBibleRange\n";
  data += "  missing_sections []VersionBibleRange\n";
  data += "}\n\n";
  data += "type VersionData struct {\n";
  data += "  name      string\n";
  data += "  full_name string\n";
  data += "  language  string\n";
  data += "  books     []VersionBookData\n";
  data += "}\n\n";

  // Bible Data
  data += "var bibleData = [...]BookData{\n";
  for (let i = 0; i < bibleData.length; i++) {
    const bookData = bibleData[i];
    data += `  {\n`;
    data += `    name: "${bookData.name}",\n`;
    data += `    short: "${bookData.short}",\n`;
    data += `    abbreviations: []string{\n`;
    for (let j = 0; j < bookData.abbreviations.length; j++) {
      data += `      "${bookData.abbreviations[j]}",\n`;
    }
    data += `    },\n`;
    data += `  },\n`;
  }
  data += "}\n\n";

  // Version Data
  data += "var versionData = [...]VersionData{\n";
  for (let [name, version] of Object.entries(VERSION_DATA)) {
    data += `  {\n`;
    data += `    name: "${name}",\n`;
    data += `    full_name: "${version.full_name}",\n`;
    data += `    language: "${version.language}",\n`;
    data += `    books: []VersionBookData{\n`;
    for (let i = 0; i < version.books.length; i++) {
      const book = version.books[i];
      data += `      {\n`;
      data += `        book_number: ${book.book_number},\n`;
      data += `        display_name: "${book.display_name}",\n`;
      data += `        num_chapters: ${book.num_chapters},\n`;
      data += `        num_verses: []int{${book.num_verses.join(",")}},\n`;
      if ("superscripts" in book) {
        data += `        superscripts: []int{${book.superscripts.join(", ")}},\n`;
      }
      data += `        prologue: ${book.prologue ? book.prologue : 0},\n`;
      data += `        parallels: []VersionBibleRange{\n`;
      for (let j = 0; j < book.parallels.length; j++) {
        const parallel = book.parallels[j];
        data += `          {\n`;
        data += `            start: ${parallel.start},\n`;
        data += `            end: ${parallel.end},\n`;
        data += `          },\n`;
      }
      data += `        },\n`;
      data += `        missing_sections: []VersionBibleRange{\n`;
      for (let j = 0; j < book.missing_sections.length; j++) {
        const section = book.missing_sections[j];
        data += `          {\n`;
        data += `            start: ${section.start},\n`;
        data += `            end: ${section.end},\n`;
        data += `          },\n`;
      }
      data += `        },\n`;
      data += `      },\n`;
    }
    data += `    },\n`;
    data += `  },\n`;
  }
  data += "}\n";

  // Write the data
  fs.writeFileSync(__dirname + BACKEND_FILE, data);
}

export function createBibleData() {
  if (haveFilesChanged()) {
    createFrontendFile();
    createBackendFile();
  }
}
