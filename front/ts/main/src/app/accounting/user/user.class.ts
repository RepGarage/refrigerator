export class User {
  uid: string;
  displayName: string;
  photoURL: string;
  email: string;
  constructor(uid: string, displayName: string, photoURL: string, email: string) {
    this.uid = uid;
    this.displayName = displayName;
    this.photoURL = photoURL;
    this.email = email;
  }
}
