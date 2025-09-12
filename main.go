package main

import (
	"fruits-api/cmd"
)

func main() {
	cmd.Serve()
}

// package main

// import (
// 	"fmt"
// 	"fruits-api/utils"
// )

// func main() {
// 	jwt, err := utils.CreateJwt("my-secret", utils.Payload{
// 		Sub:          "24",
// 		FullName:     "Rahat",
// 		Email:        "rahat@gmail.com",
// 		IsShopeOwner: false,
// 	})

// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	// JWT টোকেন টার্মিনালে দেখানোর জন্য
// 	fmt.Println("Generated JWT:", jwt)
// }
