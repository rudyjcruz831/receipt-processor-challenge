package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rudyjcruz831/receipt-processor-challenge/model"
	"github.com/rudyjcruz831/receipt-processor-challenge/util/maputil"
)

func (h *Handler) Points(c *gin.Context) {
	// time.Sleep(6 * time.Second)

	id := c.Param("id")

	data := maputil.MyMap[id]

	// points := 0

	// if data.Retailer != "" {

	// }
	fmt.Println("data: ", data)

	// convert map to json
	jsonString, _ := json.Marshal(data)
	// fmt.Println(string(jsonString))

	s := model.Receipt{}
	json.Unmarshal(jsonString, &s)
	fmt.Println("s: ", s)

	// for _, c := range data.Retailer {
	// 	// if unicode.IsLetter(c) || unicode.IsDigit(c) {
	// 	// 	fmt.Printf("%c is alphanumeric\n", c)
	// 	// } else {
	// 	// 	fmt.Printf("%c is not alphanumeric\n", c)
	// 	// }
	// 	fmt.Printf("%c is alphanumeric\n", c)
	// }
	c.JSON(http.StatusOK, map[string]string{"Points": "processed"})
}

// One point for every alphanumeric character in the retailer name.
// 50 points if the total is a round dollar amount with no cents.
// 25 points if the total is a multiple of 0.25.
// 5 points for every two items on the receipt.
// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
// 6 points if the day in the purchase date is odd.
// 10 points if the time of purchase is after 2:00pm and before 4:00pm.

// convert json to struct
//   s := MyStruct{}
//   json.Unmarshal(jsonString, &s)
//   fmt.Println(s)
