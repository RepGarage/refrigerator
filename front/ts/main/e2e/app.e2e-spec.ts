import { AppPage } from './app.po';

describe('refrigerator App', () => {
  let page: AppPage;

  beforeEach(() => {
    page = new AppPage();
  });

  it('should get root', () => {
    page.navigateTo('/login');
    const root = page.root;
    expect(root).toBeTruthy()
  });

  it ('should get header', () => {
    const header = page.header
    expect(header).toBeTruthy()
  })

  it ('should get brand', () => {
    const brand = page.brand
    console.log(brand)
    expect(brand).toBeTruthy()
  })

  it ('should get brand title', () => {
    const brandTitle = page.brand.getText()
    console.log(brandTitle)
    expect(brandTitle).toEqual('Refrigerator control')
  })
});
