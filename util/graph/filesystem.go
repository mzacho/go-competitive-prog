package graph

type Dir struct {
	Name string
	Dirs []Dir
	Files []File
	Parent *Dir
	Size int
}

type File struct {
	Name string
	Size int
}