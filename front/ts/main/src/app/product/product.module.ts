import { RouterModule } from '@angular/router';
import { ListUnitComponent } from './list-unit/list-unit.component';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

// Components
import { ProductListComponent } from './list/list.component';
import { AddProductComponent } from './add/add.component';
import { SelectedProductComponent } from './selected/selected.component';
import { ProductsComponent } from './products.component';

// Modules
import { ReactiveFormsModule, FormsModule } from '@angular/forms';
import { MyMaterialModule } from '../my-material.module';


@NgModule({
    imports: [
        CommonModule,
        FormsModule,
        ReactiveFormsModule,
        MyMaterialModule,
        RouterModule
    ],
    declarations: [
        ProductListComponent,
        AddProductComponent,
        SelectedProductComponent,
        ProductsComponent,
        ListUnitComponent
    ],
    exports: [
        ProductsComponent,
        SelectedProductComponent
    ]
})
export class ProductModule {

}
