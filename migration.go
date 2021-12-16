package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func newCoach(db *sql.DB) {
	rows, err := db.Query(`SELECT idCoach, idGame, lastname, firstname, gender, age, wage, licenseDate FROM coach`)
	CheckErr(err)
	database, err := sql.Open("sqlite3", "./dest.sqlite")
	CheckErr(err)

	query := `
		create table if not exists coach
(
    idCoach        int not null
        constraint coach_pk
            primary key,
    idGame         int not null,
    licenseDate    text(30),
    idEmployeeData int not null
);
    `

	_, errDB := database.Prepare(query)
	CheckErr(errDB)

	var idCoach, idGame, age, wage int
	var lastname, firstname, gender, licenseDate string
	for rows.Next() {
		err := rows.Scan(&idCoach, &idGame, &lastname, &firstname, &gender, &age, &wage, &licenseDate)
		CheckErr(err)

		stmt, err1 := database.Prepare(`INSERT INTO Employee_Data(lastname,firstname,gender,age,wage) values(?,?,?,?,?)`)
		CheckErr(err1)
		result, err2 := stmt.Exec(lastname, firstname, gender, age, wage)
		CheckErr(err2)
		idEmployeeData, err := result.LastInsertId()

		stmt2, err2 := database.Prepare(`INSERT INTO coach(IdCoach,idGame,licenseDate, idEmployeeData) values(?,?,?,?)`)
		CheckErr(err1)
		_, err2 = stmt2.Exec(idCoach, idGame, licenseDate, idEmployeeData)
		CheckErr(err2)
	}
	defer database.Close()
}

func newGame(db *sql.DB) {
	rows, _ := db.Query(`SELECT idGame, Name FROM games`)
	database, err := sql.Open("sqlite3", "./dest.sqlite")
	CheckErr(err)

	query := `
	create table if not exists game
(
    idGame integer not null
        constraint game_pk
            primary key,
    name   text(30)
);

    `

	_, errDB := database.Prepare(query)
	CheckErr(errDB)

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
	rows, _ := db.Query(`SELECT * FROM tournament`)
	database, err := sql.Open("sqlite3", "./dest.sqlite")
	CheckErr(err)

	query := `
		create table if not exists tournament
		(
			idTournament integer  not null
				constraint tournament_pk
					primary key,
			idPlace      integer  not null,
			idGame       integer  not null,
			date         text(30) not null,
			duration     integer
		);
    `

	_, errDB := database.Prepare(query)
	CheckErr(errDB)

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
		_, err2 = stmt2.Exec(idPlace, idGame, date, duration)
		CheckErr(err2)

	}

	defer database.Close()
}

func newStaff(db *sql.DB) {
	rows, err := db.Query(`SELECT * FROM staff`)
	CheckErr(err)
	database, err := sql.Open("sqlite3", "./dest.sqlite")
	CheckErr(err)

	query := `
create table if not exists staff
(
    idStaff        integer not null
        constraint staff_pk
            primary key,
    idEmployeeData integer
);
    `

	_, errDB := database.Prepare(query)
	CheckErr(errDB)

	var idStaff, age, wage int
	var lastname, firstname, gender string

	for rows.Next() {
		err := rows.Scan(&idStaff, &lastname, &firstname, &gender, &age, &wage)
		CheckErr(err)

		stmt, err1 := database.Prepare(`INSERT INTO Employee_Data(lastname,firstname,gender,age,wage) values(?,?,?,?,?)`)
		CheckErr(err1)

		result, err2 := stmt.Exec(lastname, firstname, gender, age, wage)
		CheckErr(err2)
		idEmployeeData, err := result.LastInsertId()

		stmt2, err1 := database.Prepare(`INSERT INTO Staff(idStaff, idEmployeeData) values(?,?)`)
		CheckErr(err1)
		_, err2 = stmt2.Exec(idStaff, idEmployeeData)
		CheckErr(err2)

	}
	defer database.Close()
}

func newPlayer(db *sql.DB) {
	rows, _ := db.Query(`SELECT * FROM player`)
	database, err := sql.Open("sqlite3", "./dest.sqlite")
	CheckErr(err)

	query := `
create table if not exists player
(
    idPlayer       integer not null
        constraint player_pk
            primary key,
    idGame         integer not null,
    ranking        integer,
    idEmployeeData integer not null
);
    `

	_, errDB := database.Prepare(query)
	CheckErr(errDB)

	var idPlayer, idGame, age, wage, ranking int
	var lastname, firstname, gender string

	for rows.Next() {
		err := rows.Scan(&idPlayer, &idGame, &lastname, &firstname, &gender, &age, &wage, &ranking)
		CheckErr(err)

		stmt, err1 := database.Prepare(`INSERT INTO Employee_Data(lastname, firstname, gender, age, wage) values(?,?,?,?,?)`)
		CheckErr(err1)

		_, err2 := stmt.Exec(lastname, firstname, gender, age, wage)
		CheckErr(err2)
		result, err2 := stmt.Exec(lastname, firstname, gender, age, wage)
		CheckErr(err2)
		idEmployeeData, err := result.LastInsertId()
		stmt2, err1 := database.Prepare(`INSERT INTO player(idPlayer, idGame, ranking, idEmployeeData) values(?,?,?,?)`)
		CheckErr(err1)
		_, err2 = stmt2.Exec(idPlayer, idGame, ranking, idEmployeeData)
		CheckErr(err2)
	}
	defer database.Close()

}

func place(db *sql.DB) {
	rows, _ := db.Query(`SELECT placeName, address,city FROM tournament`)
	database, err := sql.Open("sqlite3", "./dest.sqlite")
	CheckErr(err)
	var name, address, city string

	query := `
create table if not exists place
(
    idPlace integer not null
        constraint place_pk
            primary key autoincrement,
    name    text(30),
    address text(30),
    city    text(30)
);
    `

	_, errDB := database.Prepare(query)
	CheckErr(errDB)

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

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	database := os.Args[1]

	sqliteDatabase, err := sql.Open("sqlite3", database)
	CheckErr(err)

	newCoach(sqliteDatabase)
	newGame(sqliteDatabase)
	newPlayer(sqliteDatabase)
	newTournament(sqliteDatabase)
	newStaff(sqliteDatabase)
	place(sqliteDatabase)

}
