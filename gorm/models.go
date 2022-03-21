package main

type Task struct {
	Idtasks     uint
	Description string
}

type Comment struct {
	Idcomments uint
	Idtasks    uint
	Text       string
}
