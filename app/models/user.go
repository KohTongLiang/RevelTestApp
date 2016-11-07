package models

import (
    "github.com/revel/revel"
    
)

type User struct {
    Id              int64   `db:"id" json:"id,omitempty"`
    Username            string  `db:"username" json:"username"`
    Password        []byte  `db:"password" json:"password"`
}


func (b *User) Validate(v *revel.Validation) {

    v.Check(b.Username,
        revel.ValidRequired())

    v.Check(b.Password,
        revel.ValidRequired())
}