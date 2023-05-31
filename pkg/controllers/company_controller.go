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

var companyCollection *mongo.Collection = configs.GetCollection(configs.DB, "companies")

func CreateCompany(c echo.Context) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    var company models.Company
    defer cancel()

    //validate the request body
    if err := c.Bind(&company); err != nil {
        return c.JSON(http.StatusBadRequest, responses.CompanyResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
    }

    //use the validator library to validate required fields
    if validationErr := validate.Struct(&company); validationErr != nil {
        return c.JSON(http.StatusBadRequest, responses.CompanyResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": validationErr.Error()}})
    }

    newCompany := models.Company{
        Id:       primitive.NewObjectID(),
        Name:     company.Name,			
    }

    result, err := companyCollection.InsertOne(ctx, newCompany)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, responses.CompanyResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
    }

    return c.JSON(http.StatusCreated, responses.CompanyResponse{Status: http.StatusCreated, Message: "success", Data: &echo.Map{"data": result}})
}


func GetACompany(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	companyId := c.Param("companyId")
	var company models.Company
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(companyId)

	err := companyCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&company)

	if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.CompanyResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(http.StatusOK, responses.CompanyResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": company}})
}




func EditACompany(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	companyId := c.Param("companyId")
	var company models.Company
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(companyId)

	//validate the request body
	if err := c.Bind(&company); err != nil {
			return c.JSON(http.StatusBadRequest, responses.CompanyResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&company); validationErr != nil {
			return c.JSON(http.StatusBadRequest, responses.CompanyResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": validationErr.Error()}})
	}

	update := bson.M{"name": company.Name}

	result, err := companyCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})

	if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.CompanyResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	//get updated company details
	var updateCompany models.Company
	if result.MatchedCount == 1 {
			err := companyCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updateCompany)

			if err != nil {
					return c.JSON(http.StatusInternalServerError, responses.CompanyResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
			}
	}

	return c.JSON(http.StatusOK, responses.CompanyResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": updateCompany}})
}




func DeleteACompany(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	companyId := c.Param("companyId")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(companyId)

	result, err := companyCollection.DeleteOne(ctx, bson.M{"id": objId})

	if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.CompanyResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	if result.DeletedCount < 1 {
			return c.JSON(http.StatusNotFound, responses.CompanyResponse{Status: http.StatusNotFound, Message: "error", Data: &echo.Map{"data": "Company with specified ID not found!"}})
	}

	return c.JSON(http.StatusOK, responses.CompanyResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": "Company successfully deleted!"}})
}


func GetAllCompanies(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var companies []models.Company
	defer cancel()

	results, err := companyCollection.Find(ctx, bson.M{})

	if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.CompanyResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
			var singleCompany models.Company
			if err = results.Decode(&singleCompany); err != nil {
					return c.JSON(http.StatusInternalServerError, responses.CompanyResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
			}

			companies = append(companies, singleCompany)
	}

	return c.JSON(http.StatusOK, responses.CompanyResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": companies}})
}