package tdd_kata

type game struct {
	slots       [21]int
	currentSlot int
}

func (g *game) Roll(pins int) {
	g.slots[g.currentSlot] = pins
	g.currentSlot++
}

func (g *game) Score() int {
	score := 0
	frameIndex := 0
	for frame := 0; frame < 10; frame++ {
		if g.isStrike(frameIndex) {
			score += 10 + g.strikeBonus(frameIndex)
			frameIndex++
		} else if g.isSpare(frameIndex) {
			score += 10 + g.spareBonus(frameIndex)
			frameIndex += 2
		} else {
			score += g.twoBalls(frameIndex)
			frameIndex += 2
		}
	}
	return score
}

func (g *game) twoBalls(frameIndex int) int {
	return g.slots[frameIndex] + g.slots[frameIndex+1]
}

func (g *game) spareBonus(frameIndex int) int {
	return g.slots[frameIndex+2]
}

func (g *game) isStrike(frameIndex int) bool {
	return g.slots[frameIndex] == 10
}

func (g *game) strikeBonus(frameIndex int) int {
	return g.slots[frameIndex+1] + g.slots[frameIndex+2]
}

func (g *game) isSpare(frameIndex int) bool {
	return g.slots[frameIndex]+g.slots[frameIndex+1] == 10
}

func NewGame() *game {
	return &game{}
}
