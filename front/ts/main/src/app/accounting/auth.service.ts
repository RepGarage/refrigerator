import { User } from './user/user.class';
import { firebaseSecret } from './../../environments/.secret';
import { Injectable } from '@angular/core';
import { AngularFireAuth } from 'angularfire2/auth';
import { AngularFireDatabase, AngularFireList } from 'angularfire2/database';
import * as firebase from 'firebase/app';
import { Auth } from 'firebase/auth';
import { Observable } from 'rxjs/Observable';
import { Subject } from 'rxjs/Subject';
import { Subscription } from 'rxjs/Subscription';
import { IRefUser, RefUser, UserQuery } from './refrigerator.user';
import { BehaviorSubject } from 'rxjs/BehaviorSubject';

export interface AccountState {
    userRef: AngularFireList<IRefUser>;
}


@Injectable()
export class AuthService {
  private _userRef: AngularFireList<UserQuery>;
  private subscriptions: Array<Subscription> = new Array();
  private session: Observable<Auth>;

  constructor(
      private $afa: AngularFireAuth,
      private $afd: AngularFireDatabase
    ) {
    this.session = this.$afa.authState
      .switchMap((u: User) => {
        return Observable.of(u);
      });
    this.$afa.authState
      .subscribe((u: User) => {
        if (u) {
          this.subscriptions.push(
            Observable.merge(
              this.$afd.object(`/users/${u.uid}`).snapshotChanges(),
              this.$afd.object(`/users/${u.uid}`).valueChanges()
            ).pairwise()
            .subscribe(value => {
              const user = value[1];
              if (!user) {
                this.$afd.object(`/users/${u.uid}`).set(
                  new User(u.uid, u.displayName, u.photoURL, u.email)
                );
              }
            })
          );
        } else {
          this.subscriptions.map(s => s.unsubscribe());
        }
      });
    }

    authenticate(): Observable<Auth> {
      return Observable.fromPromise(this.$afa.auth.signInWithPopup(new firebase.auth.GoogleAuthProvider()));
    }

    logout(): Observable<Auth> {
      return Observable.fromPromise(this.$afa.auth.signOut());
    }

    fetchSession(): Observable<Auth> {
      return this.session;
    }

    addSubscriptions(state: User) {
    }

    deleteSubscriptions() {
        this.subscriptions.map(s => {
            s.unsubscribe();
            this.subscriptions = this.subscriptions.slice(1, this.subscriptions.length);
        });
    }

}
