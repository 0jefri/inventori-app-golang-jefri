package common

import "fmt"

func CountTotalPage(total, perPage int) int {
	if (total % perPage) > 0 {
		fmt.Println(total, perPage)
		return (total / perPage) + 1
          
	}

	

	return total / perPage

	// if perPage <= 0 {
	// 	perPage = 10
	//     }
	//     if total <= 0 {

	// 	total = 1
	//     }
	//     return (total + perPage - 1) / perPage

}
