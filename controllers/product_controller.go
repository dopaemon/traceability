package controllers

import (
    "net/http"
    "time"
    "traceability/models"
    "traceability/services"

    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
)

func UploadProduct(c *gin.Context) {
    name := c.PostForm("name")
    farmer := c.PostForm("farmer")
    location := c.PostForm("location")
    // harvest := c.PostForm("harvest_date")
    file, _ := c.FormFile("file")

    // Save temp file
    tempPath := "./tmp/" + file.Filename
    c.SaveUploadedFile(file, tempPath)

    ipfsHash, _ := services.UploadToIPFS(tempPath)
    txHash, _ := services.WriteToBlockchain(ipfsHash)

    product := models.Product{
        ID:          time.Now().Format("20060102150405"),
        Name:        name,
        Farmer:      farmer,
        Location:    location,
        HarvestDate: time.Now(),
        IPFSHash:    ipfsHash,
        TxHash:      txHash,
        CreatedAt:   time.Now(),
    }

    _, err := services.ProductCollection.InsertOne(c, product)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "DB failed"})
        return
    }

    c.JSON(http.StatusOK, product)
}

// Truy xuất 1 sản phẩm
func GetProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	err := services.ProductCollection.FindOne(c, bson.M{"_id": id}).Decode(&product)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// Xoá 1 sản phẩm
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	_, err := services.ProductCollection.DeleteOne(c, bson.M{"_id": id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Delete failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
