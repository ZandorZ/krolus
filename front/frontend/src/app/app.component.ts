import { DOCUMENT } from '@angular/common';
import { Component, Inject, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { filter } from 'rxjs/operators';
import { getPath, LoadingDictionary, NodeModel, TreexNode } from 'src/treex/model';
import { TreexNodeHeader } from 'src/treex/state/store';
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
    treeSelectedHeaders$: Observable<TreexNodeHeader[]>;
    isSelected$: Observable<boolean>;
    dragged$: Observable<TreexNode>;
    item: ItemModel;

    constructor(
        @Inject(DOCUMENT) private document: any,
        private treeStore: TreexStore,
        private feedStore: FeedStore,
        private itemStore: ItemStore) {

        this.treeSelected$ = this.treeStore.getSelected();
        this.treeSelectedHeaders$ = this.treeStore.getSelectedHeaders();
        this.isSelected$ = this.feedStore.isSelected();
        this.loading$ = this.treeStore.getLoading();
        this.dragged$ = this.treeStore.getDragged();
        this.feedStore.getSelected()
            .pipe(
                filter(item => !!item)
            ).subscribe(item => this.onSelectedChange(item));
    }

    private async onSelectedChange(item: ItemModel) {
        let reloadTree = item.New;
        this.item = await this.itemStore.fetchItem(item.ID, item.New);
        if (reloadTree) {
            //TODO: check if sub is loaded in tree to update newItemsCount
        }
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

    async onOpenItem(id: string) {
        //@ts-ignore
        await window.backend.ItemStore.OpenItem(id); //TODO: wrong access
    }

    async onSelectHeader(header: TreexNodeHeader) {

        await this.treeStore.loadAncestors(header.id, header.leaf);
        if (header.path) {
            let node: TreexNode = { id: header.id, label: header.label };
            if (!header.leaf) {
                (<NodeModel>node).nodes_count = 0;
            }
            this.treeStore.updateSelected(node, header.path);
        } else {
            setTimeout(() => {
                const path = getPath(this.treeStore.state.root, header.id).slice(1);
                this.treeStore.updateSelected({ id: header.id, label: header.label }, path);
            }, 100);
        }
    }

    async onSelectItemSub(header: TreexNodeHeader) {
        await this.treeStore.loadAncestors(header.id, header.leaf);
        setTimeout(() => {
            const path = getPath(this.treeStore.state.root, header.id).slice(1);
            this.treeStore.updateSelected({ id: header.id, label: header.label }, path);
        }, 100);

    }

}
