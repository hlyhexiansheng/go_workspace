package logfile

import (
	"os"
	"io"
	"log"
	"regexp"
	"fmt"
)

const Buffer_Size = 1024 * 1024

type TailFile2 struct {
	NewLineStartMatchMode  int
	NewLineStartPreRegular string //标识新的一行开始的正则
	FileNode               string
	FileName               string
	Belong2Dir             string
	Offset                 int64
	MaxLineSize            int    //一行最大的字节数
	MaxOnceReadCount       int

	AppName                string
	Domain                 string
	Topic                  string
	LogType                string

	fileFd                 *os.File
	buffer                 []byte
	bufferCapacity         int
	bufferPosition         int
}

func (this *TailFile2) String() string {
	return fmt.Sprintf("tailfile2=[NewLineStartPreAscii:%d], [FileNode:%s],[FileName:%s],[Belong2Dir:%s],[Offset:%d],[MaxLineSize:%d],[UseDefaultLineStart:%v],[appname:%s],[domain:%s],[topic:%s],[logtype:%s]>",
		this.NewLineStartPreRegular, this.FileNode, this.FileName, this.Belong2Dir, this.Offset, this.MaxLineSize, this.AppName, this.Domain, this.Topic, this.LogType)
}

type LineInfo struct {
	Content string
	IsMatch bool
	Offset  int
	Size    int
}

func NewTailFile2(fileNode, fileName, belong2Dir string,
offset int64,
newLineStartMatchMode int,
newLineStartPreRegular string,
maxLineSize int,
maxOnceReadCount int,
appName string,
domain string,
topic string,
logType string) *TailFile2 {
	tailFile := &TailFile2{FileNode:fileNode,
		FileName:fileName,
		Belong2Dir:belong2Dir,
		Offset:offset,
		NewLineStartMatchMode:newLineStartMatchMode,
		NewLineStartPreRegular:newLineStartPreRegular,
		MaxLineSize:maxLineSize,
		MaxOnceReadCount:maxOnceReadCount,
		AppName:appName,
		Domain:domain,
		Topic:topic,
		LogType:logType,
	}
	tailFile.buffer = make([]byte, Buffer_Size)
	tailFile.fileFd, _ = os.OpenFile(tailFile.FileName, os.O_RDONLY, 0444)
	return tailFile
}

func (this *TailFile2) ReadMultiLine() ([]string, []int64) {
	//1.先读取一大块内容到buffer中
	this.readChunk()
	//2.一行一行从buffer中读取出来
	lines := make([]LineInfo, 0)
	for ; ; {
		lineInfo := this.nextLine()
		if lineInfo.Content == "" {
			break
		}
		lines = append(lines, lineInfo)
	}
	//3.相等的时候，说明是buffer读满了，所以可能是最后一行那里被截断，这里把最后那行丢弃掉，留到下次读
	if this.bufferCapacity == Buffer_Size {
		for i := len(lines) - 1; i >= 0; i-- {
			if lines[i].IsMatch {
				this.bufferPosition = lines[i].Offset
				lines = lines[0:i]
				break
			}
		}
	}

	//4.把读到行，拼接起来
	resultLine := make([]string, 0)
	resultOffset := make([]int64, 0)
	length := len(lines)
	line := ""
	for i := 0; i < length; i++ {
		if lines[i].IsMatch {
			line = lines[i].Content
			resultOffset = append(resultOffset, this.Offset + int64(lines[i].Offset))
		} else {
			line = line + lines[i].Content
		}
		if i + 1 == length || lines[i + 1].IsMatch {
			resultLine = append(resultLine, line)
			line = ""
		}
	}
	return resultLine, resultOffset
}

func (this *TailFile2) Commit() {
	this.Offset += int64(this.bufferPosition)
}

func (this *TailFile2) nextLine() (LineInfo) {

	lineInfo := LineInfo{Content:"", IsMatch:false, Offset:this.bufferPosition}
	tempArray := make([]byte, 0)
	for ; ; {
		if len(tempArray) >= this.MaxLineSize {
			break
		}

		c := this.readByte()
		if c == -1 {
			this.backStep(len(tempArray))
			return lineInfo
		} else if c == Enter_Ascii {
			tempArray = append(tempArray, byte(c))
			break
		} else {
			tempArray = append(tempArray, byte(c))
		}
	}

	if len(tempArray) > 0 {
		var isMatch bool
		if this.NewLineStartMatchMode == 0 {
			isMatch = true
		} else if this.NewLineStartMatchMode == 1 {
			isMatch = (tempArray[0] == this.NewLineStartPreRegular[0])
		} else {
			isMatch, _ = regexp.MatchString(this.NewLineStartPreRegular, string(tempArray))
		}

		if len(tempArray) >= this.MaxLineSize {
			//当一行超过最大限制的值时，为了防止下一步拼接，直接认为它是匹配正则的.
			isMatch = true
		}
		lineInfo.IsMatch = isMatch
		lineInfo.Content = string(tempArray)
		lineInfo.Size = len(tempArray)
	}
	return lineInfo

}

func (this *TailFile2) readByte() int8 {
	if this.bufferPosition >= this.bufferCapacity {
		return -1
	}
	c := this.buffer[this.bufferPosition]
	this.bufferPosition++
	return int8(c)
}

func (this *TailFile2) backStep(count int) {
	if count > this.bufferPosition {
		count = this.bufferPosition
	}
	this.bufferPosition -= count
}

func (this *TailFile2) readChunk() {
	size, err := this.fileFd.ReadAt(this.buffer, this.Offset)
	if err != nil && err != io.EOF {
		log.Println("read file error", err)
	}
	this.bufferPosition = 0
	this.bufferCapacity = size
}