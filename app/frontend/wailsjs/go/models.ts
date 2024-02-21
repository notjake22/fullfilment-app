export namespace database {
	
	export class CustomUPCModel {
	    itemName: string;
	    imageUri: string;
	    upc: string;
	
	    static createFrom(source: any = {}) {
	        return new CustomUPCModel(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.itemName = source["itemName"];
	        this.imageUri = source["imageUri"];
	        this.upc = source["upc"];
	    }
	}
	export class ItemDataModel {
	    id: string;
	    itemName: string;
	    imageUri: string;
	    upc: string;
	    ownerName: string;
	    ownerID: string;
	    currentlyStored: boolean;
	
	    static createFrom(source: any = {}) {
	        return new ItemDataModel(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.itemName = source["itemName"];
	        this.imageUri = source["imageUri"];
	        this.upc = source["upc"];
	        this.ownerName = source["ownerName"];
	        this.ownerID = source["ownerID"];
	        this.currentlyStored = source["currentlyStored"];
	    }
	}
	export class ScanInRequest {
	    upc: string;
	    clientID: string;
	
	    static createFrom(source: any = {}) {
	        return new ScanInRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.upc = source["upc"];
	        this.clientID = source["clientID"];
	    }
	}
	export class ScanInResponse {
	    ID: string;
	    itemName: string;
	    imageUri: string;
	
	    static createFrom(source: any = {}) {
	        return new ScanInResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.itemName = source["itemName"];
	        this.imageUri = source["imageUri"];
	    }
	}
	export class ServiceSettings {
	    itemPrice: number;
	
	    static createFrom(source: any = {}) {
	        return new ServiceSettings(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.itemPrice = source["itemPrice"];
	    }
	}
	export class User {
	    id: string;
	    name: string;
	    phone: string;
	    balance: number;
	    currentlyStoredItems: number;
	    totalOrderHistory: number;
	
	    static createFrom(source: any = {}) {
	        return new User(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.phone = source["phone"];
	        this.balance = source["balance"];
	        this.currentlyStoredItems = source["currentlyStoredItems"];
	        this.totalOrderHistory = source["totalOrderHistory"];
	    }
	}

}

