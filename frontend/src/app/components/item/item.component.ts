import { Component, EventEmitter, Input, OnChanges, Output, ViewEncapsulation } from '@angular/core';
import { DomSanitizer, SafeHtml } from '@angular/platform-browser';
import { ItemModel } from 'src/app/models/item.model';
import { MediaStore } from 'src/app/services/state/media.store';


@Component({
    selector: 'app-item',
    templateUrl: './item.component.html',
    styleUrls: ['./item.component.scss'],
    encapsulation: ViewEncapsulation.None,
})
export class ItemComponent implements OnChanges {

    @Input()
    model: ItemModel;

    @Output()
    close = new EventEmitter<void>();

    @Output()
    open = new EventEmitter<string>();

    content: SafeHtml;

    constructor(private sanitizer: DomSanitizer, private mediaStore: MediaStore) {



    }


    ngOnChanges(): void {
        if (!!this.model) this.model.New = false;
        this.content = "";
    }


    openLink() {
        this.open.emit(this.model.Link);
    }

    async donwloadItem() {
        //@ts-ignore
        const cont = await this.mediaStore.downloadItem(this.model.ID);

        const style = `
        <style>
            img {                 
                max-width: 80%; 
            }
            a {
                color: orange;
            }
        </style>`;

        this.content = this.sanitizer.bypassSecurityTrustHtml(style + cont);
    }

}
