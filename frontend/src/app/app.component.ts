import { DOCUMENT } from '@angular/common';
import { Component, Inject, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { filter } from 'rxjs/operators';
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
    treeSelected$: Observable<TreexNode>;
    isSelected$: Observable<boolean>;
    dragged$: Observable<TreexNode>;
    item: ItemModel;

    constructor(
        @Inject(DOCUMENT) private document: any,
        private treeStore: TreexStore,
        private feedStore: FeedStore,
        private itemStore: ItemStore) {

        this.treeSelected$ = this.treeStore.getSelected();
        this.isSelected$ = this.feedStore.isSelected();
        this.loading$ = this.treeStore.getLoading();
        this.dragged$ = this.treeStore.getDragged();
        this.feedStore.getSelected()
            .pipe(
                filter(item => !!item)
            ).subscribe(item => this.onSelectedChange(item));
    }

    private async onSelectedChange(item: ItemModel) {
        let loadRoot = item.New;
        this.item = await this.itemStore.fetchItem(item.ID, item.New);
        //TODO: patch tree by sub path
        if (loadRoot) this.treeStore.loadChildren("", "");
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
