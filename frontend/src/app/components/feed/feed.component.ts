import { ChangeDetectionStrategy, Component, ElementRef, EventEmitter, Input, OnChanges, Output, SimpleChanges, ViewChild } from '@angular/core';
import { MatPaginator, PageEvent } from '@angular/material/paginator';
import { Observable } from 'rxjs';
import { FilterRequest, ItemCollection, ItemModel, PaginatedRequest } from 'src/app/models/item.model';
import { FeedStore } from 'src/app/services/state/feed.store';
import { ItemStore } from 'src/app/services/state/item.store';
import { LeafModel, TreexNode } from 'src/treex/model';
import { TreexStore } from 'src/treex/state/treex.store';

@Component({
    selector: 'app-feed',
    templateUrl: './feed.component.html',
    styleUrls: ['./feed.component.scss'],
    changeDetection: ChangeDetectionStrategy.OnPush,
})
export class FeedComponent implements OnChanges {

    @ViewChild('scrollcont') private myScrollContainer: ElementRef;


    @Output() hidemenu: EventEmitter<boolean>;

    hiddenmenu = false;

    @Input()
    node: TreexNode;

    typeGrid = false;

    selected$: Observable<ItemModel>;
    label$: Observable<string>;
    items$: Observable<ItemCollection>;
    total$: Observable<number>;
    request: PaginatedRequest = {
        ItemsPerPage: 30,
        Page: 0,
        LeafIDs: [],
        NodeID: "",
    }

    @ViewChild('paginator')
    paginator: MatPaginator;

    constructor(private store: FeedStore, private treeStore: TreexStore, private itemStore: ItemStore) {
        this.label$ = this.store.getLabel();
        this.items$ = this.store.getItems();
        this.total$ = this.store.getTotal();
        this.selected$ = this.store.getSelected();
        this.hidemenu = new EventEmitter();
    }

    toggleHideMenu() {
        this.hiddenmenu = !this.hiddenmenu;
        this.hidemenu.emit(this.hiddenmenu);
    }

    async ngOnChanges(changes: SimpleChanges) {
        if (!!this.node) {
            this.request.Page = 0;
            this.request.LeafIDs = [];
            this.request.NodeID = "";
            this.paginator.pageIndex = 0;
            await this.loadMoreItems();
        }
    }

    async loadMoreItems() {
        await this.store.loadMoreItems(this.node, this.request);
        this.myScrollContainer.nativeElement.scrollTo(0, 0);
    }

    async onChangePage(event: PageEvent) {
        this.request.Page = event.pageIndex;
        this.request.ItemsPerPage = event.pageSize;
        await this.loadMoreItems();
    }

    async onFavoriteItem(item: ItemModel) {
        await this.itemStore.favoriteItem(item.ID);
    }

    async onChangeFilter(filter: FilterRequest) {
        this.request.Filter = filter;
        this.request.Page = 0;
        this.paginator.pageIndex = 0;
        await this.loadMoreItems();
    }

    onSelecItem(item: ItemModel) {
        this.store.selectItem(item);
    }

    onSelectLeaf(sub: LeafModel) {
        this.treeStore.updateSelected(sub);
    }

}
