import { DomSanitizer, SafeUrl } from '@angular/platform-browser';
import { Pipe, PipeTransform } from '@angular/core';

@Pipe({name: 'safe'})
export class SafePipe implements PipeTransform {
  constructor(private $dom: DomSanitizer) {}

  transform(v: string): SafeUrl {
    return this.$dom.bypassSecurityTrustUrl(v);
  }
}
