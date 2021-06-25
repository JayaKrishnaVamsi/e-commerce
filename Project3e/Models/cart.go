//Models/User.go
package Models

import (
	"Project3e/Config"

	_ "github.com/go-sql-driver/mysql"
)

//GetAllUsers Fetch all user data
func GetCart(cart *[]Cart) (err error) {
	if err = Config.DB.Find(cart).Error; err != nil {
		return err
	}
	return nil
}

func AddCart(cart *Cart) (err error) {
	if err = Config.DB.Create(cart).Error; err != nil {
		return err
	}

	return nil
}

//GetUserByID ... Fetch only one user by Id
/*func GetProductByID(product *Product, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}*/

//UpdateUser ... Update user
/*func UpdateProduct(product *Product, id string) (err error) {
	fmt.Println(product)
	Config.DB.Save(product)
	return nil
}*/

//DeleteUser ... Delete user
/*func DeleteProduct(product *Product, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(product)
	return nil
}
*/
