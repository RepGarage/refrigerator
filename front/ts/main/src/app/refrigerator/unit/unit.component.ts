import { AbstractControl, FormControl } from '@angular/forms';
import { Refrigerator } from './../refrigerator';
import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-ref-unit',
  templateUrl: './unit.component.html',
  styleUrls: ['./unit.component.sass']
})
export class RefUnitComponent implements OnInit {
  @Input() ref: Refrigerator;

  photoCtrl: AbstractControl;
  photo: string;

  ngOnInit() {
    this.photoCtrl = new FormControl();
    this.photoCtrl.valueChanges.subscribe(console.log);
  }

  photoFileChanged(event) {
    const reader = new FileReader();
    reader.onload = (file) => {
      console.log(file);
      this.photo = file.target['result'];
    };
    reader.readAsDataURL(event.target.files[0]);
  }

  getProductsLength(products: object) {
    if (!products) { return 0; }
    return Object.keys(products).length;
  }
}
