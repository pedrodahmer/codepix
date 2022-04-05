package model

import (
	"time"
	uuid "github.com/satori/go.uuid"
	"github.com/asaskevich/govalidator"
)

type PixKeyRepositoryInterface interface {
	RegisterKey(pixKey *PixKey) (*PixKey, error)
	FindKeyByKind(key string, kind string) (*PixKey, error)
	AddBank(bank *Bank) error
	AddAccount(account *Account) error
	FindAccount(id string) (*Account, error)
}

type PixKey struct {
	Base	`valid:"required"`
	Kind			string		`json:"kind" valid:"notnull"`
	Key				string 		`json:"key" valid:"notnull"`
	AccountId	string		`json:"account_id" valid:"notnull"`
	Account		*Account	`valid:"-"`
	Status		string 		`json:"status valid:"notnull"`
}

func (pixKey *PixKey) isValid() error {
	_, err := govalidator.ValidateStruct(pixKey)

	if pixKey.Kind != "email" && pixKey.Kind != "cpf" {
		return errors.New("invalid key type")
	}

	if pixKey.Status != "active" && pixKey.Status != "inactive" {
		return errors.New("invalid key status")
	}

	if err != nil {
		return err
	}
	return nil
}

func NewPixKey(kind string, account *Account, key string) (*PixKey, error) {
	pixKey := PixKey {
		Kind:			kind,
		Account:	account,
		Status:		"active",
		Key: 			key,
	}

	pixKey.ID = uuid.NewV4().String()
	pixKey.CreatedAt = time.Now()

	err := pixKey.isValid()
	if err != nil {
		return nil, err
	}
	
	return &pixKey, nil
}
