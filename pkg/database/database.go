package database

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"backend/pkg/config"

	_ "github.com/go-sql-driver/mysql"
)

var Handle *sql.DB
var HandleSeq *sql.DB

func InitDB() {
	dsn := config.Data.Get("DBuser") + ":" + config.Data.Get("DBpass") + "@tcp(" + config.Data.Get("DBhost") + ":" + config.Data.Get("DBport") + ")/" + config.Data.Get("DBname")
	handle, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Panic(err)
	}

	if err = handle.Ping(); err != nil {
		log.Panic(err)
	}
	Handle = handle
}

func LogSQL(query string, args ...interface{}) {
	var buffer bytes.Buffer
	nArgs := len(args)
	// Break the string by question marks, iterate over its parts and for each
	// question mark - append an argument and format the argument according to
	// it's type, taking into consideration NULL values and quoting strings.
	for i, part := range strings.Split(query, "?") {
		buffer.WriteString(part)
		if i < nArgs {
			switch a := args[i].(type) {
			case int:
				buffer.WriteString(fmt.Sprintf("%d", a))
			case int64:
				buffer.WriteString(fmt.Sprintf("%d", a))
			case bool:
				buffer.WriteString(fmt.Sprintf("%t", a))
			case sql.NullBool:
				if a.Valid {
					buffer.WriteString(fmt.Sprintf("%t", a.Bool))
				} else {
					buffer.WriteString("NULL")
				}
			case sql.NullInt64:
				if a.Valid {
					buffer.WriteString(fmt.Sprintf("%d", a.Int64))
				} else {
					buffer.WriteString("NULL")
				}
			case sql.NullString:
				if a.Valid {
					buffer.WriteString(fmt.Sprintf("%q", a.String))
				} else {
					buffer.WriteString("NULL")
				}
			case sql.NullFloat64:
				if a.Valid {
					buffer.WriteString(fmt.Sprintf("%f", a.Float64))
				} else {
					buffer.WriteString("NULL")
				}
			default:
				buffer.WriteString(fmt.Sprintf("%q", a))
			}
		}
	}
	log.Print(buffer.String())
}
