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
    constructor(
        private $productService: ProductService
    ) {}

    ngOnInit() {
        this.state = 'visible';
    }

    destroy() {
        this.state = 'hidden';
        this.$productService.selectProduct(null);
    }
}
