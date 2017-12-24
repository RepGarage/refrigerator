import { RefrigeratorService } from './../refrigerator/refrigerator.service';
import { User } from './../accounting/user/user.class';
import { Refrigerator } from './../refrigerator/refrigerator';
import { Injectable, Inject } from '@angular/core';
import { Http } from '@angular/http';
import { Observable } from 'rxjs/Observable';
import { Subject } from 'rxjs/Subject';
import { Subscription } from 'rxjs/Subscription';
import { Product } from './product';
import {
    AngularFireDatabase,
    AngularFireList,
    AngularFireObject
} from 'angularfire2/database';
import { AuthService } from '../accounting/auth.service';
import { BehaviorSubject } from 'rxjs/BehaviorSubject';
import { HttpClient } from '@angular/common/http';

@Injectable()
export class ProductService {
  readonly ADD_PRODUCT_ACTIVE = 'add_product_active';
  private _refrigeratorRef: Observable<AngularFireObject<Refrigerator>>;
  private user: User;
  private selectedRefrigerator: Refrigerator;
  private selectedProduct: BehaviorSubject<Product> = new BehaviorSubject(null);
  private addProductActive = new BehaviorSubject(Boolean(localStorage.getItem(this.ADD_PRODUCT_ACTIVE)));

  constructor(
      private $afd: AngularFireDatabase,
      private $auth: AuthService,
      private $refService: RefrigeratorService,
      private $http: HttpClient,
      @Inject('PRODUCTS_BASE_URL') private productsBaseUrl: string
  ) {
    this.$auth.fetchSession()
      .subscribe((u: User) => {
        this.user = u;
      });
    this.$refService.selectedRefrigerator.subscribe((r: Refrigerator) => this.selectedRefrigerator = r);
  }

  /**
   * Returns products list reference
   * @param ref_id
   */
  fetchProdudctsList(ref_id: string): Observable<AngularFireList<Product>> {
    if (this.user) {
      return Observable.of(this.$afd.list(`/refrigerators/${this.user.uid}/${ref_id}/products`));
    } else {
      return null;
    }
  }

  fetchAddProductActive() {
    return this.addProductActive;
  }

  /**
   * Fetch products
   */
  fetchProducts(): Observable<Array<Product>> {
    if (this.user) {
      return this.$refService.selectedRefrigerator
        .switchMap((r: Refrigerator) => {
          return this.$afd.object(`/refrigerators/${this.user.uid}/${r.key}/products`).valueChanges()
          .map(value => {
            if (value) {
              const result = [];
              const obj = <Object>value;
              Object.keys(obj).map((key, i) => {
                result.push(
                  new Product({
                    key: key,
                    photoUrl: value[key]['photoUrl'],
                    name: value[key]['name'],
                    created: value[key]['created'],
                    shelf: value[key]['shelf'],
                    exp: value[key]['exp'],
                    moved: value[key]['moved']
                  })
                );
              });
              return result;
            } else {
              return [];
            }
          });
        });
    } else {
      return Observable.of(undefined);
    }
  }

  /**
   * Trigger add product state
   */
  triggerAddProductState() {
    this.addProductActive.next(!this.addProductActive.getValue());
    this.addProductActive.getValue()
      ? localStorage.setItem(this.ADD_PRODUCT_ACTIVE, 'true')
      : localStorage.removeItem(this.ADD_PRODUCT_ACTIVE);
  }

  /**
   * Set certain add product state
   * @param s state
   */
  setAddProductActive(s: boolean) {
    this.addProductActive.next(s);
    s ? localStorage.setItem(this.ADD_PRODUCT_ACTIVE, 'true') : localStorage.removeItem(this.ADD_PRODUCT_ACTIVE);
  }

  /**
   * Add new product to firebase
   * @param product Product object
   */
  addProduct(product: Product): void {
    if (this.user) {
      this.$afd.list(`/refrigerators/${this.user.uid}/${this.selectedRefrigerator.key}/products`).push(product);
      this.addProductActive.next(false);
    } else {
      return;
    }
  }

  /**
   * Remove product from firebase
   * @param product Product object
   */
  removeProduct(product: Product): void {
    if (this.user) {
      this.$afd.list(`/refrigerators/${this.user.uid}/${this.selectedRefrigerator.key}/products`).remove(product.key);
      this.selectProduct(undefined);
    } else {
      return;
    }
  }

  updateProduct(product: Product): void {
    if (this.user) {
      this.$afd.object(`/refrigerators/${this.user.uid}/${this.selectedRefrigerator.key}/products/${product.key}`)
        .update(product);
    } else {
      return;
    }
  }

  selectProduct(p: Product) {
    this.selectedProduct.next(p);
  }

  fetchSelectedProduct(): Observable<Product> {
    return this.selectedProduct;
  }

  /**
   * Search for products
   * @param name product query
   */
  fetchProductsListFromApi(name: string): Observable<Array<Product>> {
    return this.$http.get(
      `${this.productsBaseUrl}api/get/products?name=${name}`
    ).map((r: Array<{_id: number, name: string}>) => {
      return r.map(v => new Product({name: v.name, product_id: v._id.toString()}));
    })
    .catch(e => Observable.of([]));
  }

  /**
   * Get product image
   * @param p product
   * @param side size
   */
  fetchProductImageFromAPI(p: Product, side: number): Observable<string> {
    return this.$http.get(
      `${this.productsBaseUrl}api/get/product/image?product_id=${p.product_id}&side=${side}`)
      .map((response: { result: string}) => <string>response.result)
      .catch(() => Observable.of(null));
  }

  fetchProductShelflife(id: string): Observable<number> {
    return this.$http.get(`${this.productsBaseUrl}api/get/product/shelflife?product_id=${id}`)
      .map((response: { result: number }) => response.result)
      .catch(() => Observable.of(0));
  }

}
