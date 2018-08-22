package db


// Interface that needs to be fulfilled by any database handler
type DatabaseHandler interface {
	AddEvent(Event) ([]byte, error)
	FindEvent([]byte) (Event, error)
	FindEventByName(string) (Event, error)
	FindAllAvailableEvents() ([]Event, error)
	DeleteEvent(Event) error
}
