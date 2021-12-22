package main

import (
	"os"
	"os/exec"
)

func main() {

	subs := make([]*exec.Cmd, 5)
	for i := 0; i < 5; i++ {
		c := exec.Command("ping", "8.8.8.8")
		c.Stdout = os.Stdout
		c.Stderr = os.Stdout
		subs[i] = c
		c.Start()
	}

	for _, sub := range subs {
		sub.Wait()
	}

	/*
		for i := 1; i < 5; i++ {
			dir := fmt.Sprintf("./%v", i)
			c := exec.Command("go", "run", dir+"/main.go")
			c.Stdout = os.Stdout
			c.Stderr = os.Stdout
			fmt.Println(c.Run())

			fmt.Println(c.Path, c.Args, c.Dir)

		} */
}
