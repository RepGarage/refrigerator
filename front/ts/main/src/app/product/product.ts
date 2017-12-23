import { SafeUrl } from '@angular/platform-browser';
export class Product {
  photoUrl: SafeUrl;
  name: string;
  created: string;
  moved: string;
  shelf: string;
  exp: string;
  product_id: string;
  key: string;

  constructor({
    photoUrl = <SafeUrl>'',
    name = '',
    created = '',
    moved = '',
    shelf = '',
    exp = '',
    product_id = '',
    key = ''
  }) {
    this.photoUrl = photoUrl;
    this.name = name;
    this.created = created;
    this.moved = moved;
    this.shelf = shelf;
    this.exp = exp;
    this.product_id = product_id;
    this.key = key;
  }
}
