package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	bulletSize  = 32
	bulletSpeed = 1
)

type bullet struct {
	tex    *sdl.Texture
	x, y   float64
	active bool
	angle  float64
}

func newBullet(renderer *sdl.Renderer) *element {
	bullet := &element{}

	sr := newSpriteRenderer(bullet, renderer, "sprites/player_bullet.bmp")
	bullet.addComponent(sr)

	mover := newBulletMover(bullet, bulletSpeed)
	bullet.addComponent(mover)

	bullet.active = false

	return bullet
}

// this a bad practice in web dev but it's common in gaming dev (GLOBAL VARIABLES)
var bulletPool []*element

func initBulletPool(renderer *sdl.Renderer) {
	for i := 0; i < 30; i++ {
		bul := newBullet(renderer)
		elements = append(elements, bul)
		bulletPool = append(bulletPool, bul)
	}
}

func bulletFromPool() (*element, bool) {
	for _, bul := range bulletPool {
		if !bul.active {
			return bul, true
		}
	}
	return nil, false
}
