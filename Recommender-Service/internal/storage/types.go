package storage

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type Config struct {
	Host    string `yaml:"host"`
	AlgHost string `yaml:"alg_host"`
	Store   string `yaml:"store"`
}

type Server struct {
	Cfg          Config
	UserStore    Storable
	TicketsStore Storable
	App          *fiber.App
}

type Storable interface {
	RegistrationUser(u UserInfo) (string, bool, error)  // return internal id
	AuthorizationUser(u UserInfo) (string, bool, error) // return internal id
	GetTickets(object TicketsRequest) ([]Ticket, error) // return array of tickets
	GetUserName(id string) (string, error)              // return name
}

type UserInfo struct {
	Login    string `json:"login" form:"login"`
	Password string `json:"password" form:"password"`
}

type TicketsRequest struct {
	From string `json:"from" form:"from"`
	To   string `json:"to" form:"to"`
	Date string `json:"date" form:"date"`
}

type TicketsResponse struct {
	Tickets []Ticket `json:"tickets" form:"tickets"`
}

type Ticket struct {
	From                  string `json:"from" form:"from"`
	To                    string `json:"to" form:"to"`
	Date                  string `json:"date" form:"date"`
	DepartureTime         string `json:"departureTime" form:"departureTime"`
	ArrivalTime           string `json:"arrivalTime" form:"arrivalTime"`
	Transfer              string `json:"transfer" form:"transfer"`
	TransferDepartureTime string `json:"transferDepartureTime" form:"transferDepartureTime"`
	TransferTime          string `json:"transferTime" form:"transferTime"`
}

type AuthResponse struct {
	Token string `json:"token" form:"token"`
}

type RecommendRequest struct {
	Name   string `json:"name,omitempty" form:"name"`
	City   string `json:"city" form:"city"`
	InTime string `json:"in_time" form:"in_time"`
}

type RecommendResponse struct {
	ArtsList map[string]string `json:"arts_list" form:"arts_list"`
}

type Points struct {
	Name   string  `json:"name" form:"name"`
	Art    string  `json:"art" form:"art"`
	Points float64 `json:"points" form:"points"`
}

type Arts struct {
	Art    string `json:"art" form:"art"`
	InTime string `json:"in_time" form:"in_time"`
}

type Claims struct {
	jwt.StandardClaims
	ID string `json:"id"`
}
