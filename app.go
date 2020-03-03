package main

import (
    "./types"
    "fmt"
    "strconv"
)


func main() {
    valor := types.RequestLoop{
        Input: 10,
    }

    nomes := makeRequest(valor)
    fmt.Println(nomes)
}



func makeRequest(input types.RequestLoop) []string {
    arrayName := []string{}
    for loop := 0; loop < input.Input; loop++ {
        arrayName = append(arrayName, "V"+strconv.Itoa(loop))
    }
    return arrayName
}

