package rules

import (
    "testing"
    "fmt"
)

func TestPointsForReceipt(t *testing.T) {
    receipt := Receipt{
        retailer: "Target",
        purchaseDate: "2022-01-01",
        purchaseTime: "13:01",
        total: "35.35",
        items: []Item{
            Item{"Mountain Dew 12PK", "6.49"},
            Item{"Emils Cheese Pizza", "12.25"},
            Item{"Knorr Creamy Chicken", "1.26"},
            Item{"Doritos Nacho Cheese", "3.35"},
            Item{"   Klarbrunn 12-PK 12 FL OZ  ", "12.00"},
        },
    }
    fmt.Println(PointsForReceipt(receipt))
}

func TestRoundDollarFalse(t *testing.T) {
    total := "11.23"
    roundDollar := PointsForRoundDollar(total)
    if roundDollar > 0 {
        t.Fatalf(`11.23 is not round`)
    }
}

func TestRoundDollarTrue(t *testing.T) {
    total := "11.00"
    roundDollar := PointsForRoundDollar(total)
    if roundDollar == 0 {
        t.Fatalf(`11.00 is round`)
    }
}

func TestTwentyFiveFalse(t *testing.T) {
    test_totals := []string{"11.99", "11.01", "11.24", "11.26", "11.49", "11.51", "11.74", "11.76"}
    for _, total := range test_totals {
        roundDollar := PointsForMultipleTwentyFiveCents(total)
        if roundDollar > 0 {
            t.Fatalf(`%v is not multiple of .25`, total)
        }
    }
}

func TestTwentyFiveTrue(t *testing.T) {
    test_totals := []string{"11.00", "11.25", "11.50", "11.75"}
    for _, total := range test_totals {
        roundDollar := PointsForMultipleTwentyFiveCents(total)
        if roundDollar == 0 {
            t.Fatalf(`%v is a multiple of .25`, total)
        }
    }
}

func TestPurchaseDateOdd(t *testing.T) {
    date := "2023-08-19"
    if PointsForPurchaseDateOdd(date) == 0 {
        t.Fatalf(`%v is an odd date and should have points`, date)
    }
}

func TestPurchaseDateEven(t *testing.T) {
    date := "2023-08-18"
    if PointsForPurchaseDateOdd(date) > 0 {
        t.Fatalf(`%v is an even date and should NOT have points`, date)
    }
}

func TestPurchaseTimeInside14And16(t *testing.T) {
    times := []string{"14:59", "15:37"}
    for _, time := range times {
        points := PointsForPurchaseTime14And16(time)
        if points == 0 {
            t.Fatalf(`%v is inside 14 and 16`, time)
        }
    }
}

func TestPurchaseTimeOutside14And16(t *testing.T) {
    times := []string{"13:59", "16:01"}
    for _, time := range times {
        points := PointsForPurchaseTime14And16(time)
        if points > 0 {
            t.Fatalf(`%v is outside 14 and 16`, time)
        }
    }
}

func TestPointsForItems(t *testing.T) {
    items := []Item{
        Item{"Mountain Dew 12PK", "6.49"},
        Item{"Emils Cheese Pizza", "12.25"},
        Item{"Knorr Creamy Chicken", "1.26"},
        Item{"Doritos Nacho Cheese", "3.35"},
        Item{"   Klarbrunn 12-PK 12 FL OZ  ", "12.00"},
    }
    if PointsForItems(items[:1]) != 0 {
        t.Fatalf(`One item should give zero points`)
    }
    if PointsForItems(items[:2]) != 5 {
        t.Fatalf(`Two items should give five points`)
    }
    if PointsForItems(items[:3]) != 5 {
        t.Fatalf(`Three items should give five points`)
    }
    if PointsForItems(items[:4]) != 10 {
        t.Fatalf(`Four items should give ten points`)
    }
}

func TestPointsForItemName(t *testing.T) {
    if PointsForItemName("Emils Cheese Pizza", "12.25") != 3 {
        t.Fatalf(`18 characters should give 3 points with a 12.25 price`)
    }
    if PointsForItemName("   Klarbrunn 12-PK 12 FL OZ  ", "12.00") != 3 {
        t.Fatalf(`24 characters should give 3 points with a 12.00 price`)
    }
}

func TestPointsForRetailerName(t *testing.T) {
    if PointsForRetailerName("abc123$!#") != 6 {
        t.Fatalf(`abc123$!# should have 6 points`)
    }
    if PointsForRetailerName("Target") != 6 {
        t.Fatalf(`Target should have 6 points`)
    }
    if PointsForRetailerName("M&M Corner Market") != 14 {
        t.Fatalf(`M&M Corner Market should have 14 points`)
    }
}
