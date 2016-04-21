// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package amos

import (
	"math"
	"math/rand"
	"strconv"
	"testing"

	"github.com/gonum/mathext/internal/amos/amoslib"
)

type input struct {
	x    []float64
	is   []int
	kode int
	id   int
	yr   []float64
	yi   []float64
	n    int
	tol  float64
}

func randnum(rnd *rand.Rand) float64 {
	r := 2e2 // Fortran has infinite loop if this is set higher than 2e3
	if rnd.Float64() > 0.99 {
		return 0
	}
	return rnd.Float64()*r - r/2
}

func randInput(rnd *rand.Rand) input {
	x := make([]float64, 8)
	for j := range x {
		x[j] = randnum(rnd)
	}
	is := make([]int, 3)
	for j := range is {
		is[j] = rand.Intn(1000)
	}
	kode := rand.Intn(2) + 1
	id := rand.Intn(2)
	n := rand.Intn(5) + 1
	yr := make([]float64, n+1)
	yi := make([]float64, n+1)
	for j := range yr {
		yr[j] = randnum(rnd)
		yi[j] = randnum(rnd)
	}
	tol := 1e-14

	return input{
		x, is, kode, id, yr, yi, n, tol,
	}
}

const nInputs = 100000

func TestAiry(t *testing.T) {
	rnd := rand.New(rand.NewSource(1))
	for i := 0; i < nInputs; i++ {
		in := randInput(rnd)
		zairytest(t, in.x, in.kode, in.id)
	}
}

func TestZacai(t *testing.T) {
	rnd := rand.New(rand.NewSource(1))
	for i := 0; i < nInputs; i++ {
		in := randInput(rnd)
		zacaitest(t, in.x, in.is, in.tol, in.n, in.yr, in.yi, in.kode)
	}
}

func TestZbknu(t *testing.T) {
	rnd := rand.New(rand.NewSource(1))
	for i := 0; i < nInputs; i++ {
		in := randInput(rnd)
		zbknutest(t, in.x, in.is, in.tol, in.n, in.yr, in.yi, in.kode)
	}
}

func TestZasyi(t *testing.T) {
	rnd := rand.New(rand.NewSource(1))
	for i := 0; i < nInputs; i++ {
		in := randInput(rnd)
		zasyitest(t, in.x, in.is, in.tol, in.n, in.yr, in.yi, in.kode)
	}
}

func TestZseri(t *testing.T) {
	rnd := rand.New(rand.NewSource(1))
	for i := 0; i < nInputs; i++ {
		in := randInput(rnd)
		zseritest(t, in.x, in.is, in.tol, in.n, in.yr, in.yi, in.kode)
	}
}

func TestZmlri(t *testing.T) {
	rnd := rand.New(rand.NewSource(1))
	for i := 0; i < nInputs; i++ {
		in := randInput(rnd)
		zmlritest(t, in.x, in.is, in.tol, in.n, in.yr, in.yi, in.kode)
	}
}

func TestZkscl(t *testing.T) {
	rnd := rand.New(rand.NewSource(1))
	for i := 0; i < nInputs; i++ {
		in := randInput(rnd)
		zkscltest(t, in.x, in.is, in.tol, in.n, in.yr, in.yi)
	}
}

func TestZuchk(t *testing.T) {
	rnd := rand.New(rand.NewSource(1))
	for i := 0; i < nInputs; i++ {
		in := randInput(rnd)
		zuchktest(t, in.x, in.is, in.tol)
	}
}

func TestZs1s2(t *testing.T) {
	rnd := rand.New(rand.NewSource(1))
	for i := 0; i < nInputs; i++ {
		in := randInput(rnd)
		zs1s2test(t, in.x, in.is)
	}
}

func zs1s2test(t *testing.T, x []float64, is []int) {

	type data struct {
		ZRR, ZRI, S1R, S1I, S2R, S2I float64
		NZ                           int
		ASCLE, ALIM                  float64
		IUF                          int
	}

	input := data{
		x[0], x[1], x[2], x[3], x[4], x[5],
		is[0],
		x[6], x[7],
		is[1],
	}

	impl := func(input data) data {
		zrr, zri, s1r, s1i, s2r, s2i, nz, ascle, alim, iuf :=
			Zs1s2(input.ZRR, input.ZRI, input.S1R, input.S1I, input.S2R, input.S2I, input.NZ, input.ASCLE, input.ALIM, input.IUF)
		return data{zrr, zri, s1r, s1i, s2r, s2i, nz, ascle, alim, iuf}
	}

	comp := func(input data) data {
		zrr, zri, s1r, s1i, s2r, s2i, nz, ascle, alim, iuf :=
			amoslib.Zs1s2Fort(input.ZRR, input.ZRI, input.S1R, input.S1I, input.S2R, input.S2I, input.NZ, input.ASCLE, input.ALIM, input.IUF)
		return data{zrr, zri, s1r, s1i, s2r, s2i, nz, ascle, alim, iuf}
	}

	oi := impl(input)
	oc := comp(input)

	sameF64(t, "zs1s2 zrr", oc.ZRR, oi.ZRR)
	sameF64(t, "zs1s2 zri", oc.ZRI, oi.ZRI)
	sameF64(t, "zs1s2 s1r", oc.S1R, oi.S1R)
	sameF64(t, "zs1s2 s1i", oc.S1I, oi.S1I)
	sameF64(t, "zs1s2 s2r", oc.S2R, oi.S2R)
	sameF64(t, "zs1s2 s2i", oc.S2I, oi.S2I)
	sameF64(t, "zs1s2 ascle", oc.ASCLE, oi.ASCLE)
	sameF64(t, "zs1s2 alim", oc.ALIM, oi.ALIM)
	sameInt(t, "iuf", oc.IUF, oi.IUF)
	sameInt(t, "nz", oc.NZ, oi.NZ)
}

func zuchktest(t *testing.T, x []float64, is []int, tol float64) {
	YR := x[0]
	YI := x[1]
	NZ := is[0]
	ASCLE := x[2]
	TOL := tol

	YRfort, YIfort, NZfort, ASCLEfort, TOLfort := amoslib.ZuchkFort(YR, YI, NZ, ASCLE, TOL)
	YRamos, YIamos, NZamos, ASCLEamos, TOLamos := Zuchk(YR, YI, NZ, ASCLE, TOL)

	sameF64(t, "zuchk yr", YRfort, YRamos)
	sameF64(t, "zuchk yi", YIfort, YIamos)
	sameInt(t, "zuchk nz", NZfort, NZamos)
	sameF64(t, "zuchk ascle", ASCLEfort, ASCLEamos)
	sameF64(t, "zuchk tol", TOLfort, TOLamos)
}

func zkscltest(t *testing.T, x []float64, is []int, tol float64, n int, yr, yi []float64) {
	ZRR := x[0]
	ZRI := x[1]
	FNU := x[2]
	NZ := is[1]
	ELIM := x[3]
	ASCLE := x[4]
	RZR := x[6]
	RZI := x[7]

	yrfort := make([]float64, len(yr))
	copy(yrfort, yr)
	yifort := make([]float64, len(yi))
	copy(yifort, yi)
	ZRRfort, ZRIfort, FNUfort, Nfort, YRfort, YIfort, NZfort, RZRfort, RZIfort, ASCLEfort, TOLfort, ELIMfort :=
		amoslib.ZksclFort(ZRR, ZRI, FNU, n, yrfort[1:], yifort[1:], NZ, RZR, RZI, ASCLE, tol, ELIM)
	YRfort2 := make([]float64, len(yrfort))
	YRfort2[0] = yrfort[0]
	copy(YRfort2[1:], YRfort)
	YIfort2 := make([]float64, len(yifort))
	YIfort2[0] = yifort[0]
	copy(YIfort2[1:], YIfort)

	yramos := make([]float64, len(yr))
	copy(yramos, yr)
	yiamos := make([]float64, len(yi))
	copy(yiamos, yi)
	ZRRamos, ZRIamos, FNUamos, Namos, YRamos, YIamos, NZamos, RZRamos, RZIamos, ASCLEamos, TOLamos, ELIMamos :=
		Zkscl(ZRR, ZRI, FNU, n, yramos, yiamos, NZ, RZR, RZI, ASCLE, tol, ELIM)

	sameF64(t, "zkscl zrr", ZRRfort, ZRRamos)
	sameF64(t, "zkscl zri", ZRIfort, ZRIamos)
	sameF64(t, "zkscl fnu", FNUfort, FNUamos)
	sameInt(t, "zkscl n", Nfort, Namos)
	sameInt(t, "zkscl nz", NZfort, NZamos)
	sameF64(t, "zkscl rzr", RZRfort, RZRamos)
	sameF64(t, "zkscl rzi", RZIfort, RZIamos)
	sameF64(t, "zkscl ascle", ASCLEfort, ASCLEamos)
	sameF64(t, "zkscl tol", TOLfort, TOLamos)
	sameF64(t, "zkscl elim", ELIMfort, ELIMamos)

	sameF64S(t, "zkscl yr", YRfort2, YRamos)
	sameF64S(t, "zkscl yi", YIfort2, YIamos)
}

func zmlritest(t *testing.T, x []float64, is []int, tol float64, n int, yr, yi []float64, kode int) {
	ZR := x[0]
	ZI := x[1]
	FNU := x[2]
	KODE := kode
	NZ := is[1]

	yrfort := make([]float64, len(yr))
	copy(yrfort, yr)
	yifort := make([]float64, len(yi))
	copy(yifort, yi)
	ZRfort, ZIfort, FNUfort, KODEfort, Nfort, YRfort, YIfort, NZfort, TOLfort :=
		amoslib.ZmlriFort(ZR, ZI, FNU, KODE, n, yrfort[1:], yifort[1:], NZ, tol)
	YRfort2 := make([]float64, len(yrfort))
	YRfort2[0] = yrfort[0]
	copy(YRfort2[1:], YRfort)
	YIfort2 := make([]float64, len(yifort))
	YIfort2[0] = yifort[0]
	copy(YIfort2[1:], YIfort)

	yramos := make([]float64, len(yr))
	copy(yramos, yr)
	yiamos := make([]float64, len(yi))
	copy(yiamos, yi)
	ZRamos, ZIamos, FNUamos, KODEamos, Namos, YRamos, YIamos, NZamos, TOLamos :=
		Zmlri(ZR, ZI, FNU, KODE, n, yramos, yiamos, NZ, tol)

	sameF64(t, "zmlri zr", ZRfort, ZRamos)
	sameF64(t, "zmlri zi", ZIfort, ZIamos)
	sameF64(t, "zmlri fnu", FNUfort, FNUamos)
	sameInt(t, "zmlri kode", KODEfort, KODEamos)
	sameInt(t, "zmlri n", Nfort, Namos)
	sameInt(t, "zmlri nz", NZfort, NZamos)
	sameF64(t, "zmlri tol", TOLfort, TOLamos)

	sameF64S(t, "zmlri yr", YRfort2, YRamos)
	sameF64S(t, "zmlri yi", YIfort2, YIamos)
}

func zseritest(t *testing.T, x []float64, is []int, tol float64, n int, yr, yi []float64, kode int) {
	ZR := x[0]
	ZI := x[1]
	FNU := x[2]
	KODE := kode
	NZ := is[1]
	ELIM := x[3]
	ALIM := x[4]

	yrfort := make([]float64, len(yr))
	copy(yrfort, yr)
	yifort := make([]float64, len(yi))
	copy(yifort, yi)
	ZRfort, ZIfort, FNUfort, KODEfort, Nfort, YRfort, YIfort, NZfort, TOLfort, ELIMfort, ALIMfort :=
		amoslib.ZseriFort(ZR, ZI, FNU, KODE, n, yrfort[1:], yifort[1:], NZ, tol, ELIM, ALIM)
	YRfort2 := make([]float64, len(yrfort))
	YRfort2[0] = yrfort[0]
	copy(YRfort2[1:], YRfort)
	YIfort2 := make([]float64, len(yifort))
	YIfort2[0] = yifort[0]
	copy(YIfort2[1:], YIfort)

	yramos := make([]float64, len(yr))
	copy(yramos, yr)
	yiamos := make([]float64, len(yi))
	copy(yiamos, yi)
	ZRamos, ZIamos, FNUamos, KODEamos, Namos, YRamos, YIamos, NZamos, TOLamos, ELIMamos, ALIMamos :=
		Zseri(ZR, ZI, FNU, KODE, n, yramos, yiamos, NZ, tol, ELIM, ALIM)

	sameF64(t, "zseri zr", ZRfort, ZRamos)
	sameF64(t, "zseri zi", ZIfort, ZIamos)
	sameF64(t, "zseri fnu", FNUfort, FNUamos)
	sameInt(t, "zseri kode", KODEfort, KODEamos)
	sameInt(t, "zseri n", Nfort, Namos)
	sameInt(t, "zseri nz", NZfort, NZamos)
	sameF64(t, "zseri tol", TOLfort, TOLamos)
	sameF64(t, "zseri elim", ELIMfort, ELIMamos)
	sameF64(t, "zseri elim", ALIMfort, ALIMamos)

	sameF64S(t, "zseri yr", YRfort2, YRamos)
	sameF64S(t, "zseri yi", YIfort2, YIamos)
}

func zasyitest(t *testing.T, x []float64, is []int, tol float64, n int, yr, yi []float64, kode int) {
	ZR := x[0]
	ZI := x[1]
	FNU := x[2]
	KODE := kode
	NZ := is[1]
	ELIM := x[3]
	ALIM := x[4]
	RL := x[5]

	yrfort := make([]float64, len(yr))
	copy(yrfort, yr)
	yifort := make([]float64, len(yi))
	copy(yifort, yi)
	ZRfort, ZIfort, FNUfort, KODEfort, Nfort, YRfort, YIfort, NZfort, RLfort, TOLfort, ELIMfort, ALIMfort :=
		amoslib.ZasyiFort(ZR, ZI, FNU, KODE, n, yrfort[1:], yifort[1:], NZ, RL, tol, ELIM, ALIM)
	YRfort2 := make([]float64, len(yrfort))
	YRfort2[0] = yrfort[0]
	copy(YRfort2[1:], YRfort)
	YIfort2 := make([]float64, len(yifort))
	YIfort2[0] = yifort[0]
	copy(YIfort2[1:], YIfort)

	yramos := make([]float64, len(yr))
	copy(yramos, yr)
	yiamos := make([]float64, len(yi))
	copy(yiamos, yi)
	ZRamos, ZIamos, FNUamos, KODEamos, Namos, YRamos, YIamos, NZamos, RLamos, TOLamos, ELIMamos, ALIMamos :=
		Zasyi(ZR, ZI, FNU, KODE, n, yramos, yiamos, NZ, RL, tol, ELIM, ALIM)

	sameF64(t, "zasyi zr", ZRfort, ZRamos)
	sameF64(t, "zasyi zr", ZIfort, ZIamos)
	sameF64(t, "zasyi fnu", FNUfort, FNUamos)
	sameInt(t, "zasyi kode", KODEfort, KODEamos)
	sameInt(t, "zasyi n", Nfort, Namos)
	sameInt(t, "zasyi nz", NZfort, NZamos)
	sameF64(t, "zasyi rl", RLfort, RLamos)
	sameF64(t, "zasyi tol", TOLfort, TOLamos)
	sameF64(t, "zasyi elim", ELIMfort, ELIMamos)
	sameF64(t, "zasyi alim", ALIMfort, ALIMamos)

	sameF64S(t, "zasyi yr", YRfort2, YRamos)
	sameF64S(t, "zasyi yi", YIfort2, YIamos)
}

func zbknutest(t *testing.T, x []float64, is []int, tol float64, n int, yr, yi []float64, kode int) {
	ZR := x[0]
	ZI := x[1]
	FNU := x[2]
	KODE := kode
	NZ := is[1]
	ELIM := x[3]
	ALIM := x[4]

	yrfort := make([]float64, len(yr))
	copy(yrfort, yr)
	yifort := make([]float64, len(yi))
	copy(yifort, yi)
	ZRfort, ZIfort, FNUfort, KODEfort, Nfort, YRfort, YIfort, NZfort, TOLfort, ELIMfort, ALIMfort :=
		amoslib.ZbknuFort(ZR, ZI, FNU, KODE, n, yrfort[1:], yifort[1:], NZ, tol, ELIM, ALIM)
	YRfort2 := make([]float64, len(yrfort))
	YRfort2[0] = yrfort[0]
	copy(YRfort2[1:], YRfort)
	YIfort2 := make([]float64, len(yifort))
	YIfort2[0] = yifort[0]
	copy(YIfort2[1:], YIfort)

	yramos := make([]float64, len(yr))
	copy(yramos, yr)
	yiamos := make([]float64, len(yi))
	copy(yiamos, yi)
	ZRamos, ZIamos, FNUamos, KODEamos, Namos, YRamos, YIamos, NZamos, TOLamos, ELIMamos, ALIMamos :=
		Zbknu(ZR, ZI, FNU, KODE, n, yramos, yiamos, NZ, tol, ELIM, ALIM)

	sameF64(t, "zbknu zr", ZRfort, ZRamos)
	sameF64(t, "zbknu zr", ZIfort, ZIamos)
	sameF64(t, "zbknu fnu", FNUfort, FNUamos)
	sameInt(t, "zbknu kode", KODEfort, KODEamos)
	sameInt(t, "zbknu n", Nfort, Namos)
	sameInt(t, "zbknu nz", NZfort, NZamos)
	sameF64(t, "zbknu tol", TOLfort, TOLamos)
	sameF64(t, "zbknu elim", ELIMfort, ELIMamos)
	sameF64(t, "zbknu alim", ALIMfort, ALIMamos)

	sameF64S(t, "zbknu yr", YRfort2, YRamos)
	sameF64S(t, "zbknu yi", YIfort2, YIamos)
}

func zairytest(t *testing.T, x []float64, kode, id int) {
	ZR := x[0]
	ZI := x[1]
	KODE := kode
	ID := id

	AIRfort, AIIfort, NZfort := amoslib.ZairyFort(ZR, ZI, ID, KODE)
	AIRamos, AIIamos, NZamos := Zairy(ZR, ZI, ID, KODE)

	sameF64(t, "zairy air", AIRfort, AIRamos)
	sameF64(t, "zairy aii", AIIfort, AIIamos)
	sameInt(t, "zairy nz", NZfort, NZamos)
}

func zacaitest(t *testing.T, x []float64, is []int, tol float64, n int, yr, yi []float64, kode int) {
	ZR := x[0]
	ZI := x[1]
	FNU := x[2]
	KODE := kode
	NZ := is[1]
	MR := is[2]
	ELIM := x[3]
	ALIM := x[4]
	RL := x[5]

	yrfort := make([]float64, len(yr))
	copy(yrfort, yr)
	yifort := make([]float64, len(yi))
	copy(yifort, yi)
	ZRfort, ZIfort, FNUfort, KODEfort, MRfort, Nfort, YRfort, YIfort, NZfort, RLfort, TOLfort, ELIMfort, ALIMfort :=
		amoslib.ZacaiFort(ZR, ZI, FNU, KODE, MR, n, yrfort[1:], yifort[1:], NZ, RL, tol, ELIM, ALIM)
	YRfort2 := make([]float64, len(yrfort))
	YRfort2[0] = yrfort[0]
	copy(YRfort2[1:], YRfort)
	YIfort2 := make([]float64, len(yifort))
	YIfort2[0] = yifort[0]
	copy(YIfort2[1:], YIfort)

	yramos := make([]float64, len(yr))
	copy(yramos, yr)
	yiamos := make([]float64, len(yi))
	copy(yiamos, yi)
	ZRamos, ZIamos, FNUamos, KODEamos, MRamos, Namos, YRamos, YIamos, NZamos, RLamos, TOLamos, ELIMamos, ALIMamos :=
		Zacai(ZR, ZI, FNU, KODE, MR, n, yramos, yiamos, NZ, RL, tol, ELIM, ALIM)

	sameF64(t, "zacai zr", ZRfort, ZRamos)
	sameF64(t, "zacai zi", ZIfort, ZIamos)
	sameF64(t, "zacai fnu", FNUfort, FNUamos)
	sameInt(t, "zacai kode", KODEfort, KODEamos)
	sameInt(t, "zacai mr", MRfort, MRamos)
	sameInt(t, "zacai n", Nfort, Namos)
	sameInt(t, "zacai nz", NZfort, NZamos)
	sameF64(t, "zacai rl", RLfort, RLamos)
	sameF64(t, "zacai tol", TOLfort, TOLamos)
	sameF64(t, "zacai elim", ELIMfort, ELIMamos)
	sameF64(t, "zacai elim", ALIMfort, ALIMamos)

	sameF64S(t, "zacai yr", YRfort2, YRamos)
	sameF64S(t, "zacai yi", YIfort2, YIamos)
}

func sameF64(t *testing.T, str string, c, native float64) {
	if math.IsNaN(c) && math.IsNaN(native) {
		return
	}
	if c == native {
		return
	}
	cb := math.Float64bits(c)
	nb := math.Float64bits(native)
	t.Errorf("Case %s: Float64 mismatch. c = %v, native = %v\n cb: %v, nb: %v\n", str, c, native, cb, nb)
}

func sameInt(t *testing.T, str string, c, native int) {
	if c != native {
		t.Errorf("Case %s: Int mismatch. c = %v, native = %v.", str, c, native)
	}
}

func sameF64S(t *testing.T, str string, c, native []float64) {
	if len(c) != len(native) {
		panic(str)
	}
	for i, v := range c {
		sameF64(t, str+"_idx_"+strconv.Itoa(i), v, native[i])
	}
}