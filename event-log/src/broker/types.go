// Events and other types used by the broker

package broker

type PartyEvent struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (p PartyEvent) TopicName() string {
	return "event.party"
}

type Message struct {
	Key    []byte
	Value  []byte
	Offset int64 // Offset where the data is stored
}
