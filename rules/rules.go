package rules

import (
    "math"
	"strconv"
	"strings"
	"unicode"
)

type Item struct {
    ShortDescription string `json:"shortDescription"`
    Price string `json:"price"`
}

type Receipt struct {
    Retailer string `json:"retailer"`
    PurchaseDate string `json:"purchaseDate"`
    PurchaseTime string `json:"purchaseTime"`
    Total string `json:"total"`
    Items []Item `json:"items"`
}

func PointsForReceipt(receipt Receipt) float64 {
    return PointsForRetailerName(receipt.Retailer) +
        PointsForRoundDollar(receipt.Total) +
        PointsForMultipleTwentyFiveCents(receipt.Total) +
        PointsForItems(receipt.Items) +
        PointsForItemNames(receipt.Items) +
        PointsForPurchaseDateOdd(receipt.PurchaseDate) +
        PointsForPurchaseTime14And16(receipt.PurchaseTime)
}

func PointsForRoundDollar(total string) float64 {
    points := 50.0
    if isRoundDollar(total) {
        return points
    }
    return 0.0
}

func isRoundDollar(total string) bool {
    // Assumption: all `total` variables have a decimal point and two cents digits
    cents := total[len(total)-2:]
    return cents == "00"
}

func PointsForMultipleTwentyFiveCents(total string) float64 {
    points := 25.0
    if isMultipleTwentyFiveCents(total) {
        return points
    }
    return 0.0
}

func isMultipleTwentyFiveCents(total string) bool {
    // Assumption: all `total` variables have a decimal point and two cents digits
    centsString := total[len(total)-2:]
    // Assumption: all totals are valid numbers, cents are valid integers
    cents, err := strconv.Atoi(centsString)
    if !(err == nil) {
        return false
    }
    return cents%25 == 0
}

func PointsForPurchaseDateOdd(purchaseDate string) float64 {
    points := 6.0
    if isPurchaseDateOdd(purchaseDate) {
        return points
    }
    return 0.0
}

func isPurchaseDateOdd(purchaseDate string) bool {
    dayLastDigitString := purchaseDate[len(purchaseDate)-1:]
    day, err := strconv.Atoi(dayLastDigitString)
    if !(err == nil) {
        return false
    }
    return !(day%2 == 0);
}

func PointsForPurchaseTime14And16(purchaseTime string) float64 {
    points := 10.0
    if isPurchaseTimeBetween14And16(purchaseTime) {
        return points
    }
    return 0.0
}

func isPurchaseTimeBetween14And16(purchaseTime string) bool {
    // Assumption: time is always in 24-hour format
    hourString := purchaseTime[0:2]
    hour, err := strconv.Atoi(hourString)
    if !(err == nil) {
        return false
    }
    return hour >= 14 && hour < 16
}

func PointsForItems(items []Item) float64 {
    return float64(len(items) / 2) * 5
}

func PointsForItemNames(items []Item) float64 {
    total := 0.0
    for _, item := range items {
        total += PointsForItemName(item.ShortDescription, item.Price)
    }
    return total
}

func PointsForItemName(name string, priceString string) float64 {
    trimmedName := strings.TrimSpace(name);
    if len(trimmedName)%3 != 0 {
        return 0;
    }
    price, err := strconv.ParseFloat(priceString, 64)
    if !(err == nil) {
        return 0
    }
    points := price * 0.2
    return math.Ceil(points)
}

func PointsForRetailerName(name string) float64 {
    points := 0.0
    for _, char := range name {
        if unicode.IsLetter(char) || unicode.IsNumber(char) {
            points += 1
        }
    }
    return points
}
