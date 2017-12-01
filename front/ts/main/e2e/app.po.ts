import { browser, by, element } from 'protractor';

export class AppPage {
  get browser () {
    return browser;
  }

  get root () {
    return element(by.css('app-root'))
  }
  get header () {
    return this.root.element(by.css('navbar'))
  }
  get brand () {
    return this.header.element(by.css('.navbar-brand'))
  }
  navigateTo(path: string) {
    return browser.get(path);
  }
}
