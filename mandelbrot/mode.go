// /home/krylon/go/src/github.com/blicero/buddhabrot/mandelbrot/mode.go
// -*- mode: go; coding: utf-8; -*-
// Created on 25. 11. 2026 by Benjamin Walkenhorst
// (c) 2026 Benjamin Walkenhorst
// Time-stamp: <2026-11-25 10:32:16 krylon>

package mandelbrot

//go:generate stringer -type=Mode

// Mode defines what kind of calculation we perform
type Mode uint8

// Mandelbrot means we render the good old Mandelbrot set.
//
// Buddhabrot means we render the Buddhabrot.
//
// Nebulabrot means we render the Buddhabrot multiple times, in different
// colors, and then stack the results.
const (
	Mandelbrot Mode = iota
	Buddhabrot
	Nebulabrot
)
