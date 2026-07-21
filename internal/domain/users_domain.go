package domain

import (
	"time"
)

type User struct { 
  Id int64
  Email string 
  Password string 
  CreatedAt time.Time
}

type CreateUserRequest struct { 
  Id int64 `json:"id"`
  Email string `json:"email"`
  Password string `json:"password"`
}

type LoginRequest struct { 
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type LoginResponse struct { 
  Id int64
  Email string 
  CreatedAt time.Time
}

// type Repository interface { 
//   Create(ctx context.Context ,req *CreateUserRequest) (*User, error) 
//   FindByEmail(ctx context.Context, req *LoginRequest) (*User, error)
// }