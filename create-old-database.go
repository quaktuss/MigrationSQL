package main

import (
	"github.com/bxcodec/faker"
	"math/rand"
	"time"
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

func getFakeTournament() Tournament {
	rand.Seed(time.Now().UnixNano())
	return Tournament{
		IdTournament: rand.Intn(90-01) + 1,
		IdGame:       Games{}.IdGame,
		Date:         faker.Date(),
		Duration:     rand.Intn(240-35) + 1,
		PlaceName:    faker.Word(),
		City:         faker.FirstNameFemale(),
	}
}

func getFakeGames() Games {
	rand.Seed(time.Now().UnixNano())
	return Games{
		IdGame: rand.Intn(90-01) + 1,
		Name:   faker.Word(),
	}
}

func getFakeStaff() Staff {
	rand.Seed(time.Now().UnixNano())
	return Staff{
		IdStaff:   rand.Intn(90-01) + 1,
		Lastname:  faker.LastName(),
		Firstname: faker.FirstName(),
		Gender:    faker.Gender(),
		Age:       rand.Intn(25-17) + 1,
		Wage:      rand.Intn(2500-1500) + 1,
	}
}

func getFakePlayer() Player {
	rand.Seed(time.Now().UnixNano())
	return Player{
		IdPlayer:  rand.Intn(90-01) + 1,
		IdGame:    Games{}.IdGame,
		Lastname:  faker.LastName(),
		Firstname: faker.FirstName(),
		Gender:    faker.Gender(),
		Age:       rand.Intn(25-17) + 1,
		Wage:      rand.Intn(2500-1500) + 1,
	}
}

func getFakeCoach() Coach {
	rand.Seed(time.Now().UnixNano())
	return Coach{
		IdCoach:     rand.Intn(99-01) + 1,
		IdGame:      Games{}.IdGame,
		Lastname:    faker.LastName(),
		Firstname:   faker.FirstName(),
		Gender:      faker.Gender(),
		Age:         rand.Intn(25-17) + 1,
		Wage:        rand.Intn(2500-1500) + 1,
		LicenseDate: faker.Word(),
	}
}
