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
	"sync"
)

var (
	lock = new(sync.RWMutex)
)

func GetAllFile(dir string, filter string, recursion bool) []*FileStruct {

	lock.Lock()
	defer lock.Unlock()

	fileList := make([]*FileStruct, 0)

	on_each_file := func(path string, fileInfo os.FileInfo, err error) error {
		if fileInfo == nil {
			return nil
		}
		isMatch, _ := regexp.MatchString(filter, fileInfo.Name())
		if (!isMatch) {
			return nil
		}
		stat, isMatch := fileInfo.Sys().(*syscall.Stat_t)
		if (!isMatch || fileInfo.IsDir()) {
			return nil
		}
		if !recursion && (dir + "/" + fileInfo.Name()) != path { //是否递归
			return nil
		}

		file_struct := &FileStruct{FileNode:strconv.FormatUint(stat.Ino, 10), FullName:path, Belong2dir:dir}
		fileList = append(fileList, file_struct)

		return nil
	}
	filepath.Walk(dir, on_each_file)
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



