package main

import "fmt"

type customSet struct{
    set map[string] struct{}
}

func makeSet() *customSet{
    return &customSet{set:make(map[string] struct{}), }
}

func (myset *customSet)Add(s string){
    myset.set[s] = struct{}{}
}

func (myset *customSet) Exists(key string) bool{
    val, exists := myset.set[key]
    return exists
}

func (myset *customSet) Remove(key string) error{
    val, exists := myset.set[key]
    fmt.Printf("Value in the myset for the key", key, val)
    if !exists{
        return fmt.Errorf("Key does not exists")
    }
    
    delete(myset.set, key)
    return fmt.Errorf("Success")
    
}


func main(){
    customSet := makeSet()
    customSet.Add("A")
    customSet.Add("B")
    fmt.Println(customSet.Remove("B"))
    
}
