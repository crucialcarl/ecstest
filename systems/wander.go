package systems

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/crucialcarl/ecstest/entity"
)

func WanderEntities(e *entity.Entity) {
	seed := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(seed)

	dir := []string{"North", "South", "East", "West"}
	chosen := randGen.Intn(4)
	if e.HasComponent("wander") {
		fmt.Printf("WANDER: %s wanders to the %s\n", e.Name, dir[chosen])
	}
}
