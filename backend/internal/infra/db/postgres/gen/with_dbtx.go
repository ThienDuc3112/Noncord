package gen

func (q *Queries) WithDbtx(tx DBTX) *Queries {
	return &Queries{
		db: tx,
	}
}
