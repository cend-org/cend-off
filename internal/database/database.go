package database

import (
	"database/sql"
	"github.com/cend-org/duval/internal/configuration"
	"github.com/cend-org/duval/internal/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/iancoleman/strcase"
	"github.com/jmoiron/sqlx"
	"time"
)

const (
	defaultDriver        = "mysql"
	maxOpenConnexion     = 120
	maxIdleConnexion     = 8
	maxConnexionLifeTime = time.Minute
)

var client *sqlx.DB

func init() {
	var err error
	client, err = sqlx.Connect(defaultDriver, configuration.App.DatabaseConnexionString)
	if err != nil {
		panic(err)
	}

	client.SetMaxOpenConns(maxOpenConnexion)
	client.SetMaxIdleConns(maxIdleConnexion)
	client.SetConnMaxLifetime(maxConnexionLifeTime)
	client.MapperFunc(strcase.ToSnake)
}

func CloseConnexion() {
	client.Close()
}

func Insert(T any) (lastId int64, err error) {
	var result sql.Result
	result, err = client.Exec(db.I(T))
	if err != nil {
		return 0, err
	}

	lastId, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastId, err
}

func Select(R any, Q string, A ...any) (err error) {
	err = client.Select(R, Q, A...)
	if err != nil {
		return err
	}

	return err
}

func Get(R any, Q string, A ...any) (err error) {
	err = client.Get(R, Q, A...)
	if err != nil {
		return err
	}
	return err
}

func InsertOne(T any) (id int, err error) {
	lastId, err := Insert(T)
	if err != nil {
		return 0, err
	}

	return int(lastId), err
}

func Update(T any) (err error) {
	_, err = client.Exec(db.U(T))
	if err != nil {
		return err
	}
	return err
}

func Delete(T any) (err error) {
	_, err = client.Exec(db.D(T))
	if err != nil {
		return err
	}

	return err
}

func Exec(Q string, A ...any) (err error) {
	_, err = client.Exec(Q, A...)
	if err != nil {
		return err
	}
	return err
}

func InsertMany(T []any) (err error) {
	for i := 0; i < len(T); i++ {
		_, err = InsertOne(T[i])
		if err != nil {
			return err
		}
	}
	return err
}

func GetMany(R interface{}, Q string, A ...interface{}) (err error) {
	err = client.Select(R, Q, A...)
	if err != nil {
		return err
	}
	return err
}
