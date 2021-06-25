//Models/Customer.go
package Models

import (
	"fmt"

	"Project3e/Config"

	_ "github.com/go-sql-driver/mysql"
)

//GetAllCustomers Fetch all user data
func GetAllCustomers(customer *[]Customer) (err error) {
	if err = Config.DB.Find(customer).Error; err != nil {
		return err
	}
	return nil
}

//CreateCustomer ... Insert New data
func CreateCustomer(customer *Customer) (err error) {
	if err = Config.DB.Create(customer).Error; err != nil {
		return err
	}
	return nil
}

//GetCustomerByID ... Fetch only one user by Id
func GetCustomerByID(customer *Customer, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(customer).Error; err != nil {
		return err
	}
	return nil
}

//UpdateCustomer ... Update user
func UpdateCustomer(customer *Customer, id string) (err error) {
	fmt.Println(customer)
	Config.DB.Save(customer)
	return nil
}

//DeleteCustomer ... Delete user
func DeleteCustomer(customer *Customer, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(customer)
	return nil
}
