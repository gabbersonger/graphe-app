import { versionData } from "@/lib/Scripture/data";
import type { BibleVersion } from "@/lib/Scripture/types";
import { GrapheLog } from "../utils";
import type { ScriptureRange } from "!/graphe/internal/scripture";
import { createRef } from "@/lib/Scripture/ref";

/**
 * Determine if the given string is a valid version name
 * @param {string} version - The string to check
 */
export function isValidVersion(version: string): boolean {
  if (version == null) return false;
  return version in versionData;
}

/**
 * Get the index of a book in a version (-1 if not in version)
 * @throws Will throw if `version` is not valid
 * @param {BibleVersion} version - The version to check against
 * @param {number} book - The book to check for the index of
 */
export function getVersionBookIndex(version: BibleVersion, book: number) {
  if (!isValidVersion(version))
    GrapheLog(
      "error",
      `Invalid version (${version}) passed to \`getVersionBookIndex\``,
    );
  return versionData[version].books.findIndex((b) => b.book_number == book);
}

/**
 * Get a range that encapsulates the full data of a version (start to end book)
 * @throws Will throw if `version` is not valid
 * @param {BibleVersion} version - The version to get a range for
 */
export function createVersionRange(version: BibleVersion): ScriptureRange {
  if (!isValidVersion(version))
    GrapheLog(
      "error",
      `Invalid version (${version}) passed to \`createVersionRange\``,
    );

  const startBookData = versionData[version].books[0];
  const endBookData =
    versionData[version].books[versionData[version].books.length - 1];

  return {
    version: version,
    start: createRef(
      version,
      startBookData.book_number,
      "prologue" in startBookData ? 0 : 1,
      "start",
    ),
    end: createRef(
      version,
      endBookData.book_number,
      endBookData.num_chapters,
      "end",
    ),
  };
}
