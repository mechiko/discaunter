package processing

import (
	"discaunter/domain"
)

// const startSSCC = "1462709225" // gs1 rus id zapivkom для памяти запивком

type Processing struct {
	Boxes map[string][]*Record
	domain.Apper
	warnings []string
	errors   []string
}

func New(app domain.Apper) (*Processing, error) {
	pr := &Processing{
		Apper:    app,
		Boxes:    make(map[string][]*Record),
		warnings: make([]string, 0),
		errors:   make([]string, 0),
	}
	return pr, nil
}

func (k *Processing) AddWarn(warn string) {
	k.warnings = append(k.warnings, warn)
}

func (k *Processing) Warnings() []string {
	return k.warnings
}

func (k *Processing) AddError(err string) {
	k.errors = append(k.errors, err)
}

func (k *Processing) Errors() []string {
	return k.errors
}

func (k *Processing) Reset() {
	k.errors = make([]string, 0)
	k.warnings = make([]string, 0)
}
