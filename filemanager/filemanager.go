package filemanager

import (
	"fmt"
	"os"
	"github.com/ortizdavid/appgen/helpers"
)

type FileManager struct {}

func (fm *FileManager) CreateFolder(folderName string) {
	folder := os.Mkdir(folderName, 0777)
	if folder != nil {
		fmt.Println(folder.Error())
	} else {
		fmt.Printf(helpers.FOLDER_CREATED, folderName)
	}
}

func (fm *FileManager) CreateFolderAll(folderName string) {
	folder := os.MkdirAll(folderName, 0777)
	if folder != nil {
		fmt.Println(folder.Error())
	} else {
		fmt.Printf(helpers.FOLDER_CREATED, folderName)
	}
}

func (fm *FileManager) CreateFile(dirName string, fileName string) {
	file, err := os.Create(dirName +"/"+ fileName)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf(helpers.FILE_CREATED, file.Name())
	}
}

func (fm *FileManager) CreateFileAll(dirName string, files ...string) {
	for _, file := range files {
		fm.CreateFile(dirName, file)
	}
}

func (fm *FileManager) MoveFile(fileName string, origin string, destination string) {
	err := os.Rename(origin +"/"+ fileName, destination +"/"+ fileName)
	if err != nil {
		fmt.Println(err)
	}
}

func (fm *FileManager) WriteFile(folderName string, fileName string, content string) {
	file, err := os.OpenFile(folderName +"/"+ fileName, os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	} else {
		file.WriteString(content)
	}
	file.Close()
}