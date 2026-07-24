package domain

import (
	"time"
)

type User struct { 
  Id int64 `json:"id"`
  Email string `json:"email"`
  FullName string `json:"fullname"`
  Picture *string `json:"picture"`
  Password string `json:"-"`
  CreatedAt time.Time `json:"created_at"`
}

type CreateUserRequest struct { 
  // Id int64 `json:"id"`
  FullName string `json:"fullname"`
  Email string `json:"email"`
  Password string `json:"password"`
}

type LoginRequest struct { 
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type UserFilter struct { 
  email string 
  name string
}

type PatchUserRequest struct { 
  Email *string `json:"email"`                                              
  FullName *string `json:"fullname"` 
}

type UploadPicturesProfileRequest struct {
  Picture string `json:"picture"`
}

type LoginResponse struct { 
  Id int64
  Email string 
  CreatedAt time.Time
}


