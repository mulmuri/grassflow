package model

import "time"

type File struct {
	Name    string
	Path    string
	Parent  string
	ModTime time.Time
}


type Task struct {
	Message  string
	Branch   string
	Dir    []string
	Match  []string	
}