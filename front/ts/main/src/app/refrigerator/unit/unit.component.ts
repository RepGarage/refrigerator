import { Observable } from 'rxjs/Observable';
import { RefrigeratorService } from './../refrigerator.service';
import { AbstractControl, FormControl } from '@angular/forms';
import { Refrigerator } from './../refrigerator';
import { Component, Input, OnInit, ViewChild, ElementRef } from '@angular/core';

@Component({
  selector: 'app-ref-unit',
  templateUrl: './unit.component.html',
  styleUrls: ['./unit.component.sass']
})
export class RefUnitComponent implements OnInit {
  @Input() ref: Refrigerator;
  photo: string;
  selected: Observable<boolean>;

  @ViewChild('photoInput')
  private photoInputRef: ElementRef;

  /**
   * CONSTRUCTOR
   */
  constructor(
    private $rs: RefrigeratorService
  ) {}

  ngOnInit() {
    this.selected = this.$rs.selectedRefrigerator.map(r => r ? r.key === this.ref.key : false);
  }

  /**
   * FUNCTIONS
   */
  triggerPhotoInput() {
    this.photoInputRef.nativeElement.click();
  }
  /**
   * Photo file changed
   * @param event event
   */
  photoFileChanged(event) {
    const reader = new FileReader();
    reader.onload = (file) => {
      const photo = file.target['result'];
      this.$rs.updateRefPhoto(photo, this.ref.key);
    };
    reader.readAsDataURL(event.target.files[0]);
  }

  select() {
    this.$rs.selectRefrigerator(this.ref);
  }

  getProductsLength(products: object) {
    if (!products) { return 0; }
    return Object.keys(products).length;
  }
}
