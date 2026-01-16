package reverse

import (
	"crypto/rand"
	"io"

	"github.com/xtls/xray-core/common/dice"
)

func (c *Control) FillInRandom(padding uint32) {
	randomLength := dice.Roll(int(padding))
	randomLength++
	c.Random = make([]byte, randomLength)
	io.ReadFull(rand.Reader, c.Random)
}
