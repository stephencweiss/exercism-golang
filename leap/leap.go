/*
Package leap is a small utility to determine if a year is a leap year.
The rules for a leap year are:
1. The year is divisible by four
2. Not divisible by 100,
3. Unless it is divisible by 400
*/
package leap

// IsLeapYear takes in a year (int) and returns a boolean value of whether the year qualifies as a leap year.
func IsLeapYear(year int) bool {

	divisibleByFour := IsEvenlyDivisible(year, 4)
	divisibleByOneHundred := IsEvenlyDivisible(year, 100)
	divisibleByFourHundred := IsEvenlyDivisible(year, 400)

	if divisibleByFour && !divisibleByOneHundred {
		return true
	}
	if divisibleByFour && divisibleByOneHundred && divisibleByFourHundred {
		return true
	}

	return false
}

// IsEvenlyDivisible takes in two numbers and returns true if the modulo operator returns 0.
func IsEvenlyDivisible(numerator, denominator int) bool {
	if 0 == numerator%denominator {
		return true
	}
	return false

}
