// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import {Create as $Create} from "@wailsio/runtime";

export class SettingsValues {
    "general": SettingsValues_General;
    "appearence": SettingsValues_Appearence;
    "shortcuts": SettingsValues_Shortcuts;
    "version": SettingsValues_Version;
    "formatting": SettingsValues_Formatting;
    "search": SettingsValues_Search;
    "instantDetails": SettingsValues_InstantDetails;

    /** Creates a new SettingsValues instance. */
    constructor($$source: Partial<SettingsValues> = {}) {
        if (!("general" in $$source)) {
            this["general"] = (new SettingsValues_General());
        }
        if (!("appearence" in $$source)) {
            this["appearence"] = (new SettingsValues_Appearence());
        }
        if (!("shortcuts" in $$source)) {
            this["shortcuts"] = (new SettingsValues_Shortcuts());
        }
        if (!("version" in $$source)) {
            this["version"] = (new SettingsValues_Version());
        }
        if (!("formatting" in $$source)) {
            this["formatting"] = (new SettingsValues_Formatting());
        }
        if (!("search" in $$source)) {
            this["search"] = (new SettingsValues_Search());
        }
        if (!("instantDetails" in $$source)) {
            this["instantDetails"] = (new SettingsValues_InstantDetails());
        }

        Object.assign(this, $$source);
    }

    /**
     * Creates a new SettingsValues instance from a string or object.
     */
    static createFrom($$source: any = {}): SettingsValues {
        const $$createField0_0 = $$createType0;
        const $$createField1_0 = $$createType1;
        const $$createField2_0 = $$createType2;
        const $$createField3_0 = $$createType3;
        const $$createField4_0 = $$createType4;
        const $$createField5_0 = $$createType5;
        const $$createField6_0 = $$createType6;
        let $$parsedSource = typeof $$source === 'string' ? JSON.parse($$source) : $$source;
        if ("general" in $$parsedSource) {
            $$parsedSource["general"] = $$createField0_0($$parsedSource["general"]);
        }
        if ("appearence" in $$parsedSource) {
            $$parsedSource["appearence"] = $$createField1_0($$parsedSource["appearence"]);
        }
        if ("shortcuts" in $$parsedSource) {
            $$parsedSource["shortcuts"] = $$createField2_0($$parsedSource["shortcuts"]);
        }
        if ("version" in $$parsedSource) {
            $$parsedSource["version"] = $$createField3_0($$parsedSource["version"]);
        }
        if ("formatting" in $$parsedSource) {
            $$parsedSource["formatting"] = $$createField4_0($$parsedSource["formatting"]);
        }
        if ("search" in $$parsedSource) {
            $$parsedSource["search"] = $$createField5_0($$parsedSource["search"]);
        }
        if ("instantDetails" in $$parsedSource) {
            $$parsedSource["instantDetails"] = $$createField6_0($$parsedSource["instantDetails"]);
        }
        return new SettingsValues($$parsedSource as Partial<SettingsValues>);
    }
}

export class SettingsValues_Appearence {
    "theme": string;
    "font": SettingsValues_Appearence_Font;
    "zoom": number;

    /** Creates a new SettingsValues_Appearence instance. */
    constructor($$source: Partial<SettingsValues_Appearence> = {}) {
        if (!("theme" in $$source)) {
            this["theme"] = "";
        }
        if (!("font" in $$source)) {
            this["font"] = (new SettingsValues_Appearence_Font());
        }
        if (!("zoom" in $$source)) {
            this["zoom"] = 0;
        }

        Object.assign(this, $$source);
    }

    /**
     * Creates a new SettingsValues_Appearence instance from a string or object.
     */
    static createFrom($$source: any = {}): SettingsValues_Appearence {
        const $$createField1_0 = $$createType7;
        let $$parsedSource = typeof $$source === 'string' ? JSON.parse($$source) : $$source;
        if ("font" in $$parsedSource) {
            $$parsedSource["font"] = $$createField1_0($$parsedSource["font"]);
        }
        return new SettingsValues_Appearence($$parsedSource as Partial<SettingsValues_Appearence>);
    }
}

export class SettingsValues_Appearence_Font {
    "system": string;
    "greek": string;
    "hebrew": string;
    "english": string;

    /** Creates a new SettingsValues_Appearence_Font instance. */
    constructor($$source: Partial<SettingsValues_Appearence_Font> = {}) {
        if (!("system" in $$source)) {
            this["system"] = "";
        }
        if (!("greek" in $$source)) {
            this["greek"] = "";
        }
        if (!("hebrew" in $$source)) {
            this["hebrew"] = "";
        }
        if (!("english" in $$source)) {
            this["english"] = "";
        }

        Object.assign(this, $$source);
    }

    /**
     * Creates a new SettingsValues_Appearence_Font instance from a string or object.
     */
    static createFrom($$source: any = {}): SettingsValues_Appearence_Font {
        let $$parsedSource = typeof $$source === 'string' ? JSON.parse($$source) : $$source;
        return new SettingsValues_Appearence_Font($$parsedSource as Partial<SettingsValues_Appearence_Font>);
    }
}

export class SettingsValues_Formatting {

    /** Creates a new SettingsValues_Formatting instance. */
    constructor($$source: Partial<SettingsValues_Formatting> = {}) {

        Object.assign(this, $$source);
    }

    /**
     * Creates a new SettingsValues_Formatting instance from a string or object.
     */
    static createFrom($$source: any = {}): SettingsValues_Formatting {
        let $$parsedSource = typeof $$source === 'string' ? JSON.parse($$source) : $$source;
        return new SettingsValues_Formatting($$parsedSource as Partial<SettingsValues_Formatting>);
    }
}

export class SettingsValues_General {

    /** Creates a new SettingsValues_General instance. */
    constructor($$source: Partial<SettingsValues_General> = {}) {

        Object.assign(this, $$source);
    }

    /**
     * Creates a new SettingsValues_General instance from a string or object.
     */
    static createFrom($$source: any = {}): SettingsValues_General {
        let $$parsedSource = typeof $$source === 'string' ? JSON.parse($$source) : $$source;
        return new SettingsValues_General($$parsedSource as Partial<SettingsValues_General>);
    }
}

export class SettingsValues_InstantDetails {

    /** Creates a new SettingsValues_InstantDetails instance. */
    constructor($$source: Partial<SettingsValues_InstantDetails> = {}) {

        Object.assign(this, $$source);
    }

    /**
     * Creates a new SettingsValues_InstantDetails instance from a string or object.
     */
    static createFrom($$source: any = {}): SettingsValues_InstantDetails {
        let $$parsedSource = typeof $$source === 'string' ? JSON.parse($$source) : $$source;
        return new SettingsValues_InstantDetails($$parsedSource as Partial<SettingsValues_InstantDetails>);
    }
}

export class SettingsValues_Search {

    /** Creates a new SettingsValues_Search instance. */
    constructor($$source: Partial<SettingsValues_Search> = {}) {

        Object.assign(this, $$source);
    }

    /**
     * Creates a new SettingsValues_Search instance from a string or object.
     */
    static createFrom($$source: any = {}): SettingsValues_Search {
        let $$parsedSource = typeof $$source === 'string' ? JSON.parse($$source) : $$source;
        return new SettingsValues_Search($$parsedSource as Partial<SettingsValues_Search>);
    }
}

export class SettingsValues_Shortcuts {
    "aboutGraphe": string;
    "checkForUpdates": string;
    "openSettings": string;
    "openWorkspace": string;
    "openDataDirectory": string;
    "openLogDirectory": string;
    "purgeLogs": string;
    "passageMode": string;
    "searchMode": string;
    "openAnalytics": string;
    "openFunctions": string;
    "chooseVersion": string;
    "chooseText": string;
    "zoomIn": string;
    "zoomOut": string;
    "zoomReset": string;
    "changeTheme": string;

    /** Creates a new SettingsValues_Shortcuts instance. */
    constructor($$source: Partial<SettingsValues_Shortcuts> = {}) {
        if (!("aboutGraphe" in $$source)) {
            this["aboutGraphe"] = "";
        }
        if (!("checkForUpdates" in $$source)) {
            this["checkForUpdates"] = "";
        }
        if (!("openSettings" in $$source)) {
            this["openSettings"] = "";
        }
        if (!("openWorkspace" in $$source)) {
            this["openWorkspace"] = "";
        }
        if (!("openDataDirectory" in $$source)) {
            this["openDataDirectory"] = "";
        }
        if (!("openLogDirectory" in $$source)) {
            this["openLogDirectory"] = "";
        }
        if (!("purgeLogs" in $$source)) {
            this["purgeLogs"] = "";
        }
        if (!("passageMode" in $$source)) {
            this["passageMode"] = "";
        }
        if (!("searchMode" in $$source)) {
            this["searchMode"] = "";
        }
        if (!("openAnalytics" in $$source)) {
            this["openAnalytics"] = "";
        }
        if (!("openFunctions" in $$source)) {
            this["openFunctions"] = "";
        }
        if (!("chooseVersion" in $$source)) {
            this["chooseVersion"] = "";
        }
        if (!("chooseText" in $$source)) {
            this["chooseText"] = "";
        }
        if (!("zoomIn" in $$source)) {
            this["zoomIn"] = "";
        }
        if (!("zoomOut" in $$source)) {
            this["zoomOut"] = "";
        }
        if (!("zoomReset" in $$source)) {
            this["zoomReset"] = "";
        }
        if (!("changeTheme" in $$source)) {
            this["changeTheme"] = "";
        }

        Object.assign(this, $$source);
    }

    /**
     * Creates a new SettingsValues_Shortcuts instance from a string or object.
     */
    static createFrom($$source: any = {}): SettingsValues_Shortcuts {
        let $$parsedSource = typeof $$source === 'string' ? JSON.parse($$source) : $$source;
        return new SettingsValues_Shortcuts($$parsedSource as Partial<SettingsValues_Shortcuts>);
    }
}

export class SettingsValues_Version {

    /** Creates a new SettingsValues_Version instance. */
    constructor($$source: Partial<SettingsValues_Version> = {}) {

        Object.assign(this, $$source);
    }

    /**
     * Creates a new SettingsValues_Version instance from a string or object.
     */
    static createFrom($$source: any = {}): SettingsValues_Version {
        let $$parsedSource = typeof $$source === 'string' ? JSON.parse($$source) : $$source;
        return new SettingsValues_Version($$parsedSource as Partial<SettingsValues_Version>);
    }
}

// Private type creation functions
const $$createType0 = SettingsValues_General.createFrom;
const $$createType1 = SettingsValues_Appearence.createFrom;
const $$createType2 = SettingsValues_Shortcuts.createFrom;
const $$createType3 = SettingsValues_Version.createFrom;
const $$createType4 = SettingsValues_Formatting.createFrom;
const $$createType5 = SettingsValues_Search.createFrom;
const $$createType6 = SettingsValues_InstantDetails.createFrom;
const $$createType7 = SettingsValues_Appearence_Font.createFrom;
