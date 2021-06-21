import { Injectable } from '@angular/core';
import { Store } from 'rxjs-observable-store';
import { ItemCollection, ItemModel, PaginatedItemCollection, PaginatedRequest } from 'src/app/models/item.model'
import { Observable, Subject } from 'rxjs';
import { distinctUntilChanged, map } from 'rxjs/operators';
import { isNode, TreexNode } from 'src/treex/model';
import { SubscriptionReadMap } from 'src/app/models/subscription.model';



export class FeedState {
    PaginatedItems: PaginatedItemCollection
    Selected: ItemModel
    Loading: boolean
}

@Injectable({
    providedIn: "root"
})
export class FeedStore extends Store<FeedState> {

    // alerts$: BehaviorSubject<SubscriptionModel>;
    update$: Subject<void>;

    currentReq: PaginatedRequest;

    constructor() {
        super(new FeedState);

        this.update$ = new Subject();
        //  this.alerts$ = new BehaviorSubject<SubscriptionModel>(null);
        //@ts-ignore                
        wails.Events.On("feed.update", _ => this.update$.next());
        //@ts-ignore                
        //wails.Events.On("feed.alert", (sub: SubscriptionModel) => this.alerts$.next(sub));
    }

    async loadMoreItems(selected: TreexNode, req: PaginatedRequest) {

        if (isNode(selected)) {
            req.NodeID = selected.id;
        } else {
            req.LeafIDs = [selected.id];
        }
        this.currentReq = req
        await this.loadRemote();
    }

    private async loadRemote() {
        this.setLoading(true);
        //@ts-ignore
        const itens = await window.backend.FeedStore.LoadMoreItems(this.currentReq);
        this.patchState(itens, "PaginatedItems");
        this.setLoading(false);
    }


    selectItem(item: ItemModel) {
        this.patchState(item, "Selected");
    }

    async markAllRead() {

        let ids: SubscriptionReadMap = {};
        for (let i in this.state.PaginatedItems.Items) {
            if (this.state.PaginatedItems.Items[i].New) {
                const sub = this.state.PaginatedItems.Items[i].Subscription;
                if (!ids[sub]) {
                    ids[sub] = [];
                }
                ids[sub].push(this.state.PaginatedItems.Items[i].ID);
            }
        }

        if (Object.keys(ids).length == 0) {
            return
        }

        await this.markAsRead(ids);
    }

    async markAsRead(ids: SubscriptionReadMap) {
        try {
            // @ts-ignore
            await window.backend.TreeStore.MarkAllRead(ids);
            await this.loadRemote();

        } catch (e: any) {
            console.error(e);
        }
    }


    unSelectItem() {
        this.patchState(undefined, "Selected");
    }

    // getAlerts(throttle: number): Observable<SubscriptionModel> {
    //     return this.alerts$
    //         .asObservable()
    //         .pipe(
    //             throttleTime(throttle)
    //         );
    // }

    getUpdate(): Observable<void> {
        return this.update$.asObservable();
    }

    getTotal(): Observable<number> {
        return this.onChanges("PaginatedItems", "Total")
            .pipe(
                distinctUntilChanged()
            );
    }

    getItems(): Observable<ItemCollection> {
        return this.onChanges("PaginatedItems", "Items");
    }

    isSelected(): Observable<boolean> {
        return this.onChanges("Selected")
            .pipe(
                map(item => !!item)
            );
    }

    getSelected(): Observable<ItemModel> {
        return this.onChanges("Selected")
            .pipe(
                // filter(item => !!item),
                // distinctUntilKeyChanged("ID")
            );
    }

    isLoading(): Observable<boolean> {
        return this.onChanges('Loading');
    }

    setLoading(flag: boolean) {
        this.patchState(flag, 'Loading');
    }

}