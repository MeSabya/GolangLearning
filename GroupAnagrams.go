//Group anagrams 

package main

import ("fmt"
        "strings"
        "sort"
    )

func sortString(str string) string{
    s := strings.Split(str, "")
    sort.Strings(s)
    return strings.Join(s, "")
}
func groupAllAnagrams(strs []string) [][]string{
    
    dict := make(map[string][]string, 0)
    for _, str := range strs{
        s := sortString(str)
        dict[s] = append(dict[s], str)
        
    }
    result := make([][]string, 0)
    for _, val := range dict{
	result = append(result, val)
	
    }
    return result    

}
func main(){
    strs := []string{"cat", "dog", "act", "god", "ream"}
    results := groupAllAnagrams(strs)
    for _, result := range results{
         fmt.Println(result)
    }
}
