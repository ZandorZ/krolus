import { Component, Inject, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { debounceTime, filter, tap } from 'rxjs/operators';
import { SubscriptionModel } from 'src/app/models/subscription.model';

@Component({
    selector: 'app-leaf-dialog-form',
    templateUrl: './leaf-dialog-form.component.html',
    styleUrls: ['./leaf-dialog-form.component.scss']
})
export class LeafDialogFormComponent implements OnInit {

    formG: FormGroup

    querying = false;
    errorQuery = "";

    constructor(
        private fb: FormBuilder,
        public dialogRef: MatDialogRef<LeafDialogFormComponent>,
        @Inject(MAT_DIALOG_DATA) public data?: SubscriptionModel) {

        const urlRegex = /^(?:http(s)?:\/\/)?[\w.-]+(?:\.[\w\.-]+)+[\w\-\._~:/?#[\]@!\$&'\(\)\*\+,;=.]+$/;

        this.formG = this.fb.group({
            XURL: new FormControl(null, [Validators.required, Validators.pattern(urlRegex)]),
            Title: new FormControl(null, [Validators.required,]),
            Description: new FormControl(null, [Validators.required,]),
            URL: new FormControl(null, [Validators.required]),
        });

        if (this.data == undefined) {
            const urlField = this.formG.get('XURL');
            urlField.valueChanges
                .pipe(
                    debounceTime(400),
                    tap(_ => {
                        this.errorQuery = "";
                        this.formG.get("Title").reset();
                        this.formG.get("Description").reset();
                        this.formG.get("URL").reset();
                    }),
                    filter(_ => urlField.valid),
                ).subscribe(value => this.onChangeURL(value));
        } else {
            this.formG.get("XURL").disable({ emitEvent: false });
            this.formG.patchValue(this.data, { emitEvent: false });
        }
    }

    async onChangeURL(url: string) {
        //TODO: should be here?

        this.querying = true;
        this.formG.disable({ emitEvent: false });
        try {
            //@ts-ignore          
            const sub: SubscriptionModel = await window.backend.FeedStore.LoadSub(url);
            sub.XURL = url;
            this.formG.patchValue(sub, { emitEvent: false });

        } catch (e) {
            this.errorQuery = e;
        }
        this.querying = false;
        this.formG.enable({ emitEvent: false });
    }

    ngOnInit(): void {
    }

    onCancel(): void {
        this.dialogRef.close();
    }

    onSend(): void {
        if (this.data) {
            this.data.LastUpdate = undefined;
        }
        this.dialogRef.close({ ...this.data, ...this.formG.value });
    }

    getErrorMessage(field: string): string {
        //TODO: fix this
        if (this.formG.get(field).invalid) {
            return "Invalid field"
        }
        return "";
    }

}
