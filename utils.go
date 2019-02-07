package main

import "strings"

func getJSON(solutions int) (res []byte) {
	var builder strings.Builder

	solutionLen := len(getSolution())
	totalBytes := solutionLen * solutions
	// fmt.Printf("solution length %d", solutionLen)
	// fmt.Printf("growing builder to %d", totalBytes)
	builder.Grow(totalBytes)
	builder.WriteString("[")
	for i := 0; i < solutions; i++ {
		builder.WriteString(getSolution())
		if i < solutions-1 {
			builder.WriteString(",")
		}
	}
	builder.WriteString("]")

	res = []byte(builder.String())
	return
}

func getSolution() string {

	return `
	  {
		"_id": "5c5b8c8ce734ee28a2ba2157",
		"index": 0,
		"guid": "47732255-82cd-4789-b3c5-ecba645e9db9",
		"isActive": false,
		"balance": "$2,556.68",
		"picture": "http://placehold.it/32x32",
		"age": 39,
		"eyeColor": "blue",
		"name": "Rasmussen Rios",
		"gender": "male",
		"company": "MEDIFAX",
		"email": "rasmussenrios@medifax.com",
		"phone": "+1 (880) 547-3952",
		"address": "566 Troutman Street, Falconaire, Vermont, 1264",
		"about": "Minim ipsum do non proident est mollit dolore Lorem aliquip proident. Sint ad ad id proident nulla duis incididunt commodo nostrud. Adipisicing aute excepteur adipisicing nostrud deserunt anim cillum commodo anim labore culpa dolor. Elit voluptate proident nostrud nostrud consectetur quis enim id adipisicing. Aliqua dolore nisi velit deserunt commodo non occaecat excepteur adipisicing pariatur labore sit ad.\r\n",
		"registered": "2018-07-21T02:21:35 +04:00",
		"latitude": -70.998467,
		"longitude": 84.291804,
		"tags": [
		  "cupidatat",
		  "aliqua",
		  "pariatur",
		  "excepteur",
		  "consectetur",
		  "quis",
		  "sit"
		],
		"friends": [
		  {
			"id": 0,
			"name": "Letha Mcintyre",
			"parents": [
			  {
				"phone": 0,
				"name": "Jeanne Ray"
			  },
			  {
				"phone": 1,
				"name": "Antoinette Wilson"
			  }
			]
		  },
		  {
			"id": 1,
			"name": "Alexandria Guy",
			"parents": [
			  {
				"phone": 0,
				"name": "Ellis Galloway"
			  },
			  {
				"phone": 1,
				"name": "Watson Fowler"
			  }
			]
		  },
		  {
			"id": 2,
			"name": "Larsen Hicks",
			"parents": [
			  {
				"phone": 0,
				"name": "Teresa Dickson"
			  },
			  {
				"phone": 1,
				"name": "Suzanne Dorsey"
			  }
			]
		  }
		],
		"greeting": "Hello, Rasmussen Rios! You have 8 unread messages.",
		"favoriteFruit": "strawberry"
	  }
	`

}
