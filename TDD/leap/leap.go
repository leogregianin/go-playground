package leap

const testVersion = 3

func IsLeapYear(i int) bool {
    return ((i % 400) == 0) || (((i % 4) == 0) && ((i % 100) != 0))
}