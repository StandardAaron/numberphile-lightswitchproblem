package main

import "fmt"

// calculate the number of factors that a given int has 
func howManyFactors(person int) int {
    var factors []int
    for i := 1; i <= person; i++ {
        if person % i == 0 {
            factors = append(factors, i)
        }
    }
    return len(factors)
}

func main() {
    for i :=0; i<=100; i++{
        factors := howManyFactors(i)
        //it turns out, for a light's state to be opposite it's
        //starting state, it needs to be toggled an odd number of
        //times, and that directly corelates to whether or not the 
        //light's position number has an odd number of factors.
        //The only time this is true, since factors come in pairs,
        //is when the number is an even square where a factor is
        //duplicated (i.e. intA * intA)
        if factors % 2 != 0 {fmt.Printf("%d is ON \n", i)}
    }
}
