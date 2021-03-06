import { Store } from 'rxjs-observable-store';
import { Injectable, NgZone } from '@angular/core';
import { Observable } from 'rxjs';
import { delay, distinctUntilChanged, distinctUntilKeyChanged, filter, map } from 'rxjs/operators';
import { getHeadersFromPath, getPath, isNode, LeafModel, LoadingDictionary, NodeModel, TreexNode } from 'src/treex/model';
import { TreexNodeHeader, TreexState } from 'src/treex/state/store';
import * as Wails from '@wailsapp/runtime';


@Injectable({
    providedIn: "root"
})
export class TreeStore extends Store<TreexState> {

    constructor(private zone: NgZone) {
        super(new TreexState);

        Wails.Store.New("TreeState").subscribe((state: any) => {
            this.zone.run(() => this.onChange(state));
        });

    }

    private async onChange(data: NodeModel) {
        // console.log("Tree state updated");
        // console.log(data);
        // root first load
        if (data.children == null && data.leaves == null) {
            this.updateSelected(data, "/");
            await this.loadChildren(data.id, "/");
        }
        this.patchState(data, "root");
    }

    async loadChildren(id: string, path: string) {
        this.patchState(true, "loading", id);
        //@ts-ignore
        await window.backend.TreeStore.LoadNode(id);
        this.patchState(false, "loading", id);
    }

    async unloadChildren(id: string, path: string) {
        this.patchState(true, "loading", id);
        //@ts-ignore
        await window.backend.TreeStore.UnLoadNode(id);
        this.patchState(false, "loading", id);
    }

    async loadAncestors(id: string, isLeaf: boolean) {
        //@ts-ignore
        await window.backend.TreeStore.LoadAncestors(id, isLeaf);
    }


    updateSelected(node: TreexNode, path: string): void {
        this.patchState(node, "selected");
        this.patchState('root.' + path, "selectedPath");
    }

    clearSelected(): void {
        this.patchState(undefined, "selected");
        this.patchState(undefined, "selectedPath");
    }

    async collapseAll() {
        //@ts-ignore
        await window.backend.TreeStore.UnLoadAll();
    }

    async toggleFavorites() {
        //@ts-ignore
        await window.backend.TreeStore.FilterFavorites(!this.state.favorite);
        this.patchState(!this.state.favorite, "favorite");
    }

    async toggleLeafFavorite(leaf: LeafModel) {
        //@ts-ignore
        await window.backend.TreeStore.ToggleLeafFavorite(leaf.id);
        leaf.favorite = !leaf.favorite;
    }

    isLoading(id: string): Observable<boolean> {
        return this.onChanges('loading', id).pipe(
            filter(loading => loading != undefined),
            distinctUntilChanged()
        );
    }

    isSelected(id: string): Observable<boolean> {
        return this.onChanges('selected').pipe(
            map(selected => !!selected && selected == id),
            distinctUntilChanged()
        );
    }

    getRoot(): Observable<TreexNode> {
        return this.onChanges('root');
    }

    getSelected(): Observable<TreexNode> {
        return this.onChanges('selected').pipe(
            filter(selected => !!selected),
            distinctUntilKeyChanged("id")
        );
    }

    getSelectedPath(): Observable<string> {
        return this.onChanges('selectedPath').pipe(
            filter(path => !!path),
            distinctUntilChanged()
        );
    }

    getSelectedHeaders(): Observable<TreexNodeHeader[]> {
        return this.getSelectedPath().pipe(
            delay(200),
            map(path => getHeadersFromPath(this.state.root as NodeModel, path))
        );
    }

    getLoading(): Observable<LoadingDictionary> {
        return this.onChanges('loading');
    }

    startDrag(dragged: TreexNode): void {
        this.patchState(dragged, "dragged");
    }

    endDrag(): void {
        this.patchState(undefined, "dragged");
    }

    getDragged(): Observable<TreexNode> {
        return this.onChanges('dragged');
    }

    async moveLeaf(id: string, target: string) {

        try {
            //@ts-ignore
            await window.backend.TreeStore.MoveLeaf(id, target);

            //moved is selected
            if (this.state.selectedPath.includes(id)) {
                const targetPath = getPath(this.state.root, target);
                //update selectedPath
                this.patchState(targetPath + '.leaves.' + id, "selectedPath");
            }

        } catch (e) { }

    }

    async moveNode(id: string, target: string) {
        //@ts-ignore
        await window.backend.TreeStore.MoveNode(id, target);

        //moved is selected
        if (this.state.selectedPath.includes(id)) {
            const targetPath = getPath(this.state.root, target);
            //update selectedPath
            this.patchState(targetPath + '.children.' + id + this.state.selectedPath.split(id)[1], "selectedPath");
        }
    }

    async addNode(node: TreexNode, target: string) {
        if (isNode(node)) {
            //@ts-ignore
            await window.backend.TreeStore.AddNode(node, target);
        } else {
            //@ts-ignore
            await window.backend.TreeStore.AddSubscription(node, target);
        }
    }

    async editNode(node: TreexNode) {
        if (isNode(node)) {
            //@ts-ignore
            await window.backend.TreeStore.EditNode(node);
        } else {
            //@ts-ignore
            await window.backend.TreeStore.EditSubscription(node as LeafModel);
        }
    }

    async removeNode(node: TreexNode) {
        if (isNode(node)) {

        } else {
            //@ts-ignore
            await window.backend.TreeStore.RemoveSubscription(node.id);
        }
    }

}