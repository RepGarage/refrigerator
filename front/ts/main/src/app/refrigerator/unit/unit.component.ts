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

  @ViewChild('photoInput')
  private photoInputRef: ElementRef;

  constructor(
    private $refService: RefrigeratorService
  ) {}

  ngOnInit() {
  }

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
      console.log(file);
      this.$refService.updateRef(
          {
            key: this.ref.key,
            name: this.ref.name,
            products: this.ref.products,
            archivedProducts: this.ref.archivedProducts,
            iconAssetUrl: this.ref.iconAssetUrl,
            photo
          }
      );
    };
    reader.readAsDataURL(event.target.files[0]);
  }

  getProductsLength(products: object) {
    if (!products) { return 0; }
    return Object.keys(products).length;
  }
}
