package main

import "html/template"

type Blog struct {
	Id      int           `json:"id"`
	Author  string        `json:"author"`
	Title   string        `json:"title"`
	Content template.HTML `json:"content"`
	Posted  string        `json:"posted"`
}
