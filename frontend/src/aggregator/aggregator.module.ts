import { NgModule, NO_ERRORS_SCHEMA } from "@angular/core";
import { FormsModule } from "@angular/forms";
import { MatButtonModule } from "@angular/material/button";
import { MatCardModule } from "@angular/material/card";
import { MatDialogModule } from "@angular/material/dialog";
import { MatIconModule } from "@angular/material/icon";
import { MatInputModule } from "@angular/material/input";
import { MatListModule } from "@angular/material/list";
import { MatPaginatorModule } from "@angular/material/paginator";
import { MatSidenavModule } from "@angular/material/sidenav";
import { MatSliderModule } from "@angular/material/slider";
import { MatToolbarModule } from "@angular/material/toolbar";
import { MatTooltipModule } from "@angular/material/tooltip";

@NgModule({
    declarations: [
    ],
    imports: [
        MatSliderModule,
        MatIconModule,
        MatTooltipModule,
        MatDialogModule,
        MatButtonModule,
        MatCardModule,
        MatToolbarModule,
        MatInputModule,
        MatListModule,
        FormsModule,
        MatPaginatorModule,
        MatSidenavModule,
    ],
    providers: [],
    bootstrap: [],
    schemas: [NO_ERRORS_SCHEMA] //
})
export class AggregatorModule { }
