package service

import (
	"init/entity"
)

func Find(c *entity.Customer) (err error) {

	c.ID = "dslfjslkfjldjs234l3j4l2j4"
	c.Age = 20
	c.Code = "C0001"
	c.Name = "ABC"
	c.Email = "abc@gmail.com"

	return nil
}
