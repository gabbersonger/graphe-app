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
	export class ScriptureBlock {
	    range: ScriptureRange;
	    verses: ScriptureVerse[];
	
	    static createFrom(source: any = {}) {
	        return new ScriptureBlock(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.range = this.convertValues(source["range"], ScriptureRange);
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
	    range: ScriptureRange;
	    blocks: ScriptureBlock[];
	
	    static createFrom(source: any = {}) {
	        return new ScriptureSection(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.range = this.convertValues(source["range"], ScriptureRange);
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
	
	
	
	export class ScriptureWordData_Dictionary {
	    form: string;
	    gloss: string;
	    strong: string;
	    grammar: string;
	    count: number;
	
	    static createFrom(source: any = {}) {
	        return new ScriptureWordData_Dictionary(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.form = source["form"];
	        this.gloss = source["gloss"];
	        this.strong = source["strong"];
	        this.grammar = source["grammar"];
	        this.count = source["count"];
	    }
	}
	export class ScriptureWordData {
	    ref: number;
	    word_number: number;
	    text: string;
	    translit: string;
	    english: string;
	    dictionary: ScriptureWordData_Dictionary[];
	    inflected_count: number;
	
	    static createFrom(source: any = {}) {
	        return new ScriptureWordData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ref = source["ref"];
	        this.word_number = source["word_number"];
	        this.text = source["text"];
	        this.translit = source["translit"];
	        this.english = source["english"];
	        this.dictionary = this.convertValues(source["dictionary"], ScriptureWordData_Dictionary);
	        this.inflected_count = source["inflected_count"];
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

