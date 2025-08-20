package main

import (
	"discaunter/app"
	"discaunter/config"
	"discaunter/processing"
	"discaunter/xmltmpl"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mechiko/utility"
)

var fileExe string
var dir string

// если local true то папка создается локально
var home = flag.Bool("home", false, "")
var file = flag.String("file", "", "file to parse xlsx")

func init() {
	flag.Parse()
	fileExe = os.Args[0]
	var err error
	dir, err = filepath.Abs(filepath.Dir(fileExe))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get absolute path: %v\n", err)
		os.Exit(1)
	}
	if err := os.Chdir(dir); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to change directory: %v\n", err)
		os.Exit(1)
	}
}

func errMessageExit(title string, errDescription string) {
	utility.MessageBox(title, errDescription)
	os.Exit(-1)
}

func main() {
	cfg, err := config.New("", *home)
	if err != nil {
		errMessageExit("ошибка конфигурации", err.Error())
	}

	// var logsOutConfig = map[string][]string{
	// 	"logger": {"stdout", filepath.Join(cfg.LogPath(), config.Name)},
	// }
	// zl, err := zaplog.New(logsOutConfig, true)
	// if err != nil {
	// 	errMessageExit("ошибка создания логера", err.Error())
	// }
	// defer zl.Shutdown()

	// lg, err := zl.GetLogger("logger")
	// if err != nil {
	// 	zl.Shutdown()
	// 	errMessageExit("ошибка получения логера", err.Error())
	// }
	// loger := lg.Sugar()
	// loger.Debug("zaplog started")
	// loger.Infof("mode = %s", config.Mode)
	// if cfg.Warning() != "" {
	// 	loger.Infof("pkg:config warning %s", cfg.Warning())
	// }

	errProcessExit := func(title string, errDescription string) {
		// loger.Errorf("%s %s", title, errDescription)
		// zl.Shutdown()
		errMessageExit(title, errDescription)
	}
	// создаем приложение с опциями из конфига и логером основным
	app := app.New(cfg, dir)
	// инициализируем пути необходимые приложению
	app.CreatePath()

	fileXLSX := *file
	if fileXLSX == "" {
		fileXLSX, err = utility.DialogOpenFile([]utility.FileType{utility.Excel, utility.All}, "", "")
		if err != nil {
			// zl.Shutdown()
			errMessageExit("ошибка", err.Error())
		}
	}

	proc, err := processing.New(app)
	if err != nil {
		// zl.Shutdown()
		errProcessExit("запуск обработки с ошибкой ", err.Error())
	}

	err = proc.ReadXlsx(fileXLSX)
	if err != nil {
		// zl.Shutdown()
		errMessageExit("ошибка чтения файла", err.Error())
	}
	xml, err := xmltmpl.NewTemplate(app).StringXML(proc)
	if err != nil {
		// zl.Shutdown()
		errMessageExit("ошибка чтения файла", err.Error())
	}

	fileXml := "XML_" + filepath.Base(fileXLSX)
	fileXml = fileXml[:len(fileXml)-len(filepath.Ext(fileXLSX))]
	fileXml = utility.TimeFileName(fileXml) + ".xml"
	fileXml, err = utility.DialogSaveFile(utility.All, fileXml, ".")
	if err != nil {
		// zl.Shutdown()
		errMessageExit("ошибка диалога выбора файла для записи", err.Error())
	}
	if fileXml != "" {
		if !strings.HasSuffix(fileXml, ".xml") {
			fileXml = fmt.Sprintf("%s.xml", fileXml)
		}
		err := os.WriteFile(fileXml, xml, 0644) // 0644 represents file permissions (rw-r--r--)
		if err != nil {
			// zl.Shutdown()
			errMessageExit("ошибка диалога выбора файла для записи", err.Error())
		}
	}
}
