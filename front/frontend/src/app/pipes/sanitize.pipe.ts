import { Pipe, PipeTransform } from '@angular/core';
import { DomSanitizer } from '@angular/platform-browser';

@Pipe({ name: 'sanitize', pure: false })
export class sanitizePipe implements PipeTransform {
    constructor(private sanitizer: DomSanitizer) {
    }

    transform(content: string, type: string) {

        console.log('pipe: ', type)

        if (type == 'html')
            return this.sanitizer.bypassSecurityTrustHtml(content);
        if (type == 'url')
            return this.sanitizer.bypassSecurityTrustUrl(content);
        if (type == 'resourceUrl')
            return this.sanitizer.bypassSecurityTrustResourceUrl(content);
        return '';
    }
}