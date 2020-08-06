package user

type User struct {
	ID int `json:"CharacterID"`
	Name string `json:"CharacterName"`
	OwnerHash string `json:"CharacterOwnerHash"`
}