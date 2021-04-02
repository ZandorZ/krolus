import { Directive, ElementRef, Input, OnChanges, SimpleChanges } from '@angular/core';

@Directive({
    selector: '[appLoadingMask]'
})
export class LoadingMaskDirective implements OnChanges {

    @Input('appLoadingMask') enabled: boolean;

    mask: HTMLDivElement;

    constructor(private el: ElementRef) {

        this.mask = document.createElement('div');
        this.mask.style.transition = '0.2s all linear';
        this.mask.style.backgroundColor = 'black';
        this.mask.style.position = 'absolute';
        this.mask.style.width = '100%';
        this.mask.style.height = '100%';
        this.mask.style.zIndex = '1000';
        this.mask.style.opacity = '0.3';
        this.mask.style.display = 'none';

        this.el.nativeElement.appendChild(this.mask);
        this.el.nativeElement.style.transition = 'filter 0.2s linear';
    }

    ngOnChanges(changes: SimpleChanges): void {
        if (this.enabled) {
            this.addMask();
        } else {
            this.removeMask();
        }
    }

    private addMask() {

        this.el.nativeElement.style.filter = 'blur(4px)';
        this.mask.style.display = 'block';


    }

    private removeMask() {
        this.el.nativeElement.style.filter = 'none';
        this.mask.style.display = 'none';
    }

}
