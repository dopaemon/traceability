package services

import (
    "context"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "traceability/config"
)

var ProductCollection *mongo.Collection

func InitMongo() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoURI))
    if err != nil {
        log.Fatal(err)
    }

    ProductCollection = client.Database(config.DBName).Collection("products")
}
