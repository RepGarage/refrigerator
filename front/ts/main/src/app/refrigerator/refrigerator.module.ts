import { RefMenuComponent } from './menu/menu.component';
import { MyMaterialModule } from './../my-material.module';
import { RefUnitComponent } from './unit/unit.component';
import { NgModule } from '@angular/core';
import { RefrigeratorsComponent } from './refrigerators.component';
import { RefrigeratorService } from './refrigerator.service';
import { AddRefrigeratorComponent } from './add/add.component';
import { RefrigeratorsListComponent } from './list/list.component';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';

// Modules
import { ReactiveFormsModule, FormsModule } from '@angular/forms';

@NgModule({
    declarations: [
        RefrigeratorsComponent,
        AddRefrigeratorComponent,
        RefrigeratorsListComponent,
        RefUnitComponent,
        RefMenuComponent
    ],
    imports: [
        FormsModule,
        ReactiveFormsModule,
        CommonModule,
        RouterModule,
        MyMaterialModule
    ],
    exports: [
        RefrigeratorsComponent,
        RefMenuComponent
    ],
    providers: [
        RefrigeratorService
    ]
})
export class RefrigeratorModule {

}
