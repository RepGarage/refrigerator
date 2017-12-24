import { Component, Input } from '@angular/core';
import {
  trigger,
  state,
  style,
  animate,
  transition
} from '@angular/animations';

@Component({
  selector: 'app-modal-view',
  templateUrl: './modal.component.html',
  styleUrls: ['./modal.component.sass'],
  animations: [
    trigger('state', [
      state('visible', style({
        transform: 'translateY(0)'
      })),
      state('hidden', style({
        transform: 'translateY(100%)'
      })),
      transition('* => *', animate('200ms ease-in'))
    ]),
    trigger('modalState', [
      state('visible', style({
        opacity: '0.7'
      })),
      state('hidden', style({
        opacity: '0'
      })),
      transition('* => *', animate('200ms ease-in-out'))
    ])
  ]
})
export class AppModalViewComponent {
  @Input()
  destroy: VoidFunction;
  @Input()
  state: string;
  @Input()
  modalState: string;
  @Input()
  context;
}
