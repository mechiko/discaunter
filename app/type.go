package app

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"discaunter/config"
	"discaunter/domain"
	"time"

	"github.com/google/uuid"
)

// type IApp interface {
// 	Options() *config.Configuration
// 	SaveOptions(string, any) error
// 	Logger() *zap.SugaredLogger
// }

type app struct {
	uuid    string // идентификатор для уникальности формы
	config  *config.Config
	options *config.Configuration // копия config.Configuration
	// loger   *zap.SugaredLogger
	pwd    string
	output string
}

var _ domain.Apper = (*app)(nil)

// const modError = "app"

func New(cfg *config.Config, pwd string) *app {
	newApp := &app{}
	newApp.pwd = pwd
	// newApp.loger = logger
	newApp.config = cfg
	newApp.options = cfg.Configuration()
	newApp.uuid = uuid.New().String()
	return newApp
}

func (a *app) NowDateString() string {
	n := time.Now()
	return fmt.Sprintf("%4d.%02d.%02d %02d:%02d:%02d", n.Local().Year(), n.Local().Month(), n.Local().Day(), n.Local().Hour(), n.Local().Minute(), n.Local().Second())
}

func (a *app) Pwd() string {
	return a.pwd
}

func (a *app) Output() string {
	return a.output
}

func (a *app) Config() *config.Config {
	return a.config
}

// func (a *app) Logger() *zap.SugaredLogger {
// 	return a.loger
// }

// выдаем адрес структуры опций программы чтобы править по месту
func (a *app) Options() *config.Configuration {
	return a.options
}

// записываем ключ и его значение только в пакет config
// изменения записываются в файл конфигурации
func (a *app) SaveOptions(key string, value any) error {
	a.config.SetInConfig(key, value)
	if err := a.config.Save(); err != nil {
		return fmt.Errorf("save in config error %w", err)
	}
	return nil
}

// изменения записываются в файл конфигурации
func (a *app) SaveAllOptions() error {
	if err := a.config.Save(); err != nil {
		return fmt.Errorf("save all in config error %w", err)
	}
	return nil
}

// создаем по необходимости пути программы
func (a *app) CreatePath() error {
	// создаем папку вывода если не пустое значение
	// в папке запуска программы только или если она задана абсолютным значением пути
	if a.options == nil {
		return fmt.Errorf("опции программы не инициализированы")
	}
	if a.options.Output != "" {
		if output, err := createPath(a.options.Output, ""); err != nil {
			return fmt.Errorf("ошибка создания каталога %w", err)
		} else {
			a.options.Output = output
		}
		// a.loger.Infof("путь output приложения %s", a.options.Output)
	}
	return nil
}

// создаем путь в каталоге программы или home
func createPath(path string, home string) (string, error) {
	fullPath := filepath.Join(home, path)
	if filepath.IsAbs(path) {
		fullPath = path
	}
	if err := pathCreate(fullPath); err != nil && !errors.Is(err, fs.ErrExist) {
		return "", fmt.Errorf("cannot create path %s: %w", fullPath, err)
	}
	return filepath.Abs(fullPath)
}

func pathCreate(path string) error {
	if path != "" {
		// if err := os.MkdirAll(path, os.ModePerm); err != nil { // создает весь путь
		if err := os.Mkdir(path, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

func (a *app) ConfigPath() string {
	if a.config != nil {
		return a.config.ConfigPath()
	}
	return ""
}

func (a *app) DbPath() string {
	if a.config != nil {
		return a.config.DbPath()
	}
	return ""
}

func (a *app) LogPath() string {
	if a.config != nil {
		return a.config.LogPath()
	}
	return ""
}
