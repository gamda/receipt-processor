package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "ReceiptProcessor/rules"
)

var receiptPoints = map[string]float64{}

func processReceipt(c *gin.Context){
    var newReceipt rules.Receipt
    // Call BindJSON to bind the received JSON to newReceipt.
    if err := c.BindJSON(&newReceipt); err != nil {
        return
    }
    // calculate points
    points := rules.PointsForReceipt(newReceipt)
    // Generate ID
    id := uuid.NewString()
    // store points
    receiptPoints[id] = points
    c.IndentedJSON(http.StatusCreated, map[string]string{
        "id": id,
    })
}

func main() {
    router := gin.Default()
    router.POST("/receipts/process", processReceipt)

    router.Run("localhost:8080")
}