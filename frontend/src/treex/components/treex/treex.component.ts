import { Component, Input, OnInit } from '@angular/core';
import { TreexStore } from 'src/treex/state/treex.store';
import { isDescendent, isNode, LeafModel, LoadingDictionary, NodeModel, TreexNode } from 'src/treex/model';
import { DndDropEvent } from 'ngx-drag-drop';
import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import { NodeDialogFormComponent } from 'src/app/components/node-dialog-form/node-dialog-form.component';
import { filter } from 'rxjs/operators';
import { LeafDialogFormComponent } from 'src/app/components/leaf-dialog-form/leaf-dialog-form.component';


@Component({
    selector: 'treex',
    templateUrl: './treex.component.html',
    styleUrls: ['./treex.component.scss'],
})
export class TreexComponent implements OnInit {

    @Input() path = "";
    @Input() depth = 0;
    @Input() model: NodeModel;
    @Input() selected: TreexNode;
    @Input() dragged: TreexNode;
    @Input() loading: LoadingDictionary;


    constructor(public store: TreexStore, public dialog: MatDialog) { }

    async ngOnInit() {
        if (this.model == undefined && this.depth == 0) {
            this.store
                .onChanges("root").pipe(
                    filter(data => !!data),
                )
                .subscribe(root => this.model = root);
            //first load
            this.store.loadChildren("", this.path);
        }
    }

    clearDir() {
        this.store.clearSelected();
    }

    subPath(n: NodeModel): string {
        if (this.depth == 0) {
            return 'children.' + n.id;
        }
        return this.path + '.children.' + n.id;
    }

    expand(event: Event) {
        event.stopImmediatePropagation();

        if (!!this.model && (!!this.model.children && !!this.model.leaves)) {
            this.store.unloadChildren(this.model.id, this.path); //this.path
        } else {
            this.store.loadChildren(this.model.id, this.path);
        }
    }

    isLoading(): boolean {
        return !!this.loading && !!this.loading[this.model.id];
    }

    select(event: Event) {
        event.stopPropagation();
        this.store.updateSelected(this.model);
    }

    selectLeaf(leaf: LeafModel) {
        this.store.updateSelected(leaf);
    }

    isSelected(): boolean {
        return !!this.selected && this.model.id == this.selected.id;
    }

    isLeafSelected(leaf: LeafModel): boolean {
        return !!this.selected && leaf.id == this.selected.id;
    }

    isExpansable(): boolean {
        return this.model.nodes_count > 0;
    }

    options(event: Event) {
        event.stopImmediatePropagation();
    }

    async onRemoveLeaf(leaf: LeafModel) {
        await this.store.removeNode(leaf);
    }

    async onDrop(event: DndDropEvent) {
        const moved = event.data as TreexNode;
        const target = this.model as NodeModel;
        if (isNode(moved)) {
            await this.store.moveNode(moved.id, target.id);
        } else {
            await this.store.moveLeaf(moved.id, target.id);
        }
    }

    onDragStart() {
        this.store.startDrag(this.model);
    }

    onDragStartLeaf(model: LeafModel) {
        this.store.startDrag(model);
    }

    onDragEnd() {
        this.store.endDrag();

        setTimeout(() => {
            //clean element class
            document.querySelectorAll('.dndDragover').forEach(item => {
                item.classList.remove('dndDragover');
            });
        }, 300);
    }

    isDroppagle(): boolean {

        //leaf
        if (!!this.dragged && !!this.model.leaves && this.model.leaves.length > 0 && !!this.model.leaves.find(leaf => leaf.id == this.dragged.id)) {
            return false;
        }

        //node
        if (!!this.dragged && isNode(this.dragged)) {
            //same parent
            if (!!this.model.children && this.model.children && !!this.model.children.find(node => node.id == this.dragged.id)) {
                return false;
            }
            return !isDescendent(this.dragged, this.model.id);
        }
        return true;
    }


    //TODO: decouple
    async showNodeDialogForm(isNew: boolean) {

        let options: MatDialogConfig<NodeModel> = {
            disableClose: true,
            panelClass: 'custom-modalbox-directory',
        };

        if (!isNew) {
            options.data = { ...this.model };
        }

        const data = await this.dialog
            .open(NodeDialogFormComponent, options)
            .afterClosed()
            .toPromise();


        if (!!data) {
            if (!!data.id) {
                await this.store.editNode(data as NodeModel);
            } else {
                await this.store.addNode(data as NodeModel, this.model.id);
            }
        }

    }

    //TODO: decouple
    async showLeafDialogForm(leaf?: LeafModel) {

        let options: MatDialogConfig<LeafModel> = {
            disableClose: true,
            panelClass: 'custom-modalbox-subscription',
        };

        if (leaf) {
            //TODO: decouple
            //@ts-ignore 
            const sub = await window.backend.FeedStore.GetSub(leaf.id);
            options.data = { ...sub };
        }

        const data = await this.dialog
            .open(LeafDialogFormComponent, options)
            .afterClosed()
            .toPromise();

        if (!!data) {
            if (!!data.ID) {
                await this.store.editNode(data);
            } else {
                await this.store.addNode(data, this.model.id);
            }
        }
    }

}
