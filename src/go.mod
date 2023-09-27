module iknowbook.com/iknow

go 1.20

replace iknowbook.com/handler => ./handler

replace iknowbook.com/data => ./data

replace iknowbook.com/service => ./service

require (
	github.com/gorilla/mux v1.8.0
	iknowbook.com/data v0.0.0
	iknowbook.com/handler v0.0.0
	iknowbook.com/service v0.0.0
)

require (
	github.com/jmoiron/sqlx v1.3.5 // indirect
	github.com/lib/pq v1.10.9 // indirect
)
