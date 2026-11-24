// /home/krylon/go/src/github.com/blicero/buddhabrot/logdomain/id.go
// -*- mode: go; coding: utf-8; -*-
// Created on 01. 02. 2023 by Benjamin Walkenhorst
// (c) 2023 Benjamin Walkenhorst
// Time-stamp: <2026-11-24 12:48:40 krylon>

// Package logdomain provides symbolic constants for the various facilities
// in the application that might want to do some logging.
package logdomain

//go:generate stringer -type=ID

type ID uint8

const (
	Common ID = iota
	Mandelbrot
	Render
)
