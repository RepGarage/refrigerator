import { DomSanitizer } from '@angular/platform-browser';
import { RefrigeratorService } from './../refrigerator/refrigerator.service';
import { Refrigerator } from './../refrigerator/refrigerator';
import { Router, ActivatedRoute } from '@angular/router';
import { AngularFireList, AngularFireDatabase, AngularFireObject } from 'angularfire2/database';
import { Component, OnInit, OnDestroy } from '@angular/core';
import { Product } from './product';
import { Observable } from 'rxjs/Observable';
import { ProductService } from './product.service';
import { Subscription } from 'rxjs/Subscription';
import { Subject } from 'rxjs/Subject';
import 'rxjs/add/operator/pluck';

@Component({
    selector: 'app-products-component',
    templateUrl: './products.component.html',
    styleUrls: [
        './products.component.sass'
    ]
})
export class ProductsComponent implements OnInit, OnDestroy {
  private ref_id: string;
  private refIdSubscription: Subscription;
  private productsRef: AngularFireList<Product>;

  products: Observable<Array<Product>>;
  selectedProduct: Observable<Product>;
  addProductActive: Observable<boolean>;
  selectedRefrigerator: Observable<Refrigerator>;

  constructor(
    private $route: ActivatedRoute,
    private $afd: AngularFireDatabase,
    private $rs: RefrigeratorService,
    private $dom: DomSanitizer,
    private $productService: ProductService,
    private $refService: RefrigeratorService
  ) {}

  ngOnInit() {
    this.selectedRefrigerator = this.$rs.selectedRefrigerator;
    this.products = this.$productService.fetchProducts()
      .map(
        (prods: Array<Product>) => prods.map(
          p => { p.photoUrl = this.$dom.bypassSecurityTrustUrl(<string>p.photoUrl); return p; }));
    this.selectedProduct = this.$productService.fetchSelectedProduct();
    this.addProductActive = this.$productService.addProductActive;
  }

  ngOnDestroy() {
    // delete subscriptions here
  }

  addProduct(product: Product) {
    this.$productService.addProduct(product);
  }

  selectProduct(product: Product) {
    this.$productService.selectProduct(product);
  }

  removeProduct(product: Product) {
    this.$productService.removeProduct(product);
    this.$productService.selectProduct(undefined);
  }

}
