package creature

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitialHealth(t *testing.T) {
	var creature = Creature{Health: 10}

	assert.Equal(t, 10, creature.Health)
}

func TestDamage(t *testing.T) {
	var creature = Creature{Health: 10}

	creature = creature.Damage(1)

	assert.Equal(t, 9, creature.Health)
}



