import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import { ConfirmDialogComponent } from 'src/app/components/confirm-dialog/confirm-dialog.component';
import { LeafModel } from 'src/treex/model';

@Component({
    selector: 'treex-item',
    templateUrl: './treex-item.component.html',
    styleUrls: ['../treex/treex.component.scss']
})
export class TreexItemComponent implements OnInit {

    @Input()
    model: LeafModel;

    @Input()
    depth: number;

    @Input()
    selected: boolean;

    @Output()
    select: EventEmitter<void> = new EventEmitter();

    @Output()
    dragStart: EventEmitter<void> = new EventEmitter();

    @Output()
    dragEnd: EventEmitter<void> = new EventEmitter();

    @Output()
    removeLeaf: EventEmitter<LeafModel> = new EventEmitter();

    @Output()
    editLeaf: EventEmitter<LeafModel> = new EventEmitter();

    @Output()
    addFavorite: EventEmitter<LeafModel> = new EventEmitter();


    constructor(public dialog: MatDialog) { }

    ngOnInit(): void {
    }

    onSelect(event: Event): void {
        event.stopImmediatePropagation();
        this.select.emit();
    }

    onDragStart() {
        this.dragStart.emit();
    }

    onDragEnd() {
        this.dragEnd.emit();
    }

    options(event: Event) {
        event.stopImmediatePropagation();
    }

    async showConfirmRemoveLeaf() {
        let options: MatDialogConfig<string> = {
            disableClose: true,
            panelClass: 'custom-modalbox-directory',
            data: `Would you like to remove "${this.model.label}" ?? `,
        };

        const dialogRef = this.dialog.open(ConfirmDialogComponent, options);

        const result = await dialogRef.afterClosed().toPromise();

        if (!!result) {
            this.removeLeaf.emit(this.model);
        }
    }

    clickEdit() {
        this.editLeaf.emit(this.model);
    }

    clickFav() {
        this.addFavorite.emit(this.model);
    }

}
