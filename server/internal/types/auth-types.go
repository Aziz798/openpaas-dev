package types

import (
	"github.com/jackc/pgtype"
)

type UserSignupWithEmail struct {
	FirstName       pgtype.Text `db:"first_name" json:"first_name" validate:"required"`
	LastName        pgtype.Text `db:"last_name" json:"last_name" validate:"required"`
	Email           pgtype.Text `db:"email" json:"email" validate:"required,email"`
	Password        pgtype.Text `db:"password" json:"password" validate:"required,min=8"`
	ConfirmPassword pgtype.Text `json:"confirm_password" validate:"required,min=8,eqfield=Password"`
}
