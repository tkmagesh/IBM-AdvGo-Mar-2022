package db

import "fmt"

type DBConnection struct {
	ID int
}

//implements the io.Closer interface
func (c *DBConnection) Close() error {
	fmt.Printf("Closing Resource # %d\n", c.ID)
	return nil
}
