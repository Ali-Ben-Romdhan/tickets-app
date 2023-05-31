package controllers

import (
	"echo-mongo-api/pkg/configs"
	"echo-mongo-api/pkg/models"
	"echo-mongo-api/pkg/responses"
	"net/http"
	"time"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2/bson"
)

var departmentCollection *mongo.Collection = configs.GetCollection(configs.DB, "departments")

func CreateDepartment(c echo.Context) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    var department models.Department
    defer cancel()

    //validate the request body
    if err := c.Bind(&department); err != nil {
        return c.JSON(http.StatusBadRequest, responses.DepartmentResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
    }

    //use the validator library to validate required fields
    if validationErr := validate.Struct(&department); validationErr != nil {
        return c.JSON(http.StatusBadRequest, responses.DepartmentResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": validationErr.Error()}})
    }

    newDepartment := models.Department{
        Id:       primitive.NewObjectID(),
        Name:     department.Name,			
				CompanyID:   primitive.NewObjectID(),
    }

    result, err := departmentCollection.InsertOne(ctx, newDepartment)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, responses.DepartmentResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
    }

    return c.JSON(http.StatusCreated, responses.DepartmentResponse{Status: http.StatusCreated, Message: "success", Data: &echo.Map{"data": result}})
}


func GetADepartment(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	departmentId := c.Param("departmentId")
	var department models.Department
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(departmentId)

	err := departmentCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&department)

	if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.DepartmentResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(http.StatusOK, responses.DepartmentResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": department}})
}




func EditADepartment(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	departmentId := c.Param("departmentId")
	var department models.Department
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(departmentId)

	//validate the request body
	if err := c.Bind(&department); err != nil {
			return c.JSON(http.StatusBadRequest, responses.DepartmentResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&department); validationErr != nil {
			return c.JSON(http.StatusBadRequest, responses.DepartmentResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": validationErr.Error()}})
	}

	update := bson.M{"name": department.Name}

	result, err := departmentCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})

	if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.DepartmentResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	//get updated department details
	var updateDepartment models.Department
	if result.MatchedCount == 1 {
			err := departmentCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updateDepartment)

			if err != nil {
					return c.JSON(http.StatusInternalServerError, responses.DepartmentResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
			}
	}

	return c.JSON(http.StatusOK, responses.DepartmentResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": updateDepartment}})
}




func DeleteADepartment(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	departmentId := c.Param("departmentId")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(departmentId)

	result, err := departmentCollection.DeleteOne(ctx, bson.M{"id": objId})

	if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.DepartmentResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	if result.DeletedCount < 1 {
			return c.JSON(http.StatusNotFound, responses.DepartmentResponse{Status: http.StatusNotFound, Message: "error", Data: &echo.Map{"data": "Department with specified ID not found!"}})
	}

	return c.JSON(http.StatusOK, responses.DepartmentResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": "Department successfully deleted!"}})
}


func GetAllDepartments(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var departments []models.Department
	defer cancel()

	results, err := departmentCollection.Find(ctx, bson.M{})

	if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.DepartmentResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
			var singleDepartment models.Department
			if err = results.Decode(&singleDepartment); err != nil {
					return c.JSON(http.StatusInternalServerError, responses.DepartmentResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
			}

			departments = append(departments, singleDepartment)
	}

	return c.JSON(http.StatusOK, responses.DepartmentResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": departments}})
}