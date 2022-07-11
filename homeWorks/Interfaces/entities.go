package main

type dog struct {
	weight    int
	name      string
	edibility bool
}

func (d dog) feedPerMonth() int {
	const feedPerKg int = 2
	return feedPerKg * d.weight
}

func (d dog) weightOf() int {
	return d.weight
}

func (d dog) nameOf() string {
	return d.name
}

func (d dog) edibilityOf() bool {
	return d.edibility
}

func (d dog) typeOf() string {
	return "dog"
}

type cat struct {
	weight    int
	name      string
	edibility bool
}

func (c cat) feedPerMonth() int {
	const feedPerKg int = 7
	return feedPerKg * c.weight
}

func (c cat) weightOf() int {
	return c.weight
}

func (c cat) nameOf() string {
	return c.name
}

func (c cat) edibilityOf() bool {
	return c.edibility
}
func (c cat) typeOf() string {
	return "cat"
}

type cow struct {
	weight    int
	name      string
	edibility bool
}

func (c cow) feedPerMonth() int {
	const feedPerKg int = 25
	return feedPerKg * c.weight
}

func (c cow) weightOf() int {
	return c.weight
}

func (c cow) nameOf() string {
	return c.name
}

func (c cow) edibilityOf() bool {
	return c.edibility
}

func (c cow) typeOf() string {
	return "cow"
}
