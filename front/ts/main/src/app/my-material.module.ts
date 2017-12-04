import { NgModule } from '@angular/core';

// Imports

import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import {
    MatFormFieldModule,
    MatToolbarModule,
    MatCardModule,
    MatGridListModule,
    MatExpansionModule,
    MatInputModule,
    MatButtonModule,
    MatDatepickerModule,
    MatNativeDateModule,
    MatListModule,
    MatMenuModule,
    MatIconModule
} from '@angular/material';

@NgModule({
    imports: [
        BrowserAnimationsModule,
        MatFormFieldModule,
        MatToolbarModule,
        MatCardModule,
        MatGridListModule,
        MatExpansionModule,
        MatInputModule,
        MatButtonModule,
        MatDatepickerModule,
        MatNativeDateModule,
        MatListModule,
        MatMenuModule,
        MatIconModule
    ],
    exports: [
        BrowserAnimationsModule,
        MatFormFieldModule,
        MatToolbarModule,
        MatCardModule,
        MatGridListModule,
        MatExpansionModule,
        MatInputModule,
        MatButtonModule,
        MatDatepickerModule,
        MatNativeDateModule,
        MatListModule,
        MatMenuModule,
        MatIconModule
    ]
})
export class MyMaterialModule {

}
