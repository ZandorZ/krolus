import { Component, Input, OnChanges, SimpleChanges } from '@angular/core';
import { ItemTypes } from 'src/app/models/item.model';

@Component({
    selector: 'app-item-icon',
    templateUrl: './item-icon.component.html',
    styleUrls: ['./item-icon.component.scss']
})
export class ItemIconComponent implements OnChanges {


    @Input()
    ItemType: string;

    ItemIcon: string;

    constructor() {

    }

    ngOnChanges(changes: SimpleChanges): void {
        this.ItemIcon = ItemTypes[this.ItemType];
    }

}
