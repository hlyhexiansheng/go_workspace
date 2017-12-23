package logfile

import (
	"path/filepath"
	"os"
	"syscall"
	"log"
	"regexp"
	"strconv"
	"io/ioutil"
	"errors"
)

var fileList []*FileStruct
var filterKey string

func GetAllFile(dir string, filter string) []*FileStruct {

	if filter == "" {
		filterKey = "(.log$|.txt$|.tmp$)"
	} else {
		filterKey = filter
	}

	//clear the list
	fileList = make([]*FileStruct, 0)

	filepath.Walk(dir, on_each_file)

	//add Belong2dir info
	for _, file := range fileList {
		file.Belong2dir = dir
	}
	return fileList
}

func GetFileLastOffset(fileName string) int64 {
	if !FileIsExist(fileName) || IsDir(fileName) {
		return 0
	}
	file, err := os.Open(fileName)
	if err != nil {
		return 0
	}
	info, _ := file.Stat()
	file.Close()
	return info.Size()
}

func ReadFullFileContent(filename string) (string, bool) {
	byteArray, err := ioutil.ReadFile(filename)
	if (err != nil) {
		return "", false
	}
	content := string(byteArray[:])
	return content, true
}

func FileIsExist(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func OpenFileAndReadAll(filename string) ([]byte, error) {
	file, err := os.OpenFile(filename, os.O_CREATE | os.O_RDWR | os.O_SYNC, 0666)
	if err != nil {
		return nil, errors.New("open file error," + err.Error())
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.New("readall file error," + err.Error())
	}
	file.Close()
	return bytes, err
}

func OpenFileAndWriteAll(filename string, content string) (bool, error) {
	file, err := os.OpenFile(filename, os.O_CREATE | os.O_RDWR | os.O_SYNC | os.O_TRUNC, 0666)
	if err != nil {
		return false, errors.New("open file error " + err.Error())
	}

	_, err = file.WriteString(content)
	if err != nil {
		return false, errors.New("write file error" + err.Error())
	}

	file.Close()
	return true, nil
}

func IsDir(name string) bool {
	if !FileIsExist(name) {
		return false
	}
	fileInfo, err := os.Stat(name)
	if err != nil {
		return false;
	}
	return fileInfo.IsDir()
}

func IsSymlink(name string) bool {
	fileInfo, err := os.Lstat(name)
	if (err != nil) {
		log.Printf("IsSymlink error: %s", err.Error())
		return false
	}
	if (fileInfo.Mode() & os.ModeSymlink != 0) {
		return true
	}
	return false;
}

func ResolveLink(name string) (string, error) {
	file, err := os.Readlink(name)
	if (err != nil) {
		return file, err
	}
	return file, nil
}

func SelfPath() string {
	path, _ := filepath.Abs(os.Args[0])
	return path
}

func on_each_file(path string, fileInfo os.FileInfo, err error) error {
	if fileInfo == nil {
		return nil
	}

	isMatch, _ := regexp.MatchString(filterKey, fileInfo.Name())
	if (!isMatch) {
		return nil
	}

	stat, isMatch := fileInfo.Sys().(*syscall.Stat_t)
	if (!isMatch || fileInfo.IsDir()) {
		return nil
	}

	file_struct := &FileStruct{FileNode:strconv.FormatUint(stat.Ino, 10), FullName:path}
	fileList = append(fileList, file_struct)

	return nil
}

