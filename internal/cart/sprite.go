package cart

import (
	"sort"
	"strings"
)

// sprites are stored in the cart
// to do that, we need to serialize and
// deserialize sprite data on the fly
type SpriteMemory struct {
	SpriteIndex uint8
	Sprites     map[uint8]string
}

const (
	startToken = "--startSprites"
	endToken   = "--endSprites"
)

func (s *SpriteMemory) LoadSpritesFromCart(cartText string) string {
	s.Sprites = map[uint8]string{}
	for i, spr := range s.ParseSprites(cartText) {
		s.Sprites[uint8(i)] = spr
	}
	return cartText
}
func (s *SpriteMemory) SaveSpritesToCart(cartText string) string {
	spriteList := []string{}

	sorted := []int{}

	for k, _ := range s.Sprites {
		sorted = append(sorted, int(k))
	}

	sort.Ints(sorted)

	for _, k := range sorted {
		spriteList = append(spriteList, s.Sprites[uint8(k)])
	}

	combined := strings.Join(spriteList, "\n--")

	// the sprite is already in the cart and we will have to remove them
	return s.stripSpritesFromCart(cartText) +
		startToken +
		"\n--" + combined +
		"\n" + endToken + "\n"
}
func (s SpriteMemory) ParseSprites(cartText string) []string {
	sprites := []string{}

	foundStart := false
	foundEnd := false
	for _, line := range strings.Split(cartText, "\n") {
		if strings.Contains(line, startToken) {
			foundStart = true
			continue
		}
		if strings.Contains(line, endToken) {
			foundEnd = true
			continue
		}
		if !foundStart {
			continue
		}
		if foundStart && foundEnd {
			break
		}
		sprite := strings.ReplaceAll(line, " ", "")
		sprite = strings.ReplaceAll(sprite, "-", "")

		sprites = append(sprites, sprite)
	}
	return sprites
}
func (SpriteMemory) Parse(spriteString string) []uint8 {
	sprite := []uint8{}
	for _, c := range strings.Split(spriteString, "") {
		n := 0

		switch c {
		case "0":
			n = 0
		case "1":
			n = 1
		case "2":
			n = 2
		case "3":
			n = 3
		case "4":
			n = 4
		case "5":
			n = 5
		case "6":
			n = 6
		case "7":
			n = 7
		case "8":
			n = 8
		case "9":
			n = 9
		case "a":
			n = 10
		case "b":
			n = 11
		case "c":
			n = 12
		case "d":
			n = 13
		case "e":
			n = 14
		case "f":
			n = 15
		}

		sprite = append(sprite, uint8(n))
	}

	return sprite
}
func (s SpriteMemory) stripSpritesFromCart(cartText string) string {
	var withoutSprites []string

	foundStart := false
	foundEnd := false
	for _, line := range strings.Split(cartText, "\n") {
		if strings.Contains(line, startToken) {
			foundStart = true
			continue
		}
		if strings.Contains(line, endToken) {
			foundEnd = true
			continue
		}
		if !foundStart {
			withoutSprites = append(withoutSprites, line)
			continue
		}
		if foundStart && !foundEnd {
			continue
		}

		withoutSprites = append(withoutSprites, line)
	}
	return strings.Join(withoutSprites, "\n")
}
func (s *SpriteMemory) StoreSprite(spriteString string) {
	s.Sprites[s.SpriteIndex] = spriteString
}
func (s *SpriteMemory) StoreSprites(spriteStrings []string) {
	for i, spr := range spriteStrings {
		s.Sprites[uint8(i)] = spr
	}
}
func (s *SpriteMemory) SetSpriteIndex(i uint8) {
	s.SpriteIndex = i
}
func (s SpriteMemory) GetSprite(i uint8) []uint8 {
	if len(s.Sprites) == 0 {
		return []uint8{}
	}
	if _, ok := s.Sprites[i]; !ok {
		return []uint8{}
	}
	return s.Parse(s.Sprites[i%uint8(len(s.Sprites))])
}
