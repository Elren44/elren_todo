package model

import "github.com/Elren44/elren_todo/internal/forms"

type TemplateData struct {
	Title     string
	Tasks     []Task
	SCRFToken string
	StringMap map[string]string
	Form      *forms.Form
	Data      map[string]interface{}
	Flash     string
	Warning   string
	Error     string
}
