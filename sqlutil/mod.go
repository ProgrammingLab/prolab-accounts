package sqlutil

import (
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type QueryMods []qm.QueryMod

func (m QueryMods) Apply(q *queries.Query) {
	qm.Apply(q, m...)
}
