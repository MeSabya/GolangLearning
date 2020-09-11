package main
import "fmt"

func main() {
   fmt.Printf("Merge Sort in Golang\n")
   slice := generateSlice(5)
   fmt.Println("Slice is ", slice)
   fmt.Println(mergeSort(slice))
}

func mergeSort(items []int) []int{
    nums := len(items)
    if nums == 1{
        return items
    }
    middle := int(nums/2)
    left := make([]int, middle)
    right := make([]int, nums-middle)
    for i := 0; i<nums; i++{
        if i < middle{
            left[i] = items[i]
        }else{
            right[i-middle] = items[i]
            
        }
    }
    
    return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) []int{
    
    result := make([]int, len(left)+len(right))
    i := 0
    
    for len(left) > 0 && len(right) > 0{
        if (left[0] < right[0]){
            result[i] = left[0]
            left = left[1:]
        }else{
            result[i] = right[0]
            right = right[1:]
        }
        i++
    }
        
    for j := 0; j < len(left); j++{
        result[i] = left[j]
        i++
    }
    
    for j := 0; j < len(right); j++{
        result[i] = right[j]
        i++
    }
    
    return result
    
}

func generateSlice(size int) []int{
    //Can we return short hand declaration for the param which will be returned from the function , yes 
    slice := []int {1, 9, 7,6, 6}
    return slice
    
    
}