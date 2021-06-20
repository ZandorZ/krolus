import { Component, EventEmitter, Input, OnChanges, Output, ViewEncapsulation } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';
import { DomSanitizer, SafeHtml } from '@angular/platform-browser';
import { ItemModel } from 'src/app/models/item.model';
import { MediaStore } from 'src/app/services/state/media.store';
import { TreexNodeHeader } from 'src/treex/state/store';


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

    @Output()
    selectSub = new EventEmitter<TreexNodeHeader>();

    content: SafeHtml;
    loading = false;

    constructor(private sanitizer: DomSanitizer, private mediaStore: MediaStore, private _snackBar: MatSnackBar) {
    }

    ngOnChanges(): void {
        if (!!this.model) {
            this.model.New = false;
        }
        this.content = this.sanitizer.bypassSecurityTrustHtml(this.model.Description);
    }

    openLink() {
        this.open.emit(this.model.ID);
    }

    async donwloadItem() {

        this.loading = true;
        try {
            const cont = await this.mediaStore.downloadItem(this.model.ID);
            this.content = this.sanitizer.bypassSecurityTrustHtml(cont);
        } catch (error) {
            this._snackBar.open(error, 'OK', { verticalPosition: 'top' });
        }
        this.loading = false;
    }

}
