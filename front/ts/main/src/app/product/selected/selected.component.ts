import { AbstractControl, FormControl } from '@angular/forms';
import { ProductService } from './../product.service';
import { Observable } from 'rxjs/Observable';
import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';
import { Product } from '../product';
import { DatePipe } from '@angular/common';

@Component({
    selector: 'app-selected-product',
    templateUrl: './selected.component.html',
    styleUrls: [
        './selected.component.sass'
    ],
    providers: [
      DatePipe
    ]
})
export class SelectedProductComponent implements OnInit {
    @Input()
    product: Product;

    createdControl: AbstractControl;
    shelfControl:   AbstractControl;
    movedControl:   AbstractControl;

    createdChanging = false;
    shelfChanging   = false;
    movedChanging   = false;

    constructor(
        private $productService: ProductService
    ) {}

    ngOnInit() {
        this.createdControl = new FormControl(new Date(this.product.created));
        this.shelfControl   = new FormControl(new Date(this.product.shelf));
        this.movedControl   = new FormControl(new Date(this.product.moved));
    }

    updateProduct(from: string) {
        const created = this.createdControl.value;
        const shelf = this.shelfControl.value;
        const moved = this.movedControl.value;
        const key = this.product.key;
        const photoUrl = this.product.photoUrl;
        const product_id = this.product.product_id;
        const name = this.product.name;
        const newProduct = new Product({
            photoUrl,
            name,
            created,
            moved,
            shelf,
            product_id,
            key
        });
        this.product = newProduct;
        this.$productService.updateProduct(newProduct);
        this.triggerChangingState(from);
    }

    removeProduct() {
        this.$productService.removeProduct(this.product);
    }

    triggerChangingState(name: string) {
        switch (name) {
            case 'created':
                this.createdChanging = !this.createdChanging;
                break;
            case 'shelf':
                this.shelfChanging = !this.shelfChanging;
                break;
            case 'moved':
                this.movedChanging = !this.movedChanging;
                break;
        }
    }

    fadeToBlack() {
        this.$productService.selectProduct(null);
    }
}
