import { Injectable } from "@angular/core";
import { ItemModel } from "src/app/models/item.model";

@Injectable({
    providedIn: "root"
})
export class ItemStore {

    constructor() { }

    async fetchItem(id: string, isNew: boolean): Promise<ItemModel> {
        //@ts-ignore
        return window.backend.ItemStore.FetchItem(id, isNew);
    }


}