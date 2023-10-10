package brew

import (
	"archive/tar"

	"github.com/bketelsen/fleekgenbrew/bling"
	"github.com/bketelsen/fleekgenbrew/brew/brewfile"
	"github.com/bketelsen/fleekgenbrew/brew/config"
	"github.com/bketelsen/fleekgenbrew/brew/hook"
)

type Global struct {
	Bling *bling.Bling
	Hook  string
}

func FromBling(b *bling.Bling) *Global {
	c := &Global{}
	c.Bling = b
	return c
}

func (g *Global) Files() (map[string][]byte, error) {
	files := make(map[string][]byte)
	scripts, err := config.FromBling(g.Bling)
	if err != nil {
		return files, err
	}
	brewfile, err := brewfile.FromBling(g.Bling)
	if err != nil {
		return files, err
	}
	initsh, err := hook.FromBling(g.Bling)
	if err != nil {
		return files, err
	}
	zshrc, err := hook.Zshrc(g.Bling)
	if err != nil {
		return files, err
	}
	bashrc, err := hook.Bashrc(g.Bling)
	if err != nil {
		return files, err
	}
	configfiles := g.Bling.ConfigFiles()
	for name, content := range configfiles {
		files[".fleek/configs/"+name] = []byte(content)
	}
	files[".fleek/Brewfile"] = brewfile
	files[".fleek/scripts.sh"] = scripts
	files[".fleek/init.sh"] = initsh
	files[".fleek/zsh/.zshrc"] = zshrc
	files[".fleek/bash/.bashrc"] = bashrc
	return files, nil
}

func (g *Global) Write(files map[string][]byte, w *tar.Writer) error {

	for name, content := range files {
		hdr := &tar.Header{
			Name: name,
			Mode: 0644,
			Size: int64(len(content)),
		}
		if err := w.WriteHeader(hdr); err != nil {
			return err
		}
		if _, err := w.Write([]byte(content)); err != nil {
			return err
		}
	}
	return nil
}
