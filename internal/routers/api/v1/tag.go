package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Tag struct {

}
func NewTag() Tag{
	return Tag{}
}
func (t Tag) Get(c *gin.Context){}
func (t Tag) List(c *gin.Context){}
func (t Tag) Create(c *gin.Context) {
	fmt.Println("Tag Create")
}
func (t Tag) Update(c *gin.Context)  {
}
func (t Tag) Delete(c *gin.Context){
}
