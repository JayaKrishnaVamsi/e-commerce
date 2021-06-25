//Routes/Routes.go
package Routes

import (
	"Project3e/Controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/customer")
	{
		grp1.GET("all_customers", Controllers.GetCustomers)
		grp1.POST("customers", Controllers.CreateCustomer)
		grp1.GET("customers/:id", Controllers.GetCustomerByID)
		grp1.PUT("customers/:id", Controllers.UpdateCustomer)
		grp1.DELETE("customers/:id", Controllers.DeleteCustomer)
	}
	grp2 := r.Group("/product")
	{
		grp2.GET("all_products", Controllers.GetProducts)
		grp2.GET("products/:id", Controllers.GetProductByID)

	}
	grp3 := r.Group("/mycart")
	{
		grp3.GET("allitems", Controllers.GetCart)
		grp3.POST("additem", Controllers.AddCart)

	}
	grp4 := r.Group("/payment")
	{
		grp4.GET("allpayments", Controllers.GetPayment)
		grp4.POST("pay", Controllers.CreatePayment)

	}

	return r
}
