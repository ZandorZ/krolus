import { Input, SimpleChanges } from '@angular/core';
import { Component, OnInit } from '@angular/core';

@Component({
    selector: 'app-preload-img',
    templateUrl: './preload-img.component.html',
    styleUrls: ['./preload-img.component.scss']
})
export class PreloadImgComponent implements OnInit {


    @Input()
    src: string;

    loading = false;
    source: string;
    filter = 'blur(4px)';
    opacity = '0.3';

    constructor() { }


    ngOnInit(): void {

    }


    ngOnChanges(changes: SimpleChanges): void {

        this.loading = false;
        this.source = undefined;
        if (changes.src.currentValue) {
            this.loadImage(changes.src.currentValue);
        }
    }


    private loadImage(src: string) {

        this.opacity = '0.3'
        this.filter = 'blur(5px)';
        this.loading = true;
        const img = new Image();
        img.src = src;
        img.onload = () => {
            this.loading = false;
            this.source = src;
            setTimeout(() => {
                this.filter = 'none';
                this.opacity = '1';
            }, 200);
        };
    }

}
