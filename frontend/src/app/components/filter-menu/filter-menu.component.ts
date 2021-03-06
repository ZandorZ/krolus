import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import { MatDialog, MatDialogConfig, MatDialogRef } from '@angular/material/dialog';
import { FilterRequest } from 'src/app/models/item.model';
import { FilterDialogFormComponent } from './dialog/filter-dialog-form';

@Component({
    selector: 'app-filter-menu',
    templateUrl: './filter-menu.component.html',
    styleUrls: ['./filter-menu.component.scss']
})
export class FilterMenuComponent implements OnInit {


    @Output()
    filter: EventEmitter<FilterRequest>;

    filterRequest: FilterRequest = {};

    constructor(public dialog: MatDialog) {
        this.filter = new EventEmitter<FilterRequest>();
    }

    ngOnInit(): void {
    }

    async showDialog() {
        const config: MatDialogConfig<FilterRequest> = {
            panelClass: 'custom-modalbox-filter',
            data: this.filterRequest
        }
        const dialogRef: MatDialogRef<FilterDialogFormComponent, FilterRequest> = this.dialog.open(FilterDialogFormComponent, config);
        const data = await dialogRef.afterClosed().toPromise();
        if (!!data) {
            this.filterRequest = data;
            this.filter.emit(this.filterRequest);
        }
    }

    hasFilter(): boolean {
        return this.filterRequest.Favorite != undefined || this.filterRequest.New != undefined;
    }

}
