package mathgl

import "math"

// Epsilon is some tiny value that determines how precisely equal we want our floats to be

// FloatEqual is a safe utility function to compare floats.
// It's Taken from http://floating-point-gui.de/errors/comparison/
//
// It is slightly altered to not call Abs when not needed.
// Keep in mind that it expects float32s to be converted to float64s before being passed in, because they have to be converted for Abs anyway
func FloatEqual(a, b float64) bool {
	var epsilon float64 = 1e-15
	if a == b { // Handles the case of inf or shortcuts the loop when no significant error has accumulated
		return true
	} else if a*b == 0 { // If a or b are 0
		return math.Abs(a-b) < epsilon*epsilon
	}

	// Else compare difference
	return math.Abs(a-b)/(math.Abs(a)+math.Abs(b)) < epsilon
}

func FloatEqual32(a, b float32) bool {
	var epsilon float32 = 1e-15
	if a == b { // Handles the case of inf or shortcuts the loop when no significant error has accumulated
		return true
	} else if a*b == 0 { // If a or b is 0
		//return math.Abs(float64(a-b)) < epsilon*epsilon
		return Fabs32(a-b) < epsilon*epsilon
	}

	// Else compare difference
	return Fabs32(a-b)/(Fabs32(a)+Fabs32(b)) < epsilon
}

func FloatEqualFunc(epsilon float64) func(float64, float64) bool {
	return func(a, b float64) bool {
		return FloatEqualThreshold(a, b, epsilon)
	}
}

func FloatEqual32Func(epsilon float32) func(float32, float32) bool {
	return func(a, b float32) bool {
		return FloatEqualThreshold32(a, b, epsilon)
	}
}

// FloatEqualThreshold is a utility function to compare floats.
// It's Taken from http://floating-point-gui.de/errors/comparison/
//
// It is slightly altered to not call Abs when not needed.
// Keep in mind that it expects float32s to be converted to float64s before being passed in, because they have to be converted for Abs anyway
//
// This differs from FloatEqual in that it lets you pass in your comparison threshold, so that you can adjust the comparison value to your specific needs
func FloatEqualThreshold(a, b, epsilon float64) bool {
	if a == b { // Handles the case of inf or shortcuts the loop when no significant error has accumulated
		return true
	} else if a*b == 0 { // If a or b is 0
		return math.Abs(a-b) < epsilon*epsilon
	}

	// Else compare difference
	return math.Abs(a-b)/(math.Abs(a)+math.Abs(b)) < epsilon
}

func FloatEqualThreshold32(a, b, epsilon float32) bool {
	if a == b { // Handles the case of inf or shortcuts the loop when no significant error has accumulated
		return true
	} else if a*b == 0 { // If a or b is 0
		return Fabs32(a-b) < Fsqr32(epsilon)
	}

	// Else compare difference
	return Fabs32(a-b)/(Fabs32(a)+Fabs32(b)) < epsilon
}

func Clampf(a, t1, t2 float32) float32 {
	if a < t1 {
		return t1
	} else if a > t2 {
		return t2
	}

	return a
}

func ClampfFunc(t1, t2 float32) func(float32) {
	return func(a float32) {
		Clampf(a, t1, t2)
	}
}

func ClampdFunc(t1, t2 float64) func(float64) {
	return func(a float64) {
		Clampd(a, t1, t2)
	}
}

func Clampd(a, t1, t2 float64) float64 {
	if a < t1 {
		return t1
	} else if a > t2 {
		return t2
	}

	return a

}

/* The IsClamped functions use strict equality (meaning: not the FloatEqual function)
there shouldn't be any major issues with this since clamp is often used to fix minor errors*/

func IsClampedf(a, t1, t2 float32) bool {
	return a >= t1 && a <= t2
}

func IsClampedd(a, t1, t2 float64) bool {
	return a >= t1 && a <= t2
}

func SetMinf(a, b *float32) {
	if *b < *a {
		*a = *b
	}
}

func SetMaxf(a, b *float32) {
	if *a < *b {
		*a = *b
	}
}

func SetMind(a, b *float64) {
	if *b < *a {
		*a = *b
	}
}

func SetMaxd(a, b *float64) {
	if *a < *b {
		*a = *b
	}
}
