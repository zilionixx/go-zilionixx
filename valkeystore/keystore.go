package valkeystore

import (
	"github.com/zilionixx/go-zilionixx/inter/validatorpk"
	"github.com/zilionixx/go-zilionixx/valkeystore/encryption"
)

type RawKeystoreI interface {
	Has(pubkey validatorpk.PubKey) bool
	Add(pubkey validatorpk.PubKey, key []byte, auth string) error
	Get(pubkey validatorpk.PubKey, auth string) (*encryption.PrivateKey, error)
}

type KeystoreI interface {
	RawKeystoreI
	Unlock(pubkey validatorpk.PubKey, auth string) error
	Unlocked(pubkey validatorpk.PubKey) bool
	GetUnlocked(pubkey validatorpk.PubKey) (*encryption.PrivateKey, error)
}
