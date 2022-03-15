package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/jroimartin/gocui"
)

func layout(g *gocui.Gui) error {
	var views = []string{outputView, inputView}
	maxX, maxY := g.Size()
	for _, view := range views {
		x0, y0, x1, y1 := viewPositions[view].getCoordinates(maxX, maxY)
		//logger.Println(x0, y0, x1, y1)
		if v, err := g.SetView(view, x0, y0, x1, y1); err != nil {
			logger.Println(err)
			v.SelFgColor = gocui.ColorBlack
			v.SelBgColor = gocui.ColorGreen

			v.Title = " " + view + " "

			if view == inputView {
				v.Editable = true
				v.Wrap = true
			}

			if err != gocui.ErrUnknownView {
				return err
			}
		}
	}

	_, err := g.SetCurrentView(inputView)
	if err != nil {
		log.Fatal("failed to set current view: ", err)
	}
	return nil
}

func resetInput(g *gocui.Gui, v *gocui.View) error {
	m := map[string]interface{}{
		"from_user":    "a",
		"to_user":      "user1",
		"content_type": 1,
		"content":      "hello",
	}
	b, err := json.Marshal(&m)
	if err != nil {
		logger.Println(err)
		return err
	}

	buf := bytes.NewReader(b)
	size := buf.Size()

	req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:8071/gateway/service/v1/send_msg", buf)
	if err != nil {
		logger.Println(err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Length", strconv.FormatInt(size, 10))
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Println(err)
		return err
	}
	logger.Println(rsp.StatusCode)
	_ = rsp.Body.Close()
	v.Clear()
	v.SetCursor(v.Origin())
	g.Update(func(gg *gocui.Gui) error {
		v, err1 := gg.View("output")
		if err1 != nil {
			logger.Println("update err:", err1)
			return err1
		}
		fmt.Fprintln(v, "------")
		fmt.Fprintf(v, "Send|From:%v|Tp:%v|Content:%v\n", m["from_user"], m["content_type"], m["content"])
		return nil
	})
	return nil
}
