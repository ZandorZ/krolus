import { ChangeDetectionStrategy, Component, EventEmitter, Input, OnInit, Output } from '@angular/core';

@Component({
    selector: 'app-pagesize',
    templateUrl: './pagesize.component.html',
    styleUrls: ['./pagesize.component.scss'],
    changeDetection: ChangeDetectionStrategy.OnPush,
})
export class PagesizeComponent implements OnInit {

    @Input()
    disabled = false;

    @Output()
    change = new EventEmitter<number>();

    counter = 40;
    min = 10;
    max = 120;
    step = 10;

    constructor() { }

    ngOnInit(): void {
    }

    add() {
        if (this.counter < this.max) {
            this.counter += this.step;
            this.change.emit(this.counter);
        }
    }

    minus() {
        if (this.counter > this.min) {
            this.counter -= this.step;
            this.change.emit(this.counter);
        }
    }

}
