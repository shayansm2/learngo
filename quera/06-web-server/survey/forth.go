package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := NewServer(1234)
	server.Start()
}

type Server struct {
	port      int
	ginEngine *gin.Engine
}

func NewServer(port int) *Server {
	core := NewSurvey()
	return &Server{
		port:      port,
		ginEngine: getEngine(core),
	}
}

func (server *Server) Start() {
	server.ginEngine.Run(fmt.Sprintf(":%d", server.port))
}

func handlerBuilder(f func(*gin.Context, *Survey), survey *Survey) gin.HandlerFunc {
	return func(c *gin.Context) {
		f(c, survey)
	}
}

func getEngine(core *Survey) *gin.Engine {
	e := gin.Default()
	e.GET("/", index)
	e.POST("flights", handlerBuilder(addFlights, core))
	e.POST("tickets", handlerBuilder(addTickets, core))
	e.POST("comments", handlerBuilder(addComments, core))
	e.GET("comments", handlerBuilder(getAllComments, core))
	e.GET("comments/:flight", handlerBuilder(getComments, core))
	return e
}

func index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "OK",
	})
}

type AddFlightsRequest struct {
	Flight string `json:"Name" binding:"required"`
}

func addFlights(c *gin.Context, core *Survey) {
	request := new(AddFlightsRequest)
	err := c.BindJSON(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "cannot work with json data in body of the request",
		})
		return
	}
	err = core.AddFlight(request.Flight)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"Message": "OK",
	})
}

type AddTicketRequest struct {
	Flight    string `json:"FlightName" binding:"required"`
	Passenger string `json:"PassengerName" binding:"required"`
}

func addTickets(c *gin.Context, core *Survey) {
	request := new(AddTicketRequest)
	err := c.BindJSON(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "cannot work with json data in body of the request",
		})
		return
	}
	err = core.AddTicket(request.Flight, request.Passenger)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"Message": "OK",
	})
}

type AddCommentRequest struct {
	Flight    string `json:"FlightName" binding:"required"`
	Passenger string `json:"PassengerName" binding:"required"`
	Score     int    `json:"Score" binding:"required"`
	Text      string `json:"Text" binding:"required"`
}

func addComments(c *gin.Context, core *Survey) {
	request := new(AddCommentRequest)
	err := c.BindJSON(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "cannot work with json data in body of the request",
		})
		return
	}
	err = core.AddComment(request.Flight, request.Passenger, Comment{request.Score, request.Text})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"Message": "OK",
	})
}

func getAllComments(c *gin.Context, core *Survey) {
	average := c.DefaultQuery("average", "false")
	if average == "true" {
		c.JSON(http.StatusOK, gin.H{
			"Message":  "OK",
			"Averages": core.GetAllCommentsAverage(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Message": "OK",
			"Texts":   core.GetAllComments(),
		})
	}
}

func getComments(c *gin.Context, core *Survey) {
	flight := c.Param("flight")
	average := c.DefaultQuery("average", "false")
	if average == "true" {
		avg, err := core.GetCommentsAverage(flight)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"Message": "OK",
			"Average": avg,
		})
	} else {
		comments, err := core.GetComments(flight)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"Message": "OK",
			"Texts":   comments,
		})
	}
}
