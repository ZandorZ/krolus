import { Component, Inject, OnInit } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { MatSlideToggleChange } from '@angular/material/slide-toggle';
import { FilterRequest } from 'src/app/models/item.model';


@Component({
    selector: 'app-filter-dialog-form',
    templateUrl: './filter-dialog-form.html',
    styleUrls: ['./filter-dialog-form.scss']
})
export class FilterDialogFormComponent implements OnInit {


    data: FilterRequest;

    favoriteOption = false;
    newOption = false;

    constructor(public dialogRef: MatDialogRef<FilterDialogFormComponent>,
        @Inject(MAT_DIALOG_DATA) _data?: FilterRequest) {

        this.data = { ..._data };
        this.newOption = this.data.New != undefined;
        this.favoriteOption = this.data.Favorite != undefined;
    }

    ngOnInit() {
    }

    onClear(): void {
        this.favoriteOption = false;
        this.newOption = false;
        this.data = {};
    }

    isFilled(): boolean {
        //return this.data.label.trim().length > 0 && this.data.description.trim().length > 0;
        return false;
    }

    onChangeItem(event: MatSlideToggleChange, field: string) {
        if (field == 'New') {
            if (event.checked) {
                this.data.New = true;
            } else {
                delete this.data.New;
            }
        }
        if (field == 'Favorite') {
            if (event.checked) {
                this.data.Favorite = true;
            } else {
                delete this.data.Favorite;
            }
        }
    }

}


