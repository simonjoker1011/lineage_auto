package utils

import (
	"github.com/go-vgo/robotgo"
	"log"
)

func MS_init_listener() {

	go monitor_l()
	//	TODO: concurrent event
	//	go monitor_r()
	go listen_mouse_movement()
}

func monitor_l() {
	for {
		l_eve := robotgo.AddEvent("mleft")
		if l_eve == 0 {
			log.Println("mleft")
			continue
		}
	}
}
func monitor_r() {
	for {
		r_eve := robotgo.AddEvent("mright")
		if r_eve == 0 {
			log.Println("mright")
			continue
		}
	}
}

func listen_mouse_movement() {
	x0, y0 := robotgo.GetMousePos()
	for {
		x, y := robotgo.GetMousePos()
		if x0 != x || y0 != y {
			log.Println(x0, y0)
			x0, y0 = x, y
		}
	}
}
