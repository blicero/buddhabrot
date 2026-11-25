// /home/krylon/go/src/github.com/blicero/buddhabrot/mandelbrot/mandelbrot.go
// -*- mode: go; coding: utf-8; -*-
// Created on 24. 11. 2026 by Benjamin Walkenhorst
// (c) 2026 Benjamin Walkenhorst
// Time-stamp: <2026-11-25 10:43:52 krylon>

package mandelbrot

import "math/cmplx"

const maxAbs = 4.0

// Pixel defines a point in the complex plane, and its Mandelbrot number.
type Pixel struct {
	Pt  complex128
	Val int
}

// RenderSpec defines the parameters for rendering the Mandelbrot set.
type RenderSpec struct {
	Xmin       float64
	Xmax       float64
	Ymin       float64
	Ymax       float64
	Resolution [2]uint32
	MaxIter    uint32
	Mode       Mode
}

// CalcMandelbrotPoint calculates if a number n is part of the Mandelbrot set.
// If it is, the function returns 0. If not, it returns a positive integer
// to show how much <n> is not a part of the set.
func (r *RenderSpec) CalcMandelbrotPoint(n complex128) int {
	var p = n

	for i := 0; i < int(r.MaxIter); i++ {
		if cmplx.Abs(p) > maxAbs {
			return i
		}

		p = p*p + n
	}

	return 0
} // func CalcMandelbrotPoint(n complex128) int

func (r *RenderSpec) CalcMandelbrotRow(yOffset float64) []Pixel {
	var (
		row  = make([]Pixel, r.Resolution[0])
		step = (r.Xmax - r.Xmin) / float64(r.Resolution[0])
		xval = r.Xmin
	)

	for i := 0; i < int(r.Resolution[0]); i++ {
		var pt = Pixel{
			Pt: complex(xval, yOffset),
		}

		pt.Val = r.CalcMandelbrotPoint(pt.Pt)
		row[i] = pt
		xval += step
	}

	return row
} // func (r *RenderSpec) CalcMandelbrotRow(yOffset float64) []int
