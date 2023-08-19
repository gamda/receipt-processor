package rules

import (
    "math"
	"strconv"
)

func PointsForRoundDollar(total string) int {
    points := 50
    if isRoundDollar(total) {
        return points
    }
    return 0
}

func isRoundDollar(total string) bool {
    // Assumption: all `total` variables have a decimal point and two cents digits
    cents := total[len(total)-2:]
    return cents == "00"
}

func PointsForMultipleTwentyFiveCents(total string) int {
    points := 25
    if isMultipleTwentyFiveCents(total) {
        return points
    }
    return 0
}

func isMultipleTwentyFiveCents(total string) bool {
    // Assumption: all `total` variables have a decimal point and two cents digits
    centsString := total[len(total)-2:]
    // Assumption: all totals are valid numbers, cents are valid integers
    cents, err := strconv.ParseFloat(centsString, 64)
    if (err == nil) {
        return math.Mod(cents, 25) == 0
    }
    return false
}

func PointsForPurchaseDateOdd(purchaseDate string) int {
    points := 6
    if isPurchaseDateOdd(purchaseDate) {
        return points
    }
    return 0
}

func isPurchaseDateOdd(purchaseDate string) bool {
    dayLastDigitString := purchaseDate[len(purchaseDate)-1:]
    day, err := strconv.Atoi(dayLastDigitString)
    if !(err == nil) {
        return false
    }
    return !(day%2 == 0);
}

func PointsForPurchaseTime14And16(purchaseTime string) int {
    points := 10
    if isPurchaseTimeBetween14And16(purchaseTime) {
        return points
    }
    return 0
}

func isPurchaseTimeBetween14And16(purchaseTime string) bool {
    hourString := purchaseTime[0:2]
    hour, err := strconv.Atoi(hourString)
    if !(err == nil) {
        return false
    }
    return hour >= 14 && hour < 16
}

