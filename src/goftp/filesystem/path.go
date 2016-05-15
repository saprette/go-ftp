package filesystem

type Path struct {
	path string
}

func (p *Path) GetPath() string {
	return p.path
}

func (p *Path) setPath(pathName string) {
	p.path = pathName
}

func NewPath(pathName string) *Path {
	p := new(Path)
	p.path = pathName
	return p
}
