package config

import (
	"bytes"
	"embed"
	"text/template"

	"github.com/bketelsen/fleekgenbrew/bling"
)

var (
	tmplNewBuf = bytes.NewBuffer(make([]byte, 0, 4096))
)

func FromBling(bling *bling.Bling) ([]byte, error) {
	var err error
	tmplNewBuf.Reset()
	// TODO: cache template parsing
	tmpl, err := template.ParseFS(tmplFS, "tmpl/scripts.sh.tmpl")
	if err != nil {
		return []byte{}, err
	}
	err = tmpl.Execute(tmplNewBuf, scripts)
	if err != nil {
		return []byte{}, err
	}
	return tmplNewBuf.Bytes(), nil
}

//go:embed tmpl/*
var tmplFS embed.FS
