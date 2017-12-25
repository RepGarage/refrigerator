import { RefrigeratorService } from './../refrigerator.service';
import { Component, OnInit } from '@angular/core';
import { Refrigerator } from '../refrigerator';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

@Component({
    selector: 'app-add-refrigerator',
    templateUrl: './add.component.html',
    styleUrls: [ './add.component.sass' ]
})
export class AddRefrigeratorComponent implements OnInit {
    fGroup: FormGroup;

    constructor(
        private fb: FormBuilder,
        private $rs: RefrigeratorService
    ) {}

    ngOnInit() {
        this.fGroup = this.fb.group({
            name: ['', Validators.required]
        });
    }

    addRefrigerator() {
      if (!this.fGroup.valid) {
          return;
      }
      const name = this.fGroup.get('name').value;
      const newRefrigerator = new Refrigerator({ name });
      this.$rs.addRefrigerator(newRefrigerator);
      this.fGroup.reset();
    }

    fadeToBlack() {
        this.$rs.addRefrigeratorActive.next(false);
    }
}
