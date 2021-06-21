import { ChangeDetectionStrategy, Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { ItemCollection, ItemModel } from 'src/app/models/item.model';
import { SubscriptionReadMap } from 'src/app/models/subscription.model';
import { LeafModel } from 'src/treex/model';

@Component({
    selector: 'app-grid',
    templateUrl: './grid.component.html',
    styleUrls: ['./grid.component.scss'],
    changeDetection: ChangeDetectionStrategy.OnPush,
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

    @Output()
    markRead: EventEmitter<SubscriptionReadMap> = new EventEmitter();


    @Output()
    favorite: EventEmitter<ItemModel> = new EventEmitter();

    constructor() { }

    ngOnInit(): void {
    }

    onChange(item: ItemModel) {
        this.select.emit(item);
    }

    onSelectSub(event: Event, sub: LeafModel) {
        event.stopPropagation();
        this.selectSub.emit(sub);
    }

    markAsRead(event: Event, item: ItemModel) {
        event.stopImmediatePropagation();
        event.preventDefault();
        let idMap: SubscriptionReadMap = {};
        idMap[item.Subscription] = [item.ID];
        this.markRead.emit(idMap);
    }


    setFavorite(event: Event, item: ItemModel) {
        event.stopImmediatePropagation();
        event.preventDefault();
        this.favorite.emit(item);
        item.Favorite = !item.Favorite;
    }

}
