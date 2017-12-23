import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-modal-view',
  templateUrl: './modal.component.html',
  styleUrls: ['./modal.component.sass']
})
export class AppModalViewComponent {
  @Input()
  destroy: VoidFunction;
  @Input()
  modalState: string;
}
