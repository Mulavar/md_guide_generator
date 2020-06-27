package main

import (
	"io/ioutil"
	"os"
)

var (
	prefix []byte
)

func init() {
	prefix = make([]byte, 128)
	for i, _ := range prefix {
		prefix[i] = 32
	}
}

func main() {
	file, err := os.Create("Guide.md")
	if err != nil {
		panic(err)
	}

	scan("data", 0, file)
}

func scan(path string, level int, target *os.File) error {
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		return nil
	}

	dirs := make([]os.FileInfo, 0)
	files := make([]os.FileInfo, 0)
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			dirs = append(dirs, fileInfo)
		} else {
			files = append(files, fileInfo)
		}
	}

	for _, dir := range dirs {
		target.Write(prefix[:4*level])
		target.Write([]byte("[" + dir.Name() + "](" + path + "/" + dir.Name() + ")\n"))
		scan(path+"/"+dir.Name(), level+1, target)
	}

	for _, file := range files {
		target.Write(prefix[:4*level])
		target.Write([]byte("[" + file.Name()[:len(file.Name())-3] + "](" + path + "/" + file.Name() + ")\n"))
	}

	return nil
}
