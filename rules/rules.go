package rules

func RoundDollar(total string) int {
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