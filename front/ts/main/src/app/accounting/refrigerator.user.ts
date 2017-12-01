import { IRefrigerator } from '../refrigerator/refrigerator';

export interface IRefUser {
    key: string;
    refrigeratorsIds?: Object;
    email: string;
    displayName: string;
}

export class RefUser implements IRefUser {
    key: string;
    refrigeratorsIds?: Object;
    email: string;
    displayName: string;

    constructor(u: IRefUser) {
        this.refrigeratorsIds = u.refrigeratorsIds || [];
        this.email = u.email;
        this.displayName = u.displayName;
    }
}

export interface UserQuery {
    key?: string;
    refrigeratorsIds?: Object;
    email?: string;
    displayName?: string;
}
