package requesthandlers

import (
	"net/http"

	"github.com/cfindlayisme/whatcanimake/llm"
	"github.com/cfindlayisme/whatcanimake/model"
	"github.com/gin-gonic/gin"
)

const ApiPath = "/api/v1"

func GetRecipes(c *gin.Context) {
	query := model.QueryParameters{}
	c.Bind(&query)

	if query.Count == nil {
		defaultCount := 5
		query.Count = &defaultCount
	}

	if query.Ingredients == nil {
		c.String(http.StatusNotAcceptable, "Please provide ingredients")
	}

	if (query.PantryStaples != nil) && *query.PantryStaples {
		*query.Ingredients = *query.Ingredients + ", " + model.PantryStaples
	}

	recipes, _ := llm.GetRecipes(*query.Ingredients, *query.Count)

	c.JSON(http.StatusOK, recipes)
}
