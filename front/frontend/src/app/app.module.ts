import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { NgModule, NO_ERRORS_SCHEMA } from '@angular/core';
import { MatSliderModule } from '@angular/material/slider';
import { MatIconModule } from '@angular/material/icon';
import { MatRadioModule } from '@angular/material/radio';
import { MatProgressBarModule } from '@angular/material/progress-bar';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatTooltipModule } from '@angular/material/tooltip';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatMenuModule } from '@angular/material/menu';
import { TreexModule } from 'src/treex/treex.module';

import { AppComponent } from './app.component';
import { sanitizePipe } from './pipes/sanitize.pipe';
import { TreexStore } from 'src/treex/state/treex.store';
import { TreeStore } from './services/state/tree.store';
import { TimelineComponent } from './components/timeline/timeline.component';
import { FeedComponent } from './components/feed/feed.component';
import { ItemComponent } from './components/item/item.component';
import { MatPaginatorModule } from '@angular/material/paginator';
import { MatListModule } from '@angular/material/list';
import { MatDialogModule } from '@angular/material/dialog';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { MatInputModule } from '@angular/material/input';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { IframePlayerComponent } from './components/iframe-player/iframe-player.component';
import { NodeDialogFormComponent } from './components/node-dialog-form/node-dialog-form.component';
import { LeafDialogFormComponent } from './components/leaf-dialog-form/leaf-dialog-form.component';
import { GridComponent } from './components/grid/grid.component';
import { ConfirmDialogComponent } from './components/confirm-dialog/confirm-dialog.component';
import { CustomDatePipe } from './pipes/custom-date.pipe';
import { FilterMenuComponent } from './components/filter-menu/filter-menu.component';
import { FilterDialogFormComponent } from './components/filter-menu/dialog/filter-dialog-form';
import { PagesizeComponent } from './components/pagesize/pagesize.component';
import { LoadingMaskDirective } from './directives/loading-mask.directive';
import { PreloadImgComponent } from './components/preload-img/preload-img.component';
import { LoadingComponent } from './components/loading/loading.component';
import { ItemIconComponent } from './components/item-icon/item-icon.component';


@NgModule({
    declarations: [
        AppComponent,
        TimelineComponent,
        FeedComponent,
        ItemComponent,
        sanitizePipe,
        IframePlayerComponent,
        NodeDialogFormComponent,
        LeafDialogFormComponent,
        FilterDialogFormComponent,
        GridComponent,
        ConfirmDialogComponent,
        CustomDatePipe,
        FilterMenuComponent,
        PagesizeComponent,
        LoadingMaskDirective,
        PreloadImgComponent,
        LoadingComponent,
        ItemIconComponent,
    ],
    imports: [
        BrowserModule,
        BrowserAnimationsModule,
        FormsModule,
        ReactiveFormsModule,
        MatSliderModule,
        MatIconModule,
        MatTooltipModule,
        MatDialogModule,
        MatMenuModule,
        MatButtonModule,
        MatCardModule,
        MatRadioModule,
        MatSlideToggleModule,
        MatToolbarModule,
        MatInputModule,
        MatProgressBarModule,
        MatListModule,
        MatSnackBarModule,
        MatPaginatorModule,
        MatSidenavModule,
        {
            ngModule: TreexModule,
            providers: [
                { provide: TreexStore, useExisting: TreeStore },
            ]
        },
    ],
    providers: [],
    bootstrap: [AppComponent],
    schemas: [NO_ERRORS_SCHEMA] //
})
export class AppModule { }
