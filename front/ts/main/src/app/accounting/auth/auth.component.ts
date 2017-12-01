import { AuthService } from './../auth.service';
import { Component, OnInit } from '@angular/core';
import * as firebase from 'firebase/app';
import { User } from 'firebase/auth';

import { Observable } from 'rxjs/Observable';

@Component({
    selector: 'app-auth',
    templateUrl: './auth.component.html',
    styleUrls: [
        './auth.component.sass'
    ]
})
export class AuthComponent implements OnInit {
    constructor(
      private $authService: AuthService
    ) {}

    ngOnInit() {
    }

    auth() {
        this.$authService.authenticate();
    }

    logout() {
        this.$authService.logout();
    }
}
