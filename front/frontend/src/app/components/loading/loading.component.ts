import { Component, Input, OnInit } from '@angular/core';

@Component({
    selector: 'app-loading',
    templateUrl: './loading.component.svg',
    styleUrls: ['./loading.component.scss']
})
export class LoadingComponent implements OnInit {


    @Input()
    diameter = 64;

    @Input()
    color = '#003E80';

    @Input()
    opacity = 1;

    constructor() { }

    ngOnInit() { }

}
