package common

import (
	"encoding/json"
	"os"
	"github.com/natefinch/lumberjack"
	log "github.com/sirupsen/logrus"
	"io"
)

const (
	StatusCodeUnknown = -1
	StatusCodeNotFound = 404
	StatusCodeBadRequest = 400
	StatusCodeOK      = 1000
	StatusMismatch = 10
	UnautohrizeErrorMessage = "Unauthorize Access"
	InternalServerErrorMessage = "Internal Server Error"
	WrongUsernameOrPassword = "Wrong username or password"
	UserCreated = "User created"
	UserAlreadyExist = "User already exist"
	UserNotExist = "User not exist"
	IDIsNotValid = "ID Is not valid"
	EmptyString = ""
	DeleteSuccess = "Delete success"
)

const (
	ErrUsernameEmpty      = "Username is empty"
	ErrPasswordEmpty  = "Password is empty"
	ErrNotObjectIDHex = "String is not a valid hex representation of an ObjectId"
)

const (
	ColUsers  = "users"
	ColMovies = "movies"
)


type Configuration struct {
	Port string `json:"port"`
	EnableGinConsoleLog bool `json:"enableGinConsoleLog"`
	EnableGinFileLog bool `json:"enableGinFileLog"`

	LogFileName string `json:"logFilename"`
	LogMaxSize int `json:"logMaxSize"`
	LogMaxBackups int `json:"logMaxBackups"`
	LogMaxAge int `json:"logMaxAge"`

	MgAddrs string `json:"mgAddrs"`
	MgDbName string `json:"mgDbName"`
	MgDbUsername string `json:"mgDbUsername"`
	MgDbPassword string `json:"mgDbPassword"`
	ClusterName string `json:"clusterName"`

	JwtSecretPassword string `json:"jwtSecretPassword"`
	Issuer string `json:"issuer"`
}

var (
	Config *Configuration
)

func LoadConfig() error {
	file, err := os.Open("config/config.json")
	if err != nil {
		return err
	}
	Config = new(Configuration)
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)	
	logFile := &lumberjack.Logger{
		Filename: Config.LogFileName,
		MaxSize: Config.LogMaxSize,
		MaxBackups: Config.LogMaxBackups,
		MaxAge: Config.LogMaxAge,
	}

	log.SetOutput(io.MultiWriter(os.Stdout, logFile))
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.JSONFormatter{})

	return nil
}