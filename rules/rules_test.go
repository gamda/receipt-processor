package rules

import (
    "testing"
)

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
    for i, total := range test_totals {
        roundDollar := PointsForMultipleTwentyFiveCents(total)
        if roundDollar > 0 {
            t.Fatalf(`Test %v: %v is not multiple of .25`, i, total)
        }
    }
}

func TestTwentyFiveTrue(t *testing.T) {
    test_totals := []string{"11.00", "11.25", "11.50", "11.75"}
    for i, total := range test_totals {
        roundDollar := PointsForMultipleTwentyFiveCents(total)
        if roundDollar == 0 {
            t.Fatalf(`Test %v: %v is a multiple of .25`, i, total)
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
