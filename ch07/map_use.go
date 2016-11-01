package main

import "fmt"

func main() {
	hist := make(map[string]int, 6)
	hist["Jan"] = 100
	hist["Feb"] = 445
	hist["Mar"] = 514
	hist["Apr"] = 233
	hist["May"] = 321
	hist["Jun"] = 644
	hist["Jul"] = 113
	save(hist, "Aug", 734)
	save(hist, "Sep", 553)
	save(hist, "Oct", 344)
	save(hist, "Nov", 831)
	save(hist, "Dec", 312)
	save(hist, "Dec0", 332)
	remove(hist, "Dec0")

	for key, val := range hist {
		adjVal := int(float64(val) * 0.100)
		fmt.Printf("%s (%d):", key, val)
		for i := 0; i < adjVal; i++ {
			fmt.Print(".")
		}
		fmt.Println()
	}
}

// saves into map, blows up upon dup key
func save(store map[string]int, key string, value int) {
	val, ok := store[key]
	if !ok {
		store[key] = value
	} else {
		panic(fmt.Sprintf("Slot %d taken", val))
	}
}

// removes an entry, error if not found
func remove(store map[string]int, key string) error {
	_, ok := store[key]
	if !ok {
		return fmt.Errorf("Key not found")
	}
	delete(store, key)
	return nil
}
