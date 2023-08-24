package main

type CrudDao interface {
	QueryWithLimit()
	Insert()
	Update()
	Delete()
}
