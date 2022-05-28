package storage

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"log"
	"net/http"
	_ "recommendersystem/docs"
)

// Registration godoc
// @Description  registration handler
// @Tags         All
// @Accept       json
// @Produce      json
// @Param        request body UserInfo true "Request structure"
// @Success      200	{object} AuthResponse
// @Failure      404
// @Failure      409
// @Failure      507
// @Router       /api/user/registration [post]
func (s *Server) Registration(c *fiber.Ctx) error {
	var u UserInfo
	err := c.BodyParser(&u)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	token, exist, err := s.UserStore.RegistrationUser(u)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	if exist {
		return c.SendStatus(http.StatusConflict)
	}

	jwt := GenerateToken(token)
	return c.Status(http.StatusOK).JSON(AuthResponse{Token: jwt})
}

// Login godoc
// @Description  authorization handler
// @Tags         All
// @Accept       json
// @Produce      json
// @Param        request body UserInfo true "Request structure"
// @Success      200	{object} AuthResponse
// @Failure      404
// @Failure      401
// @Failure      507
// @Router       /api/user/auth [post]
func (s *Server) Login(c *fiber.Ctx) error {
	var u UserInfo
	err := c.BodyParser(&u)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	uid, ok, err := s.UserStore.AuthorizationUser(u)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	if !ok {
		return c.SendStatus(http.StatusUnauthorized)
	}

	jwt := GenerateToken(uid)
	return c.Status(http.StatusOK).JSON(AuthResponse{Token: jwt})
}

// Tickets godoc
// @Description  handler for tickets
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param 		 Authorization header string true "Insert your access token" default(<Add access token here>)
// @Param        request body TicketsRequest true "Request structure"
// @Success      200	{object} TicketsResponse
// @Failure      403
// @Failure      404
// @Failure      507
// @Router       /api/tickets/ [post]
func (s *Server) Tickets(c *fiber.Ctx) error {
	// _______Auth check_______
	_, err := CheckToken(c.Get("Authorization"))
	if err != nil {
		return c.SendStatus(http.StatusForbidden)
	}
	// _______Auth check_______

	var req TicketsRequest
	err = c.BodyParser(&req)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	tickets, err := s.TicketsStore.GetTickets(req)
	if err != nil {
		log.Println(err)
		return c.SendStatus(http.StatusInsufficientStorage)
	}
	return c.Status(http.StatusOK).JSON(TicketsResponse{Tickets: tickets})
}

// LeisureList godoc
// @Description  handler for recommendation
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param 		 Authorization header string true "Insert your access token" default(<Add access token here>)
// @Param        request body  algoStorage.RecommendRequest true "Request structure"
// @Success      200	{object}  algoStorage.RecommendResponse
// @Failure      403
// @Failure      404
// @Failure      507
// @Failure      500
// @Router       /api/leisure/list [post]
func (s *Server) LeisureList(c *fiber.Ctx) error {
	// _______Auth check_______
	claim, err := CheckToken(c.Get("Authorization"))
	if err != nil {
		return c.SendStatus(http.StatusForbidden)
	}
	// _______Auth check_______

	var req RecommendRequest
	err = c.BodyParser(&req)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	name, err := s.UserStore.GetUserName(claim.ID)
	if err != nil {
		return c.SendStatus(http.StatusInsufficientStorage)
	}

	req.Name = name
	ss, _ := json.Marshal(req)
	buf := new(bytes.Buffer)
	buf.WriteString(string(ss))

	resp, err := http.DefaultClient.Post("http://"+s.Cfg.AlgHost, "application/json; charset=utf-8", buf)
	if err != nil {
		log.Println(err)
		return c.SendStatus(http.StatusInternalServerError)
	}

	if resp.StatusCode == http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		var res RecommendResponse
		json.Unmarshal(body, &res)
		return c.Status(http.StatusOK).JSON(res)
	}
	return c.SendStatus(resp.StatusCode)

}
