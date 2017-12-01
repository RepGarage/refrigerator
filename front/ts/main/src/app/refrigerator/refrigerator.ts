import { Product } from '../product/product';

export interface IRefrigerator {
    key: string;
    name: string;
    products: Array<Product>;
    archivedProducts: Array<Product>;
}

export class Refrigerator implements IRefrigerator {
    key: string;
    name: string;
    products: Array<Product>;
    archivedProducts: Array<Product>;
    iconAssetUrl: string;

    constructor({
        key = '',
        name = '',
        products = [],
        archivedProducts = [],
        iconAssetUrl = ''
    }) {
        this.key = key;
        this.name = name;
        this.products = products;
        this.archivedProducts = archivedProducts;
        this.iconAssetUrl = iconAssetUrl;
    }
}
