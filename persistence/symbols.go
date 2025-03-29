package persistence

import (
	"encoding/json"
	"log"
	"strconv"

	"ctraderapi/messages/github.com/Carlosokumu/messages"

	"gorm.io/gorm"

	badger "github.com/dgraph-io/badger/v3"
)

func CreateBadgerConnection() (*badger.DB, error) {
	db, err := badger.Open(badger.DefaultOptions("/tmp/badger"))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InsertSymbolData(Symbol messages.ProtoOASymbol) error {
	db, err := badger.Open(badger.DefaultOptions("/swingwizards/symbols"))
	if err != nil {
		log.Fatal(err)
	}
	err = db.Update(func(txn *badger.Txn) error {
		json, _ := json.Marshal(&Symbol)
		err := txn.Set([]byte(strconv.Itoa(int(*Symbol.SymbolId))), []byte(json))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	defer db.Close()
	return nil
}

func InsertLightSymbol(symbolId int64, protoLightSymbol messages.ProtoOALightSymbol) error {
	db, err := badger.Open(badger.DefaultOptions("/swingwizards/lightsymbols"))
	if err != nil {
		log.Fatal(err)
	}
	err = db.Update(func(txn *badger.Txn) error {
		json, _ := json.Marshal(&protoLightSymbol)
		err := txn.Set([]byte(strconv.Itoa(int(symbolId))), []byte(json))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	defer db.Close()
	return nil
}

func ReadSymbolData(SymbolId int64) (*messages.ProtoOASymbol, error) {
	var symbolEntity messages.ProtoOASymbol
	db, err := badger.Open(badger.DefaultOptions("/swingwizards/symbols"))
	if err != nil {
		return nil, err
	}
	err = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(strconv.Itoa(int(SymbolId))))
		if err != nil {
			return err
		}
		item.Value(func(val []byte) error {

			err = json.Unmarshal(val, &symbolEntity)
			if err != nil {
				return err
			}
			return nil
		})

		return nil
	})
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return &symbolEntity, nil
}

func ReadLightSymbolData(SymbolId int64) (*messages.ProtoOALightSymbol, error) {
	var symbolEntity messages.ProtoOALightSymbol
	db, err := badger.Open(badger.DefaultOptions("/swingwizards/lightsymbols"))
	if err != nil {
		return nil, err
	}
	err = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(strconv.Itoa(int(SymbolId))))
		if err != nil {
			return err
		}
		item.Value(func(val []byte) error {

			err = json.Unmarshal(val, &symbolEntity)
			if err != nil {
				return err
			}
			return nil
		})

		return nil
	})
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return &symbolEntity, nil
}

type SymbolData struct {
	Name string `json:"name"`
}

type SymbolModel struct {
	gorm.Model
	ConversionSymbols []SymbolModel `gorm:"foreignkey:ParentID"`
	ID                int64
	ParentID          int64
}
