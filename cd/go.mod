module iknowbook.com/cd

go 1.20

replace iknowbook.com/repository => ../repository

replace iknowbook.com/data => ../data

replace iknowbook.com/db => ../db

require (
	github.com/jmoiron/sqlx v1.3.5
	iknowbook.com/data v0.0.0
	iknowbook.com/db v0.0.0
	iknowbook.com/repository v0.0.0
)

require github.com/lib/pq v1.2.0 // indirect
