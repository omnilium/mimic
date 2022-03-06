package main

import (
	"fmt"
	template "github.com/omnilium/mimic/internal/template"
	"os"
)

func main() {
	templateFile, err := os.OpenFile("bin/test_template.html", os.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}
	defer templateFile.Close()

	templateData := make([]byte, 1024)
	_, err = templateFile.Read(templateData)
	if err != nil {
		panic(err)
	}

	t, err := template.NewTemplate("TestTemplate", nil, nil, string(templateData))
	if err != nil {
		panic(err)
	}

	for _, token := range t.NodeList {
		fmt.Println(token.Representation(false))
	}
}
