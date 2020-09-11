package main

import "fmt"
import "math"

func myAtoi(str string) int {
    sign, i, result := 1, 0, 0

    for ; i < len(str) && str[i] == ' '; i++{}  
    
    fmt.Printf("ith value is %d\n", i)
    if i < len(str) && (str[i] == '-' || str[i] == '+'){        
        if str[i] == '-'{
	   sign = -1
	} 
 	i++	           
    }    
    
    for ;i < len(str) && '0' <= str[i] && str[i] <= '9'; i++{        
            result = result * 10 + int(str[i]-'0')
            if sign*result < math.MinInt32 {
		return math.MinInt32
	    }
	    if sign*result > math.MaxInt32 {
		return math.MaxInt32
	    }        
        
    }
    
    return sign * result
    
}

func main() {
	fmt.Printf("%d\n", myAtoi("word and 1234"))
	fmt.Printf("Expected = -42 Output = %d\n", myAtoi("   -42"))
	fmt.Printf("%d \n", myAtoi("9234 and number"))
	fmt.Printf("%d\n", myAtoi("-12345"))
	fmt.Printf("Expected = 92 Output = %d\n", myAtoi("   +92"))
	fmt.Printf("%d\n", myAtoi("     "))
	fmt.Printf("%d\n", myAtoi("-"))
	fmt.Printf("%d\n", myAtoi("+"))
	fmt.Printf("%d\n", myAtoi("9223372036854775808"))
}