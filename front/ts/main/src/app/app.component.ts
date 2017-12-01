import { User } from './accounting/user/user.class';
import { ProductService } from './product/product.service';
import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs/Observable';
import { Router, NavigationStart } from '@angular/router';
import { AppService } from './app.service';
import { AuthService } from './accounting/auth.service';
import { Product } from './product/product';
import {
  trigger,
  state,
  style,
  animate
} from '@angular/animations';

import 'rxjs/add/operator/map';
import 'rxjs/add/operator/filter';
import 'rxjs/add/operator/switchMap';
import 'rxjs/add/operator/distinct';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.sass'],
  animations: [
    trigger('selectedProductState', [
      state('selected', style({
        gridTemplateRows: '70% 30%'
      })),
      state('none', style({
        gridTemplateRows: '100% 1fr'
      }))
    ])
  ]
})
export class AppComponent implements OnInit {
  selectedProduct: Observable<Product>;
  authState: Observable<User>;
  selectedProductState = 'none';

  constructor(
    private router: Router,
    private $appService: AppService,
    private $authService: AuthService,
    private $productService: ProductService
  ) {}

  ngOnInit() {
    this.authState = this.$authService.fetchSession();
    this.addSubscriptions();
  }

  addSubscriptions() {
    this.$appService.watchForLoginConponent();
    this.selectedProduct = this.$productService.fetchSelectedProduct();
    this.selectedProduct.subscribe((p: Product) => {
      if (p) {
        this.triggerSelectedProduct('selected');
      } else {
        this.triggerSelectedProduct('none');
      }
    });
  }

  logout() {
    this.$authService.logout();
  }

  triggerSelectedProduct(__state: string) {
    this.selectedProductState = __state;
  }

  removeProduct(p: Product) {
    this.$productService.removeProduct(p);
  }
}
