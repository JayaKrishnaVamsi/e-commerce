//Controllers/User.go
package Controllers

import (
	"Project3e/Config"
	"Project3e/Models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//GetUsers ... Get all users
func GetPayment(c *gin.Context) {
	var payment []Models.Payment
	err := Models.GetPayment(&payment)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, payment)
	}
}

//CreateUser ... Create User
func CreatePayment(c *gin.Context) {
	var payment Models.Payment
	//var cart Models.Cart
	c.BindJSON(&payment)
	err := Models.CreatePayment(&payment)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, payment)
	}
	DoPayment(&payment)

	//fmt.Println(payment)
}

func DoPayment(payment *Models.Payment) {
	//fmt.Println(payment)
	po := payment.OId
	var result []Models.Cart
	Config.DB.Raw("SELECT * FROM cart WHERE order_id = ?", po).Scan(&result)
	fmt.Println(result)
	var total int = 0
	var r1 Models.Product
	for _, v := range result {
		Config.DB.Raw("SELECT * FROM product WHERE id = ?", v.PId).Scan(&r1)
		total = total + (int(r1.Price) * int(v.PQty))
	}
	fmt.Println("Total Bill")
	fmt.Println(total)
	Config.DB.Exec("UPDATE payment SET bill = ? WHERE o_id = ?", total, po)
	fmt.Println("Updated bill in db")

	for _, v := range result {
		Config.DB.Exec("UPDATE product SET quantity = ? WHERE id = ?", gorm.Expr("quantity - ?", v.PQty), v.PId)

	}
	fmt.Println("Quantity decremented in db")
}

/*
//GetUserByID ... Get the user by id
func GetProductByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var product Models.Product
	err := Models.GetProductByID(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

//UpdateUser ... Update the user information
func UpdateProduct(c *gin.Context) {
	var product Models.Product
	id := c.Params.ByName("id")
	err := Models.GetProductByID(&product, id)
	if err != nil {
		c.JSON(http.StatusNotFound, product)
	}
	c.BindJSON(&product)
	err = Models.UpdateProduct(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}
*/
//DeleteUser ... Delete the user
/*func DeleteProduct(c *gin.Context) {
	var product Models.Product
	id := c.Params.ByName("id")
	err := Models.DeleteProduct(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}*/
