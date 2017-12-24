import * as moment from 'moment';
import { Subject } from 'rxjs/Subject';
import { BehaviorSubject } from 'rxjs/BehaviorSubject';
import { Observable } from 'rxjs/Observable';
import {
  Component,
  OnInit,
  Output,
  EventEmitter,
  ChangeDetectorRef,
  ElementRef,
  AfterViewInit,
  ViewChild,
  OnDestroy
} from '@angular/core';
import { FormGroup, FormBuilder, Validators, FormControl, AbstractControl } from '@angular/forms';
import { Product } from '../product';
import { ProductService } from '../product.service';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { SafeUrl, DomSanitizer } from '@angular/platform-browser';
import { AddState } from './add.enum';

import {
  trigger,
  state,
  style,
  animate,
  transition
} from '@angular/animations';
import { switchMap } from 'rxjs/operator/switchMap';

@Component({
    selector: 'app-add-product',
    templateUrl: './add.component.html',
    styleUrls: ['./add.component.sass'],
    animations: [
      trigger('state', [
        state('visible', style({
          transform: 'translateY(0)'
        })),
        state('hidden', style({
          transform: 'translateY(100%)'
        })),
        transition('* => *', animate('200ms ease-in'))
      ]),
      trigger('modalState', [
        state('visible', style({
          opacity: '0.7'
        })),
        state('hidden', style({
          opacity: '0'
        })),
        transition('* => *', animate('200ms ease-in-out'))
      ])
    ]
})
export class AddProductComponent implements OnInit, OnDestroy, AfterViewInit {
  readonly LAST_QUERY_RESULT = 'last_query_result';
  readonly QUERY_STATE = 'query_state';
  readonly ADD_PRODUCT_STAE = 'add_product_state';
  readonly SELECTED_PRODUCT = 'add_product_selected_product';
  readonly addState = AddState;
  query: AbstractControl;
  created: AbstractControl;
  exp = 'не рассчитано';
  photoURL: string;
  queryResult = new BehaviorSubject(<Array<Product>>JSON.parse(localStorage.getItem(this.LAST_QUERY_RESULT)));
  queryState = new BehaviorSubject(localStorage.getItem(this.QUERY_STATE));
  state = AddState.SEARCH;
  selectedProduct = new BehaviorSubject<Product>(JSON.parse(localStorage.getItem(this.SELECTED_PRODUCT)));
  animationState = 'hidden';
  modalState = 'hidden';

  set viewState(s: string) {
    this.animationState = s;
    this.modalState = s;
  }
  @ViewChild('queryEl')
  private queryRef: ElementRef;
  /**
   * CONSTRUCTOR
   */
  constructor(
    private $ps: ProductService,
    private $fb: FormBuilder,
    private $dom: DomSanitizer
  ) {}

  backToSearch() {
    this.selectedProduct.next(null);
  }

  triggerState() {
    this.animationState === 'hidden' ? this.animationState = 'visible' : this.animationState = 'hidden';
  }

  selectProduct(p: Product) {
    this.selectedProduct.next(p);
  }

  destroy() {
    this.$ps.setAddProductActive(false);
    localStorage.removeItem(this.LAST_QUERY_RESULT);
    localStorage.removeItem(this.ADD_PRODUCT_STAE);
    localStorage.removeItem(this.SELECTED_PRODUCT);
  }

  addProduct() {
    const product = this.selectedProduct.getValue();
    product.photoUrl = this.photoURL;
    this.$ps.addProduct(product);
    this.destroy();
  }

  subscribeCreated() {
    this.created.valueChanges
      .subscribe(v => {
        this.exp = this.calculateExp(v);
        const product = this.selectedProduct.getValue();
        product.exp = v;
        this.selectedProduct.next(product);
      });
  }

  calculateExp(created: string): string {
    const from = moment(new Date(created).toISOString());
    const to = moment();
    const diff = to.diff(from, 'days');
    const product = this.selectedProduct.getValue();
    const shelf = product.shelf ? product.shelf : null;
    if (!shelf) { return; }
    const exp = shelf - diff;
    return exp.toString();
  }

  /**
   * LIFECYCLE HOOKS
   */
  ngAfterViewInit() {
    if (this.queryRef) { this.queryRef.nativeElement.focus(); }
  }

  ngOnDestroy() {
    this.viewState = 'hidden';
  }

  /**
   * ON INIT
   */
  ngOnInit() {
    /**
     * Form initialization
     */
    const product = this.selectedProduct.getValue();
    this.query = new FormControl();
    this.created = new FormControl(product ? new Date(product.exp) : '', Validators.compose([Validators.required]));
    this.subscribeCreated();
    this.queryResult.subscribe(
      v => v ? v.map(
        p => this.$ps.fetchProductImageFromAPI(p, 112).subscribe(image => p.photoUrl = this.$dom.bypassSecurityTrustUrl(image))) : '');
    Observable.timer(200).take(1)
        .subscribe(() => {
          this.query.valueChanges
            .debounceTime(200)
            .throttleTime(200)
            .switchMap(v => {
              return this.$ps.fetchProductsListFromApi(v);
            })
            .subscribe(v => {
              console.log(v);
              this.queryResult.next(v);
              localStorage.setItem(this.LAST_QUERY_RESULT, JSON.stringify(v));
            });
        });
    this.queryState.subscribe(v => v ? this.query.disable() : this.query.enable());
    this.viewState = 'visible';
    this.selectedProduct.subscribe(p => {
      this.state = p ? AddState.ADD : AddState.SEARCH;
      localStorage.setItem(this.SELECTED_PRODUCT, JSON.stringify(p));
      if (p) {
        this.exp = p.exp ? this.calculateExp(p.exp) : 'не рассчитано';
        p.photoUrl = null;
        this.$ps.fetchProductImageFromAPI(p, 112).subscribe(image => {
          p.photoUrl = this.$dom.bypassSecurityTrustUrl(image);
          this.photoURL = image;
        });
        this.$ps.fetchProductShelflife(p.product_id).subscribe(shelf => p.shelf = shelf);
      }
      console.log(p);
    });
  }

}
