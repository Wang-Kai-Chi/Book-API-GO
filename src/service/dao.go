package service

type CrudDao interface {
	QueryWithLimit()
	Insert()
	Update()
	Delete()
}
