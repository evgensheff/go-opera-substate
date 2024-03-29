package substate

import (
	"github.com/syndtr/goleveldb/leveldb/opt"

	"github.com/Fantom-foundation/Substate/db"
	"github.com/Fantom-foundation/Substate/substate"
)

var (
	staticSubstateDB db.SubstateDB
	RecordReplay     bool = false
)

func NewSubstateDB(path string) error {
	var err error
	staticSubstateDB, err = db.NewSubstateDB(path, &opt.Options{ReadOnly: false}, nil, nil)
	return err
}

func CloseSubstateDB() error {
	return staticSubstateDB.Close()
}

func PutSubstate(ss *substate.Substate) error {
	return staticSubstateDB.PutSubstate(ss)
}
