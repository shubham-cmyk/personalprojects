package controllers

import (
	"context"
	"log"
	"net/http"
	"restaurentServiceProject/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"helm.sh/helm/v3/pkg/time"
)

// var menuCollection *mongo.Collection = database.OpenCollection(database.Client,"menu")
// Already Declared the menuCollection in the food Controller

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Creating a empty context background with timeout
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		// Now we can use Find to find the complete list of the menus and store it on the collection
		result, err := menuCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{"error": err.Error()})
			return
		}
		// This time we don't decode it on the json instead we use .
		// All on the Result and make use the pointer to allmenu to decode on it
		var allMenus []bson.M
		if err = result.All(ctx, &allMenus); err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, allMenus)

	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		menuID := c.Param("menu_id")
		var menu models.Menu
		err := menuCollection.FindOne(ctx, bson.M{"menu_id": menuID}).Decode(&menu)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Menu was not found in database"})
		}
		defer cancel()

		c.JSON(http.StatusOK, menu)

	}
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

		var menu models.Menu

		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "error while binding json to food struct"})
			return
		}

		validationError := validate.Struct(menu)
		if validationError != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationError.Error()})
			return
		}

		menu.Created_at, _ = time.Parse(time.RFC1123, time.Now().Format(time.RFC1123))
		menu.Updated_at, _ = time.Parse(time.RFC1123, time.Now().Format(time.RFC1123))

		menu.ID = primitive.NewObjectID()
		menu.Menu_id = menu.ID.Hex()

		result, insertErr := menuCollection.InsertOne(ctx, menu)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "menu item was not created"})
		}

		defer cancel()

		c.JSON(http.StatusOK, result)
	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

		// JSON IS BINDED ON THE var Menu
		var menu models.Menu

		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validate the struct
		validationError := validate.Struct(menu)
		if validationError != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationError.Error()})
		}

		var menuID = c.Param("menu_id")
		filter := bson.M{"menu_id": menuID}

		// Primivite Object is of type D and E type will be store in it and we would return this type.
		var updateObj primitive.D

		if menu.Start_Date != nil && menu.End_Date != nil {
			if !inTimeSpan(*menu.Start_Date, *menu.End_Date, time.Now()) {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "please retype the date"})
			}
			defer cancel()
			return
		}

		updateObj = append(updateObj, bson.E{"start_date", menu.Start_Date})
		updateObj = append(updateObj, bson.E{"end_date", menu.End_Date})

		menu.Created_at, _ = time.Parse(time.RFC1123, time.Now().Format(time.RFC1123))
		menu.Updated_at, _ = time.Parse(time.RFC1123, time.Now().Format(time.RFC1123))

		upsert := true

		opt := options.UpdateOptions{Upsert: &upsert}

		// menuCollection.UpdateOne(ctx,filter,bson.D{"$set",updateObj},&opt,)
		result, err := menuCollection.UpdateOne(ctx, filter, updateObj, &opt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "menu update failed"})
			defer cancel()
			return
		}

		c.JSON(http.StatusOK, result)
		defer cancel()

	}
}

func inTimeSpan(start, end, check time.Time) bool {
	return start.After(time.Now()) && end.After(start)
}
