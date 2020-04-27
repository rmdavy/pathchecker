package main

import  (
		"fmt"
		"os"
		"strings"
)

func main() {

	//Banner
	fmt.Printf("PPC v0.1 - Path Privesc Checker\n")
	fmt.Printf("By Richard Davy @rd_pentest\n\n")

	fmt.Printf("Current Path Items\n\n")

	//Get path details
	s := os.Getenv("path")
	//Separate them by semi-colon
	ss := strings.FieldsFunc(s, func(r rune) bool {
		if r == ';' {
			return true
		}
		return false
	})
	
	//Loop through list of paths
	for i, s := range ss {
    	//Display position name then path
    	fmt.Println(i, s)

    	values := []string{}
    	values = append(values, s)
    	
    	//Check the last character of the path for a \ 
    	if string(s[len(s)-1:]) == "\\" {
        	values = append(values, "test.txt")
    	} else {
        	values = append(values, "\\test.txt")
    	}

    	//Concat the path directory with the text.txt string
		result := strings.Join(values, "")
	
		//Try to create the file
	    f, err := os.Create(result)
	
		//On error do nothing
	    if err != nil {
	        //fmt.Println(err)
	        //return
		}else{
		//If we successfully write the file
		//close the handle
		//display a message that it's a success
		//delete the file we created.
			f.Close()
			//fmt.Println(result)
			fmt.Println("	[!] Successfully wrote file test.txt to this folder!!, double check it was cleaned up...")
			os.Remove(result)
		}

	}
	
}