package types

import (
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

type UserRole string
type UserInProjectRole string
type UserLoginProvider string

var (
	RoleUser    UserRole = "user"
	RoleAdmin   UserRole = "admin"
	RoleCompany UserRole = "company"
)

type User struct {
	ID               uuid.UUID        `db:"id" json:"id"`
	CompanyID        pgtype.UUID      `db:"company_id" json:"company_id"`
	FirstName        pgtype.Text      `db:"first_name" json:"first_name"`
	LastName         pgtype.Text      `db:"last_name" json:"last_name"`
	Email            pgtype.Text      `db:"email" json:"email"`
	Password         pgtype.Text      `db:"password" json:"password"`
	IsPremium        pgtype.Bool      `db:"is_premium" json:"is_premium"`
	IsActive         pgtype.Bool      `db:"is_active" json:"is_active"`
	PremiumStartDate pgtype.Date      `db:"premium_start_date" json:"premium_start_date"`
	PremiumEndDate   pgtype.Date      `db:"premium_end_date" json:"premium_end_date"`
	LoginProvider    pgtype.Text      `db:"login_provider" json:"login_provider"`
	OtpSecret        pgtype.Text      `db:"otp_secret" json:"otp_secret"`
	CreatedAt        pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt        pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	Role             pgtype.Text      `db:"role" json:"role"`
}

type Company struct {
	ID   uuid.UUID   `db:"id" json:"id"`
	Name pgtype.Text `db:"name" json:"name"`
}

type Project struct {
	ID          uuid.UUID        `db:"id" json:"id"`
	CompanyID   pgtype.UUID      `db:"company_id" json:"company_id"`
	Name        pgtype.Text      `db:"name" json:"name"`
	Description pgtype.Text      `db:"description" json:"description"`
	CreatedAt   pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt   pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type ProjectMember struct {
	ID        uuid.UUID        `db:"id" json:"id"`
	ProjectID pgtype.UUID      `db:"project_id" json:"project_id"`
	UserID    pgtype.UUID      `db:"user_id" json:"user_id"`
	Role      pgtype.Text      `db:"role" json:"role"`
	CreatedAt pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}
