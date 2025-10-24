package entities

import (
	"crypto/rand"
	"log"
	"time"

	"github.com/google/uuid"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func MustRandomString(n int) string {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		// You can choose to panic(err) or log.Fatal(err)
		log.Panicf("token: could not generate random bytes: %v", err)
	}
	for i := range b {
		b[i] = letters[int(b[i])%len(letters)]
	}
	return string(b)
}

func RandomToken() string {
	return MustRandomString(32)
}

type Session struct {
	Id            uuid.UUID
	RotationCount int32
	CreatedAt     time.Time
	UpdatedAt     time.Time
	ExpiresAt     time.Time
	UserId        UserId
	UserAgent     string
	Token         string
}

func (s *Session) Validate() error {
	if s.ExpiresAt.Compare(time.Now()) <= 0 {
		return NewError(ErrCodeValidationError, "expired token", nil)
	}
	if len(s.Token) != 32 {
		return NewError(ErrCodeValidationError, "invalid token form", nil)
	}
	if s.RotationCount <= 0 {
		return NewError(ErrCodeValidationError, "invalid rotation count", nil)
	}

	return nil
}

func NewSession(uid UserId, expiresAt time.Time, userAgent string) *Session {
	return &Session{
		Id:            uuid.New(),
		RotationCount: 1,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		ExpiresAt:     expiresAt,
		UserId:        uid,
		UserAgent:     userAgent,
		Token:         RandomToken(),
	}
}
