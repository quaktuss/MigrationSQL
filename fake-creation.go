package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func fakeTableGames() {
	db, err := sql.Open("sqlite3", "old-database.sqlite")
	CheckErr(err)

	// insert
	stmt, err1 := db.Prepare("INSERT INTO games(IDGAME, NAME) VALUES (?,?)")
	CheckErr(err1)
	_, err2 := stmt.Exec(getFakeGames().IdGame, getFakeGames().Name)
	CheckErr(err2)
}
func fakeTableCoach()  {
	db, err := sql.Open("sqlite3", "old-database.sqlite")
	CheckErr(err)

	// insert
	stmt, err1 := db.Prepare("INSERT INTO main.coach(idCoach, idGame, lastname, firstname, gender, age, wage, licenseDate) values(?,?,?,?,?,?,?,?)")
	CheckErr(err1)
	_, err2 := stmt.Exec(getFakeCoach().IdCoach, getFakeGames().IdGame, getFakeCoach().Lastname, getFakeCoach().Firstname, getFakeCoach().Gender, getFakeCoach().Age, getFakeCoach().Wage, getFakeCoach().LicenseDate)
	CheckErr(err2)
}
func fakeTablePlayer()  {
	db, err := sql.Open("sqlite3", "old-database.sqlite")
	CheckErr(err)

	// insert
	stmt, err1 := db.Prepare("INSERT INTO main.player(idPlayer, idGame, lastname, firstname, gender, age, wage, ranking) values(?,?,?,?,?,?,?,?)")
	CheckErr(err1)
	_, err2 := stmt.Exec(getFakePlayer().IdPlayer, getFakeGames().IdGame, getFakePlayer().Lastname, getFakePlayer().Firstname, getFakePlayer().Gender, getFakePlayer().Age, getFakePlayer().Wage, getFakePlayer().Ranking)
	CheckErr(err2)
}

func fakeTableStaff()  {
	db, err := sql.Open("sqlite3", "old-database.sqlite")
	CheckErr(err)

	// insert
	stmt, err1 := db.Prepare("INSERT INTO main.staff(idStaff, lastname, firstname, gender, age, wage) values(?,?,?,?,?,?)")
	CheckErr(err1)
	_, err2 := stmt.Exec(getFakeStaff().IdStaff, getFakeGames().IdGame, getFakeStaff().Lastname, getFakeStaff().Firstname, getFakeStaff().Gender, getFakeStaff().Age, getFakeStaff().Wage)
	CheckErr(err2)
}
func fakeTableTournament() {
	db, err := sql.Open("sqlite3", "old-database.sqlite")
	CheckErr(err)

	// insert
	stmt, err1 := db.Prepare("INSERT INTO main.tournament(idTournament, idGame, date, duration, placeName, address, city) values(?,?,?,?,?,?,?)")
	CheckErr(err1)
	_, err2 := stmt.Exec(getFakeTournament().IdTournament, getFakeGames().IdGame, getFakeTournament().Date, getFakeTournament().Duration, getFakeTournament().PlaceName, getFakeTournament().Address, getFakeTournament().City)
	CheckErr(err2)
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fakeTableCoach()
	fakeTableTournament()
	fakeTablePlayer()
	//fakeTableStaff()
	fakeTableGames()
}
