package cq

import (
	"encoding/binary"
	"io"
	"math"
	"math/cmplx"
)

// Math Utils

func mean(fs []float64) float64 {
	s := 0.0
	for _, v := range fs {
		s += v
	}
	return s / float64(len(fs))
}

func maxidx(row []complex128) int {
	idx, max := 0, cmplx.Abs(row[0])
	for i, v := range row {
		vAbs := cmplx.Abs(v)
		if vAbs > max {
			idx, max = i, vAbs
		}
	}
	return idx
}

func complexTimes(c complex128, f float64) complex128 {
	return complex(real(c)*f, imag(c)*f)
}

// IsPowerOf2 returns true if x is a power of 2, else false.
func isPowerOf2(x int) bool {
	return x&(x-1) == 0
}

// NextPowerOf2 returns the next power of 2 >= x.
func nextPowerOf2(x int) int {
	if isPowerOf2(x) {
		return x
	}
	return int(math.Pow(2, math.Ceil(math.Log2(float64(x)))))
}

func round(x float64) int {
	return int(x + 0.5)
}
func roundUp(x float64) int {
	return int(math.Ceil(x))
}
func roundDown(x float64) int {
	return int(x)
}

func unsafeShift(s int) int {
	return 1 << uint(s)
}

func clampUnit(v float64) float64 {
	switch {
	case v > 1.0:
		return 1.0
	case v < -1.0:
		return -1.0
	default:
		return v
	}
}

func realParts(values []complex128) []float64 {
	n := len(values)
	reals := make([]float64, n, n)
	for i, c := range values {
		reals[i] = real(c)
	}
	return reals
}

func minInt(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func calcGcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return calcGcd(b, a%b)
}

func bessel0(x float64) float64 {
	b := 0.0
	for i := 0; i < 20; i++ {
		b += besselTerm(x, i)
	}
	return b
}

func besselTerm(x float64, i int) float64 {
	if i == 0 {
		return 1.0
	}
	f := float64(factorial(i))
	return math.Pow(x/2.0, float64(i)*2.0) / (f * f)
}

func factorial(i int) int {
	if i == 0 {
		return 1
	}
	return i * factorial(i-1)
}

// IO Utils

func WriteComplexBlock(w io.Writer, block [][]complex128) {
	WriteInt32(w, int32(len(block)))
	for _, b := range block {
		WriteComplexArray(w, b)
	}
}

func WriteComplexArray(w io.Writer, array []complex128) {
	WriteInt32(w, int32(len(array)))
	for _, c := range array {
		WriteComplex(w, c)
	}
}

func WriteComplex(w io.Writer, c complex128) {
	WriteFloat32(w, float32(real(c)))
	WriteFloat32(w, float32(imag(c)))
}

func WriteInt32(w io.Writer, i int32) {
	binary.Write(w, binary.LittleEndian, i)
}

func WriteFloat32(w io.Writer, f float32) {
	binary.Write(w, binary.LittleEndian, f)
}
