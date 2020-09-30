package filter

type Filter interface {
	Process(srcPath, dstPath string) error
	Name() string
}
