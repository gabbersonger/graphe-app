export namespace app {
	
	export class ScriptureVerse {
	    ref: number;
	    words: string[];
	
	    static createFrom(source: any = {}) {
	        return new ScriptureVerse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ref = source["ref"];
	        this.words = source["words"];
	    }
	}
	export class ScriptureRange {
	    start: number;
	    end: number;
	
	    static createFrom(source: any = {}) {
	        return new ScriptureRange(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
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
	    version: string;
	    range: ScriptureRange;
	    blocks: ScriptureBlock[];
	
	    static createFrom(source: any = {}) {
	        return new ScriptureSection(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.version = source["version"];
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
	
	export class ScriptureWordValue_Strongs {
	    num: string;
	    grammar: string;
	
	    static createFrom(source: any = {}) {
	        return new ScriptureWordValue_Strongs(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.num = source["num"];
	        this.grammar = source["grammar"];
	    }
	}
	export class ScriptureWordValue_Dictionary {
	    form: string;
	    gloss: string;
	
	    static createFrom(source: any = {}) {
	        return new ScriptureWordValue_Dictionary(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.form = source["form"];
	        this.gloss = source["gloss"];
	    }
	}
	export class ScriptureWordValue {
	    translit: string;
	    english: string;
	    conjoin_word: string;
	    sub_meaning: string;
	    dictionary: ScriptureWordValue_Dictionary[];
	    strongs: ScriptureWordValue_Strongs[];
	
	    static createFrom(source: any = {}) {
	        return new ScriptureWordValue(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.translit = source["translit"];
	        this.english = source["english"];
	        this.conjoin_word = source["conjoin_word"];
	        this.sub_meaning = source["sub_meaning"];
	        this.dictionary = this.convertValues(source["dictionary"], ScriptureWordValue_Dictionary);
	        this.strongs = this.convertValues(source["strongs"], ScriptureWordValue_Strongs);
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

