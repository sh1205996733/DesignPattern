package main

import "fmt"

// 外观模式（门面模式、过程模式）封装底层实现，提供一组更加简单易用的高层接口
type Facade interface {
	Ready()
	Play()
	Pause()
	End()
}

type homeTheaterFacade struct {
	dvd          *dvdPlayer
	popcorn      *popcorn
	theaterLight *theaterLight
	screen       *screen
	projector    *projector
	stereo       *stereo
}

func NewHomeTheaterFacade() *homeTheaterFacade {
	return &homeTheaterFacade{
		dvd:          new(dvdPlayer),
		popcorn:      new(popcorn),
		theaterLight: new(theaterLight),
		screen:       new(screen),
		projector:    new(projector),
		stereo:       new(stereo),
	}
}

func (h *homeTheaterFacade) Ready() {
	h.popcorn.on()
	h.popcorn.pop()
	h.screen.down()
	h.projector.on()
	h.stereo.on()
	h.dvd.on()
	h.theaterLight.dim()
}
func (h *homeTheaterFacade) Play() {
	h.dvd.play()
}
func (h *homeTheaterFacade) Pause() {
	h.dvd.play()
}
func (h *homeTheaterFacade) End() {
	h.popcorn.off()
	h.theaterLight.bright()
	h.screen.up()
	h.projector.off()
	h.stereo.off()
	h.dvd.off()
}

type dvdPlayer string
type popcorn string
type theaterLight string
type screen string
type projector string
type stereo string

func (d *dvdPlayer) on() {
	fmt.Println(" dvd on ")
}
func (d *dvdPlayer) off() {
	fmt.Println(" dvd off ")
}
func (d *dvdPlayer) play() {
	fmt.Println(" dvd play ")
}
func (d *dvdPlayer) pause() {
	fmt.Println(" dvd play ")
}
func (p *popcorn) on() {
	fmt.Println(" popcorn on ")
}
func (p *popcorn) off() {
	fmt.Println(" popcorn off ")
}
func (p *popcorn) pop() {
	fmt.Println(" popcorn pop ")
}
func (s *screen) down() {
	fmt.Println(" screen down ")
}
func (s *screen) up() {
	fmt.Println(" screen up ")
}
func (s *stereo) on() {
	fmt.Println(" screen on ")
}
func (s *stereo) off() {
	fmt.Println(" screen off ")
}
func (p *projector) on() {
	fmt.Println(" projector on ")
}
func (p *projector) off() {
	fmt.Println(" projector off ")
}
func (t *theaterLight) dim() {
	fmt.Println(" theaterLight dim ")
}
func (t *theaterLight) bright() {
	fmt.Println(" theaterLight bright ")
}
