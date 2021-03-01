import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { ItemCollection, ItemModel } from 'src/app/models/item.model';
import { LeafModel } from 'src/treex/model';

@Component({
    selector: 'app-grid',
    templateUrl: './grid.component.html',
    styleUrls: ['./grid.component.scss']
})
export class GridComponent implements OnInit {
    @Input()
    items: ItemCollection;

    @Input()
    selected: ItemModel;

    @Output()
    select: EventEmitter<ItemModel> = new EventEmitter();

    @Output()
    selectSub: EventEmitter<LeafModel> = new EventEmitter();

    constructor() { }

    ngOnInit(): void {
    }

    onChange(item: ItemModel) {
        this.select.emit(item);
        item.New = false;
    }

    onSelectSub(event: Event, sub: LeafModel) {
        event.stopPropagation();
        this.selectSub.emit(sub);
    }

}
