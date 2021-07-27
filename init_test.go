package main_test

import (
	"github.com/cucumber/godog"
	"github.com/spf13/pflag"
	"runtime"
)

var opts godog.Options

func init() {
	godog.BindCommandLineFlags("godog.", &opts)
	pflag.Parse()
	opts.Paths = pflag.Args()
	if opts.Concurrency < runtime.NumCPU() {
		runtime.GOMAXPROCS(opts.Concurrency)
	}
}
