package controllers

import "github.com/gofiber/fiber/v2"

// UserSignUp method for creating a new user.
// @Description create a new user.
// @Summary create a new user.
// @Tags User
// @Accept json
// @Produce json
// @Param email body string true "Email"
// @Param password body string true "Password"
// @Param user_role body string true "User role"
// @Success 200 {object} model.User
// @Router /v1/user/sign/up [post]
func UserSignUp(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
