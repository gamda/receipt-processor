package rules

import (
    "testing"
)

func TestRoundDollarFalse(t *testing.T) {
    total := "11.23"
    roundDollar := RoundDollar(total)
    if roundDollar {
        t.Fatalf(`11.23 is not round`)
    }
}

func TestRoundDollarTrue(t *testing.T) {
    total := "11.23"
    roundDollar := RoundDollar(total)
    if !roundDollar {
        t.Fatalf(`11.00 is round`)
    }
}
