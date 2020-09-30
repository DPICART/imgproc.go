package task

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sync"

	"imgproc.go/filter"
)

type WaitGroupTask struct {
	dirCtx // embedded dirCtx
	Filter filter.Filter
}

func NewWaitGroupTask(srcDir, dstDir string, filter filter.Filter) Tasker {
	return &WaitGroupTask{
		Filter: filter,
		dirCtx: dirCtx{
			SrcDir: srcDir,
			DstDir: dstDir,
			files:  buildFileList(srcDir),
		},
	}
}

func (w *WaitGroupTask) Process() error {
	var wg sync.WaitGroup
	size := len(w.files)

	dstDir := path.Join(w.DstDir, w.Filter.Name())
	err := os.MkdirAll(dstDir, os.ModePerm)
	if nil != err {
		fmt.Printf("failed to create dstDir=%v", dstDir)
		os.Exit(1)
	}

	for i, f := range w.files {
		filename := filepath.Base(f)
		dst := path.Join(w.DstDir, w.Filter.Name(), filename)
		wg.Add(1)

		go w.applyFilter(f, dst, &wg, i+1, size)
	}
	wg.Wait()
	fmt.Println("Done processing file(s)!")
	return nil
}

func (w *WaitGroupTask) applyFilter(src string, dst string, wg *sync.WaitGroup, i int, size int) {
	w.Filter.Process(src, dst)
	fmt.Printf("Processed [%d/%d] %v => %v\n", i, size, src, dst)
	wg.Done()

}
