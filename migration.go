package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func newCoach(db *sql.DB) {
	rows, _ := db.Query(`SELECT idCoach, idGame, lastname, firstname, gender, age, wage, licenseDate FROM coach`)
	database, err := sql.Open("sqlite3", "./dest.sqlite")
	CheckErr(err)

	var idCoach, idGame, age, wage int
	var lastname, firstname, gender, licenseDate string
	for rows.Next() {
		err := rows.Scan(&idCoach, &idGame, &lastname, &firstname, &gender, &age, &wage, &licenseDate)
		CheckErr(err)

		stmt, err1 := database.Prepare(`INSERT INTO coach(idCoach, idGame, licenseDate, idEmployeeData) values(?,?,?,?)`)
		CheckErr(err1)
		_, err2 := stmt.Exec(idCoach, idGame, licenseDate, 5)
		CheckErr(err2)
	}
	defer database.Close()
}

func newGame(db *sql.DB) {
	rows, _ := db.Query(`SELECT idGame, Name FROM games`)
	database, err := sql.Open("sqlite3", "./dest.sqlite")
	CheckErr(err)

	var idGame int
	var name string

	for rows.Next() {
		err3 := rows.Scan(&idGame, &name)
		CheckErr(err3)

		stmt, err1 := database.Prepare(`INSERT INTO game(idGame, name) values(?,?)`)
		CheckErr(err1)

		_, err2 := stmt.Exec(idGame, name)
		CheckErr(err2)
	}
	defer database.Close()
}

func newTournament(db *sql.DB, newdb *sql.DB) {
	rows, _ := db.Query(`SELECT * FROM tournament`)
	database, err := sql.Open("sqlite3", "./dest.sqlite")
	CheckErr(err)

	var idTournament, idGame, duration int
	var date, address, city, placeName string

	for rows.Next() {
		err := rows.Scan(&idTournament, &idGame, &date, &duration, &placeName, &address, &city)
		CheckErr(err)

		stmt, err1 := database.Prepare(`INSERT INTO place(name,address,city) values(?,?,?)`)
		CheckErr(err1)

		result, err2 := stmt.Exec(placeName, address, city)
		CheckErr(err2)
		idPlace, err := result.LastInsertId()

		stmt2, err2 := database.Prepare(`INSERT INTO tournament(IdPlace,idGame,date,duration) values(?,?,?,?)`)
		CheckErr(err1)
		fmt.Println(idPlace, idGame, date, duration)
		_, err2 = stmt2.Exec(idPlace, idGame, date, duration)
		CheckErr(err2)

	}

	defer database.Close()
}

func newStaff(db *sql.DB) {
	rows, err := db.Query(`SELECT idStaff FROM staff`)
	CheckErr(err)
	database, err := sql.Open("sqlite3", "./dest.sqlite")
	CheckErr(err)

	var idStaff int

	for rows.Next() {
		err := rows.Scan(&idStaff)
		CheckErr(err)

		stmt, err1 := database.Prepare(`INSERT INTO staff(idStaff, idEmployeeData) values(?,?)`)
		CheckErr(err1)

		_, err2 := stmt.Exec(idStaff, 1234)
		CheckErr(err2)
	}
	defer database.Close()
}

func newPlayer(db *sql.DB) {
	rows, _ := db.Query(`SELECT idPlayer, idGame, ranking FROM player`)
	database, err := sql.Open("sqlite3", "./dest.sqlite")
	CheckErr(err)

	var idPlayer, idGame, ranking int

	for rows.Next() {
		err := rows.Scan(&idPlayer, &idGame, &ranking)
		CheckErr(err)

		stmt, err1 := database.Prepare(`INSERT INTO main.player(idPlayer, idGame, ranking, idEmployeeData) values(?,?,?,?)`)
		CheckErr(err1)

		_, err2 := stmt.Exec(idPlayer, idGame, ranking, 6)
		CheckErr(err2)
	}
	defer database.Close()

}

func place(db *sql.DB) {
	rows, _ := db.Query(`SELECT placeName, address,city FROM tournament`)
	database, err := sql.Open("sqlite3", "./dest.sqlite")
	CheckErr(err)
	var name, address, city string

	for rows.Next() {
		err := rows.Scan(&name, &address, &city)
		CheckErr(err)
		stmt, err1 := database.Prepare(`INSERT INTO place(name, address, city) values(?,?,?)`)
		CheckErr(err1)
		_, err2 := stmt.Exec(name, address, city)
		CheckErr(err2)
	}
	defer database.Close()

}

func employeeData(db *sql.DB) {
	rowsStaff, _ := db.Query(`SELECT lastname, firstname,gender,age,wage FROM staff`)
	database, err := sql.Open("sqlite3", "./dest.sqlite")
	CheckErr(err)
	var lastname, firstname, gender string
	var age, wage int

	for rowsStaff.Next() {
		err := rowsStaff.Scan(&lastname, &firstname, &gender, &age, &wage)
		CheckErr(err)
		stmt, err1 := database.Prepare(`INSERT INTO employee_data(lastname, firstname,gender,age,wage) values(?,?,?,?,?)`)
		CheckErr(err1)
		_, err2 := stmt.Exec(lastname, firstname, gender, age, wage)
		CheckErr(err2)
	}
	defer database.Close()

	rowsPlayer, _ := db.Query(`SELECT lastname, firstname,gender,age,wage FROM player`)
	CheckErr(err)
	var lastnameP, firstnameP, genderP string
	var ageP, wageP int

	for rowsPlayer.Next() {
		err := rowsPlayer.Scan(&lastnameP, &firstnameP, &genderP, &ageP, &wageP)
		CheckErr(err)
		stmt, err1 := database.Prepare(`INSERT INTO employee_data(lastname, firstname,gender,age,wage) values(?,?,?,?,?)`)
		CheckErr(err1)
		_, err2 := stmt.Exec(lastnameP, firstnameP, genderP, ageP, wageP)
		CheckErr(err2)
	}
	defer database.Close()

	rowsCoach, _ := db.Query(`SELECT lastname, firstname,gender,age,wage FROM coach`)
	CheckErr(err)
	var lastnameC, firstnameC, genderC string
	var ageC, wageC int

	for rowsCoach.Next() {
		err := rowsCoach.Scan(&lastnameC, &firstnameC, &genderC, &ageC, &wageC)
		CheckErr(err)
		stmt, err1 := database.Prepare(`INSERT INTO employee_data(lastname, firstname,gender,age,wage) values(?,?,?,?,?)`)
		CheckErr(err1)
		_, err2 := stmt.Exec(lastnameC, firstnameC, genderC, ageC, wageC)
		CheckErr(err2)
	}
	defer database.Close()

}

func idPlaceInTournament(db *sql.DB) {
	rows, _ := db.Query(`SELECT idPlace FROM place`)
	database, err := sql.Open("sqlite3", "./dest.sqlite")
	CheckErr(err)
	var idplace int

	for rows.Next() {
		err := rows.Scan(&idplace)
		CheckErr(err)
		stmt, err1 := database.Prepare(`INSERT INTO tournament(idPlace) values(?)`)
		CheckErr(err1)
		_, err2 := stmt.Exec(idplace)
		CheckErr(err2)
	}
	defer database.Close()
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	sqliteDatabase, err := sql.Open("sqlite3", "./old-database.sqlite")
	newsqliteDatabase, err := sql.Open("sqlite3", "./dest.sqlite")
	CheckErr(err)

	//newCoach(sqliteDatabase)
	//newGame(sqliteDatabase)
	//newPlayer(sqliteDatabase)
	newTournament(sqliteDatabase, newsqliteDatabase)
	//newStaff(sqliteDatabase)
	//employeeData(sqliteDatabase)
	//place(sqliteDatabase)
	//idPlaceInTournament(newsqliteDatabase)

}
