import { AuthService } from './auth.service';
import { Injectable } from '@angular/core';
import { CanActivate, Router, NavigationEnd, GuardsCheckStart } from '@angular/router';
import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/pairwise';

@Injectable()
export class AuthGuard implements CanActivate {
  constructor(
    private $authService: AuthService,
    private $router: Router
  ) {}

  canActivate(): Observable<boolean> {
    return this.$authService.fetchSession()
      .switchMap(user => {
        if (user) {
          return Observable.of(true);
        } else {
          this.$router.navigate(['/login']);
          return Observable.of(false);
        }
      });
  }
}
