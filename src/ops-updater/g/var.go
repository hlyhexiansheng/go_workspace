package g

import (
	"toolkits/file"
)

var SelfDir string

func InitGlobalVariables() {
	SelfDir = file.SelfDir()
}
