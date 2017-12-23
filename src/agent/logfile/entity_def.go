package logfile

import "fmt"


type FileStruct struct {
	Belong2dir string        `json:"-"`
	FullName   string       `json:"filename"`
	FileNode   string        `json:"inode"`
	Offset     int64        `json:"offset"`
}

func (this *FileStruct) String() string {
	return fmt.Sprintf("<[Belong2dir:%s], [FullName:%s], [FileNode:%s]>, [Offset:%d]>\n", this.Belong2dir, this.FullName, this.FileNode, this.Offset)
}

