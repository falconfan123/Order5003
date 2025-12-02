package api

import (
    "Order5003/internal/handlers"
    "github.com/gin-gonic/gin"
)

func RegisterMenuRoutes(r *gin.Engine, h *handlers.MenuHandler) {
    r.GET("/api/menu", func(c *gin.Context) { h.GetAllMenuItems(c.Writer, c.Request) })
    r.Any("/api/menu-item", func(c *gin.Context) {
        switch c.Request.Method {
        case "GET":
            h.GetMenuItemByID(c.Writer, c.Request)
        case "POST":
            h.CreateMenuItem(c.Writer, c.Request)
        case "PUT":
            h.UpdateMenuItem(c.Writer, c.Request)
        case "DELETE":
            h.DeleteMenuItem(c.Writer, c.Request)
        default:
            c.Status(405)
        }
    })
}

