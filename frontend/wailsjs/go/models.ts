export namespace app {
	
	export class EnvironmentInfo {
	    arch: string;
	    buildType: string;
	    platform: string;
	    version: string;
	    homeDirectory: string;
	    dataDirectory: string;
	    logDirectory: string;
	
	    static createFrom(source: any = {}) {
	        return new EnvironmentInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.arch = source["arch"];
	        this.buildType = source["buildType"];
	        this.platform = source["platform"];
	        this.version = source["version"];
	        this.homeDirectory = source["homeDirectory"];
	        this.dataDirectory = source["dataDirectory"];
	        this.logDirectory = source["logDirectory"];
	    }
	}

}

export namespace data {
	
	export class ScriptureVerseDetail {
	    type: number;
	    data?: string;
	
	    static createFrom(source: any = {}) {
	        return new ScriptureVerseDetail(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.type = source["type"];
	        this.data = source["data"];
	    }
	}
	export class ScriptureWordDetail {
	    type: number;
	    position?: boolean;
	    data?: string;
	
	    static createFrom(source: any = {}) {
	        return new ScriptureWordDetail(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.type = source["type"];
	        this.position = source["position"];
	        this.data = source["data"];
	    }
	}
	export class ScriptureWord {
	    word_num: number;
	    text: string;
	    pre: string;
	    post: string;
	    details?: ScriptureWordDetail[];
	    no_instant_details?: boolean;
	
	    static createFrom(source: any = {}) {
	        return new ScriptureWord(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.word_num = source["word_num"];
	        this.text = source["text"];
	        this.pre = source["pre"];
	        this.post = source["post"];
	        this.details = this.convertValues(source["details"], ScriptureWordDetail);
	        this.no_instant_details = source["no_instant_details"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ScriptureVerse {
	    ref: number;
	    words: ScriptureWord[];
	    details?: ScriptureVerseDetail[];
	    continuation?: boolean;
	
	    static createFrom(source: any = {}) {
	        return new ScriptureVerse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ref = source["ref"];
	        this.words = this.convertValues(source["words"], ScriptureWord);
	        this.details = this.convertValues(source["details"], ScriptureVerseDetail);
	        this.continuation = source["continuation"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ScriptureBlock {
	    range: scripture.ScriptureRange;
	    verses: ScriptureVerse[];
	
	    static createFrom(source: any = {}) {
	        return new ScriptureBlock(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.range = this.convertValues(source["range"], scripture.ScriptureRange);
	        this.verses = this.convertValues(source["verses"], ScriptureVerse);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ScriptureSection {
	    range: scripture.ScriptureRange;
	    blocks: ScriptureBlock[];
	
	    static createFrom(source: any = {}) {
	        return new ScriptureSection(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.range = this.convertValues(source["range"], scripture.ScriptureRange);
	        this.blocks = this.convertValues(source["blocks"], ScriptureBlock);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	
	
	export class ScriptureWordDataField {
	    name: string;
	    data: any;
	
	    static createFrom(source: any = {}) {
	        return new ScriptureWordDataField(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.data = source["data"];
	    }
	}
	export class ScriptureWordData {
	    version: string;
	    ref: number;
	    word_number: number;
	    text: string;
	    fields: ScriptureWordDataField[];
	
	    static createFrom(source: any = {}) {
	        return new ScriptureWordData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.version = source["version"];
	        this.ref = source["ref"];
	        this.word_number = source["word_number"];
	        this.text = source["text"];
	        this.fields = this.convertValues(source["fields"], ScriptureWordDataField);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	

}

export namespace menu {
	
	export class Menu {
	
	
	    static createFrom(source: any = {}) {
	        return new Menu(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}

}

export namespace scripture {
	
	export class ScriptureRange {
	    version: string;
	    start: number;
	    end: number;
	
	    static createFrom(source: any = {}) {
	        return new ScriptureRange(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.version = source["version"];
	        this.start = source["start"];
	        this.end = source["end"];
	    }
	}

}

export namespace settings {
	
	export class SettingsValues_InstantDetails {
	
	
	    static createFrom(source: any = {}) {
	        return new SettingsValues_InstantDetails(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}
	export class SettingsValues_Search {
	
	
	    static createFrom(source: any = {}) {
	        return new SettingsValues_Search(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}
	export class SettingsValues_Formatting {
	
	
	    static createFrom(source: any = {}) {
	        return new SettingsValues_Formatting(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}
	export class SettingsValues_Version {
	
	
	    static createFrom(source: any = {}) {
	        return new SettingsValues_Version(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}
	export class SettingsValues_Shortcuts {
	    aboutGraphe?: string;
	    checkForUpdates?: string;
	    openSettings?: string;
	    openWorkspace?: string;
	    openDataDirectory?: string;
	    openLogDirectory?: string;
	    purgeLogs?: string;
	    passageMode?: string;
	    searchMode?: string;
	    openAnalytics?: string;
	    openFunctions?: string;
	    chooseVersion?: string;
	    chooseText?: string;
	    zoomIn?: string;
	    zoomOut?: string;
	    zoomReset?: string;
	    changeTheme?: string;
	
	    static createFrom(source: any = {}) {
	        return new SettingsValues_Shortcuts(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.aboutGraphe = source["aboutGraphe"];
	        this.checkForUpdates = source["checkForUpdates"];
	        this.openSettings = source["openSettings"];
	        this.openWorkspace = source["openWorkspace"];
	        this.openDataDirectory = source["openDataDirectory"];
	        this.openLogDirectory = source["openLogDirectory"];
	        this.purgeLogs = source["purgeLogs"];
	        this.passageMode = source["passageMode"];
	        this.searchMode = source["searchMode"];
	        this.openAnalytics = source["openAnalytics"];
	        this.openFunctions = source["openFunctions"];
	        this.chooseVersion = source["chooseVersion"];
	        this.chooseText = source["chooseText"];
	        this.zoomIn = source["zoomIn"];
	        this.zoomOut = source["zoomOut"];
	        this.zoomReset = source["zoomReset"];
	        this.changeTheme = source["changeTheme"];
	    }
	}
	export class SettingsValues_Appearence_Font {
	    system?: string;
	    greek?: string;
	    hebrew?: string;
	    english?: string;
	
	    static createFrom(source: any = {}) {
	        return new SettingsValues_Appearence_Font(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.system = source["system"];
	        this.greek = source["greek"];
	        this.hebrew = source["hebrew"];
	        this.english = source["english"];
	    }
	}
	export class SettingsValues_Appearence {
	    theme?: string;
	    font: SettingsValues_Appearence_Font;
	    zoom?: number;
	
	    static createFrom(source: any = {}) {
	        return new SettingsValues_Appearence(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.theme = source["theme"];
	        this.font = this.convertValues(source["font"], SettingsValues_Appearence_Font);
	        this.zoom = source["zoom"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class SettingsValues_General {
	
	
	    static createFrom(source: any = {}) {
	        return new SettingsValues_General(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}
	export class SettingsValues {
	    // Go type: SettingsValues_General
	    general: any;
	    appearence: SettingsValues_Appearence;
	    shortcuts: SettingsValues_Shortcuts;
	    // Go type: SettingsValues_Version
	    version: any;
	    // Go type: SettingsValues_Formatting
	    formatting: any;
	    // Go type: SettingsValues_Search
	    search: any;
	    // Go type: SettingsValues_InstantDetails
	    instantDetails: any;
	
	    static createFrom(source: any = {}) {
	        return new SettingsValues(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.general = this.convertValues(source["general"], null);
	        this.appearence = this.convertValues(source["appearence"], SettingsValues_Appearence);
	        this.shortcuts = this.convertValues(source["shortcuts"], SettingsValues_Shortcuts);
	        this.version = this.convertValues(source["version"], null);
	        this.formatting = this.convertValues(source["formatting"], null);
	        this.search = this.convertValues(source["search"], null);
	        this.instantDetails = this.convertValues(source["instantDetails"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	

}

