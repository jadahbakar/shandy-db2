package main

import (
	"database/sql"
	"fmt"

	_ "github.com/ibmdb/go_ibm_db"
)

func Create_Con(con string) *sql.DB {
	db, err := sql.Open("go_ibm_db", con)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return db
}

func execquery(st *sql.Stmt) error {
	rows, err := st.Query()
	if err != nil {
		return err
	}
	cols, _ := rows.Columns()
	fmt.Printf("%s    %s   \n", cols[0], cols[1])
	fmt.Println("-------------------------------------")
	defer rows.Close()
	for rows.Next() {
		var t, x string
		err = rows.Scan(&t, &x)
		if err != nil {
			return err
		}
		fmt.Printf("%v  %v   \n", t, x)
	}
	return nil
}

func display(db *sql.DB) error {
	st, err := db.Prepare("select * from mahasiswa")
	if err != nil {
		return err
	}
	err = execquery(st)
	if err != nil {
		return err
	}
	return nil
}
func main() {
	con := "HOSTNAME=localhost;DATABASE=TESTDB;PORT=50000;UID=DB2INST1;PWD=password"

	type Db *sql.DB
	var re Db
	re = Create_Con(con)

	err := display(re)
	if err != nil {
		fmt.Println(err)
	}
}
