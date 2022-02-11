package controller

import (
    "context"
    "net/http"
    "time"

    "github.com/programmerug/fibermongo/config"
    "github.com/programmerug/fibermongo/model"
    "github.com/programmerug/fibermongo/response"

    "github.com/go-playground/validator/v10"
    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = config.GetCollection(config.DB, "user")
var validate = validator.New()

func CreateUser(c *fiber.Ctx) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    var user model.User
    defer cancel()

    // validate the request body
    if err := c.BodyParser(&user); err != nil {
        return c.Status(http.StatusBadRequest).JSON(response.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
    }

    // use the validator library to validate required fields
    if validationErr := validate.Struct(&user); validationErr != nil {
        return c.Status(http.StatusBadRequest).JSON(response.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
    }

    newUser := model.User{
        ID:       primitive.NewObjectID(),
        Name:     user.Name,
        Location: user.Location,
        Title:    user.Title,
    }

    result, err := userCollection.InsertOne(ctx, newUser)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(response.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
    }

    return c.Status(http.StatusCreated).JSON(response.UserResponse{Status: http.StatusCreated, Message: "success", Data: result})
}