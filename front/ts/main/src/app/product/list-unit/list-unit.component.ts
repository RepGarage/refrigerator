import { Subscription } from 'rxjs/Subscription';
import { Product } from './../../product/product';
import { Component, Input, OnInit, OnDestroy } from '@angular/core';
import { DomSanitizer } from '@angular/platform-browser';
import { ProductService } from '../product.service';

@Component({
    selector: 'app-prod-list-unit',
    templateUrl: './list-unit.component.html',
    styleUrls: [ './list-unit.component.sass' ]
})
export class ListUnitComponent {
    p: Product;
    subscription: Subscription;

    constructor(
    ) {}

    @Input()
    set product(product: Product) {
        if (product.photoUrl) {
            product.photoUrl = product.photoUrl;
        } else {
            product.photoUrl = '/assets/icons/groceries.svg';
        }
        this.p = product;
    }

    @Input()
    textWhite = false;

    get product() {
        return this.p;
    }
}
