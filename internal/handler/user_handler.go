package handler

import (
	"koda-b8-backend1/internal/domain"
	"koda-b8-backend1/internal/libs"
	"koda-b8-backend1/internal/service"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct { 
  service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler { 
  return &UserHandler{service: service}
}

// Create User godoc
// @Summary Register New User
// @Description  Register User
// @Accept json
// @Param request body domain.CreateUserRequest true "JSON"
// @Produce json
// @Success 201 {object} domain.User
// @Failure 400 {object} map[string]string
// @Security BearerAuth
// @Tags  users
// @Router /users [post]
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

// LoginAccount godoc
// @Summary Login User
// @Description  Login User
// @Accept json
// @Param request body domain.LoginRequest true "LoginJSON"
// @Produce json
// @Success 201 {object} domain.User
// @Failure 400 {object} map[string]string
// @Tags  auth
// @Router /auth/sign-in [post]
func (h *UserHandler) Login(c *gin.Context) {
	var req domain.LoginRequest
	// if err := c.ShouldBind(&req); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message": err.Error(),
	// 	})
	// 	return
	// }
	// 
	if err := c.ShouldBindJSON(&req); err != nil { 
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
	token, err := libs.GenerateToken(user.Id)
	c.JSON(http.StatusOK, token)
}

// CreateAccount godoc
// @Summary Register New User
// @Description  Register User
// @Accept json
// @Param request body domain.CreateUserRequest true "JSON"
// @Produce json
// @Success 201 {object} domain.User
// @Failure 400 {object} map[string]string
// @Tags  auth
// @Router /auth/sign-up [post]
func (h *UserHandler) Register(c *gin.Context) { 
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

// GetAllUsers godoc
// @Summary GetAllUsers
// @Description  Get User
// @Accept json
// @Param request true "LoginJSON"
// @Produce json
// @Success 201 {object} domain.User
// @Failure 400 {object} map[string]string
// @Security		BearerAuth
// @Tags  users
// @Router /users [get]
func (h *UserHandler) GetUsers(c *gin.Context) {
  page, err := strconv.ParseInt(c.Query("page"), 10, 64)
  limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
  search := c.QueryMap("search")
  users, err := h.service.GetUsers(c.Request.Context(), search, (page), limit)
  if err != nil { 
    c.JSON(http.StatusInternalServerError, gin.H{ 
      "message": "error",
    })
    return 
  }
  c.JSON(http.StatusOK, gin.H {
    "results": users,
  })
}

// FindById godoc
// @Summary Find By id user
// @Accept json
// @Produce json
// @Success 201 {object} domain.User
// @Failure 400 {object} map[string]string
// @Security		BearerAuth
// @Tags  users
// @Router /users/{id} [get]
func (h *UserHandler) GetUserByID(c *gin.Context) { 
  id, err := strconv.ParseInt(c.Param("id"), 10, 64) 
  if err != nil { 
    c.JSON(http.StatusBadRequest, gin.H{ 
      "message": "invalid id", 
    })
  }

  users, err := h.service.GetUserByID(id, c.Request.Context())
  if err != nil { 
    c.JSON(http.StatusNotFound, gin.H{ 
      "message": err.Error(),
    })
    return
  }
  c.JSON(http.StatusOK, users)
}

// DeleteUser godoc
// @Summary Delete User
// @Accept json
// @Param request body domain.PatchUserRequest true "PatchUser"
// @Produce json
// @Success 201 {object} domain.User
// @Failure 400 {object} map[string]string
// @Security		BearerAuth
// @Tags  users
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUsers(c *gin.Context) { 
  id, err := strconv.ParseInt(c.Param("id"), 10, 64) 
  if err != nil { 
    c.JSON(http.StatusBadRequest, gin.H{ 
      "message": "invalid id", 
    })
  }

  err = h.service.DeleteUser(id, c.Request.Context())
  if err != nil { 
    c.JSON(http.StatusNotFound, gin.H{ 
      "message": err.Error(),
    })
    return
  }
  c.JSON(http.StatusOK, gin.H{ 
    "message": "Berhasil",
  })
}

// EditUsers godoc
// @Summary Edit User
// @Accept json
// @Param request body domain.PatchUserRequest true "PatchUser"
// @Produce json
// @Success 201 {object} domain.User
// @Failure 400 {object} map[string]string
// @Security		BearerAuth
// @Tags  users
// @Router /users/{id} [patch]
func (h *UserHandler) PatchUser(c *gin.Context) { 
  id, err := strconv.ParseInt(c.Param("id"), 10, 64) 
  if err != nil { 
    c.JSON(http.StatusBadRequest, gin.H{ 
      "message": "invalid id", 
    })
  }

  var req domain.PatchUserRequest

  if err := c.ShouldBindJSON(&req); err != nil { 
    c.JSON(http.StatusBadRequest, gin.H{ 
      "message": err.Error(),
    })
    return
  }
  
  user, err := h.service.Patch(id, &req, c.Request.Context())
  if err != nil { 
    c.JSON(http.StatusInternalServerError, gin.H{ 
      "message": err.Error(), 
    })
    return
  }
  c.JSON(200, user)
}

// UploadPictureUser godoc
// @Summary Upload Picture User
// @Accept json
// @Param request body domain.PatchUserRequest true "PatchUser"
// @Produce json
// @Success 201 {object} domain.User
// @Failure 400 {object} map[string]string
// @Security		BearerAuth
// @Tags  users
// @Router /users/{id}/picture [patch]
func (h *UserHandler) UploadPictureProfile(c *gin.Context) {
  id, err := strconv.ParseInt(c.Param("id"), 10, 64) 
  file, err := c.FormFile("profile_picture")
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  dst := filepath.Join("uploads/", filepath.Base(file.Filename))
  c.SaveUploadedFile(file, dst)
  
  req := domain.UploadPicturesProfileRequest{
    Picture: dst,
  }
  
  if err := c.ShouldBind(&req); err != nil { 
    c.JSON(http.StatusBadRequest, gin.H{ 
      "message": err.Error(),
    })
    return
  }
 
  user, err := h.service.UploadPictureProfile(id, &req, c.Request.Context())
  if err != nil { 
    c.JSON(http.StatusInternalServerError, gin.H{ 
      "message": err.Error(), 
    })
    return
  }
  c.JSON(http.StatusOK, user)
}