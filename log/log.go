package log

import(
	"fmt"
	"github.com/sirupsen/logrus"
	"runtime"
	"strings"
)

//initialize logrus custom format log
func init(){
	logrus.SetReportCaller(true)
	format:=&logrus.TextFormatter{
		FullTimestamp: true,
		DisableLevelTruncation: true,
		CallerPrettyfier:func(f *runtime.Frame)(string,string){
			pathArr:=strings.Split(f.File,"/")
			//get filename
			return "",fmt.Sprintf("%s:%d",pathArr[len(pathArr)-1],f.Line)
		},
	}
	logrus.SetFormatter(format)
}
