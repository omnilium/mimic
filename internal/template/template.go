package template

import (
	"github.com/omnilium/mimic/internal/lexer"
	"github.com/omnilium/mimic/internal/types"
)

type Origin struct {
	Name         string
	TemplateName string
	Loader       interface{}
}

type Template struct {
	Name     string
	Origin   interface{}
	Engine   interface{}
	Source   string
	NodeList []types.Token
}

func NewTemplate(name string, origin interface{}, engine interface{}, source string) (*Template, error) {
	t := &Template{
		Name:     name,
		Origin:   origin,
		Engine:   engine,
		Source:   source,
		NodeList: nil,
	}

	err := t.compileNodeList()

	return t, err
}

func (t *Template) Render(data interface{}) (string, error) {
	// TODO: Implement Render functionality
	return "", nil
}

func (t *Template) compileNodeList() error {
	lexer, err := lexer.NewLexer(t.Source)
	if err != nil {
		return err
	}

	t.NodeList = lexer.Tokenize()

	return nil

	// TODO: Implement Parser functionality
}
