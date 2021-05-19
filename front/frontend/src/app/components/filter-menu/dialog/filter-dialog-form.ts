import { Component, Inject, OnInit } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { MatSlideToggleChange } from '@angular/material/slide-toggle';
import { FilterRequest, ItemType, ItemTypes } from 'src/app/models/item.model';


@Component({
    selector: 'app-filter-dialog-form',
    templateUrl: './filter-dialog-form.html',
    styleUrls: ['./filter-dialog-form.scss']
})
export class FilterDialogFormComponent implements OnInit {


    data: FilterRequest;
    favoriteOption = false;
    newOption = false;
    typesOption = false;
    itemTypes: ItemType;

    constructor(public dialogRef: MatDialogRef<FilterDialogFormComponent>,
        @Inject(MAT_DIALOG_DATA) _data?: FilterRequest) {

        this.data = { ..._data };
        this.newOption = this.data.New != undefined;
        this.favoriteOption = this.data.Favorite != undefined;
        this.typesOption = this.data.Type != undefined;
        this.itemTypes = ItemTypes;
    }

    ngOnInit() {
    }

    onClear(): void {
        this.favoriteOption = this.newOption = this.typesOption = false;
        this.data = {};
    }

    isFilled(): boolean {
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
        if (field == 'Types') {
            if (event.checked) {
                this.data.Type = [];
            } else {
                delete this.data.Type;
            }
        }
    }

}


