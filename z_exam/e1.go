package zexam

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"uniqueIndex"`
	Role      string `gorm:"default:'user'"` // "admin" 或 "user"
	DeletedAt gorm.DeletedAt
}

func GetCurrentUserID(c *gin.Context) string {
	return c.GetString("currentUserID") // 从token解析
}

// DELETE /api/v1/users/:id
func DeleteUser(c *gin.Context) {
	userID := c.Param("id")
	currentUserID := GetCurrentUserID(c)

	var currentUser User
	db.Where("id = ?", currentUserID).First(&currentUser)
	if currentUser.Role != "admin" {
		c.JSON(403, gin.H{"error": "permission denied"})
		return
	}

	db.Delete(&User{}, userID)
	c.JSON(200, gin.H{"message": "deleted"})
}

// DELETE /api/v1/users/batch
func BatchDeleteUsers(c *gin.Context) {
	var req struct {
		UserIDs []string `json:"user_ids"`
	}
	c.BindJSON(&req)

	for _, id := range req.UserIDs {
		DeleteUser(c)
	}
}
