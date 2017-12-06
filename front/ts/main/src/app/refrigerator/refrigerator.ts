import { Product } from '../product/product';

export class Refrigerator {
    key: string;
    name: string;
    products: Array<Product>;
    archivedProducts: Array<Product>;
    iconAssetUrl: string;
    photo: string;

    constructor({
        key = '',
        name = '',
        products = [],
        archivedProducts = [],
        iconAssetUrl = '',
        photo = ''
    }) {
        this.key = key;
        this.name = name;
        this.products = products;
        this.archivedProducts = archivedProducts;
        this.iconAssetUrl = iconAssetUrl;
        this.photo = photo;
    }
}
