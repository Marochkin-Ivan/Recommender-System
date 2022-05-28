package algoStorage

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	_ "recommenderalgo/docs"
)

// Recommend godoc
// @Description  handler for recommendation
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body RecommendRequest true "Request structure"
// @Success      200	{object} RecommendResponse
// @Failure      404
// @Failure      507
// @Router       / [post]
func (s *AlgoServer) Recommend(c *fiber.Ctx) error {
	var req RecommendRequest
	err := c.BodyParser(&req)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	reqTup := RecommendTuple{Req: req, Name: req.Name}
	resChan := make(chan map[string]string)
	reqTup.Out = resChan

	s.AlgoChan <- reqTup
	res := <-reqTup.Out
	if res == nil {
		return c.SendStatus(http.StatusInsufficientStorage)
	}

	return c.Status(http.StatusOK).JSON(RecommendResponse{ArtsList: res})
}
