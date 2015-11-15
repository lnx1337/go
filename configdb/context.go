package configdb

import (
	"upper.io/db"
)

type ctx struct {
	col db.Collection
}

func Context(col db.Collection) *ctx {
	return &ctx{col}
}

func (c *ctx) Collection() db.Collection {
	return c.col
}