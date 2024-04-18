export namespace app {
	
	export class ScriptureWord {
	    word_num: number;
	    text: string;
	    pre: string;
	    post: string;
	
	    static createFrom(source: any = {}) {
	        return new ScriptureWord(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.word_num = source["word_num"];
	        this.text = source["text"];
	        this.pre = source["pre"];
	        this.post = source["post"];
	    }
	}
	export class ScriptureVerse {
	    ref: number;
	    words: ScriptureWord[];
	
	    static createFrom(source: any = {}) {
	        return new ScriptureVerse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ref = source["ref"];
	        this.words = this.convertValues(source["words"], ScriptureWord);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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
		    if (a.slice) {
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
		    if (a.slice) {
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
	
	
	export class ScriptureWordData_Strongs {
	    num: string;
	    grammar: string;
	
	    static createFrom(source: any = {}) {
	        return new ScriptureWordData_Strongs(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.num = source["num"];
	        this.grammar = source["grammar"];
	    }
	}
	export class ScriptureWordData_Dictionary {
	    form: string;
	    gloss: string;
	
	    static createFrom(source: any = {}) {
	        return new ScriptureWordData_Dictionary(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.form = source["form"];
	        this.gloss = source["gloss"];
	    }
	}
	export class ScriptureWordData {
	    ref: number;
	    word_number: number;
	    text: string;
	    translit: string;
	    english: string;
	    dictionary: ScriptureWordData_Dictionary[];
	    strongs: ScriptureWordData_Strongs[];
	
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
	        this.strongs = this.convertValues(source["strongs"], ScriptureWordData_Strongs);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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

