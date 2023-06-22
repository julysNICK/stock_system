package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("Invalid token")
	ErrExpiredToken = errors.New("Expired token")
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IdStore   int64     `json:"id_store"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(idStore int64, username string, duration time.Duration) (*Payload, error) {

	tokenId, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenId,
		Username:  username,
		IdStore:   idStore,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}

func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
