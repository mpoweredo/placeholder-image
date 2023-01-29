package handlers

import (
	"bytes"
	"github.com/fogleman/gg"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func GetImage(c *fiber.Ctx) error {
	r := c.Params("resolution")

	t := c.Query("text")

	parts := strings.Split(r, "x")

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

	fontSize := float64((width + height) / (width/15 + height/15))

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

	if err := dc.EncodePNG(buf); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Something went wrong while encoding PNG!"})
	}

	c.Attachment(text)
	c.Type("png")
	c.Send(buf.Bytes())

	return nil
}
