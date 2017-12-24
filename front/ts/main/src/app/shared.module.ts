import { AppModalViewComponent } from './modal/modal.component';
import { NgModule, ModuleWithProviders } from '@angular/core';
import 'rxjs/add/operator/debounceTime';
import 'rxjs/add/operator/throttleTime';
import 'rxjs/add/observable/fromEvent';
import 'rxjs/add/operator/catch';
import 'rxjs/add/operator/switch';
import 'rxjs/add/operator/share';
import 'rxjs/add/observable/timer';
import 'rxjs/add/operator/switchMap';
import 'rxjs/add/observable/of';

@NgModule({
  declarations: [
    AppModalViewComponent
  ],
  exports: [
    AppModalViewComponent
  ]
})
export class SharedModule {
  static forRoot(): ModuleWithProviders {
    return {
      ngModule: SharedModule
    };
  }
}
