package client_core;

import (
	"bufio"
	"fmt"
	"os"
	"encoding/json"
	"github.com/topTalent1212/SQSClientServer/internal/sqs"
)

const inputCommandTemplateString = 
`
------------------------------
New Round
------------------------------
Type Number to Select Command

0 - Add Item
1 - Remove Item
2 - Get Item
3 - Get All Items
others - Exit
`;

const inputItemTemplateString = 
`
-------------------------
%s 

`

func StartClient(){

	for(true){
			message, err := scan();
			if(err != nil){
				continue;
			}
			sqs.SendMessage(message);
		}
}

func scan() (string, error){
	command := [2]string{"0", ""};
	
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(inputCommandTemplateString)
	scanner.Scan()
	// Holds the string that was scanned
	text := scanner.Text()
	switch text{
	case "0", "1", "2", "3":
		command[0] = text;
		placeholder := "Input Item to add"
		if (text == "0" || text == "1") {
			if (text == "1"){
				placeholder = "Input Item index to delete"
			} 
			fmt.Println(fmt.Sprintf(inputItemTemplateString, placeholder));
			scanner.Scan()
			command[1] = scanner.Text();
		}
		break;
	default:
		os.Exit(1);
	}

	// handle error
	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
		return "", scanner.Err();
	}

	message, err:= json.Marshal(command);
	if(err != nil){
		fmt.Println(err)
		return "", err
	}

	return string(message), nil;
}

