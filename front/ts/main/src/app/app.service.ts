import { Router } from '@angular/router';
import { Injectable } from '@angular/core';
import { AuthService } from './accounting/auth.service';
import * as io from 'socket.io-client';

@Injectable()
export class AppService {
  constructor(
    private $accService: AuthService,
    private $router: Router
  ) {
    this.pollSocket();
  }

  watchForLoginConponent() {
    return this.$accService.fetchSession()
      .subscribe(user => {
        if (!user) {
          this.$router.navigate(['/login']);
        } else {
          const { url } = this.$router.routerState.snapshot;
          switch (url) {
            case '/login':
              this.$router.navigate(['/']);
              break;
          }
        }
      });
  }

  pollSocket() {}
}
