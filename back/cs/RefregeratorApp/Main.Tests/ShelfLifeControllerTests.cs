using MainApp.Controllers;
using Microsoft.AspNetCore.Mvc;
using Xunit;

namespace Main.Tests
{
   public class ShelfLifeControllerTests
    {
        [Fact]
        public void GetAsyncInsertAndDeleteTest()
        {
            string url = "https://www.perekrestok.ru/catalog/moloko-syr-yaytsa/syry/tender-syr-rossiyskiy-1kg--304011";
            int productId = new System.Random().Next(0,100000);
            ShelfLifeController controller = new ShelfLifeController();

            var resultFirst = controller.GetAsync(productId, url).Result;
            Assert.IsType<JsonResult>(resultFirst);

            var resultSecond = controller.GetAsync(productId, "aaa").Result;
            Assert.IsType<JsonResult>(resultFirst);

            var res = controller.Remove(productId);
            var resultThird = controller.GetAsync(productId, "aaa").Result;
            Assert.IsType<BadRequestResult>(resultThird);
        }
    }
}
