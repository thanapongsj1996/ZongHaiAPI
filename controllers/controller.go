package controllers

import (
	"gorm.io/gorm"
	"math"
	"strconv"
	"zonghai-api/models"

	"github.com/gin-gonic/gin"
)

type pagination struct {
	ctx     *gin.Context
	query   *gorm.DB
	records interface{}
}

func (p *pagination) paginate() *models.PagingResult {
	// 1. Get limit, page
	page, _ := strconv.Atoi(p.ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(p.ctx.DefaultQuery("limit", "12"))

	// 2. Count records
	//var count int64
	//p.query.Model(p.records).Count(&count)
	ch := make(chan int64)
	go p.countRecords(ch)
	count := <-ch

	// 3. Find records
	// limit , offset
	// EX. limit => 10
	// page => 1 , 1 -10 , offset = 0
	// page => 2 , 11 -20 , offset = 10
	// page => 3 , 21 -30 , offset = 20
	offset := (page - 1) * limit
	p.query.Limit(limit).Offset(offset).Find(p.records)

	// 4. Total page
	totalPage := int(math.Ceil(float64(count) / float64(limit)))

	// 5. Find nextPage
	var nextPage int
	if page == totalPage {
		nextPage = page
	} else {
		nextPage = page + 1
	}

	// 6. Create pagingResult
	result := models.PagingResult{
		Page:      page,
		Limit:     limit,
		Count:     count,
		PrevPage:  page - 1,
		NextPage:  nextPage,
		TotalPage: totalPage,
	}

	return &result
}

func (p *pagination) countRecords(ch chan int64) {
	var count int64
	p.query.Model(p.records).Count(&count)

	ch <- count
}
