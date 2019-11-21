package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/indeedhat/tree"
)

var connection *sql.DB
var database string

func mysqlConnect() error {
	var err error
	connection, err = sql.Open("mysql", "root@/test")
	connection.SetMaxOpenConns(1)

	return err
}

func fetchRows(rows *sql.Rows) ([]map[string]string, error) {
	var data []map[string]string

	cols, err := rows.Columns()
	if nil != err {
		return data, err
	}
	defer rows.Close()

	vals := make([]sql.RawBytes, len(cols))
	scanPtr := make([]interface{}, len(vals))
	for i := range vals {
		scanPtr[i] = &vals[i]
	}

	for rows.Next() {
		err = rows.Scan(scanPtr...)
		if nil != err {
			return data, err
		}

		row := make(map[string]string)

		for i, val := range vals {

			if nil == val {
				row[cols[i]] = "NULL"
			} else {
				row[cols[i]] = string(val)
			}
		}

		data = append(data, row)
	}

	return data, nil
}

func fetchTables(database string) ([]string, error) {
	var tables []string

	log.Printf("fetching %s\n", database)
	rows, err := query(database, "SHOW TABLES")
	if nil != err {
		return tables, err
	}

	data, err := fetchRows(rows)
	if nil != err {
		return tables, err
	}

	field := fmt.Sprintf("Tables_in_%s", database)
	for _, row := range data {
		tables = append(tables, row[field])
	}

	return tables, nil
}

func fetchDatabases() ([]string, error) {
	var databases []string

	rows, err := connection.Query("SHOW DATABASES")
	if nil != err {
		return databases, err
	}

	data, err := fetchRows(rows)
	if nil != err {
		return databases, err
	}

	for _, row := range data {
		databases = append(databases, row["Database"])
	}

	database = databases[0]
	return databases, nil
}

func query(database, query string, args ...interface{}) (*sql.Rows, error) {
	_, err := connection.Exec(fmt.Sprintf("USE %s", database))
	if nil != err {
		return nil, err
	}

	rows, err := connection.Query(query)
	if nil != err {
		return nil, err
	}

	return rows, nil
}

func selectDatabase() {
	selected := fmt.Sprintf("\033[1m%s\033[0m", database)

	for _, l := range tre.Root.Limbs {
		switch branch := l.(type) {
		case *tree.Branch:
			if branch.Key == database {
				branch.Text = selected
			} else {
				branch.Text = branch.Key
			}
		}
	}
}
