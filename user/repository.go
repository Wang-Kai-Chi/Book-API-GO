package user

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	. "iknowbook.com/app/data"
	. "iknowbook.com/app/repository"
)

type UserDao interface {
	UserRepository
	QueryEntity() []User
	ExecSql() sql.Result
}

type UserRepository struct {
	Connection *sqlx.DB
}

func (serv UserRepository) QueryEntity(sqlStr string, params ...interface{}) []User {
	return QueryEntity[User](serv.Connection, sqlStr, params...)
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return UserRepository{
		Connection: db,
	}
}

func (repo UserRepository) QueryWithLimit(limit int) []User {
	return repo.QueryEntity(NewUserSqlStr().QueryWithLimit, limit)
}

func (repo UserRepository) ExecSql(str string, us []User) sql.Result {
	return ExecSql[[]User](repo.Connection, str, us)
}

func (repo UserRepository) Insert(us User) sql.Result {
	return repo.Connection.MustExec(NewUserSqlStr().Insert,
		us.Name, us.Email, us.Phone, us.Password)
}
