module iknowbook.com/iknow

go 1.20

replace iknowbook.com/handler => ./handler

require (
	github.com/lib/pq v1.10.9
	iknowbook.com/handler v0.0.0
)

require (
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/jmoiron/sqlx v1.3.5 // indirect
)
