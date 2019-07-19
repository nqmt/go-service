package id

import (
	"github.com/rs/xid"
)

type ID struct {
	freezeID string
}

var instance *ID

func New() *ID {
	if instance == nil {
		instance = &ID{}
	}

	return instance

	//return &ID{}
}

func (i ID) Gen() string {
	if i.freezeID != "" {
		return i.freezeID
	}

	return xid.New().String()
}


func (i *ID) Freeze(id string) {
	i.freezeID = id
}
