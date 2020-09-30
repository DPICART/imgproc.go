package task

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"imgproc.go/filter"
)

type ChanTask struct {
	dirCtx
	Filter   filter.Filter
	PoolSize int
}

func NewChanTask(srcDir, dstDir string, filter filter.Filter, poolSize int) Tasker {
	return &ChanTask{
		Filter: filter,
		dirCtx: dirCtx{
			SrcDir: srcDir,
			DstDir: dstDir,
			files:  buildFileList(srcDir),
		},
		PoolSize: poolSize,
	}
}

type jobReq struct {
	src string
	dst string
}

func (c *ChanTask) Process() error {
	size := len(c.files)
	jobs := make(chan jobReq, size)
	results := make(chan string, size)

	dstDir := path.Join(c.DstDir, c.Filter.Name())
	err := os.MkdirAll(dstDir, os.ModePerm)
	if nil != err {
		fmt.Printf("failed to create dstDir=%v", dstDir)
		os.Exit(1)
	}

	// init workers
	for w := 1; w <= c.PoolSize; w++ {
		go worker(w, c, jobs, results)
	}

	for _, f := range c.files {
		filename := filepath.Base(f)
		dst := path.Join(c.DstDir, c.Filter.Name(), filename)
		jobs <- jobReq{
			src: f,
			dst: dst,
		}
	}
	close(jobs)

	for range c.files {
		fmt.Println(<-results)
	}

	fmt.Println("Done processing file(s)!")
	return nil
}

func worker(id int, chanTask *ChanTask, jobs <-chan jobReq, results chan<- string) {
	fmt.Printf("Processing %d", id)
	for j := range jobs {
		fmt.Printf("Worker %d, started job %v\n", id, j)
		chanTask.Filter.Process(j.src, j.dst)
		results <- j.dst
	}
}
