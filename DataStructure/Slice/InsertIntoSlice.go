/*
Make a function that inserts a string slice into another string slice at a certain index.

Input #
Two Slices and an integer for index

Sample input #
["M", "N", "O", "P", "Q", "R"] // slice in which another slice will be added
["A" "B" "C"] // slice to be appended
0             // index

*/

package main

// slice is the slice to which another slice will be added
// insertion is the slice that needs to be inserted.
// index is the position for insertion
func insertSlice_way1(slice, insertion []string, index int) []string {
	for idx := 0; idx < len(insertion); idx++ {
		slice = append(slice, "")
	}

	copy(slice[len(insertion)+index:], slice[index:])
	copy(slice[index:], insertion)

	return slice
}

func insertSlice(slice, insertion []string, index int) []string {
	result := make([]string, len(slice)+len(insertion))
	at := copy(result, slice[:index]) //The copy function returns number of bytes copied.
	at += copy(result[at:], insertion)
	copy(result[at:], slice[index:])
	return result
}
 
func main() {	
	insertSlice({}, {A,B,C}, 0)	
	insertSlice({M,N,O,P,Q,R},{A,B,C},0)	
	insertSlice({M,N,O,P,Q,R},{A,B,C},3)	
	insertSlice({M,N,O,P,Q,R},{},0)	
	
}
