package service

import (
	"fmt"
	"log/slog"

	"github.com/daanjo3/anweb-app/api/internal/anwb"
	"github.com/daanjo3/anweb-app/api/internal/db"
)

// TODO add pagination
// TODO add filtering options
func ListDocuments() ([]anwb.Document, error) {
	return db.GetDocuments(db.DOC_FORMAT_INDEX)
}

func GetDocument(id string) (anwb.Document, error) {
	isLatest := id == "latest"
	if isLatest {
		return db.GetDocumentLatest()
	}
	return db.GetDocumentById(id)
}

func AddDocument(checkExistence bool) (anwb.Document, error) {
	data, err := anwb.Get()
	if err != nil {
		return anwb.Document{}, err
	}
	if checkExistence {
		exists, err := db.ExistsDocument(&data)
		if err != nil {
			return anwb.Document{}, err
		}
		if exists {
			slog.Info("Did not update document as it already existed")
			return data, nil
		}
	}
	id, err := db.InsertDocument(data)
	if err != nil {
		return anwb.Document{}, err
	}
	slog.Info(fmt.Sprintf("Inserted new ANWB document with id %v", id))
	return data, nil
}
