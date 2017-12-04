using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using MainApp.Models;
using System.Diagnostics;
using System.Net.Http;
using HtmlAgilityPack;
using MainApp.context;
using System.Text.RegularExpressions;
using MainApp.Helpers;

namespace MainApp.Controllers
{
    [Route("api/[controller]")]
    public class ShelfLifeController : Controller
    {
        private readonly ShelfLifeDbContext db = new ShelfLifeDbContext();

        [HttpGet]
        public async Task<ActionResult> GetAsync(int productId, string url)
        {
  
            if (!ModelState.IsValid)
            {
                return BadRequest();
            }
            if (productId == 0 )
            {
                return BadRequest();
            }

            try
            {
                var shelfLife = await db.GetShelfLife(productId);
                if (shelfLife != null)
                {
                    return Json(shelfLife);
                }
            }
            catch
            {
                return StatusCode(500);
            }

            try
            {
                string html = await GetExpDateAsync(url);
                var doc = new HtmlDocument();
                doc.LoadHtml(html);

                var links = doc.DocumentNode.SelectNodes("//td[@class='xf-product-table__col']").Last().InnerText.Trim();
                var data = Convert.ToInt32(Regex.Match(links, @"\d+").Value);
                var newProduct = new ShelfLife { ProductId = productId, Data = data, Type = DateTypeHelper.GetDateType(links) };
                try
                {
                    await db.Create(newProduct);
                }
                catch
                {
                    return StatusCode(500);
                }
                return Json(newProduct);
            }
            catch
            {
                return BadRequest();
            }
        }

        [HttpDelete]
        public async Task Remove(int productId) => await db.Remove(productId);

        private async Task<string> GetExpDateAsync(string url)
        {
            using (HttpClient client = new HttpClient())
                return await client.GetStringAsync(url);
        }
        public IActionResult Error()
        {
            return View(new ErrorViewModel { RequestId = Activity.Current?.Id ?? HttpContext.TraceIdentifier });
        }
    }
}