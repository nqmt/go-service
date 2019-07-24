package soap

type Soap struct {
	PersonID  string `db:"PERSONID"`
	LastName  string `db:"LASTNAME"`
	FirstName string `db:"FIRSTNAME"`
	Address   string `db:"ADDRESS"`
	City      string `db:"CITY"`
}
