package main

import "github.com/fatih/color"

type directwriter struct {
	c color.Color
}

func newDirectWriter(attb color.Attribute) *directwriter {
	return &directwriter{c: *color.New(attb)}
}

func (w directwriter) Write(p []byte) (int, error) {
	return w.c.Printf("%s", p)
}
