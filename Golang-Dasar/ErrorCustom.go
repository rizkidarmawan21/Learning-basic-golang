package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"ID is required"}
	}

	if id != "Darms" {
		return &notFoundError{"Data not found"}
	}

	return nil
}

func main() {
	err := SaveData("Darms", nil)
	if err != nil {
		// if validationError, ok := err.(*validationError); ok {
		// 	fmt.Println("Validation Error: ", validationError.Message)
		// } else if notFoundError, ok := err.(*notFoundError); ok {
		// 	fmt.Println("Not Found Error: ", notFoundError.Message)
		// } else {
		// 	fmt.Println("Unknown Error: ", err.Error())
		// }

		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("Validation Error: ", finalError.Message)
		case *notFoundError:
			fmt.Println("Not Found Error: ", finalError.Message)
		default:
			fmt.Println("Unknown Error: ", finalError.Error())
		}

	} else {
		println("Data berhasil disimpan")
	}
}
