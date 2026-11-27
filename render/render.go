// /home/krylon/go/src/github.com/blicero/buddhabrot/render/render.go
// -*- mode: go; coding: utf-8; -*-
// Created on 24. 11. 2026 by Benjamin Walkenhorst
// (c) 2026 Benjamin Walkenhorst
// Time-stamp: <2026-11-27 13:08:41 krylon>

package render

import (
	"log"
	"sync"

	"github.com/blicero/buddhabrot/common"
	"github.com/blicero/buddhabrot/logdomain"
	"github.com/blicero/buddhabrot/mandelbrot"
)

type imageSlice [][]mandelbrot.Pixel

// RenderMan handles the accounting side of the job, so to speak,
// distributing the work across available workers.
type RenderMan struct {
	log            *log.Logger
	spec           *mandelbrot.RenderSpec
	path           string
	parCnt         int
	xrange, yrange float64
	xstep, ystep   float64
	wg             sync.WaitGroup
}

// New creates a fresh RenderMan.
func New(spec *mandelbrot.RenderSpec, path string, cnt int) (*RenderMan, error) {
	var (
		err error
		rm  = new(RenderMan)
	)

	if rm.log, err = common.GetLogger(logdomain.Render); err != nil {
		return nil, err
	}

	rm.spec = spec
	rm.path = path
	rm.parCnt = cnt

	return rm, nil
} // func New(spec *mandelbrot.RenderSpec, path string, cnt int) (*RenderMan, error)

func (rm *RenderMan) Initialize() {
	rm.xrange = rm.spec.Xmax - rm.spec.Xmin
	rm.yrange = rm.spec.Ymax - rm.spec.Ymin

	rm.xstep = rm.xrange / float64(rm.spec.Resolution[0])
	rm.ystep = rm.yrange / float64(rm.spec.Resolution[1])

	// Next, we need to divide the vertical range into slices for
	// our workers.

	var (
		slice  = rm.yrange / float64(rm.parCnt)
		offset = rm.spec.Ymin
		resQ   = make(chan imageSlice, 0)
	)

	for i := 0; i < rm.parCnt; i++ {

	}
} // func (rm *RenderMan) Initialize()
