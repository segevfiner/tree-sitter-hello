package tree_sitter_hello_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-hello"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_hello.Language())
	if language == nil {
		t.Errorf("Error loading Hello grammar")
	}
}
