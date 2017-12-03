import { AppRoutesModule } from './app.routes';
import { AuthGuard } from './accounting/auth.guard';
import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';

// Components import
import { AppComponent } from './app.component';
import { ProductsComponent } from './product/products.component';
import { AuthComponent } from './accounting/auth/auth.component';
import { LoginComponent } from './accounting/login/login.component';
import { RefrigeratorsComponent } from './refrigerator/refrigerators.component';

// Modules import
import { HttpModule } from '@angular/http';
import { MyMaterialModule } from './my-material.module';
import { ReactiveFormsModule, FormsModule } from '@angular/forms';
import { RouterModule, Routes } from '@angular/router';
import { ProductModule } from './product/product.module';
import { AngularFireModule } from 'angularfire2';
import { AngularFireDatabaseModule } from 'angularfire2/database';
import { AngularFireAuthModule } from 'angularfire2/auth';
import { RefrigeratorModule } from './refrigerator/refrigerator.module';

// Services import
import { ProductService } from './product/product.service';



// Environment
import { environment } from '../environments/environment';
import { AppService } from './app.service';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { AuthService } from './accounting/auth.service';

@NgModule({
  declarations: [
    AppComponent,
    AuthComponent,
    LoginComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    HttpClientModule,
    MyMaterialModule,
    ReactiveFormsModule,
    FormsModule,
    ProductModule,
    RefrigeratorModule,
    AppRoutesModule,
    AngularFireModule.initializeApp(environment.firebaseSecret),
    AngularFireDatabaseModule,
    AngularFireAuthModule
  ],
  providers: [
    ProductService,
    AuthService,
    AppService,
    {
      provide: 'PRODUCTS_BASE_URL',
      useValue: environment.productsBaseApi
    }
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
