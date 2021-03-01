import { Component, Inject, OnInit } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { NodeModel } from 'src/treex/model';

@Component({
    selector: 'app-node-dialog-form',
    templateUrl: './node-dialog-form.component.html',
    styleUrls: ['./node-dialog-form.component.scss']
})
export class NodeDialogFormComponent implements OnInit {

    constructor(public dialogRef: MatDialogRef<NodeDialogFormComponent>,
        @Inject(MAT_DIALOG_DATA) public data?: NodeModel) {
        if (this.data == undefined) {
            this.data = {
                children: null,
                nodes_count: 0,
                label: "",
                description: "",
                expanded: false
            }
        }
    }

    ngOnInit() {
    }

    onCancel(): void {
        this.dialogRef.close();
    }

    isValid(): boolean {
        return this.data.label.trim().length > 0 && this.data.description.trim().length > 0;
    }

}


