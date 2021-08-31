package valkeystore

import (
	"github.com/ethereum/go-ethereum/accounts/keystore"

	"github.com/zilionixx/go-zilionixx/valkeystore/encryption"
)

func NewDefaultFileRawKeystore(dir string) *FileKeystore {
	enc := encryption.New(keystore.StandardScryptN, keystore.StandardScryptP)
	return NewFileKeystore(dir, enc)
}

func NewDefaultMemKeystore() *SyncedKeystore {
	return NewSyncedKeystore(NewCachedKeystore(NewMemKeystore()))
}

func NewDefaultFileKeystore(dir string) *SyncedKeystore {
	return NewSyncedKeystore(NewCachedKeystore(NewDefaultFileRawKeystore(dir)))
}
