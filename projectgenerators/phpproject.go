package projectgenerators

import (
	"github.com/ortizdavid/appgen/filemanager"
	//"github.com/ortizdavid/appgen/helpers"
	//dbsamples "github.com/ortizdavid/appgen/samples/db"
	phpsamples "github.com/ortizdavid/appgen/samples/php"
	//"github.com/ortizdavid/appgen/samples/scripts"
)

type PhpProject struct {}

var phpFileManager *filemanager.FileManager
var phpImport *phpsamples.AppImport


func (php *PhpProject) GetProjectTypes() []string {
	types := [] string {"3-tier", "mvc"}
	return types
}