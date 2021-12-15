package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func newCoach(db *sql.DB) {
	database, err := sql.Open("sqlite3", "./dest.sqlite")
	CheckErr(err)

	var coach = coachMigration(db)
	stmt, err1 := database.Prepare(`INSERT INTO main.coach(idCoach, idGame, licenseDate, idEmployeeData) values(?,?,?,?)`)
	CheckErr(err1)
	_, err2 := stmt.Exec(coach.IdCoach, coach.IdGame, coach.LicenseDate, 3)
	CheckErr(err2)
	defer database.Close()
}

func newGame(db *sql.DB) {
	database, err := sql.Open("sqlite3", "./dest.sqlite")
	CheckErr(err)

	var games = gamesMigration(db)
	stmt, err1 := database.Prepare(`INSERT INTO game(idGame, name) values(?,?)`)
	CheckErr(err1)
	_, err2 := stmt.Exec(games.IdGame, games.Name)
	CheckErr(err2)
	defer database.Close()
}

func newTournament(db *sql.DB) {
	database, err := sql.Open("sqlite3", "./dest.sqlite")
	CheckErr(err)

	var tournament = tournamentMigration(db)
	stmt, err1 := database.Prepare(`INSERT INTO main.tournament(idTournament, idPlace, idGame, date, duration) values(?,?,?,?,?)`)
	CheckErr(err1)
	_, err2 := stmt.Exec(tournament.IdTournament, 0, tournament.IdGame, tournament.Date, tournament.Duration)
	CheckErr(err2)
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

func newPlayer(db *sql.DB) {
	database, err := sql.Open("sqlite3", "./dest.sqlite")
	CheckErr(err)

	var player = playerMigration(db)
	stmt, err1 := database.Prepare(`INSERT INTO main.player(idPlayer, idGame, ranking, idEmployeeData) values(?,?,?,?)`)
	CheckErr(err1)
	_, err2 := stmt.Exec(player.IdPlayer, player.IdGame, player.Ranking, 1234)
	CheckErr(err2)
	defer database.Close()

}

func newStaff(db *sql.DB) {
	database, err := sql.Open("sqlite3", "./dest.sqlite")
	CheckErr(err)

	var staff = staffMigration(db)
	stmt, err1 := database.Prepare(`INSERT INTO main.staff(idStaff, idEmployeeData) values(?,?)`)
	CheckErr(err1)
	_, err2 := stmt.Exec(staff.IdStaff, 1234)
	CheckErr(err2)
	defer database.Close()
}

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

	newCoach(sqliteDatabase)
	//newEmployeeData(sqliteDatabase)
	newGame(sqliteDatabase)
	//newPlace(sqliteDatabase)
	newPlayer(sqliteDatabase)
	newStaff(sqliteDatabase)
	newTournament(sqliteDatabase)

}

/*func createTable(db *sql.DB) {
	createStudentTableSQL := `CREATE TABLE student (
		"idStudent" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"code" TEXT,
		"name" TEXT,
		"program" TEXT
	  );` // SQL Statement for Create Table

	log.Println("Create student table...")
	statement, err := db.Prepare(createStudentTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("student table created")
}

// We are passing db reference connection from main to our method with other parameters
func insertStudent(db *sql.DB, code string, name string, program string) {

}

func displayStudents(db *sql.DB) {
	row, err := db.Query("SELECT * FROM student ORDER BY name")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var code string
		var name string
		var program string
		errScan := row.Scan(&id, &code, &name, &program)
		CheckErr(errScan)
		log.Println("Student: ", code, " ", name, " ", program)
	}
}*/
