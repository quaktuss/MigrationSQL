package main

import (
	"database/sql"
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

func newTournament(db *sql.DB) {
	rows, _ := db.Query(`SELECT idTournament, idGame, date, duration FROM tournament`)
	database, err := sql.Open("sqlite3", "./dest.sqlite")
	CheckErr(err)

	var idTournament, idGame, duration int
	var date string

	for rows.Next() {
		err := rows.Scan(&idTournament, &idGame, &date, &duration)
		CheckErr(err)

		stmt, err1 := database.Prepare(`INSERT INTO tournament(idTournament, idPlace, idGame, date, duration) values(?,?,?,?,?)`)
		CheckErr(err1)

		_, err2 := stmt.Exec(idTournament, 3, idGame, date, duration)
		CheckErr(err2)
	}
	defer database.Close()
}

func newStaff(db *sql.DB) {
	rows, _ := db.Query(`SELECT idStaff FROM staff`)
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

/*func newPlace(db *sql.DB) {
	database, err := sql.Open("sqlite3", "./dest.sqlite")
	CheckErr(err)

	var tournament = tournamentMigration(db)
	stmt, err1 := database.Prepare(`INSERT INTO place(idPlace, name, address, city) values(?,?,?,?)`)
	CheckErr(err1)
	_, err2 := stmt.Exec(tournament.IdTournament, 0, tournament.IdGame, tournament.Date, tournament.Duration)
	CheckErr(err2)
	defer database.Close()

}*/

/*func newEmployeeData(db *sql.DB) {
	database, err := sql.Open("sqlite3", "./dest.sqlite")
	CheckErr(err)

	var tournament = tournamentMigration(db)
	stmt, err1 := database.Prepare(`INSERT INTO employee_data(idEmployee, lastname, firstname, gender, age, wage) values(?,?,?,?,?,?)`)
	CheckErr(err1)
	_, err2 := stmt.Exec(tournament.IdTournament, 0, tournament.IdGame, tournament.Date, tournament.Duration)
	CheckErr(err2)
	defer database.Close()

}*/

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	sqliteDatabase, err := sql.Open("sqlite3", "./old-database.sqlite")
	CheckErr(err)

	//newCoach(sqliteDatabase)
	//newGame(sqliteDatabase)
	//newPlayer(sqliteDatabase)
	//newTournament(sqliteDatabase)
	newStaff(sqliteDatabase)
	//newEmployeeData(sqliteDatabase)
	//newPlace(sqliteDatabase)

}
