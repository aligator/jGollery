package entity

import (
	"bytes"
	"code.gitea.io/gitea/modules/log"
	"net/http"
)

type Gallery struct {
	//templ    template.Template
	Pictures PathGroup
}

func (g *Gallery) LoadPage(w http.ResponseWriter, r *http.Request) {
	pics, err := g.Pictures.GetList()
	if err != nil {
		http.NotFound(w, r)
		log.Error("Could not load gallery %s", g.Pictures.Name())
		return
	}

	buff := bytes.NewBufferString("Gallery:\n")

	for _, pic := range pics {
		if picPath, err := g.Pictures.Get(pic); err == nil {
			buff.WriteString(picPath + "\n")
		} else {
			buff.WriteString(pic + " not found\n")
		}
	}

	w.Write(buff.Bytes())
}
