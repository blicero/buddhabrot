// /home/krylon/go/src/github.com/blicero/buddhabrot/mandelbrot/mandelbrot.go
// -*- mode: go; coding: utf-8; -*-
// Created on 24. 11. 2026 by Benjamin Walkenhorst
// (c) 2026 Benjamin Walkenhorst
// Time-stamp: <2026-11-24 12:23:27 krylon>

package mandelbrot

import "math/cmplx"

const maxAbs = 4.0

// RenderSpec defines the parameters for rendering the Mandelbrot set.
type RenderSpec struct {
	Xmin       float64
	Xmax       float64
	Ymin       float64
	Ymax       float64
	Resolution [2]uint32
	MaxIter    uint32
}

func (r *RenderSpec) calcMandelbrot(n complex128) int {
	var p = n

	for i := 0; i < int(r.MaxIter); i++ {
		if cmplx.Abs(p) > maxAbs {
			return i
		}

		p = p*p + n
	}

	return 0
} // func calcMandelbrot(n complex128) int
