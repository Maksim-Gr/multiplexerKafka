package models

// Notification describes the incoming message from main topic
type Notification struct {
	Id    string `json:"id"`
	Open  bool   `json:"open"`
	BetId string `json:"betId"`
}
