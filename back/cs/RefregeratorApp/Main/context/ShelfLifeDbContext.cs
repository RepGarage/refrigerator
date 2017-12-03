using MainApp.Models;
using MongoDB.Bson;
using MongoDB.Driver;
using MongoDB.Driver.GridFS;
using System;
using System.Collections.Generic;
using System.IO;
using System.Threading.Tasks;

namespace MainApp.context
{
    public class ShelfLifeDbContext
    {
        IMongoDatabase database;
        IGridFSBucket gridFS;  

        public ShelfLifeDbContext()
        {
            string connectionString = "mongodb://localhost:27017/icebox";
            var connection = new MongoUrlBuilder(connectionString);
            MongoClient client = new MongoClient(connectionString);
            database = client.GetDatabase(connection.DatabaseName);
            gridFS = new GridFSBucket(database);
        }

        private IMongoCollection<ShelfLife> ShelfLifes
        {
            get { return database.GetCollection<ShelfLife>("ShelfLifes"); }
        }

        public async Task<ShelfLife> GetShelfLife(int id) => await ShelfLifes.Find(new BsonDocument("product_id", id)).FirstOrDefaultAsync();

        public async Task Create(ShelfLife p) => await ShelfLifes.InsertOneAsync(p);

        public async Task Update(ShelfLife p) => await ShelfLifes.ReplaceOneAsync(new BsonDocument("product_id", p.ProductId), p);
        public async Task Remove(int id) => await ShelfLifes.DeleteOneAsync(new BsonDocument("product_id", id));
    }
}
