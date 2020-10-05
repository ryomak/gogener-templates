package repository

import (
	"fmt"

	"github.com/volatiletech/sqlboiler/queries/qm"
)

func InnerJoin(joinTable, baseTable, joinTableColumn, baseTableColumn string) qm.QueryMod {
	return qm.InnerJoin(fmt.Sprintf("%s ON %s.%s = %s.%s",
		joinTable,
		joinTable,
		joinTableColumn,
		baseTable,
		baseTableColumn,
	))
}

func LeftOuterJoin(joinTable, baseTable, joinTableColumn, baseTableColumn string) qm.QueryMod {
	return qm.LeftOuterJoin(fmt.Sprintf("%s ON %s.%s = %s.%s",
		joinTable,
		joinTable,
		joinTableColumn,
		baseTable,
		baseTableColumn,
	))
}
