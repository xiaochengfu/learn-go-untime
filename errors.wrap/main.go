package main

import (
	"database/sql"
	"github.com/pkg/errors"
	"log"
	"reflect"
)

func fn() error {
	//e1 := errors.New("error")
	e2 := errors.Wrap(sql.ErrConnDone, "inner")
	e3 := errors.Wrap(e2, "middle")
	return errors.Wrap(e3, "outer")
}

func foo() error {
	return errors.Wrap(sql.ErrNoRows, "sql failed")
}

func main() {
	var err error
	err = errors.New("bar failed")

	//wrap错误包装,将foo的错误一同返回
	err = errors.Wrap(foo(), "bar failed")

	err = fn()
	log.Println(reflect.TypeOf(errors.Cause(err)), sql.ErrNoRows)
	switch errors.Cause(err) {
	case sql.ErrNoRows:
		log.Printf("data not found, %v\n", err)
	case sql.ErrConnDone:
		log.Printf("sql connectiong found, %v\n", err)
	default:
		log.Printf("%v\n", err)
	}

}
