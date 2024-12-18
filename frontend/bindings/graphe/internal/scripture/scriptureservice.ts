// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import {Call as $Call, Create as $Create} from "@wailsio/runtime";

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import * as $models from "./models.js";

export function BCV(r: $models.ScriptureRef): Promise<[number, number, number]> & { cancel(): void } {
    let $resultPromise = $Call.ByID(616843647, r) as any;
    return $resultPromise;
}

export function CreateFirstValidRef(version: $models.ScriptureVersion, book: number): Promise<$models.ScriptureRef> & { cancel(): void } {
    let $resultPromise = $Call.ByID(2368684349, version, book) as any;
    return $resultPromise;
}

export function CreateLastValidRef(version: $models.ScriptureVersion, book: number): Promise<$models.ScriptureRef> & { cancel(): void } {
    let $resultPromise = $Call.ByID(1099231067, version, book) as any;
    return $resultPromise;
}

export function CreateRef(book: number, chapter: number, verse: number): Promise<$models.ScriptureRef> & { cancel(): void } {
    let $resultPromise = $Call.ByID(3842278653, book, chapter, verse) as any;
    return $resultPromise;
}

export function DivideIntoBookRanges(rang: $models.ScriptureRange): Promise<$models.ScriptureRange[]> & { cancel(): void } {
    let $resultPromise = $Call.ByID(2064554944, rang) as any;
    let $typingPromise = $resultPromise.then(($result) => {
        return $$createType1($result);
    }) as any;
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

export function GetBibleData(): Promise<$models.BookData[]> & { cancel(): void } {
    let $resultPromise = $Call.ByID(3949367714) as any;
    let $typingPromise = $resultPromise.then(($result) => {
        return $$createType3($result);
    }) as any;
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

export function GetRefBook(r: $models.ScriptureRef): Promise<number> & { cancel(): void } {
    let $resultPromise = $Call.ByID(2286550768, r) as any;
    return $resultPromise;
}

export function GetRefChapter(r: $models.ScriptureRef): Promise<number> & { cancel(): void } {
    let $resultPromise = $Call.ByID(3734787642, r) as any;
    return $resultPromise;
}

export function GetRefVerse(r: $models.ScriptureRef): Promise<number> & { cancel(): void } {
    let $resultPromise = $Call.ByID(1574122346, r) as any;
    return $resultPromise;
}

export function GetVersionData(version: string): Promise<$models.VersionData> & { cancel(): void } {
    let $resultPromise = $Call.ByID(852716388, version) as any;
    let $typingPromise = $resultPromise.then(($result) => {
        return $$createType4($result);
    }) as any;
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

export function GetVersionLanguage(version: string): Promise<string> & { cancel(): void } {
    let $resultPromise = $Call.ByID(868042726, version) as any;
    return $resultPromise;
}

export function GetVersionLanguageHeadings(version: string): Promise<string> & { cancel(): void } {
    let $resultPromise = $Call.ByID(2122438513, version) as any;
    return $resultPromise;
}

export function GetVersionRange(version: $models.ScriptureVersion): Promise<$models.ScriptureRange> & { cancel(): void } {
    let $resultPromise = $Call.ByID(2777961761, version) as any;
    let $typingPromise = $resultPromise.then(($result) => {
        return $$createType0($result);
    }) as any;
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

export function GetVersionsBasicData(): Promise<$models.ScriptureVersionBasicInfo[]> & { cancel(): void } {
    let $resultPromise = $Call.ByID(2352453813) as any;
    let $typingPromise = $resultPromise.then(($result) => {
        return $$createType6($result);
    }) as any;
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

export function IsRangeValid(rang: $models.ScriptureRange): Promise<boolean> & { cancel(): void } {
    let $resultPromise = $Call.ByID(2090406317, rang) as any;
    return $resultPromise;
}

export function IsRefBookStart(ref: $models.ScriptureRef, version: $models.ScriptureVersion): Promise<boolean> & { cancel(): void } {
    let $resultPromise = $Call.ByID(2248902072, ref, version) as any;
    return $resultPromise;
}

export function IsRefValid(ref: $models.ScriptureRef, version: $models.ScriptureVersion): Promise<boolean> & { cancel(): void } {
    let $resultPromise = $Call.ByID(1912637581, ref, version) as any;
    return $resultPromise;
}

export function IsVersionValid(v: $models.ScriptureVersion): Promise<boolean> & { cancel(): void } {
    let $resultPromise = $Call.ByID(2537419090, v) as any;
    return $resultPromise;
}

export function RangeContains(rang: $models.ScriptureRange, ref: $models.ScriptureRef): Promise<boolean> & { cancel(): void } {
    let $resultPromise = $Call.ByID(2816720236, rang, ref) as any;
    return $resultPromise;
}

export function RefToString(ref: $models.ScriptureRef, version: $models.ScriptureVersion, format: $models.ScriptureRefStringType): Promise<string> & { cancel(): void } {
    let $resultPromise = $Call.ByID(772653793, ref, version, format) as any;
    return $resultPromise;
}

// Private type creation functions
const $$createType0 = $models.ScriptureRange.createFrom;
const $$createType1 = $Create.Array($$createType0);
const $$createType2 = $models.BookData.createFrom;
const $$createType3 = $Create.Array($$createType2);
const $$createType4 = $models.VersionData.createFrom;
const $$createType5 = $models.ScriptureVersionBasicInfo.createFrom;
const $$createType6 = $Create.Array($$createType5);
