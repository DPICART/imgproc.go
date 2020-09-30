# imgproc.go
## Crashcourse - Using channel or waitgroup with go
Small project based on the online course "Le language Go | Formation complète"

https://www.udemy.com/course/le-langage-go-formation-complete

## How to use ?
You'll need go installed on your machine.

To apply a filter on a folder of images:


### Using Go Waitgroup

Blur filter: ```go run main.go -src imgs/ -dst output -filter blur -task waitgroup```

Grayscale filter: ```go run main.go -src imgs/ -dst output -filter grayscale -task waitgroup```

### Using Go Channel

Blur filter: ```go run main.go -src imgs/ -dst output -filter blur -task channel -poolsize 2```

Grayscale filter: ```go run main.go -src imgs/ -dst output -filter grayscale -task channel -poolsize 2```

## Small Benchmark

```go test -bench=.```

|Name   | Time (ns)  | Time (s)  |
|---|---|---|
|BenchmarkGrayscaleWaitgroup  | 5148467159 ns/op | 5.148 s/op  |
|BenchmarkGrayscaleChannelPoolSize1   | 15156505624 ns/op  | 15.156 s/op  |
|BenchmarkGrayscaleChannelPoolSize2   | 9191897129 ns/op  | 9.191 s/op  |
|BenchmarkGrayscaleChannelPoolSize4   | 6083960564 ns/op  | 6.083 s/op  |
|BenchmarkGrayscaleChannelPoolSize8   | 6245621457 ns/op  | 6.245 s/op  |


## Result
Applied filters to Curiosity's lunar selfie
![Curiosity before](https://github.com/DPICART/imgproc.go/raw/master/imgs/curiosity_selfie.jpeg)
![Curiosity with blur](https://github.com/DPICART/imgproc.go/raw/master/output/blur/curiosity_selfie.jpeg)
![Curiosity with grayscale](https://github.com/DPICART/imgproc.go/raw/master/output/grayscale/curiosity_selfie.jpeg)