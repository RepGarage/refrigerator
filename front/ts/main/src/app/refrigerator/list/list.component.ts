import { Component, OnInit, OnDestroy, Input, EventEmitter, Output } from '@angular/core';
import { RefrigeratorService } from '../refrigerator.service';
import { Router } from '@angular/router';
import { Refrigerator } from '../refrigerator';
import { Observable } from 'rxjs/Observable';

import {
  trigger,
  state,
  style,
  animate,
  transition
} from '@angular/animations';

@Component({
    selector: 'app-list-refrigerators',
    templateUrl: './list.component.html',
    styleUrls: [
        './list.component.sass'
    ],
    animations: [
      trigger('expandState', [
        state('minimized', style({
          gridGap: '5px',
          gridTemplateColumns: '1fr 1fr 1fr'
        })),
        state('expanded', style({
          gridTemplateColumns: '1fr',
          gridGap: '5px',
          maxHeight: '75vh',
          overflow: 'auto'
        }))
      ])
    ]
})
export class RefrigeratorsListComponent implements OnInit, OnDestroy {
  /**
   * Inputs
   */
  @Input()
  refrigerators: Observable<Array<Refrigerator>>;
  @Input()
  expandState: string;
  /**
   * Outputs
   */
  @Output()
  onRefrigeratorClick = new EventEmitter<Refrigerator>();
  /**
   * Public values
   */
  /**
   * Constructor
   */
  constructor(
    private $rs: RefrigeratorService
  ) {}

  /**
   * On Init
   */
  ngOnInit() {
  }

  ngOnDestroy() {
    this.refrigerators = undefined;
  }

  triggerAddRefrigeratorActive() {
    this.$rs.addRefrigeratorActive.next(true);
  }

  selectRefrigerator(refrigerator: Refrigerator) {
    this.onRefrigeratorClick.emit(refrigerator);
  }
}
