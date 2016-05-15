package filesystem

import "../common"

type File struct {
	filename string
	content  []common.Chunck
}
