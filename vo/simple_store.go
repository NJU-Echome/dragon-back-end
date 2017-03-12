package vo

type Simple_store struct {
	Id        uint32  `json:"id"`
	Name      string  `json:"name"`
	Address   string  `json:"address"`
	Longitude float32 `json:"longitude"`
	Latitude  float32 `json:"latitude"`
	Tags      []Tag   `json:"tags"`
}
