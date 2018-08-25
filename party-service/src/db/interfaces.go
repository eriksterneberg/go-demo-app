package db

type DatabaseHandler interface {
	AddParty(Party) ([]byte, error)
	FindParty([]byte) (Party, error)
	FindPartyByName(string) (Party, error)
	FindAllAvailablePartys() ([]Party, error)
	DeleteParty(Party) error
}
