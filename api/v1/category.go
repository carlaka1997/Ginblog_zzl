package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//添加分类
func AddCategory(c *gin.Context){
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	code = model.CheckCategory(data.Name)
	if code ==errmsg.SUCCSE{
		model.CreateCate(&data)
	}
	if code == errmsg.ERROR_CATENAME_USED{
		code = errmsg.ERROR_CATENAME_USED
	}
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})

}

//查询单个用户


//查询分类列表
func GetCate(c *gin.Context){
	pageSize,_ := strconv.Atoi(c.Query("pagesize"))
	pageNum,_:=strconv.Atoi(c.Query("pagenum"))

	if pageSize==0{
		pageSize = -1
	}
	if pageNum==0{
		pageNum = -1
	}
	data:=model.GetCate(pageSize,pageNum)
	code=errmsg.SUCCSE

	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}

//编辑分类名
func EditCate(c *gin.Context){
	var data model.Category
	id,_ :=strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code =model.CheckCategory(data.Name)
	if code ==errmsg.SUCCSE{
		model.EditCate(id,&data)
	}
	if code ==errmsg.ERROR_CATENAME_USED{
		c.Abort()
	}

	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}

//删除分类
func DeleteCate(c *gin.Context){
	id,_ :=strconv.Atoi(c.Param("id"))

	code = model.DeleteCate(id)

	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}


