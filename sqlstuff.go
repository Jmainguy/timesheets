package main

import (
    "strings"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func addServiceSQL(name, groupname string) {

	db, err := sql.Open("mysql", "timesheet:yaya@tcp(phy01.standouthost.com)/timesheet?charset=utf8")
	check(err)

	// insert
	stmt, err := db.Prepare("INSERT services SET name=?,groupname=?")
	check(err)

	_, err = stmt.Exec(name, groupname)
    if err != nil && strings.Contains(err.Error(), "Error 1062") {
        fmt.Println("Record already exists")
        check(err)
    } else {
        check(err)
    }
	db.Close()
}

func listServices() (SL []Service) {
    var S Service
	db, err := sql.Open("mysql", "timesheet:yaya@tcp(phy01.standouthost.com)/timesheet?charset=utf8")
	check(err)
	// query
	rows, err := db.Query("SELECT * FROM services")
	check(err)
	for rows.Next() {
        err := rows.Scan(&S.Name, &S.GroupName)
        check(err)
        SL = append(SL, S)
    }
    check(err)
    rows.Close()
	db.Close()
    return SL
}

func addClientSQL(firstname, lastname, middlename string) {

	db, err := sql.Open("mysql", "timesheet:yaya@tcp(phy01.standouthost.com)/timesheet?charset=utf8")
	check(err)

	// insert
	stmt, err := db.Prepare("INSERT clients SET firstname=?,lastname=?,middlename=?")
	check(err)

	_, err = stmt.Exec(firstname, lastname, middlename)
    if err != nil && strings.Contains(err.Error(), "Error 1062") {
        fmt.Println("Record already exists")
        check(err)
    } else {
        check(err)
    }
	db.Close()
}

func listClients() (CL []Client) {
    var C Client
	db, err := sql.Open("mysql", "timesheet:yaya@tcp(phy01.standouthost.com)/timesheet?charset=utf8")
	check(err)
	// query
	rows, err := db.Query("SELECT * FROM clients")
	check(err)
	for rows.Next() {
        err := rows.Scan(&C.ID, &C.FirstName, &C.LastName, &C.MiddleName)
        check(err)
        CL = append(CL, C)
    }
    check(err)
    rows.Close()
	db.Close()
    return CL
}
    /*
	// update
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("astaxieupdate", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	// query
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}
    */

	// delete
    /*
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()
    */
