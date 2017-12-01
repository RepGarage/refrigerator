import { AuthService } from './../../accounting/auth.service';
import { RefrigeratorService } from './../refrigerator.service';
import { Refrigerator } from './../refrigerator';
import { User } from 'firebase/auth';
import { AngularFireDatabase } from 'angularfire2/database';
import { Observable } from 'rxjs/Observable';
import { Injectable } from '@angular/core';

import { Resolve, Router, ActivatedRouteSnapshot, RouterStateSnapshot } from '@angular/router';
import 'rxjs/add/observable/interval';
import 'rxjs/add/operator/take';

@Injectable()
export class RefrigeratorsResolver implements Resolve<Array<Refrigerator>> {
  constructor(
    private $refService: RefrigeratorService,
    private $auth: AuthService
  ) {}

  resolve(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): Observable<Array<Refrigerator>> {
    return this.$refService.fetchRefrigerators();
  }
}
