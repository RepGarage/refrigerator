import { firebaseSecret } from './.secret';

export const environment = {
  production: false,
  productsBaseApi: 'http://localhost:8080/',
  firebaseSecret,
  hmr: true
};
