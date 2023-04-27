package delivery

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/config"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/delivery/controller"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/delivery/middleware"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/manager"
	"github.com/sirupsen/logrus"
	"log"
)

type Server struct {
	ucManager manager.UseCaseManager
	engine    *gin.Engine
	host      string
	log       *logrus.Logger
}

func (s *Server) initController() {
	// add middleware
	s.engine.Use(middleware.LogRequestMiddleware(s.log))
	controller.NewBrandController(s.engine, s.ucManager.BrandUseCase())
	controller.NewVehicleController(s.engine, s.ucManager.VehicleUseCase())
	controller.NewCustomerController(s.engine, s.ucManager.CustomerUseCase())
	controller.NewTransactionController(s.engine, s.ucManager.TransactionUseCase())
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Printf("failed to serve config :%s", err)
	}
	infra, err := manager.NewInfraManager(cfg)
	if err != nil {
		log.Printf("failed connect to infra  :%s", err)
	}
	repo := manager.NewRepositoryManager(infra)
	uc := manager.NewUseCaseManager(repo)
	r := gin.Default()
	host := fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)
	return &Server{
		ucManager: uc,
		engine:    r,
		host:      host,
		log:       infra.Log(),
	}
}

func (s *Server) Run() {
	s.initController()
	err := s.engine.Run(s.host)
	if err != nil {
		log.Printf("failed to run server :%s", err)
	}
}
