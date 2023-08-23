package main

type CrudDao interface {
	QueryAll()
	Insert()
	Update()
	Delete()
}
