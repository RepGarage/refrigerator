import { firebaseSecret } from './.secret';

export const environment = {
  production: false,
  productsBaseApi: 'http://ref-back-main:8081/',
  firebaseSecret,
  hmr: true
};
