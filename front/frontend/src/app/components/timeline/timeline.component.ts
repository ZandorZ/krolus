import { ChangeDetectionStrategy, Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { MatSelectionListChange } from '@angular/material/list';
import { ItemCollection, ItemModel } from 'src/app/models/item.model';
import { SubscriptionReadMap } from 'src/app/models/subscription.model';
import { LeafModel } from 'src/treex/model';

@Component({
    selector: 'app-timeline',
    templateUrl: './timeline.component.html',
    styleUrls: ['./timeline.component.scss'],
    changeDetection: ChangeDetectionStrategy.OnPush,

})
export class TimelineComponent implements OnInit {

    @Input()
    items: ItemCollection;

    @Input()
    selected: ItemModel;

    @Output()
    select: EventEmitter<ItemModel> = new EventEmitter();

    @Output()
    selectSub: EventEmitter<LeafModel> = new EventEmitter();

    @Output()
    favorite: EventEmitter<ItemModel> = new EventEmitter();

    @Output()
    markRead: EventEmitter<SubscriptionReadMap> = new EventEmitter();


    constructor() { }

    ngOnInit(): void {
    }

    onChange(event: MatSelectionListChange) {
        const item = event.options[0].value as ItemModel;
        this.select.emit(item);
    }

    onSelectSub(event: Event, sub: LeafModel) {
        event.stopPropagation();
        this.selectSub.emit(sub);
    }

    setFavorite(event: Event, item: ItemModel) {
        event.stopImmediatePropagation();
        event.preventDefault();
        this.favorite.emit(item);
        item.Favorite = !item.Favorite;
    }

    markAsRead(event: Event, item: ItemModel) {
        event.stopImmediatePropagation();
        event.preventDefault();
        let idMap: SubscriptionReadMap = {};
        idMap[item.Subscription] = [item.ID];
        this.markRead.emit(idMap);
    }

}
