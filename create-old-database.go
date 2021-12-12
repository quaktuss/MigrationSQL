package main

import (
	"github.com/bxcodec/faker"
	"math/rand"
)

type Tournament struct {
	IdTournament int
	IdGame       int
	Date         string
	Duration     int
	PlaceName    string
	Address      string
	City         string
}
type Games struct {
	IdGame int
	Name   string
}
type Staff struct {
	IdStaff   int
	Lastname  string
	Firstname string
	Gender    string
	Age       int
	Wage      int
}

type Player struct {
	IdPlayer  int
	IdGame    int
	Lastname  string
	Firstname string
	Gender    string
	Age       int
	Wage      int
	Ranking   int
}
type Coach struct {
	IdCoach     int
	IdGame      int
	Lastname    string
	Firstname   string
	Gender      string
	Age         int
	Wage        int
	LicenseDate string
}

func getFakeTournament(games Games) Tournament {
	return Tournament{
		IdTournament: games.IdGame,
		IdGame:       rand.Intn(9),
		Date:         faker.Date(),
		Duration: rand.Intn(100),
		PlaceName: faker.Word(),
		City: faker.FirstNameFemale(),
	}
}

func getFakeGames() Games {
	return Games{
		IdGame: rand.Int(),
		Name: faker.Word(),
	}
}

func getFakeStaff() Staff {
	return Staff{
		IdStaff: rand.Intn(9),
		Lastname: faker.LastName(),
		Firstname: faker.FirstName(),
		Gender: faker.Gender(),
		Age: rand.Intn(50),
		Wage: rand.Intn(9999),
	}
}

func getFakePlayer(games Games) Player {
	return Player{
		IdPlayer: rand.Intn(9),
		IdGame: games.IdGame,
		Lastname: faker.LastName(),
		Firstname: faker.FirstName(),
		Gender: faker.Gender(),
		Age: rand.Intn(50),
		Wage: rand.Intn(9999),
	}
}

func getFakeCoach(games Games) Coach {
	return Coach{
		IdCoach: rand.Intn(9),
		IdGame: games.IdGame,
		Lastname: faker.LastName(),
		Firstname: faker.FirstName(),
		Gender: faker.Gender(),
		Age: rand.Intn(50),
		Wage: rand.Intn(9999),
		LicenseDate: faker.Word(),
	}
}

/*func main() {
	db, err := sql.Open("sqlite3", "old-database.sqlite")

	getFakeGames()
	getFakeTournament(Games{})
	getFakeStaff()
	getFakePlayer(Games{})
	getFakeCoach(Games{})


}*/