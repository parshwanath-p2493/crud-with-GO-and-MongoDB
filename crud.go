package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func servercheck(c *gin.Context) {
	fmt.Println("THE SERVER IS ALL SET GO `")
	c.JSON(200, gin.H{"message": "THE SERVER IS ALL SET GO "})
}

func InsertData(c *gin.Context) {
	var d Data
	if err := c.BindJSON(&d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}
	d.ID = primitive.NewObjectID()
	collection := DB.Database("project").Collection("sampledatas")
	_, err := collection.InsertOne(context.TODO(), d)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data is inserted", "data": d})
}
func GetAll(c *gin.Context) {
	collection := DB.Database("project").Collection("sampledatas")
	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	var data []Data
	for cur.Next(context.TODO()) {
		var d Data
		if err := cur.Decode(&d); err != nil {
			fmt.Println(err)
			return
		}
		data = append(data, d)
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data retrived", "data": data})
}
func DeleteData(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Format"})
		return
	}
	collection := DB.Database("project").Collection("sampledatas")
	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data deleted successfuly"})
}
func UpdateData(c *gin.Context) {
	var d Data
	idd := c.Param("id")
	if err := c.BindJSON(&d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	objectId, err := primitive.ObjectIDFromHex(idd)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID fromat"})
		return
	}
	collection := DB.Database("project").Collection("sampledatas")
	update := bson.M{"$set": d}
	_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": objectId}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Updated Successfuly", "data": update})

}
