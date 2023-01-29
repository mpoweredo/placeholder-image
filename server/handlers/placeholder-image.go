package handlers

import (
	"bytes"
	"github.com/fogleman/gg"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func GetImage(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)

	r_ := v["resolution"]
	t := r.URL.Query().Get("text")

	parts := strings.Split(r_, "x")

	var resolution [2]int
	for i, p := range parts {
		resolution[i], _ = strconv.Atoi(p)
	}

	width, height := resolution[0], resolution[1]

	if height == 0 {
		height = width
	}

	text := strings.Replace(t, "-", " ", -1)

	if text == "notext" {
		text = ""
	}

	if text == "" {
		text = strconv.Itoa(width) + "x" + strconv.Itoa(height)
	}

	fontf, _ := ioutil.ReadFile("./assets/poppins.ttf")

	font, err := freetype.ParseFont(fontf)
	if err != nil {
		log.Fatal(err)
	}

	fontSize := float64((width + height) / (width/22 + height/22))

	face := truetype.NewFace(font, &truetype.Options{Size: fontSize})

	dc := gg.NewContext(width, height)
	dc.SetFontFace(face)
	dc.SetHexColor("#2a2f3b")
	dc.DrawRectangle(0, 0, float64(width), float64(height))
	dc.Fill()
	dc.SetRGB(1, 1, 1)
	dc.DrawStringWrapped(
		text,
		float64(width/2), float64(height/2), 0.5, 0.5,
		300,
		1.5,
		gg.AlignCenter)

	buf := new(bytes.Buffer)
	dc.EncodePNG(buf)

	fileSize := buf.Len()

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Content-Disposition", "attachment; filename="+text)
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.FormatInt(int64(fileSize), 10))

	buf.WriteTo(w)

	return

}
