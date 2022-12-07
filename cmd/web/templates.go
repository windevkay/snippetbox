package main

import (
	"github.com/windevkay/snippetbox/internal/models"
)

type templateData struct {
	Snippet *models.Snippet
	Snippets []*models.Snippet
}