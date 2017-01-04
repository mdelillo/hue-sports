package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strings"
	"time"
	"errors"
	"os"
)

type Light struct {
	State LightState `json:"state"`
	Name  string `json:"name"`
}
type LightState struct {
	On        bool `json:"on"`
	Bri       int `json:"bri"`
	Hue       int `json:"hue"`
	Sat       int `json:"sat"`
	XY        []float32 `json:"xy"`
	CT        int `json:"ct"`
	ColorMode string `json:"colormode"`
}

func main() {
	if os.Getenv("HUE_USER_ID") == "" {
		fmt.Println("Must set HUE_USER_ID")
		os.Exit(1)
	}
	if os.Getenv("HUE_BRIDGE_IP") == "" {
		fmt.Println("Must set HUE_BRIDGE_IP")
		os.Exit(1)
	}

	originalLights, err := getLights()
	if err != nil {
		panic(err)
	}

	lights := make(map[string]string)
	for lightId, light := range originalLights {
		lights[light.Name] = lightId
	}

	blue := Light{
		State: LightState{
			On:             true,
			ColorMode:      "hs",
			Bri:            254,
			Sat:            254,
			Hue:            46920,
		},
	}
	red := Light{
		State: LightState{
			On:             true,
			ColorMode:      "hs",
			Bri:            254,
			Sat:            254,
			Hue:            65535,
		},
	}
	off := Light{
		State: LightState{
			On: false,
		},
	}

	// all blue <-> all red

	setLight(lights["Couch north"], blue)
	setLight(lights["Couch south"], blue)
	setLight(lights["Computer"], blue)
	setLight(lights["TV backlight"], blue)
	time.Sleep(time.Second)

	setLight(lights["Couch north"], red)
	setLight(lights["Couch south"], red)
	setLight(lights["Computer"], red)
	setLight(lights["TV backlight"], red)
	time.Sleep(time.Second)

	setLight(lights["Couch north"], blue)
	setLight(lights["Couch south"], blue)
	setLight(lights["Computer"], blue)
	setLight(lights["TV backlight"], blue)
	time.Sleep(time.Second)

	setLight(lights["Couch north"], red)
	setLight(lights["Couch south"], red)
	setLight(lights["Computer"], red)
	setLight(lights["TV backlight"], red)
	time.Sleep(time.Second)

	setLight(lights["Couch north"], blue)
	setLight(lights["Couch south"], blue)
	setLight(lights["Computer"], blue)
	setLight(lights["TV backlight"], blue)
	time.Sleep(time.Second)

	setLight(lights["Couch north"], red)
	setLight(lights["Couch south"], red)
	setLight(lights["Computer"], red)
	setLight(lights["TV backlight"], red)
	time.Sleep(time.Second)

	// Diagonal half blue half red

	setLight(lights["Couch north"], blue)
	setLight(lights["Couch south"], red)
	setLight(lights["Computer"], blue)
	setLight(lights["TV backlight"], red)
	time.Sleep(500 * time.Millisecond)

	setLight(lights["Couch north"], red)
	setLight(lights["Couch south"], blue)
	setLight(lights["Computer"], red)
	setLight(lights["TV backlight"], blue)
	time.Sleep(500 * time.Millisecond)

	setLight(lights["Couch north"], blue)
	setLight(lights["Couch south"], red)
	setLight(lights["Computer"], blue)
	setLight(lights["TV backlight"], red)
	time.Sleep(500 * time.Millisecond)

	setLight(lights["Couch north"], red)
	setLight(lights["Couch south"], blue)
	setLight(lights["Computer"], red)
	setLight(lights["TV backlight"], blue)
	time.Sleep(500 * time.Millisecond)

	setLight(lights["Couch north"], blue)
	setLight(lights["Couch south"], red)
	setLight(lights["Computer"], blue)
	setLight(lights["TV backlight"], red)
	time.Sleep(500 * time.Millisecond)

	setLight(lights["Couch north"], red)
	setLight(lights["Couch south"], blue)
	setLight(lights["Computer"], red)
	setLight(lights["TV backlight"], blue)
	time.Sleep(500 * time.Millisecond)

	setLight(lights["Couch north"], blue)
	setLight(lights["Couch south"], red)
	setLight(lights["Computer"], blue)
	setLight(lights["TV backlight"], red)
	time.Sleep(500 * time.Millisecond)

	setLight(lights["Couch north"], red)
	setLight(lights["Couch south"], blue)
	setLight(lights["Computer"], red)
	setLight(lights["TV backlight"], blue)
	time.Sleep(500 * time.Millisecond)

	// Clockwise blue

	setLight(lights["Couch north"], off)
	setLight(lights["Couch south"], off)
	setLight(lights["Computer"], off)
	setLight(lights["TV backlight"], off)
	time.Sleep(250 * time.Millisecond)

	setLight(lights["Couch north"], blue)
	setLight(lights["TV backlight"], off)
	time.Sleep(250 * time.Millisecond)
	setLight(lights["Couch south"], blue)
	setLight(lights["Couch north"], off)
	time.Sleep(250 * time.Millisecond)
	setLight(lights["Computer"], blue)
	setLight(lights["Couch south"], off)
	time.Sleep(250 * time.Millisecond)
	setLight(lights["TV backlight"], blue)
	setLight(lights["Computer"], off)
	time.Sleep(250 * time.Millisecond)

	setLight(lights["Couch north"], blue)
	setLight(lights["TV backlight"], off)
	time.Sleep(250 * time.Millisecond)
	setLight(lights["Couch south"], blue)
	setLight(lights["Couch north"], off)
	time.Sleep(250 * time.Millisecond)
	setLight(lights["Computer"], blue)
	setLight(lights["Couch south"], off)
	time.Sleep(250 * time.Millisecond)
	setLight(lights["TV backlight"], blue)
	setLight(lights["Computer"], off)
	time.Sleep(250 * time.Millisecond)

	setLight(lights["Couch north"], blue)
	setLight(lights["TV backlight"], off)
	time.Sleep(250 * time.Millisecond)
	setLight(lights["Couch south"], blue)
	setLight(lights["Couch north"], off)
	time.Sleep(250 * time.Millisecond)
	setLight(lights["Computer"], blue)
	setLight(lights["Couch south"], off)
	time.Sleep(250 * time.Millisecond)
	setLight(lights["TV backlight"], blue)
	setLight(lights["Computer"], off)
	time.Sleep(250 * time.Millisecond)

	// Counter-clockwise red

	setLight(lights["Couch north"], off)
	setLight(lights["Couch south"], off)
	setLight(lights["Computer"], off)
	setLight(lights["TV backlight"], off)
	time.Sleep(250 * time.Millisecond)

	setLight(lights["TV backlight"], red)
	setLight(lights["Couch north"], off)
	time.Sleep(250 * time.Millisecond)
	setLight(lights["Computer"], red)
	setLight(lights["TV backlight"], off)
	time.Sleep(250 * time.Millisecond)
	setLight(lights["Couch south"], red)
	setLight(lights["Computer"], off)
	time.Sleep(250 * time.Millisecond)
	setLight(lights["Couch north"], red)
	setLight(lights["Couch souch"], off)
	time.Sleep(250 * time.Millisecond)

	setLight(lights["TV backlight"], red)
	setLight(lights["Couch north"], off)
	time.Sleep(250 * time.Millisecond)
	setLight(lights["Computer"], red)
	setLight(lights["TV backlight"], off)
	time.Sleep(250 * time.Millisecond)
	setLight(lights["Couch south"], red)
	setLight(lights["Computer"], off)
	time.Sleep(250 * time.Millisecond)
	setLight(lights["Couch north"], red)
	setLight(lights["Couch souch"], off)
	time.Sleep(250 * time.Millisecond)

	setLight(lights["TV backlight"], red)
	setLight(lights["Couch north"], off)
	time.Sleep(250 * time.Millisecond)
	setLight(lights["Computer"], red)
	setLight(lights["TV backlight"], off)
	time.Sleep(250 * time.Millisecond)
	setLight(lights["Couch south"], red)
	setLight(lights["Computer"], off)
	time.Sleep(250 * time.Millisecond)
	setLight(lights["Couch north"], red)
	setLight(lights["Couch souch"], off)
	time.Sleep(250 * time.Millisecond)

	// Clockwise half blue half red

	setLight(lights["Couch north"], red)
	setLight(lights["Couch south"], red)
	setLight(lights["Computer"], blue)
	setLight(lights["TV backlight"], blue)
	time.Sleep(250 * time.Millisecond)
	setLight(lights["Couch north"], blue)
	//setLight(lights["Couch south"], red)
	setLight(lights["Computer"], red)
	//setLight(lights["TV backlight"], blue)
	time.Sleep(250 * time.Millisecond)
	//setLight(lights["Couch north"], blue)
	setLight(lights["Couch south"], blue)
	//setLight(lights["Computer"], red)
	setLight(lights["TV backlight"], red)
	time.Sleep(250 * time.Millisecond)
	setLight(lights["Couch north"], red)
	//setLight(lights["Couch south"], blue)
	setLight(lights["Computer"], blue)
	//setLight(lights["TV backlight"], red)
	time.Sleep(250 * time.Millisecond)

	//setLight(lights["Couch north"], red)
	setLight(lights["Couch south"], red)
	//setLight(lights["Computer"], blue)
	setLight(lights["TV backlight"], blue)
	time.Sleep(250 * time.Millisecond)
	setLight(lights["Couch north"], blue)
	//setLight(lights["Couch south"], red)
	setLight(lights["Computer"], red)
	//setLight(lights["TV backlight"], blue)
	time.Sleep(250 * time.Millisecond)
	//setLight(lights["Couch north"], blue)
	setLight(lights["Couch south"], blue)
	//setLight(lights["Computer"], red)
	setLight(lights["TV backlight"], red)
	time.Sleep(250 * time.Millisecond)
	setLight(lights["Couch north"], red)
	//setLight(lights["Couch south"], blue)
	setLight(lights["Computer"], blue)
	//setLight(lights["TV backlight"], red)
	time.Sleep(250 * time.Millisecond)

	//setLight(lights["Couch north"], red)
	setLight(lights["Couch south"], red)
	//setLight(lights["Computer"], blue)
	setLight(lights["TV backlight"], blue)
	time.Sleep(250 * time.Millisecond)
	setLight(lights["Couch north"], blue)
	//setLight(lights["Couch south"], red)
	setLight(lights["Computer"], red)
	//setLight(lights["TV backlight"], blue)
	time.Sleep(250 * time.Millisecond)
	//setLight(lights["Couch north"], blue)
	setLight(lights["Couch south"], blue)
	//setLight(lights["Computer"], red)
	setLight(lights["TV backlight"], red)
	time.Sleep(250 * time.Millisecond)
	setLight(lights["Couch north"], red)
	//setLight(lights["Couch south"], blue)
	setLight(lights["Computer"], blue)
	//setLight(lights["TV backlight"], red)
	time.Sleep(250 * time.Millisecond)

	//setLight(lights["Couch north"], red)
	setLight(lights["Couch south"], red)
	//setLight(lights["Computer"], blue)
	setLight(lights["TV backlight"], blue)
	time.Sleep(250 * time.Millisecond)
	setLight(lights["Couch north"], blue)
	//setLight(lights["Couch south"], red)
	setLight(lights["Computer"], red)
	//setLight(lights["TV backlight"], blue)
	time.Sleep(250 * time.Millisecond)
	//setLight(lights["Couch north"], blue)
	setLight(lights["Couch south"], blue)
	//setLight(lights["Computer"], red)
	setLight(lights["TV backlight"], red)
	time.Sleep(250 * time.Millisecond)
	setLight(lights["Couch north"], red)
	//setLight(lights["Couch south"], blue)
	setLight(lights["Computer"], blue)
	//setLight(lights["TV backlight"], red)
	time.Sleep(250 * time.Millisecond)

	time.Sleep(time.Second)

	setLights(originalLights)
}

func getLights() (map[string]Light, error) {
	resp, err := http.Get(fmt.Sprintf("http://%s/api/%s/lights", os.Getenv("HUE_BRIDGE_IP"), os.Getenv("HUE_USER_ID")))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var lights map[string]Light
	if err := json.Unmarshal(body, &lights); err != nil {
		return nil, err
	}

	return lights, nil
}

func setLights(lights map[string]Light) error {
	for lightId, light := range lights {
		if err := setLight(lightId, light); err != nil {
			return err
		}
	}
	return nil
}

func setLight(lightId string, light Light) error {
	var state string

	if !light.State.On {
		state = `{"on": false, "transitiontime": 1}`
	} else {
		switch light.State.ColorMode {
		case "ct":
			state = fmt.Sprintf(
				`{"on": true, "bri": %d, "ct": %d, "transitiontime": 1}`,
				light.State.Bri,
				light.State.CT,
			)
		case "hs":
			state = fmt.Sprintf(
				`{"on": true, "bri": %d, "hue": %d, "sat": %d, "transitiontime": 1}`,
				light.State.Bri,
				light.State.Hue,
				light.State.Sat,
			)
		case "xy":
			state = fmt.Sprintf(
				`{"on": true, "bri": %d, "xy": [%d, %d], "transitiontime": 1}`,
				light.State.Bri,
				light.State.XY[0],
				light.State.XY[1],
			)
		default:
			return errors.New("Unknown color mode: " + light.State.ColorMode)
		}
	}

	req, err := http.NewRequest(
		"PUT",
		fmt.Sprintf("http://%s/api/%s/lights/%s/state", os.Getenv("HUE_BRIDGE_IP"), os.Getenv("HUE_USER_ID"), lightId),
		strings.NewReader(state),
	)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if resp.StatusCode != 200 {
		return fmt.Errorf("%s: %s", resp.Status, resp.Body)
	}

	return nil
}
