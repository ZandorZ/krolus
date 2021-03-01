import { Store } from 'rxjs-observable-store';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { distinctUntilChanged, filter, map } from 'rxjs/operators';
import { TreexState } from './store';
import { LoadingDictionary, NodeModel, TreexNode } from '../model';
var objectPath = require("object-path");

const tree: NodeModel = {
    id: '0',
    icon: 'home',
    label: 'Root',
    nodes_count: 0,
    children: [],
    leaves: [],
    expanded: false
}

const myState = new TreexState();

@Injectable()
export class TreexStore extends Store<TreexState> {

    constructor() {
        super(myState);
    }


    loadChildren(id: string, path: string) {

        this.patchState(true, "loading", id);

        console.log(path)

        const node: NodeModel = objectPath.get(tree, path);

        this.patchState(false, "loading", id);
    }

    unloadChildren(id: string, path: string) {
        objectPath.set(this.state, 'root.' + path + '.children', null);
        objectPath.set(this.state, 'root.' + path + '.leaves', null);
    }

    updateSelected(selected: TreexNode): void {
        this.patchState(selected, "selected");
    }


    clearSelected(): void {
        this.patchState(undefined, "selected");
    }

    collapseAll(): void {

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
    getSelected(): Observable<TreexNode> {
        return this.onChanges('selected').pipe(
            distinctUntilChanged((x: TreexNode, y: TreexNode): boolean => {
                if (!x) return false;
                return x.id == y.id;
            })
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

    }

    async moveNode(id: string, target: string) {

    }

    async addNode(node: TreexNode, target: string) {

    }

    async editNode(node: TreexNode) {

    }

    async removeNode(node: TreexNode) {

    }

}