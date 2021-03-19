import { DOCUMENT } from '@angular/common';
import { Component, Inject, OnInit } from '@angular/core';
import { from, Observable } from 'rxjs';
import { filter, switchMap } from 'rxjs/operators';
import { LoadingDictionary, TreexNode } from 'src/treex/model';
import { TreexStore } from 'src/treex/state/treex.store';
import { ItemModel } from './models/item.model';
import { FeedStore } from './services/state/feed.store';
import { ItemStore } from './services/state/item.store';

@Component({
    selector: '[id="app"]',
    templateUrl: './app.component.html',
    styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {

    opened = true;

    loading$: Observable<LoadingDictionary>;
    selected$: Observable<TreexNode>;
    isSelected$: Observable<boolean>;
    dragged$: Observable<TreexNode>;
    item: ItemModel;

    constructor(
        @Inject(DOCUMENT) private document: any,
        private treeStore: TreexStore,
        private feedStore: FeedStore,
        private itemStore: ItemStore) {

        this.selected$ = this.treeStore.getSelected();
        this.isSelected$ = this.feedStore.isSelected();
        this.loading$ = this.treeStore.getLoading();
        this.dragged$ = this.treeStore.getDragged();
        this.feedStore.getSelected()
            .pipe(
                filter(item => !!item),
                switchMap(item => from(this.itemStore.fetchItem(item.ID, item.New))),
            ).subscribe(item => this.item = item);
    }

    ngOnInit(): void {
        this.document.body.classList.add('dark-theme');
    }

    onCloseItem() {
        this.item = undefined;
        this.feedStore.unSelectItem();
    }

    onToggleMenu(hidden: boolean) {
        this.opened = !hidden;
    }

    async onOpenURL(url: string) {
        //@ts-ignore
        await window.backend.ItemStore.OpenItem(url);
    }

}
