package bigquery

type tx struct {
	c *Conn
}

func newTx(c *Conn) (*tx, error) {
	return &tx{c: c}, nil
}

// Commit currently just  passes through
func (t *tx) Commit() (err error) {
	return
}

// Rollback currently just  passes through
func (t *tx) Rollback() (err error) {
	return
}
