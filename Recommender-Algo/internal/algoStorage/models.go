package algoStorage

import (
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	Host    string `yaml:"host"`
	AlgHost string `yaml:"alg_host"`
	Store   string `yaml:"store"`
}

type AlgoServer struct {
	Cfg      Config
	ArtStore ArtStorable
	App      *fiber.App
	AlgoChan chan RecommendTuple
}

type ArtStorable interface {
	GetArtsForRecommend() (map[string]map[string]float64, error)
	GetArtsByCity(object RecommendRequest) (map[string]string, error)
}

type RecommendRequest struct {
	Name   string `json:"name,omitempty" form:"name"`
	City   string `json:"city" form:"city"`
	InTime string `json:"in_time" form:"in_time"`
}

type RecommendResponse struct {
	ArtsList map[string]string `json:"arts_list" form:"arts_list"`
}

type RecommendTuple struct {
	Req  RecommendRequest
	Out  chan map[string]string
	Name string
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
