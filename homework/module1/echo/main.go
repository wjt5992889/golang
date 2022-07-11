package main

import (
	"fmt"
)

func main() {


	arr := []string{"I","am","stupid","and","weak"}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])

		if(i == 2){
			arr[i]="smart"
		}

		if(i == 4){
			arr[i]="strong"
		}
	}

}