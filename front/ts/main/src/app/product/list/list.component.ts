import { Component, OnInit, Input, Output, EventEmitter, ChangeDetectorRef } from '@angular/core';
import { ProductService } from '../product.service';
import { Product } from '../product';
import { Observable } from 'rxjs/Observable';
import { ActivatedRoute } from '@angular/router';
import { Subject } from 'rxjs/Subject';

@Component({
    selector: 'app-product-list',
    templateUrl: './list.component.html',
    styleUrls: ['./list.component.sass']
})
export class ProductListComponent implements OnInit {
  @Input()
  products: Observable<Array<Product>>;

  @Input()
  selectedProduct: Subject<Product> = new Subject();

  @Output()
  productWasSelected = new EventEmitter<Product>();

  constructor(
      private $productService: ProductService,
      private ar: ActivatedRoute,
      private $cdr: ChangeDetectorRef
  ) {}

  ngOnInit() {
  }

  select(product: Product) {
    this.productWasSelected.emit(product);
  }

  triggerAddProduct() {
    this.$productService.setAddProductActive(true);
  }
}
