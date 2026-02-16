package hello_modules2

import (
	"fmt"
	"math/rand/v2"
)

func Hello2(name string) string {
	message := fmt.Sprintf("Hello mr. %v from function hello2", name)
	return message
}

func CoolBugFacts() {
	want := false
	scarabFacts := []string{"The glorious scarab beetle grows between 0.75 inches and 1 inch long, and it appears in many different colors.", "The larvae are called grups and are yellow or white and curved. The larvae develop in soil, often at the roots of grasses.", "Its habitat is high elevation juniper forests for its food of choice is the juniper leaf. Additionally, it can camouflage easily into the juniper trees!"}
	japBeetleFacts := []string{"Japanese beetles are particularly harmful to roses, soybeans, cherries, and other plants.", "The adult beetles damage plants by skeletonizing the foliage (i.e., consuming only the material between a leaf's veins)", "The first written evidence of the insect appearing within the United States was in 1916 in a nursery near Riverton, New Jersey."}
	borerFacts := []string{"The emerald ash borer are certainly the most harmful types of green beetles because they bore into the wood of ash trees as larvae.", "Females lay eggs in bark crevices on ash trees, and larvae feed underneath the bark of ash trees to emerge as adults in one to two years.", "Adult beetles are typically bright metallic green and about 8.5 mm (0.33 in) long and 1.6 mm (0.063 in) wide. Elytra are typically a darker green, but can have copper hues."}
	var op int
	for {
		if !want {
			fmt.Println(
				"Select bug you want to learn about:\n",
				"1. Glorious scarab beetle\n",
				"2. Japanese beetle\n",
				"3. Emerald ash borer\n",
				"4. Delete task\n",
				"5. ListAllTasks",
				"6. Exit",
			)
			fmt.Scanln(&op)
		}
		switch op {
		case 1:
			index := rand.IntN(len(scarabFacts))
			coolFact := scarabFacts[index]
			fmt.Println("Glorious scarab beetle fact: \n", coolFact, "Do you want to learn more about this bug? Press 1 if you do!")
			fmt.Scanln(&op)
			if op == 1 {
				want = true
			} else {
				want = false
			}
		case 2:
			index := rand.IntN(len(japBeetleFacts))
			coolFact := japBeetleFacts[index]
			fmt.Println("Glorious scarab beetle fact: \n", coolFact, "Do you want to learn more about this bug? Press 2 if you do!")
			fmt.Scanln(&op)
			if op == 2 {
				want = true
			} else {
				want = false
			}
		case 3:
			index := rand.IntN(len(borerFacts))
			coolFact := borerFacts[index]
			fmt.Println("Glorious scarab beetle fact: \n", coolFact, "Do you want to learn more about this bug? Press 3 if you do!")
			fmt.Scanln(&op)
			if op == 3 {
				want = true
			} else {
				want = false
			}
		case 4:
			return
		}

	}

}
