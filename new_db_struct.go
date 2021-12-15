package main

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
