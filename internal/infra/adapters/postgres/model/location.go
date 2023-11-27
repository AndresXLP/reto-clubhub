package model

type Countries struct {
	ID   int64
	Name string
}

type Cities struct {
	ID        int64
	Name      string
	CountryID int64
}

type Addresses struct {
	ID      int64
	Address string
	ZipCode string
	CityID  int64
}

type Locations struct {
	Countries *Countries
	Cities    *Cities
	Addresses *Addresses
}
