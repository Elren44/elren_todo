package model

type TemplateData struct {
	Title     string
	Tasks     []Task
	ExistUser bool
	SCRFToken string
	StringMap map[string]string
}
