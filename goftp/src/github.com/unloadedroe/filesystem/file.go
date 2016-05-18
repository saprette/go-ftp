package filesystem

import "github.com/unloadedroe/common"

type File struct {
	filename string
	content  []common.Chunk
}
