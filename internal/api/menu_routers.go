package api

import (
	"Order5003/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterMenuRoutes(r *gin.Engine, h *handlers.MenuHandler) {
    r.GET("/api/menu", func(c *gin.Context) { h.GetAllMenuItems(c) })
    r.Any("/api/menu-item", func(c *gin.Context) {
        switch c.Request.Method {
        case "GET":
            h.GetMenuItemByID(c)
        case "POST":
            h.CreateMenuItem(c)
        case "PUT":
            h.UpdateMenuItem(c)
        case "DELETE":
            h.DeleteMenuItem(c)
        default:
            c.Status(405)
        }
    })
}
