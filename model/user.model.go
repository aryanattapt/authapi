package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserViewPayload struct {
	UserID    primitive.ObjectID `bson:"_id,omitempty" json:"user_id,omitempty"`
	Username  string             `bson:"username" json:"username" validate:"required,alphanum,min=3,max=50"`
	Fullname  string             `bson:"fullname" json:"fullname" validate:"required,min=3,max=50"`
	Email     string             `bson:"email" json:"email" validate:"required,email"`
	CreatedAt primitive.DateTime `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	IsPremium bool               `bson:"is_premium,omitempty" json:"is_premium,omitempty"`
	Verified  bool               `bson:"verified,omitempty" json:"verified,omitempty"`
	IsActive  bool               `bson:"isactive,omitempty" json:"isactive,omitempty"`
}

type UserInsertPayload struct {
	Password string `bson:"password,omitempty" json:"password,omitempty" validate:"required,alphanum,min=3,max=20"`
	UserViewPayload
}

type UserUpdatePayload struct {
	ID string `bson:"_id,omitempty" json:"_id,omitempty" validate:"required"`
	UserViewPayload
}
