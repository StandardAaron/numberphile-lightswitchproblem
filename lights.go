package main
    
import (
    "fmt"
    )

type lightsMap map[int]bool

//init all the lights to off
func initLights() lightsMap {
    lights := make(map[int]bool)
    for i := 1; i <= 100; i++ {
    lights[i] = false
    }
    return lights
}

//toggle the lights based on the algorithm of:
// person N toggles on/off all switches that are 
// multiples of N (inc N*1)
//Also: print the execution plan as we go.
func toggleLights(lights lightsMap) lightsMap {
    for person := 1; person <= 100; person++ {
        person_accum := person
        var person_toggles []int
        for {
            if person_accum > 100 {break}
            person_toggles = append(person_toggles, person_accum)
            person_accum += person
        }
        if person < 51 {
            fmt.Printf("Person %d is toggling the following switches: %v \n", person, person_toggles)
        }
        lights = flipSwitches(person_toggles, lights)
    }
    return lights
}

//Return the opposite state
func flipSwitches(toggleList []int, lights lightsMap) lightsMap {
    for _, element := range toggleList {
        if lights[element] == true {
            lights[element] = false
        } else {
            lights[element] = true
        }
    }
    return lights
}

//print the final state
func printState (lights lightsMap) {
    for i := 1; i <= 100; i++ {
        state := lights[i]
        if state {
            fmt.Printf("%d is ON\n", i)
        }
    }
    fmt.Println("all other lights are OFF")
}

//main func
func main() {
    lights := initLights()
    lights = toggleLights(lights)
    printState(lights)
}    
