import { Injectable } from '@angular/core';
import { Store } from 'rxjs-observable-store';
import { ItemCollection, ItemModel, PaginatedItemCollection, PaginatedRequest } from 'src/app/models/item.model'
import { Observable } from 'rxjs';
import { distinctUntilChanged, map } from 'rxjs/operators';
import { isNode, TreexNode } from 'src/treex/model';



export class FeedState {
    Label: string
    PaginatedItems: PaginatedItemCollection
    Selected: ItemModel
}

@Injectable({
    providedIn: "root"
})
export class FeedStore extends Store<FeedState> {

    //alerts$: BehaviorSubject<SubscriptionModel>;

    currentReq: PaginatedRequest;

    constructor() {
        super(new FeedState);
        //  this.alerts$ = new BehaviorSubject<SubscriptionModel>(null);
        //@ts-ignore                
        wails.Events.On("feed.update", async _ => this.dispatchUpdate());
        //@ts-ignore                
        //wails.Events.On("feed.alert", (sub: SubscriptionModel) => this.dispatchAlert(sub));
    }

    private async dispatchUpdate() {
        await this.loadRemote();
    }

    // private dispatchAlert(sub: SubscriptionModel) {
    //     console.warn('Feed Alert:', sub.Title);
    //     this.alerts$.next(sub);
    // }

    async loadMoreItems(selected: TreexNode, req: PaginatedRequest) {

        if (isNode(selected)) {
            req.NodeID = selected.id;
        } else {
            req.LeafIDs = [selected.id];
        }
        this.currentReq = req
        await this.loadRemote();
        this.patchState(selected.label, "Label");
    }

    private async loadRemote() {
        //@ts-ignore
        const itens = await window.backend.FeedStore.LoadMoreItems(this.currentReq);
        this.patchState(itens, "PaginatedItems");
    }


    // getAlerts(throttle: number): Observable<SubscriptionModel> {
    //     return this.alerts$
    //         .asObservable()
    //         .pipe(
    //             throttleTime(throttle)
    //         );
    // }

    selectItem(item: ItemModel) {
        this.patchState(item, "Selected");
    }

    unSelectItem() {
        this.patchState(undefined, "Selected");
    }

    getLabel(): Observable<string> {
        return this.onChanges("Label")
            .pipe(
                distinctUntilChanged()
            );
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

}