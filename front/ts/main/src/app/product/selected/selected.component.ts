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
    state = 'hidden';
    product: Observable<Product>;
    constructor(
        private $productService: ProductService
    ) {}

    ngOnInit() {
        this.state = 'visible';
        this.product = this.$productService.fetchSelectedProduct();
    }

    remove() {
        this.product.subscribe(p => this.$productService.removeProduct(p));
    }

    destroy() {
        this.state = 'hidden';
        this.$productService.selectProduct(null);
    }
}
