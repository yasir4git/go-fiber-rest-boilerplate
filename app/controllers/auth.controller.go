package controllers

import (
	"errors"
	"net/mail"
	"time"

	"github.com/IKHINtech/go-fiber-rest-boilerplate/app/dto"
	"github.com/IKHINtech/go-fiber-rest-boilerplate/app/mappers"
	"github.com/IKHINtech/go-fiber-rest-boilerplate/app/models"
	"github.com/IKHINtech/go-fiber-rest-boilerplate/config"
	"github.com/IKHINtech/go-fiber-rest-boilerplate/database"
	"github.com/IKHINtech/go-fiber-rest-boilerplate/helpers"
	"github.com/IKHINtech/go-fiber-rest-boilerplate/utils"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// CheckPasswordHash compare password with hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func getUserByUserID(e string) (*models.User, error) {
	db := database.DB
	var user models.User
	query := db.Preload("Role")
	if err := query.Where("id = ?", e).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func isEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// Login get user and password godoc
// @Summary Login
// @Description Login
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body dto.LoginInput true "Login"
// @Success 200 {object} utils.ResponseData
// @Router /auth/login [post]
func Login(c *fiber.Ctx) error {
	h := &utils.ResponseHandler{}

	input := new(dto.LoginInput)
	var userData dto.UserData

	if err := c.BodyParser(&input); err != nil {
		return h.BadRequest(c, []string{"Error on login request", err.Error()})
	}

	identity := input.Phone
	pass := input.Password
	usermodels, err := new(models.User), *new(error)

	if isEmail(identity) {
		usermodels, err = helpers.GetUserByEmail(identity)
	} else {
		usermodels, err = helpers.GetUserByPhone(identity)
	}

	if usermodels == nil {
		return h.Forbidden(c, []string{"User not found"})
	}

	if !usermodels.IsActive {
		return h.Forbidden(c, []string{"User is not active"})
	}

	if err != nil {
		return h.InternalServerError(c, []string{err.Error(), "Internal server error"})
	} else if usermodels == nil {
		CheckPasswordHash(pass, "")
		return h.Forbidden(c, []string{"Invalid identity or password"})
	} else {
		userData = dto.UserData{
			ID:    usermodels.ID,
			Phone: usermodels.Phone,
			Name:  usermodels.Name,
			Email: usermodels.Email,
		}
	}

	if !CheckPasswordHash(pass, usermodels.Password) {
		return h.Forbidden(c, []string{"Invalid identity or password"})
	}

	// Create Access Token
	accessString, activeUntil, err := utils.GenerateAccessToken(userData)
	if err != nil {
		return h.InternalServerError(c, []string{"Failed to create access token", err.Error()})
	}

	accessToken := dto.Token{
		Token:     accessString,
		ExpiresIn: activeUntil,
	}

	// Create Refresh Token
	refreshString, err := utils.GenerateRefreshToken(userData.ID)
	if err != nil {
		return h.InternalServerError(c, []string{"Failed to create refresh token", err.Error()})
	}

	utils.SetRefreshTokenCookie(c, refreshString)

	response := dto.ResponseLogin{
		User:        userData,
		AccessToken: accessToken,
	}
	if err != nil {
		return h.InternalServerError(c, []string{err.Error()})
	}

	return h.Ok(c, response, "Success login", nil)
}

// GetUser get user godoc
// @Summary Get User
// @Description Get User
// @Tags Auth
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ErrorResponse
// @Router /auth/me [get]
func Me(c *fiber.Ctx) error {
	h := utils.ResponseHandler{}
	userID := c.Locals("user_id").(string)
	userData, err := getUserByUserID(userID)
	if err != nil {
		return h.InternalServerError(c, []string{"Error on get user data", err.Error()})
	}
	if userData == nil {
		return h.NotFound(c, []string{"User not found"})
	}

	userDto := mappers.UserModel_ToUserDTO(userData)
	return h.Ok(c, userDto, "Success login", nil)
}

// Register handles user registration godoc
// @Summary Register
// @Description Register
// @Tags Auth
// @Accept json
// @Produce json
// @Param register body dto.RegisterInput true "Register"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ErrorResponse
// @Router /auth/register [post]
func Register(c *fiber.Ctx) error {
	input := new(dto.RegisterInput)
	h := &utils.ResponseHandler{}
	if err := c.BodyParser(&input); err != nil {
		return h.BadRequest(c, []string{err.Error(), "Invalid Input"})
	}

	// Validate input fields
	err := utils.ValidateStruct(input)
	if err := c.BodyParser(&input); err != nil {
		return h.BadRequest(c, []string{err.Error(), "Invalid Input"})
	}

	// Check if email is valid
	if input.Email == nil {
		if !isEmail(*input.Email) {
			return h.BadRequest(c, []string{"Email is not valid"})
		}
	}

	// Check if username already exists
	existingUserByPhone, _ := helpers.GetUserByPhone(input.Phone)
	if existingUserByPhone != nil {
		return h.BadRequest(c, []string{"Username already exists"})
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return h.InternalServerError(c, []string{err.Error(), "Failed to hash password"})
	}

	// Create new user
	user := models.User{
		Name:        input.Name,
		Email:       input.Email,
		Phone:       input.Phone,
		CountryCode: input.CountryCode,
		Password:    string(hashedPassword),
	}

	// Save user to database
	db := database.DB
	if err := db.Create(&user).Error; err != nil {
		return h.InternalServerError(c, []string{err.Error(), "Failed to create user"})
	}
	return h.Created(c, user, "User registered successfully")
}

// RefreshToken handles refresh token request godoc
// @Summary Refresh Token
// @Description Refresh Token
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ErrorResponse
// @Router /auth/refresh-token [get]
func RefreshToken(c *fiber.Ctx) error {
	h := &utils.ResponseHandler{}

	// Ambil refresh token dari cookie
	refreshToken := c.Cookies("refresh_token") // Ganti "refresh_token" dengan nama cookie yang sesuai

	// Pastikan token ada
	if refreshToken == "" {
		return h.Unauthorized(c, []string{"Refresh token is missing"})
	}
	// Parse and validate the refresh token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrUnauthorized
		}
		return []byte(config.AppConfig.SECRET), nil
	})

	if err != nil || !token.Valid {
		return h.Forbidden(c, []string{"Invalid or expired refresh token"})
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)

	usermodels, err := getUserByUserID(userID)
	if err != nil {
		return h.InternalServerError(c, []string{"Error on get user data", err.Error()})
	}

	if usermodels == nil {
		return h.Unauthorized(c, []string{"User not found"})
	}

	userData := dto.UserData{
		ID:    usermodels.ID,
		Name:  usermodels.Name,
		Phone: usermodels.Phone,
		Email: usermodels.Email,
	}
	// Generate new access token
	accessString, accessTime, err := utils.GenerateAccessToken(userData)
	if err != nil {
		return h.InternalServerError(c, []string{"Failed to create access token", err.Error()})
	}

	responseL := dto.ResponseLogin{
		User: userData,
		AccessToken: dto.Token{
			Token:     accessString,
			ExpiresIn: accessTime,
		},
	}

	return h.Ok(c, responseL, "Token refreshed successfully", nil)
}

// Logout handles logout request godoc
// @Summary Logout
// @Description Logout
// @Tags Auth
// @AAccept json
// @Produce json
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ErrorResponse
// @Router /auth/logout [delete]
func Logout(c *fiber.Ctx) error {
	// Menghapus refresh token dari cookie HTTP-only
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",            // Nama cookie yang disetel saat login
		Value:    "",                         // Hapus nilai cookie
		Expires:  time.Now().Add(-time.Hour), // Set waktu kedaluwarsa di masa lalu untuk menghapus cookie
		HTTPOnly: true,                       // Pastikan cookie hanya dapat diakses oleh server
		Secure:   true,                       // Hanya kirim cookie di koneksi HTTPS
		SameSite: "lax",                      // Perlindungan CSRF
	})

	// Kembalikan respons sukses tanpa token
	return c.JSON(fiber.Map{
		"message": "Successfully logged out",
	})
}
