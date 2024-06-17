package main

type Calculator struct {
	result int
}

func (c *Calculator) add(a, b int) {
	c.result = a + b
}
