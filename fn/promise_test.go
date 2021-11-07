package fn

import (
	"fmt"
	"testing"
	"time"
)

func TestAll(t *testing.T) {
	var (
		example1 = []Promise{
			func() (i interface{}, e error) {
				<-time.After(time.Second)
				fmt.Println("all", 1)
				return
			},
			func() (i interface{}, e error) {
				fmt.Println("all", 2)
				return
			},
			func() (i interface{}, e error) {
				fmt.Println("all", 3)
				return
			},
			func() (i interface{}, e error) {
				<-time.After(time.Millisecond * 100)
				fmt.Println("all", 4)
				return
			},
			func() (i interface{}, e error) {
				<-time.After(time.Millisecond * 300)
				fmt.Println("all", 5)
				return
			},
			func() (i interface{}, e error) {
				fmt.Println("all", 6)
				return
			},
		}
	)

	_, e := All(example1)
	t.Error(e)

	_, e = All(example1, 2)
	t.Error(e)
}

func TestWorkFlow(t *testing.T) {
	var (
		example1 = []string{"qonyrl67tz", "7n8jua8m9f3", "f2nhl3hz10i", "oac3136dj1", "pnrgz9u0rc", "7p7xnld3416"}
	)

	_, e := WorkFlow(example1, 2, func(i interface{}, idx int) (v interface{}, e error) {
		fmt.Println(i, idx)
		return
	})

	t.Error(e)
}
