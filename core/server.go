package core

import (
	"encoding/base64"
	"fmt"
	"izicoder/cakeauth/core/api"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/bcrypt"
)

func StartServer() error {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/fresh", api.HandleFresh)
	e.GET("/refresh", api.HandleRefresh)
	return e.Start(":9999")
}

type LoginResp struct {
	Guid         string `json:"guid"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AccessTokenPayload struct {
	jwt.Claims
	User string `json:"user"`
	IP   string `json:"ip"`
}

type RefreshTokenPlayload struct {
	jwt.Claims
	Partner AccessTokenPayload
}

func Login(e echo.Context) error {
	guid := e.QueryParam("guid")
	accessTok := AccessTokenPayload{
		IP:     e.Request().RemoteAddr,
		User:   guid,
		Claims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(time.Minute * 10)))},
	}
	access, _ := jwt.NewWithClaims(jwt.SigningMethodHS512, accessTok).SignedString(GetSecret())

	refreshTok := RefreshTokenPlayload{
		Partner: accessTok,
	}
	refresh, _ := jwt.NewWithClaims(jwt.SigningMethodHS512, refreshTok).SignedString(GetSecret())
	psswd := []byte(refresh)
	psswd = psswd[:72]
	enc, err := bcrypt.GenerateFromPassword(psswd, bcrypt.DefaultCost)
	fmt.Println(err)

	resp := LoginResp{
		Guid:         guid,
		AccessToken:  access,
		RefreshToken: base64.URLEncoding.EncodeToString(enc),
	}
	return e.JSON(200, resp)
}

func Refresh(e echo.Context) error {
	return e.JSON(200, "")
}
