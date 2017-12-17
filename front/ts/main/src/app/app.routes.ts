import { SelectedRefrigeratorResolver } from './refrigerator/resolvers/selected.refrigerator.resolver';
import { ExpandedResolver } from './refrigerator/resolvers/expanded.resolver';
import { RefrigeratorsResolver } from './refrigerator/resolvers/refrigerators.resolver';
import { LoginGuard } from './accounting/login/login.guard';
import { AuthGuard } from './accounting/auth.guard';
import { RefrigeratorsComponent } from './refrigerator/refrigerators.component';
import { RouterModule, Routes } from '@angular/router';
import { NgModule } from '@angular/core';
import { LoginComponent } from './accounting/login/login.component';
import { ProductsComponent } from './product/products.component';
import { UserResolver } from './accounting/resolvers/user.resolver';

const routes: Routes = [
  {
    path: 'login',
    component: LoginComponent,
    canActivate: [LoginGuard]
  },
  {
    path: 'r',
    canActivate: [AuthGuard],
    children: [
      {
        path: '',
        children: [
          {
            path: 'refrigerators',
            component: RefrigeratorsComponent,
            resolve: {
              user: UserResolver,
              expanded: ExpandedResolver,
              selectedRefrigerator: SelectedRefrigeratorResolver
            }
          },
          {
            path: 'refrigerators/:ref_id',
            component: ProductsComponent
          }
        ]
      }
    ]
  },
  {
    path: '**',
    redirectTo: '/r/refrigerators'
  }
];

@NgModule({
  imports: [
    RouterModule.forRoot(routes, {
      enableTracing: false,
      useHash: true
    })
  ],
  exports: [
    RouterModule
  ],
  providers: [
    AuthGuard,
    LoginGuard,
    RefrigeratorsResolver,
    ExpandedResolver,
    SelectedRefrigeratorResolver,
    UserResolver
  ]
})
export class AppRoutesModule {}
