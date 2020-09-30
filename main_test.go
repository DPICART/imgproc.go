package main

import (
	"testing"

	"imgproc.go/filter"
	"imgproc.go/task"
)

const srcDir = "imgs/"
const dstDir = "output"

func BenchmarkGrayscaleWaitgroup(b *testing.B) {
	for n := 0; n < b.N; n++ {
		f := filter.Grayscale{}
		t := task.NewWaitGroupTask(srcDir, dstDir, f)
		t.Process()
	}
}

func BenchmarkGrayscaleChannelPoolSize1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		f := filter.Grayscale{}
		t := task.NewChanTask(srcDir, dstDir, f, 1)
		t.Process()
	}
}

func BenchmarkGrayscaleChannelPoolSize2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		f := filter.Grayscale{}
		t := task.NewChanTask(srcDir, dstDir, f, 2)
		t.Process()
	}
}

func BenchmarkGrayscaleChannelPoolSize4(b *testing.B) {
	for n := 0; n < b.N; n++ {
		f := filter.Grayscale{}
		t := task.NewChanTask(srcDir, dstDir, f, 4)
		t.Process()
	}
}

func BenchmarkGrayscaleChannelPoolSize8(b *testing.B) {
	for n := 0; n < b.N; n++ {
		f := filter.Grayscale{}
		t := task.NewChanTask(srcDir, dstDir, f, 4)
		t.Process()
	}
}
