# imgproc.go
## Crashcourse - Using channel or waitgroup with go
Small project based on the online course "Le language Go | Formation complète"

https://www.udemy.com/course/le-langage-go-formation-complete

## How to use ?
You'll need go installed on your machine.

To apply a filter on a folder of images:

Blur filter: ```go run main.go -src imgs/ -dst output -filter blur -task waitgroup```
Grayscale filter: ```go run main.go -src imgs/ -dst output -filter grayscale -task waitgroup```

Blur filter: ```go run main.go -src imgs/ -dst output -filter blur -task channel -poolsize 2```
Grayscale filter: ```go run main.go -src imgs/ -dst output -filter grayscale -task channel -poolsize 2```

## Result
![Curiosity before](https://github.com/DPICART/imgproc.go/raw/master/imgs/curiosity_selfie.jpeg)
![Curiosity with blur](https://github.com/DPICART/imgproc.go/raw/master/output/blur/curiosity_selfie.jpeg)
![Curiosity with grayscale](https://github.com/DPICART/imgproc.go/raw/master/output/grayscale/curiosity_selfie.jpeg)