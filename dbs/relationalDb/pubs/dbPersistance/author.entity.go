package dbPersistance

type Author struct {
	Id          int64
	AuthorCode  string
	LastName    string
	FirstName   string
	PhoneNumber string
	Address     string
	City        string
	State       string
	ZipCode     string
	HasContract bool
}
