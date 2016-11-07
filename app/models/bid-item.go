package models

import (
    "github.com/revel/revel"
    
)

type BidItem struct {
    Id              int64   `db:"id" json:"id,omitempty"`
    Name            string  `db:"name" json:"name"`
    Entry        string  `db:"entry" json:"entry"`
}


func (b *BidItem) Validate(v *revel.Validation) {

    v.Check(b.Name,
        revel.ValidRequired())

    v.Check(b.Entry,
        revel.ValidRequired())
}