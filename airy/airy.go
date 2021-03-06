// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package airy

import "github.com/gonum/mathext/airy/internal/amos"

// Ai returns the value of the Airy function at z. The Airy function here,
// Ai(z), is one of the two linearly independent solutions to
//  y'' - y*z = 0.
// See http://mathworld.wolfram.com/AiryFunctions.html for more detailed information.
func Ai(z complex128) complex128 {
	// id specifies the order of the derivative to compute,
	// 0 for the function itself and 1 for the derivative.
	// kode specifies the scaling option. See the function
	// documentation for the exact behavior.
	id := 0
	kode := 1
	air, aii, _ := amos.Zairy(real(z), imag(z), id, kode)
	return complex(air, aii)
}

// AiDeriv returns the value of the derivative of the Airy function at z. The
// Airy function here, Ai(z), is one of the two linearly independent solutions to
//  y'' - y*z = 0.
// See http://mathworld.wolfram.com/AiryFunctions.html for more detailed information.
func AiDeriv(z complex128) complex128 {
	// id specifies the order of the derivative to compute,
	// 0 for the function itself and 1 for the derivative.
	// kode specifies the scaling option. See the function
	// documentation for the exact behavior.
	id := 1
	kode := 1
	air, aii, _ := amos.Zairy(real(z), imag(z), id, kode)
	return complex(air, aii)
}
