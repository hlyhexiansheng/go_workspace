package logfile

import (
	"os"
	"log"
	"io"
	"fmt"
)

const Enter_Ascii = 10
const BufferSize = 1024 * 1024

type TailFile struct {
	NewLineStartPreAscii int8
	FileNode             string
	FileName             string
	Belong2Dir           string
	Offset               int64
	MaxLineSize          int
	UseDefaultLineStart  bool
	MaxOnceReadCount     int

	AppName              string
	Domain               string
	Topic                string
	LogType              string

	fileFd               *os.File
	buffer               []byte
	bufferCapacity       int
	bufferPosition       int
}

func (this *TailFile) String() string {
	return fmt.Sprintf("tailfile=[NewLineStartPreAscii:%d], [FileNode:%s],[FileName:%s],[Belong2Dir:%s],[Offset:%d],[MaxLineSize:%d],[UseDefaultLineStart:%v],[appname:%s],[domain:%s],[topic:%s],[logtype:%s]>",
		this.NewLineStartPreAscii, this.FileNode, this.FileName, this.Belong2Dir, this.Offset, this.MaxLineSize, this.UseDefaultLineStart,this.AppName,this.Domain,this.Topic,this.LogType)
}

func NewTailFile(fileNode, fileName, belong2Dir string,
offset int64,
useDefaultLineStart bool,
newLineStartPreAscii int8,
maxLineSize int,
appName string,
domain string,
topic string,
logType string,
maxOnceReadCount int) *TailFile {
	tailFile := &TailFile{FileNode:fileNode,
		FileName:fileName,
		Belong2Dir:belong2Dir,
		Offset:offset,
		NewLineStartPreAscii:newLineStartPreAscii,
		MaxLineSize:maxLineSize,
		UseDefaultLineStart:useDefaultLineStart,
		AppName:appName,
		Domain:domain,
		Topic:topic,
		LogType:logType,
		MaxOnceReadCount:maxOnceReadCount,
	}
	tailFile.buffer = make([]byte, BufferSize)
	tailFile.fileFd, _ = os.OpenFile(tailFile.FileName, os.O_RDONLY, 0444)
	return tailFile
}

func (this *TailFile) ReadMultiLine() []string {
	this.batchReadChunk()
	if this.bufferCapacity == 0 {
		return nil
	}

	rs := make([]string, 0)
	for ; ; {
		line, hasNext := this.nextLine()
		if !hasNext {
			break
		}
		rs = append(rs, line)
	}

	return rs
}

func (this *TailFile) Commit() {
	this.Offset += int64(this.bufferPosition)
}

func (this *TailFile) nextLine() (string, bool) {
	if this.UseDefaultLineStart {
		return this.nextLineUseDefault()
	} else {
		return this.nextLineUseSpecificCharStart()
	}
}

func (this *TailFile) nextLineUseDefault() (string, bool) {
	tempArray := make([]byte, 0)
	for ; ; {
		if len(tempArray) >= this.MaxLineSize {
			break
		}

		c := this.readByte()
		if c == -1 {
			this.backStep(len(tempArray))
			return "", false
		} else if c == Enter_Ascii {
			tempArray = append(tempArray, byte(c))
			break
		} else {
			tempArray = append(tempArray, byte(c))
		}
	}
	if len(tempArray) > 0 {
		return string(tempArray), true
	}
	return "", false
}

func (this *TailFile) nextLineUseSpecificCharStart() (string, bool) {
	tempArray := make([]byte, 0)
	for ; ; {
		if len(tempArray) >= this.MaxLineSize {
			break
		}

		c := this.readByte()
		if c == -1 {
			this.backStep(len(tempArray))
			return "", false
		} else if c == Enter_Ascii {
			tempArray = append(tempArray, byte(c))
			c = this.readByte()
			if c == -1 {
				break
			} else if c == this.NewLineStartPreAscii {
				this.backStep(1)
				break
			} else if c == Enter_Ascii {
				this.backStep(1)
			} else {
				tempArray = append(tempArray, byte(c))
			}
		} else {
			tempArray = append(tempArray, byte(c))
		}
	}
	if len(tempArray) > 0 {
		return string(tempArray), true
	}
	return "", false
}

func (this *TailFile) readByte() int8 {
	if this.bufferPosition >= this.bufferCapacity {
		return -1
	}
	c := this.buffer[this.bufferPosition]
	this.bufferPosition++
	return int8(c)
}

func (this *TailFile) backStep(count int) {
	if count > this.bufferPosition {
		count = this.bufferPosition
	}
	this.bufferPosition -= count
}

func (this *TailFile) batchReadChunk() {
	size, err := this.fileFd.ReadAt(this.buffer, this.Offset)
	if err != nil && err != io.EOF {
		log.Println("read file error", err)
	}
	this.bufferPosition = 0
	this.bufferCapacity = size
}