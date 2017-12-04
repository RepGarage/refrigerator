import { Subscription } from 'rxjs/Subscription';
import { Observable } from 'rxjs/Observable';
import { Component, OnInit, ChangeDetectorRef, AfterViewChecked, OnDestroy } from '@angular/core';
import { IRefrigerator, Refrigerator } from './refrigerator';
import { RefrigeratorService } from './refrigerator.service';
import { Routes, ActivatedRoute, Router } from '@angular/router';
import { AngularFireObject } from 'angularfire2/database';

@Component({
    selector: 'app-refrigerators',
    templateUrl: './refrigerators.component.html',
    styleUrls: [
        './refrigerators.component.sass'
    ]
})
export class RefrigeratorsComponent implements OnInit, AfterViewChecked, OnDestroy {
    constructor(
      private $rs: RefrigeratorService,
      private $route: ActivatedRoute,
      private $router: Router
    ) {}

    private subscriptions: Array<Subscription> = [];
    /**
     * Private fields
     */
    private _expandState;
    /**
     * Public fields
     */
    refrigerators: Observable<Array<Refrigerator>>;
    selectedRefrigerator: Refrigerator;
    titleState = 'minimized';
    refZoneState = 'invisible';
    addRefrigeratorActive: Observable<boolean>;

    /**
     * Getters
     */
    get expandState() {
      return this._expandState;
    }

    /**
     * Setters
     */
    set expandState(__state: string) {
      this._expandState = __state;
      switch (__state) {
        case 'expanded':
          this.refZoneState = 'visible';
          this.titleState = 'expanded';
          break;
        case 'minimized':
          this.refZoneState = 'invisible';
          this.titleState = 'minimized';
      }
    }

    ngOnInit() {
      this.selectedRefrigerator = this.$route.snapshot.data['selectedRefrigerator'];
      this.refrigerators = this.$rs.fetchRefrigerators();
      this.expandState = this.$route.snapshot.data['expanded'] ? 'expanded' : 'minimized';
      this.$rs.selectRefrigerator(this.$route.snapshot.data['selectedRefrigerator']);
      this.$rs.selectedRefrigerator.subscribe((r: Refrigerator) => this.selectedRefrigerator = r);
      this.addRefrigeratorActive = this.$rs.addRefrigeratorActive;
    }

    /**
     * Triggrers expand state
     */
    triggerExpand(eState: string) { this.expandState = eState; }

    onRefrigeratorClick(ref: Refrigerator) {
      this.triggerExpand('expanded');
      this.$router.navigate(['/r/refrigerators/', { outlets: { 'products': [ref.key] } }]);
      this.$rs.selectRefrigerator(ref);
    }

    onHeaderTitleClick() {
      this.$router.navigate(['/r/refrigerators']);
      this.triggerExpand('minimized');
    }

    /**
     * On destroy
     */
    ngOnDestroy() {
      this.subscriptions.map(_ => _.unsubscribe());
    }

    ngAfterViewChecked() {
    }
}
