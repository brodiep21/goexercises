package structs

import "fmt"

type vehicle struct {
	doors int
	color string
}
type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func Structer() {

	var m = make(map[string][]string)

	m["Bond, James"] = []string{"Shaken, not stirred", "Martini's", "Women"}
	m["Moneypenny, Miss"] = []string{"James Bond", "literature", "Computer Science"}
	delete(m, "Bond")
	for k, v := range m {
		fmt.Printf("%s\n", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
	chevy := truck{
		vehicle: vehicle{
			doors: 3,
			color: "blue",
		},
		fourWheel: true,
	}
	pontiac := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "orange",
		},
		luxury: false,
	}

	fmt.Println(pontiac.color, pontiac.doors, chevy.color)

	programmer := struct {
		info   map[int][]string
		isgood bool
		isold  bool
	}{
		info: map[int][]string{
			21: {"Golang", "HTML", "JavaScript"},
		},
		isgood: false,
		isold:  false,
	}
	fmt.Println(programmer.info)

}
