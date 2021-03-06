import { ActivatedRoute } from '@angular/router';
import { AuthService } from './../accounting/auth.service';
import { Observable } from 'rxjs/Observable';
import { Injectable } from '@angular/core';
import { Refrigerator } from './refrigerator';
import {
    AngularFireDatabase,
    AngularFireList,
    AngularFireObject
} from 'angularfire2/database';
import { RefUser, IRefUser } from '../accounting/refrigerator.user';
import { Subscription } from 'rxjs/Subscription';
import { Subject } from 'rxjs/Subject';
import { User } from '../accounting/user/user.class';
import { assets } from '../../assets/assets';
import { BehaviorSubject } from 'rxjs/BehaviorSubject';

@Injectable()
export class RefrigeratorService {
    readonly SELECTED_REFRIGERATOR = 'selected_refrigerator';
    private refrigeratorsRef: AngularFireList<Refrigerator>;
    private _refrigeratorsObservable: Observable<Array<Refrigerator>>;
    private _refrigeratorsIds: Object;
    private user: User;
    selectedRefrigerator: BehaviorSubject<Refrigerator> = new BehaviorSubject(JSON.parse(localStorage.getItem(this.SELECTED_REFRIGERATOR)));
    subscriptions: Array<Subscription> = new Array();
    addRefrigeratorActive: BehaviorSubject<boolean> = new BehaviorSubject(false);

    get refrigeratorsIds() {
        return this._refrigeratorsIds;
    }

    refrigeratorsCount = new BehaviorSubject(localStorage.getItem('refrigerators_count'));

    constructor(
        private $afd: AngularFireDatabase,
        private $auth: AuthService,
        private $route: ActivatedRoute
    ) {
      this.$auth.fetchSession().subscribe((u: User) => {
        this.user = u;
        if (u) {
          this.refrigeratorsRef = this.$afd.list(`/refrigerators/${u.uid}`);
        } else {
          this.refrigeratorsRef = undefined;
        }
      });
    }

    selectRefrigerator(ref: Refrigerator) {
      console.log('Select refrigerator');
      this.selectedRefrigerator.next(ref);
      localStorage.setItem(this.SELECTED_REFRIGERATOR, JSON.stringify(ref));
    }

    /**
     * Fetch all refrigerators as array
     */
    fetchRefrigerators(): Observable<Array<Refrigerator>> {
      if (this.user) {
        return this.$afd.object(`/refrigerators/${this.user.uid}`).valueChanges()
        .map((value: Object) => {
          if (!value) { return []; }
          const result: Array<Refrigerator> = [];
          Object.keys(value)
            .map((k, i) => {
              result.push(
                new Refrigerator({
                    key: k,
                    name: value[k]['name'],
                    products: value[k]['products'],
                    archivedProducts: value[k]['archivedProducts'],
                    iconAssetUrl: value[k]['iconAssetUrl'],
                    photo: value[k]['photo']
                  })
              );
            });
          return result;
        });
      } else {
        return Observable.of([<Refrigerator>{}]);
      }
    }

    setRefrigeratorsCount(c: number) {
      localStorage.setItem('refrigerators_count', c.toString());
      this.refrigeratorsCount.next(c.toString());
    }

    private getRefrigerator(uid: string, ref_id: string): Observable<Refrigerator> {
      return Observable.merge(
        this.$afd.object(`/refrigerators/${uid}/${ref_id}`).snapshotChanges(),
        this.$afd.object(`/refrigerators/${uid}/${ref_id}`).valueChanges()
      ).pairwise()
      .map(value => {
        return new Refrigerator({
          key: value[0]['key'],
          name: value[1]['name'],
          products: value[1]['products'],
          archivedProducts: value[1]['archivedProducts'],
          iconAssetUrl: value[1]['iconAssetUrl'],
          photo: value[1]['photo']
        });
      });
    }

    fetchRefrigerator(ref_id: string): Observable<Refrigerator> {
      if (this.user) {
        return this.getRefrigerator(this.user.uid, ref_id);
      } else {
        return this.$auth.fetchSession()
          .take(1)
          .switchMap((u: User) => {
            if (u) {
              return this.getRefrigerator(u.uid, ref_id);
            } else {
              return Observable.of(null);
            }
          });
      }
    }

    fetchRefrigeratorRef(ref_id: string): Observable<AngularFireObject<Refrigerator>> {
      return this.$auth.fetchSession()
        .switchMap((user: User) => {
          if (user) {
            return Observable.of(this.$afd.object(`/refrigerators/${user.uid}/${ref_id}`));
          } else {
            return Observable.of(undefined);
          }
        });
    }

    updateRef(ref: Refrigerator) {
      this.$auth.fetchSession()
        .subscribe((user: User) => {
          if (user) {
            this.$afd.object(`/refrigerators/${user.uid}/${ref.key}`)
              .update(ref);
          }
        });
    }

    updateRefPhoto(photo: string, key: string) {
      this.$auth.fetchSession()
        .subscribe((user: User) => {
          if (user) {
            this.$afd.object(`/refrigerators/${user.uid}/${key}/photo`).set(photo);
          }
        });
    }

    updateRefName(name: string, key: string) {
      this.$auth.fetchSession()
        .subscribe((user: User) => {
          if (user) {
            this.$afd.object(`/refrigerators/${user.uid}/${key}/name`).set(name);
          }
        });
    }

    deleteCurrenRef(key: string) {
      this.selectRefrigerator(null);
      this.$auth.fetchSession()
        .subscribe((user: User) => {
          if (user) {
            this.$afd.object(`/refrigerators/${user.uid}/${key}`).remove();
          }
        });
    }

    addSubscriptions() {
    }

    addRefrigerator(ref: Refrigerator) {
      const rnd: number = Math.round(Math.random() * 2 + 1);
      const asset = assets.icons.refrigerators[`refrigerator${rnd}`];
      ref.iconAssetUrl = asset;
      if (this.refrigeratorsRef) {
        this.refrigeratorsRef.push(ref);
        this.addRefrigeratorActive.next(false);
      }
    }
}
