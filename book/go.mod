module iknowbook.com/book

go 1.20

replace iknowbook.com/repository => ../repository
replace iknowbook.com/data => ../data
replace iknowbook.com/product => ../product
replace iknowbook.com/db => ../db

require (
	iknowbook.com/db v0.0.0
	iknowbook.com/repository v0.0.0
	iknowbook.com/product v0.0.0
	github.com/jmoiron/sqlx v1.3.5
	iknowbook.com/data v0.0.0
)