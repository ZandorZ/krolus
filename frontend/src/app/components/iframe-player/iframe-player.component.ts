import { Component, Input, OnChanges, OnInit, SimpleChanges } from '@angular/core';
import { DomSanitizer, SafeResourceUrl } from '@angular/platform-browser';

@Component({
    selector: 'app-iframe-player',
    templateUrl: './iframe-player.component.html',
    styleUrls: ['./iframe-player.component.scss']
})
export class IframePlayerComponent implements OnChanges {

    @Input()
    url: string;
    urlSafe: SafeResourceUrl;

    constructor(private sanitizer: DomSanitizer) { }

    ngOnChanges(changes: SimpleChanges): void {
        this.urlSafe = this.sanitizer.bypassSecurityTrustResourceUrl(this.url);
    }


}
