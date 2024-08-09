export namespace main {
	
	export class Task {
	    id: number;
	    text: string;
	    dateTime: string;
	    isDone: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Task(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.text = source["text"];
	        this.dateTime = source["dateTime"];
	        this.isDone = source["isDone"];
	    }
	}

}

