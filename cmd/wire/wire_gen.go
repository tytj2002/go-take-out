// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalougata/go-take-out/api/admin"
	"github.com/kalougata/go-take-out/internal/controller/admin"
	"github.com/kalougata/go-take-out/internal/data"
	"github.com/kalougata/go-take-out/internal/server"
	"github.com/kalougata/go-take-out/internal/service/admin"
	"github.com/kalougata/go-take-out/pkg/config"
	"github.com/kalougata/go-take-out/pkg/jwt"
)

// Injectors from wire.go:

func NewApp() (*fiber.App, func(), error) {
	configConfig := config.NewConfig()
	dataData, cleanup, err := data.NewData(configConfig)
	if err != nil {
		return nil, nil, err
	}
	service := adminsrv.NewService(dataData)
	jwtJWT := jwt.NewJWT()
	employeeService := adminsrv.NewEmployeeService(service, jwtJWT)
	authController := adminctrl.NewAuthController(employeeService)
	adminAPIRouter := adminv1.NewAdminAPIRouter(authController)
	app := server.NewHTTPServer(adminAPIRouter)
	return app, func() {
		cleanup()
	}, nil
}
