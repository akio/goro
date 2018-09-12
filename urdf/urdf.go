package urdf

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"
)

type Robot struct {
	Name   string  `xml:"name,attr"`
	Links  []Link  `xml:"link"`
	Joints []Joint `xml:"joint"`
}

func (r *Robot) FindLink(s string) int {
	for i, link := range r.Links {
		if link.Name == s {
			return i
		}
	}
	return -1
}

func (r *Robot) FindJoint(s string) int {
	for i, joint := range r.Joints {
		if joint.Name == s {
			return i
		}
	}
	return -1
}

type Link struct {
	Name       string      `xml:"name,attr"`
	Inertial   *Inertial   `xml:"inertial"`
	Visuals    []Visual    `xml:"visual"`
	Collisions []Collision `xml:"collision"`
}

type Axis struct {
	X float64
	Y float64
	Z float64
}

func (a *Axis) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type axis struct {
		Xyz *string `xml:"xyz,attr"`
	}

	var value axis
	var err error
	if err = d.DecodeElement(&value, &start); err != nil {
		return err
	}
	if value.Xyz == nil {
		a.X = 1.0
		a.Y = 0.0
		a.Z = 0.0
		return nil
	}

	delimiterPred := func(r rune) bool { return unicode.IsSpace(r) }
	elements := strings.FieldsFunc(*value.Xyz, delimiterPred)
	if len(elements) != 3 {
		return fmt.Errorf("axis[@xyz] must have 3 elements but %d.  (%v)", len(elements), elements)
	}
	a.X, err = strconv.ParseFloat(elements[0], 64)
	if err != nil {
		return fmt.Errorf("Invalid float literal")
	}
	a.Y, err = strconv.ParseFloat(elements[1], 64)
	if err != nil {
		return fmt.Errorf("Invalid float literal")
	}
	a.Z, err = strconv.ParseFloat(elements[2], 64)
	if err != nil {
		return fmt.Errorf("Invalid float literal")
	}
	return nil
}

type Origin struct {
	X     float64
	Y     float64
	Z     float64
	Roll  float64
	Pitch float64
	Yaw   float64
}

func (o *Origin) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type pose struct {
		Xyz *string `xml:"xyz,attr"`
		Rpy *string `xml:"rpy,attr"`
	}

	var value pose
	var err error
	if err = d.DecodeElement(&value, &start); err != nil {
		return err
	}
	delimiterPred := func(r rune) bool { return unicode.IsSpace(r) }
	if value.Xyz == nil {
		o.X = 0.0
		o.Y = 0.0
		o.Z = 0.0
	} else {
		elements := strings.FieldsFunc(*value.Xyz, delimiterPred)
		if len(elements) != 3 {
			return fmt.Errorf("origin[@xyz] must have 3 elements but %d.  (%v)", len(elements), elements)
		}
		o.X, err = strconv.ParseFloat(elements[0], 64)
		if err != nil {
			return fmt.Errorf("Invalid float literal")
		}
		o.Y, err = strconv.ParseFloat(elements[1], 64)
		if err != nil {
			return fmt.Errorf("Invalid float literal")
		}
		o.Z, err = strconv.ParseFloat(elements[2], 64)
		if err != nil {
			return fmt.Errorf("Invalid float literal")
		}
	}
	if value.Rpy == nil {
		o.Roll = 1.0
		o.Pitch = 0.0
		o.Yaw = 0.0
	} else {
		elements := strings.FieldsFunc(*value.Rpy, delimiterPred)
		if len(elements) != 3 {
			return fmt.Errorf("origin[@rpy] must have 3 elements but %d.  (%v)", len(elements), elements)
		}
		o.Roll, err = strconv.ParseFloat(elements[0], 64)
		if err != nil {
			return fmt.Errorf("Invalid float literal")
		}
		o.Pitch, err = strconv.ParseFloat(elements[1], 64)
		if err != nil {
			return fmt.Errorf("Invalid float literal")
		}
		o.Yaw, err = strconv.ParseFloat(elements[2], 64)
		if err != nil {
			return fmt.Errorf("Invalid float literal")
		}
	}
	return nil
}

type Mass struct {
	Value float64 `xml:"value,attr"`
}

type Inertia struct {
	Ixx float64 `xml:"ixx,attr"`
	Ixy float64 `xml:"ixy,attr"`
	Ixz float64 `xml:"ixz,attr"`
	Iyy float64 `xml:"iyy,attr"`
	Iyz float64 `xml:"iyz,attr"`
	Izz float64 `xml:"izz,attr"`
}

type Inertial struct {
	Origin  Origin  `xml:"origin"`
	Mass    Mass    `xml:"mass"`
	Inertia Inertia `xml:"inertia"`
}

type Visual struct {
	Name     *string   `xml:"name,attr"`
	Origin   Origin    `xml:"origin"`
	Geometry Geometry  `xml:"geometry"`
	Material *Material `xml:"material"`
}

type Geometry struct {
	Box      *Box      `xml:"box"`
	Cylinder *Cylinder `xml:"cylinder"`
	Sphere   *Sphere   `xml:"sphere"`
	Mesh     *Resource `xml:"mesh"`
}

type Box struct {
	Size Size `xml:"size,attr"`
}

type Size struct {
	X float64
	Y float64
	Z float64
}

func (s *Size) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type vec3 struct {
		Size *string `xml:"size,attr"`
	}

	var value vec3
	var err error
	if err = d.DecodeElement(&value, &start); err != nil {
		return err
	}
	if value.Size == nil {
		s.X = 1.0
		s.Y = 0.0
		s.Z = 0.0
		return nil
	}

	delimiterPred := func(r rune) bool { return unicode.IsSpace(r) }
	elements := strings.FieldsFunc(*value.Size, delimiterPred)
	if len(elements) != 3 {
		return fmt.Errorf("box[@size] must have 3 elements but %d.  (%v)", len(elements), elements)
	}
	s.X, err = strconv.ParseFloat(elements[0], 64)
	if err != nil {
		return fmt.Errorf("Invalid float literal")
	}
	s.Y, err = strconv.ParseFloat(elements[1], 64)
	if err != nil {
		return fmt.Errorf("Invalid float literal")
	}
	s.Z, err = strconv.ParseFloat(elements[2], 64)
	if err != nil {
		return fmt.Errorf("Invalid float literal")
	}
	return nil
}

type Cylinder struct {
	Radius float64 `xml:"radius,attr"`
	Length float64 `xml:"length,attr"`
}

type Sphere struct {
	Radius float64 `xml:"radius,attr"`
}

type Resource struct {
	Filename string `xml:"filename,attr"`
}

type Material struct {
	Name    string    `xml:"name,attr"`
	Color   *Color    `xml:"color"`
	Texture *Resource `xml:"texture"`
}

type Color struct {
	R float64
	G float64
	B float64
	A float64
}

func (c *Color) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type vec4 struct {
		Rgba *string `xml:"rgba,attr"`
	}

	var value vec4
	var err error
	if err = d.DecodeElement(&value, &start); err != nil {
		return err
	}
	if value.Rgba == nil {
		c.R = 1.0
		c.G = 1.0
		c.B = 1.0
		c.A = 1.0
		return nil
	}

	delimiterPred := func(r rune) bool { return unicode.IsSpace(r) }
	elements := strings.FieldsFunc(*value.Rgba, delimiterPred)
	if len(elements) != 4 {
		return fmt.Errorf("color[@rgba] must have 4 elements but %d.  (%v)", len(elements), elements)
	}
	c.R, err = strconv.ParseFloat(elements[0], 64)
	if err != nil {
		return fmt.Errorf("Invalid float literal")
	}
	c.G, err = strconv.ParseFloat(elements[1], 64)
	if err != nil {
		return fmt.Errorf("Invalid float literal")
	}
	c.B, err = strconv.ParseFloat(elements[2], 64)
	if err != nil {
		return fmt.Errorf("Invalid float literal")
	}
	c.A, err = strconv.ParseFloat(elements[3], 64)
	if err != nil {
		return fmt.Errorf("Invalid float literal")
	}
	return nil
}

type Collision struct {
	Name     *string  `xml:"name,attr"`
	Origin   Origin   `xml:"origin"`
	Geometry Geometry `xml:"geometry"`
}

type LinkRef struct {
	Link string `xml:"link,attr"`
}

type Calibration struct {
	Rising  *float64 `xml:"rising,attr"`
	Falling *float64 `xml:"falling,attr"`
}

type Dynamics struct {
	Damping  *float64 `xml:"damping,attr"`
	Friction *float64 `xml:"friction,attr"`
}

type Limit struct {
	Lower    *float64 `xml:"lower,attr"`
	Upper    *float64 `xml:"upper,attr"`
	Effort   float64  `xml:"effort,attr"`
	Velocity float64  `xml:"velocity,attr"`
}

type Mimic struct {
	Joint      string   `xml:"joint"`
	Multiplier *float64 `xml:"multiplier"`
	Offset     *float64 `xml:"offset"`
}

type SafetyController struct {
	SoftLowerLimit *float64 `xml:"soft_lower_limit"`
	SoftUpperLimit *float64 `xml:"soft_upper_limit"`
	KPosition      *float64 `xml:"k_position"`
	KVelocity      float64  `xml:"k_velocity"`
}

type Joint struct {
	Name             string            `xml:"name,attr"`
	Type             string            `xml:"type,attr"`
	Origin           Origin            `xml:"origin"`
	Parent           LinkRef           `xml:"parent"`
	Child            LinkRef           `xml:"child"`
	Axis             Axis              `xml:"axis"`
	Calibration      *Calibration      `xml:"calibration"`
	Dyanmics         *Dynamics         `xml:"dynamics"`
	Limit            *Limit            `xml:"limit"`
	Mimic            *Mimic            `xml:"mimic"`
	SafetyController *SafetyController `xml:"safety_controller"`
}

func LoadFromString(source string) (*Robot, error) {
	robot := new(Robot)
	if err := xml.Unmarshal([]byte(source), robot); err != nil {
		fmt.Println("XML error: ", err)
		return nil, err
	}
	return robot, nil
}

func LoadFromFile(filename string) (*Robot, error) {
	xmlData, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	robot := new(Robot)
	if err := xml.Unmarshal(xmlData, robot); err != nil {
		fmt.Println("XML error: ", err)
		return nil, err
	}
	return robot, nil
}

func loadParamFromString(s string) (interface{}, error) {
	decoder := json.NewDecoder(strings.NewReader(s))
	var value interface{}
	err := decoder.Decode(&value)
	if err != nil {
		return nil, err
	}
	return value, err
}
