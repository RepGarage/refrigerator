using MongoDB.Bson.Serialization.Attributes;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace MainApp.Models
{
    public class ShelfLife
    {
        [BsonElement("product_id")]
        public int  ProductId { get; set; }
        [BsonElement("type")]
        public ShelfDegree Type { get; set; }
        [BsonElement("data")]
        public int Data { get; set; }
    }
}
