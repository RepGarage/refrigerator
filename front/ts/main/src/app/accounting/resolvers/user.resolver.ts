import { AuthService } from './../auth.service';
import { Injectable } from '@angular/core';
import { Resolve } from '@angular/router';
import { User } from '../user/user.class';
import { Observable } from 'rxjs/Observable';
@Injectable()
export class UserResolver implements Resolve<User> {
  constructor(
    private $auth: AuthService
  ) {}

  resolve(): Observable<User> {
    return this.$auth.fetchSession()
      .switchMap((u: User) => {
        return Observable.of(u);
      }).take(1);
  }
}
