package main

import "database/sql"

/*func coachMigration(db *sql.DB) Coach {

	var idCoach, idGame, age, wage int
	var lastname, firstname, gender, licenseDate string


	err := rows.Scan(&idCoach, &idGame, &lastname, &firstname, &gender, &age, &wage, &licenseDate)
	CheckErr(err)

	return Coach{IdCoach: idCoach, IdGame: idGame, Lastname: lastname, Firstname: firstname, Gender: gender, LicenseDate: licenseDate}
}*/

func gamesMigration(db *sql.DB) Games {
	rows, _ := db.Query(`SELECT idGame, Name FROM games`)

	var idGame int
	var name string

	err := rows.Scan(&idGame, &name)
	CheckErr(err)


	return Games{IdGame: idGame, Name: name}
}

func playerMigration(db *sql.DB) Player {
	rows := db.QueryRow(`SELECT idPlayer, idGame, lastname, firstname, gender, age, wage, ranking FROM player`)

	var idPlayer, idGame, age, wage, ranking int
	var lastname, firstname, gender string

	err := rows.Scan(&idPlayer, &idGame, &lastname, &firstname, &gender, &age, &wage, &ranking)
	CheckErr(err)

	return Player{IdPlayer: idPlayer, IdGame: idGame, Lastname: lastname, Firstname: firstname, Gender: gender, Wage: wage, Ranking: ranking}
}

/*func staffMigration(db *sql.DB) Staff {

	var idStaff, age, wage int
	var lastname, firstname, gender string

	err := rows.Scan(&idStaff, &lastname, &firstname, &gender, &age, &wage)
	CheckErr(err)

	return Staff{IdStaff: idStaff, Lastname: lastname, Firstname: firstname, Gender: gender, Age: age, Wage: wage}
}
*/
func tournamentMigration(db *sql.DB) Tournament {
	rows := db.QueryRow(`SELECT idTournament, idGame, date, duration, placeName, address, city FROM tournament`)

	var idTournament, idGame, duration int
	var date, placeName, address, city string

	err := rows.Scan(&idTournament, &idGame, &date, &duration, &placeName, &address, &city)
	CheckErr(err)

	return Tournament{IdTournament: idTournament, IdGame: idGame, Date: date, Duration: duration, PlaceName: placeName, Address: address, City: city}
}
