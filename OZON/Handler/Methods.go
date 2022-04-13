package Handler

import (
	"net/http"
	database "test/DataBase"
	"test/structure"

	"github.com/gin-gonic/gin"
)

func (h *Handle) CreateShortURL(ctx *gin.Context) {

	var input structure.URLS
	db := database.ConnectDB()
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, structure.Errors{Error: "Json не найден"})
		return
	}
	if input.LongURL == "" {
		ctx.JSON(http.StatusNotFound, structure.Errors{Error: "Пустая строка"})
	}

	// if err := h.rdb.Set(context.Background(), input.ShortURL, input.LongURL, 0).Err(); err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, structure.Errors{Error: "Ошибка добавления"})
	// 	return
	// }
	var count int64 = 0
	if db.Find(&structure.URLS{}).Where("long_url = ?", input.LongURL).Count(&count); count != 0 {
		//ctx.JSON(http.StatusNotFound, structure.Errors{Error: "Запись уже существует"})
		db.Where("long_url = ?", input.LongURL).Find(&input)
		ctx.JSON(http.StatusOK, gin.H{"ShortURL": input.ShortURL})
	} else {
		input.ShortURL = CreateShortString()
		db.Create(&input)
		ctx.JSON(http.StatusOK, gin.H{"ShortURL": input.ShortURL})
	}
}

func (h *Handle) GetShortURL(ctx *gin.Context) {
	db := database.ConnectDB()
	var input structure.URLS
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, structure.Errors{Error: "Json не найден"})
		return
	}
	if input.ShortURL == "" {
		ctx.JSON(http.StatusNotFound, structure.Errors{Error: "Пустая строка"})
	}
	if g := db.Find(&structure.URLS{}).First(&input, "short_url=?", input.ShortURL); g.Error != nil {
		ctx.JSON(http.StatusNotFound, structure.Errors{Error: "Запись не существует"})
	} else {
		// url, err := h.rdb.Get(context.Background(), input.ShortURL).Result()
		// if err == redis.Nil {
		// 	ctx.JSON(http.StatusNotFound, structure.Errors{Error: "Id не найден в Redis"})
		// 	return
		// }
		// if err != nil {
		// 	fmt.Println("Ошибка добавления", err)
		// 	ctx.JSON(http.StatusInternalServerError, structure.Errors{Error: "Ошибка добавления"})
		// 	return
		// }
		ctx.JSON(http.StatusOK, gin.H{"LongURL": input.LongURL})
	}
	//ctx.Redirect(http.StatusMovedPermanently, url)

}
