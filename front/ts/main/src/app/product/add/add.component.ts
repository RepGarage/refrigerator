import { Subject } from 'rxjs/Subject';
import { BehaviorSubject } from 'rxjs/BehaviorSubject';
import { Observable } from 'rxjs/Observable';
import { Component, OnInit, Output, EventEmitter, ChangeDetectorRef, ElementRef, AfterViewChecked, ViewChild, OnDestroy } from '@angular/core';
import { FormGroup, FormBuilder, Validators, FormControl, AbstractControl } from '@angular/forms';
import { Product } from '../product';
import { ProductService } from '../product.service';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { SafeUrl, DomSanitizer } from '@angular/platform-browser';
import 'rxjs/add/operator/debounceTime';
import 'rxjs/add/operator/throttleTime';
import 'rxjs/add/observable/fromEvent';
import 'rxjs/add/operator/catch';
import 'rxjs/add/operator/switch';
import 'rxjs/add/operator/share';
import 'rxjs/add/observable/timer';

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
      ])
    ]
})
export class AddProductComponent implements OnInit, OnDestroy, AfterViewChecked {
  readonly LAST_QUERY_RESULT = 'last_query_result';
  readonly QUERY_STATE = 'query_state';
  readonly ADD_PRODUCT_STATE = 'add_product_state';
  query: AbstractControl;
  queryResult = new BehaviorSubject(<Array<Product>>JSON.parse(localStorage.getItem(this.LAST_QUERY_RESULT)));
  queryState = new BehaviorSubject(localStorage.getItem(this.QUERY_STATE));
  state = new BehaviorSubject(localStorage.getItem(this.ADD_PRODUCT_STATE));
  animationState: string;
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

  /**
   * ON INIT
   */
  ngOnInit() {
    /**
     * Form initialization
     */
    this.query = new FormControl();
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
            }).subscribe(v => {
              console.log(v);
              this.queryResult.next(v);
              localStorage.setItem(this.LAST_QUERY_RESULT, JSON.stringify(v));
            });
        });
    this.queryState.subscribe(v => v ? this.query.disable() : this.query.enable());
    this.state.subscribe(s => s ? this.animationState = 'hidden' : this.animationState = 'visible');
  }

  triggerState() {
    this.animationState === 'hidden' ? this.animationState = 'visible' : this.animationState = 'hidden';
  }

  ngAfterViewChecked() {
    this.queryRef.nativeElement.focus();
  }

  ngOnDestroy() {
    localStorage.removeItem(this.LAST_QUERY_RESULT);
  }
}
