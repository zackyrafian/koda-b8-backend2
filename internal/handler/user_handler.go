package handler

import (
	"koda-b8-backend1/internal/domain"
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
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetUsers(c *gin.Context) { 
  users, err := h.service.GetUsers(c.Request.Context())
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