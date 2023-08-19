package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "ReceiptProcessor/rules"
    "strconv"
)

var receiptPoints = map[string]float64{}

func processReceipt(c *gin.Context) {
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

func getReceiptPoints(c *gin.Context) {
    id := c.Param("id")
    points, ok := receiptPoints[id]
    if !ok {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "receipt not found"})
        return
    }
    c.IndentedJSON(http.StatusOK, map[string]string{
        "points": strconv.FormatFloat(points, 'f', -1, 64),
    })
}

func main() {
    router := gin.Default()
    router.POST("/receipts/process", processReceipt)
    router.GET("/receipts/:id/points", getReceiptPoints)

    router.Run("localhost:8080")
}