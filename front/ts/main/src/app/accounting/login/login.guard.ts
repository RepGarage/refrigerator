import { Injectable } from '@angular/core';
import { Router, CanActivate } from '@angular/router';
import { Observable } from 'rxjs/Observable';
import { AuthService } from '../auth.service';
@Injectable()
export class LoginGuard implements CanActivate {
  constructor(
    private $router: Router,
    private $authService: AuthService
  ) {}

  canActivate(): Observable<boolean> {
   return this.$authService.fetchSession()
      .switchMap(user => {
        if (user) {
          this.$router.navigate(['/']);
          return Observable.of(false);
        } else {
          return Observable.of(true);
        }
      });
  }
}
