package main

import (
	"io"

	"github.com/fatih/color"
)

type directwriter struct {
	stream io.Writer
	c      color.Color
}

func newDirectWriter(stream io.Writer, attb color.Attribute) *directwriter {
	return &directwriter{stream: stream, c: *color.New(attb)}
}

func (w directwriter) Write(p []byte) (int, error) {
	return w.c.Fprintf(w.stream, "%s", p)
}
