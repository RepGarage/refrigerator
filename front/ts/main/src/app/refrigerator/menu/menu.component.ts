import { Observable } from 'rxjs/Observable';
import { Component, OnInit } from '@angular/core';
import { RefrigeratorService } from '../refrigerator.service';
import { Refrigerator } from '../refrigerator';

@Component({
  selector: 'app-ref-menu',
  templateUrl: './menu.component.html',
  styleUrls: ['./menu.component.sass']
})
export class RefMenuComponent implements OnInit {
  selectedRefrigerator: Observable<Refrigerator>;
  newName: string;
  editingName = false;
  constructor(
    private $rs: RefrigeratorService
  ) {}

  ngOnInit() {
    this.selectedRefrigerator = this.$rs.selectedRefrigerator;
    this.selectedRefrigerator.subscribe((r: Refrigerator) => this.newName = r.name);
  }

  updateName() {
    this.editingName = false;
    this.selectedRefrigerator
      .map((r: Refrigerator) => {
        this.$rs.updateRefName(this.newName, r.key);
      }).subscribe(console.log);
  }
}
