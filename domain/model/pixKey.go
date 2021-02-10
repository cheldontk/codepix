package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

const (
	PixKeyActive   string = "active"
	PixKeyInactive string = "inactive"
	PixKeyPending  string = "pending"
)

type PixKeyRepositoryInterface interface {
	RegisterKey(pixkey *PixKey) (*PixKey, error)
	FindKeyByKind(key string, kind string) (*PixKey, error)
	AddBank(bank *Bank) error
	AddAccount(account *Account) error
	FindAccount(id string) (*Account, error)
}

type PixKey struct {
	Base      `valid:"required"`
	Kind      string   `json:"kind" valid:"notnull"`
	Key       string   `json:"key" valid:"notnull"`
	AccountID string   `valid:"-"`
	Account   *Account `valid:"-"`
	Status    string   `json:"status" valid:"notnull"`
}

func (p *PixKey) isValid() error {
	_, err := govalidator.ValidateStruct(p)

	if p.Kind != "email" && p.Kind != "cpf" {
		return errors.New("invalid type of key")
	}

	if p.Status != "active" && p.Status != "inactive" {
		return errors.New("invalid status")
	}

	if err != nil {
		return err
	}

	return nil
}

func NewPixKey(account *Account, kind string, key string) (*PixKey, error) {
	pixKey := PixKey{
		Kind:      kind,
		Key:       key,
		Account:   account,
		AccountID: account.ID,
		Status:    PixKeyActive,
	}

	pixKey.ID = uuid.NewV4().String()
	pixKey.CreatedAt = time.Now()

	err := pixKey.isValid()
	if err != nil {
		return nil, err
	}

	return &pixKey, nil
}
