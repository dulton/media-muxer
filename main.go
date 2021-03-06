package main

import (
	"./media"

	"flag"
	"fmt"
	"net/url"
)

func main() {
	var driver, device string
	var display bool = false
	flag.StringVar(&driver, "driver", DRIVER, "set capture driver")
	flag.StringVar(&device, "device", DEVICE, "set capture device")
	flag.BoolVar(&display, "display", false, "enable sdl display")
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Println("no thing to mux")
		flag.Usage()
		return
	}
	uri, err := url.Parse(flag.Arg(0))
	if err != nil {
		fmt.Println(err)
		return
	}
	var format string
	if uri.Scheme == "rtsp" {
		format = "rtsp"
	}
	src := media.MediaSource{media.NewVideoSource(driver, device), &(media.AudioSource{})}
	m, err := media.NewMuxer(&src, format, flag.Arg(0))
	if err == nil {
		if display {
			m.EnableDisplay()
		}
		if m.Open() {
			m.WaitForDone()
		}
		m.Close()
	} else {
		fmt.Println(err)
	}
}

func init() {
	// C.av_log_set_level(C.AV_LOG_DEBUG)
}
