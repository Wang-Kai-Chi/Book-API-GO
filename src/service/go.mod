module iknowbook.com/service

go 1.20

replace iknowbook.com/data => ../data
require	(
	github.com/lib/pq v1.10.9
    github.com/jmoiron/sqlx v1.3.5
    iknowbook.com/data v0.0.0
)