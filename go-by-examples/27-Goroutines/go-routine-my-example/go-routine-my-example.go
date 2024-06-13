package goroutinemyexample

import (
	"fmt"
	"time"
)

func GoroutineMyExample() {
	start := time.Now()
	defer func() {
		fmt.Println("go routine my example end duration: ", time.Since(start))
	}()
	evilNinja := []string {"Evil Ninja", "Evil Ninja2", "Evil Ninja3", "Evil Ninja4"}

	for _, ninja := range evilNinja {
		go attack(ninja)
	}
	duration := float64(len(evilNinja)) * 0.5
	time.Sleep(time.Duration(duration) * time.Second)
}

func attack(ninja string) {
	start := time.Now()
	fmt.Println("Attacking", ninja)
	time.Sleep(1 * time.Second)
	defer func() {
		fmt.Print("Finished attacking ", ninja)
		fmt.Println(" time spent: ", time.Since(start))
	}()
}