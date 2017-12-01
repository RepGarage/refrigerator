import { RefrigeratorService } from './../refrigerator.service';
import { Observable } from 'rxjs/Observable';
import { Injectable } from '@angular/core';
import { Resolve, ActivatedRouteSnapshot, RouterStateSnapshot } from '@angular/router';
import { Refrigerator } from '../refrigerator';
import { AngularFireObject } from 'angularfire2/database';
import { AuthService } from '../../accounting/auth.service';
@Injectable()
export class SelectedRefrigeratorResolver implements Resolve<Refrigerator> {
  constructor(
    private $refService: RefrigeratorService,
    private $authService: AuthService
  ) {}

  resolve(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): Observable<Refrigerator> {
    const { url } = state;
    const haveChild = /\/r\/refrigerators\/\(.+\)/.test(url);
    if (haveChild) {
      const id = url.substring(27, url.length - 1);
      return this.$refService.fetchRefrigerator(id).take(1);
    } else {
      return null;
    }
  }
}
