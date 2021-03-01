import { Injectable } from "@angular/core";

@Injectable({
    providedIn: "root"
})
export class MediaStore {

    constructor() { }

    async downloadItem(id: string): Promise<string> {
        //@ts-ignore
        return window.backend.MediaStore.Download(id) as string;
    }


}