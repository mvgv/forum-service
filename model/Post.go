package model

import "time"

/*Post refere-se a abstração de um post vinculado a um topico*/
type Post struct {
	ID           uint64    `json:"id"`
	User         User      `json:"user_details"`
	CreationDate time.Time `json:"create_date"`
	LastEdition  time.Time `json:"last_edited"`
	Content      string    `json:"content"`
}
