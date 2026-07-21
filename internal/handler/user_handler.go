package handler

import (
	"koda-b8-backend1/internal/domain"
	"koda-b8-backend1/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct { 
  service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler { 
  return &UserHandler{service: service}
}

func (h *UserHandler) Create(c *gin.Context) { 
  var form domain.CreateUserRequest
  if err := c.ShouldBindJSON(&form); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
 
  user, err := h.service.Create(&form, c.Request.Context())
  if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{
          "message": err.Error(),
      })
      return
  }
  c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) Login(c *gin.Context) {
	var req domain.LoginRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	user, err := h.service.Login(&req, c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

// func (h *UserHandler) GetUsers(c *gin.Context) { 
//   users, err := h.service.GetUsers()
//   if err != nil { 
//     c.JSON(http.StatusInternalServerError, gin.H{ 
//       "message": "error",
//     })
//     return 
//   }
//   c.JSON(http.StatusOK, users)
// }