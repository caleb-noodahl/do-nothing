package models

type Step struct {
	Sequence int      `json:"seq"`
	Name     string   `json:"name"`
	Text     string   `json:"text"`
	CMDs     []string `json:"ex"`
	Err      error
}
