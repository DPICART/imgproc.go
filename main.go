package main

import (
	"flag"
	"fmt"
	"time"

	"imgproc.go/filter"
	"imgproc.go/task"
)

func main() {
	srcDir := flag.String("src", "imgs/", "Input directory")
	dstDir := flag.String("dst", "output", "Output directory")
	filterType := flag.String("filter", "grayscale", "grayscale/blur")
	taskType := flag.String("task", "waitgroup", "waitgroup/channel")
	poolSize := flag.Int("poolsize", 4, "Pool size (Only relevant if 'task' is set to channel)")
	flag.Parse()

	var f filter.Filter
	switch *filterType {
	case "grayscale":
		f = filter.Grayscale{}
	case "blur":
		f = filter.Blur{}
	}
	var t task.Tasker
	switch *taskType {
	case "waitgroup":
		t = task.NewWaitGroupTask(*srcDir, *dstDir, f)
	case "channel":
		t = task.NewChanTask(*srcDir, *dstDir, f, *poolSize)
	}

	start := time.Now()
	t.Process()
	elapsed := time.Since(start)
	fmt.Printf("Image processing took %v\n", elapsed)
}
