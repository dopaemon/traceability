package models

import "time"

type Product struct {
    ID          string    `json:"id" bson:"_id"`
    Name        string    `json:"name" bson:"name"`
    Farmer      string    `json:"farmer" bson:"farmer"`
    Location    string    `json:"location" bson:"location"`
    HarvestDate time.Time `json:"harvest_date" bson:"harvest_date"`
    IPFSHash    string    `json:"ipfs_hash" bson:"ipfs_hash"`
    TxHash      string    `json:"tx_hash" bson:"tx_hash"`
    CreatedAt   time.Time `json:"created_at" bson:"created_at"`
}
