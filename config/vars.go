package config

// используется в Config как имя файла конфиг
const Name = "discaunter"

var Mode = "development"

// This should preferably be set at build time via build scripts
// go build -ldflags="-X discaunter/config.ExeVersion=v1.0.0"
const ExeVersion string = "0.0.1"
