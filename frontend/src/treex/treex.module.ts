import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { DndModule } from 'ngx-drag-drop';



import { TreexComponent } from './components/treex/treex.component';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatMenuModule } from '@angular/material/menu';
import { MatListModule } from '@angular/material/list';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { MatTooltipModule } from '@angular/material/tooltip';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner'; import { TreexStore } from './state/treex.store';
import { TreexItemComponent } from './components/treex-item/treex-item.component';

@NgModule({
    declarations: [TreexComponent, TreexItemComponent],
    imports: [
        CommonModule,
        DndModule,
        MatButtonModule,
        MatIconModule,
        MatProgressSpinnerModule,
        MatListModule,
        MatToolbarModule,
        MatMenuModule,
        MatTooltipModule,
    ],
    exports: [TreexComponent, TreexItemComponent],
    providers: [TreexStore],
})
export class TreexModule {
}


