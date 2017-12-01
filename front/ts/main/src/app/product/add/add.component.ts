import { Subject } from 'rxjs/Subject';
import { BehaviorSubject } from 'rxjs/BehaviorSubject';
import { Observable } from 'rxjs/Observable';
import { Component, OnInit, Output, EventEmitter, ChangeDetectorRef } from '@angular/core';
import { FormGroup, FormBuilder, Validators, FormControl, AbstractControl } from '@angular/forms';
import { Product } from '../product';
import { ProductService } from '../product.service';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { SafeUrl, DomSanitizer } from '@angular/platform-browser';
import 'rxjs/add/operator/debounceTime';
import 'rxjs/add/observable/fromEvent';
import 'rxjs/add/operator/catch';
import 'rxjs/add/operator/switch';
import 'rxjs/add/operator/share';

import {
  trigger,
  state,
  style,
  animate,
  transition
} from '@angular/animations';

@Component({
    selector: 'app-add-product',
    templateUrl: './add.component.html',
    styleUrls: ['./add.component.sass']
})
export class AddProductComponent implements OnInit {
  @Output()
  newProduct = new EventEmitter<Product>();
  /**
   * Public values
   */
  productName: AbstractControl;
  availableProducts: Observable<Array<Object>>;
  fGroup: FormGroup;
  selectedProductFormGroup: FormGroup;
  selectedProduct: BehaviorSubject<Product> = new BehaviorSubject(null);
  fetchedSelectedProduct: BehaviorSubject<Product> = new BehaviorSubject(null);

  constructor(
    private $productService: ProductService,
    private fb: FormBuilder,
    private $http: HttpClient,
    private $cdr: ChangeDetectorRef,
    private $sanitizer: DomSanitizer
  ) {}

  ngOnInit() {
    this.fGroup = this.fb.group({
        name: [ '', Validators.required ],
        created: [ this.getDayRelativeToNow(-1), Validators.required ],
        shelfLife: [ this.getDayRelativeToNow(1), Validators.required ],
        movedToRef: [ this.getDayRelativeToNow(), Validators.required ]
    });
    this.productName = new FormControl();
    this.fetchPerekProducts();
  }

  /**
   * Return date, relevant to now by number of days
   * @param move number of days to switch
   */
  getDayRelativeToNow(move: number = 0) {
      return new Date(new Date().getTime() + move * 100000000);
  }

  fetchPerekProducts() {
    this.availableProducts = this.productName.valueChanges
      .debounceTime(300)
      .switchMap((name: string) => {
        return this.$productService.fetchProductsListFromApi(name)
          .map((result: Array<Product>) => {
            result.map((p: Product) => {
              this.$productService.fetchProductImageFromAPI(p, 112)
                .subscribe(v => {
                  p.photoUrl = this.$sanitizer.bypassSecurityTrustUrl(v);
                  return v;
                });
            });
            return result;
          });
      }).share();
    this.availableProducts.subscribe(console.log);
  }

  /**
   * Make product selected
   * @param product product
   */
  selectProduct(product: Product) {
    if (product) {
      console.log(product);
      this.selectedProductFormGroup = this.fb.group({
        shelf: [this.getDayRelativeToNow(1), Validators.required],
        created: [this.getDayRelativeToNow(-1), Validators.required],
        moved: [this.getDayRelativeToNow(0), Validators.required]
      });
      const selectedProduct = new Product({name: product.name, product_id: product.product_id});
      this.selectedProduct.next(product);
    } else {
      this.selectedProductFormGroup = null;
      this.selectedProduct.next(null);
    }
  }

  /**
   * Add new product
   * @param p product
   */
  addProduct(p: Product) {
    const name = p.name;
    const photoUrl = p.photoUrl['changingThisBreaksApplicationSecurity'];
    const product_id = p.product_id;
    const created = this.selectedProductFormGroup.get('created').value.toString();
    const shelf = this.selectedProductFormGroup.get('shelf').value.toString();
    const moved = this.selectedProductFormGroup.get('moved').value.toString();
    const newProduct = new Product({
      photoUrl,
      name,
      created,
      moved,
      shelf,
      product_id
    });

    this.$productService.addProduct(newProduct);
    this.selectProduct(null);
  }

  /**
   * Make add product unactive
   */
  fadeToBlack() {
    this.$productService.addProductActive.next(false);
  }
}
