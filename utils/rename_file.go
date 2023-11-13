package utils

import (
	"fmt"
	"mime/multipart"
	"strings"
)

func RenameFile(file *multipart.FileHeader, newName string) {
	sliced := strings.Split(file.Filename, ".")
	ext := sliced[len(sliced)-1]
	newName = strings.ReplaceAll(newName, " ", "")

	file.Filename = fmt.Sprintf("%v.%v", newName, ext)
}
