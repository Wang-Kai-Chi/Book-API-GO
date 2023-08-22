package main

type CrudDao interface {
	Query()
	Insert()
	Update()
	Delete()
}
