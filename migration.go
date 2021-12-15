package main

import (
	"database/sql"
)

func coachMigration(database string, coachId int) Coach {
	db, errDB := sql.Open("sqlite", database)
	CheckErr(errDB)

	rows := db.QueryRow(`SELECT * FROM coach WHERE idCoach= ?`, coachId)
	var idCoach, idGame, age, wage int
	var lastname, firstname, gender, licenseDate string
	var coach Coach
	errRows := rows.Scan(&idCoach, lastname, firstname, gender, age, wage, licenseDate)
	CheckErr(errRows)
	coach = Coach{IdCoach: idCoach, IdGame: idGame, Lastname: lastname, Firstname: firstname, Gender: gender, Age: age, Wage: wage, LicenseDate: licenseDate}
	return coach
}

func newCoach(database string)  {
	db, err := sql.Open("sqlite3", "dest.sqlite")
	CheckErr(err)
	coachMigration(database, 2)
	// insert
	stmt, err1 := db.Prepare("INSERT INTO main.coach(idCoach, idGame, licenseDate, idEmployeeData) values(?,?,?,?)")
	CheckErr(err1)
	_, err2 := stmt.Exec()
	CheckErr(err2)
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
newCoach("old-database.sqlite")
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
