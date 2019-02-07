package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"code.cloudfoundry.org/bytefmt"
	auto "github.com/JordyBaylac/go-gen/samples"
)

func main() {

	// var results auto.ResultList

	// temp := auto.Solution{
	// 	Index: 2,
	// }

	// results = append(results, temp)

	// parsedResults, _ := parseExample()

	// serializeExample(parsedResults)

	results := make(chan string)
	var wg sync.WaitGroup
	wg.Add(7)

	go func() {
		defer wg.Done()
		result := measureUnmarshalTime(100)
		results <- result
	}()

	go func() {
		defer wg.Done()
		result := measureUnmarshalTime(1000)
		results <- result
	}()

	go func() {
		defer wg.Done()
		result := measureUnmarshalTime(10000)
		results <- result
	}()

	go func() {
		defer wg.Done()
		result := measureUnmarshalTime(100000)
		results <- result
	}()

	go func() {
		defer wg.Done()
		result := measureUnmarshalTime(200000)
		results <- result
	}()

	go func() {
		defer wg.Done()
		result := measureUnmarshalTime(230000)
		results <- result
	}()

	go func() {
		defer wg.Done()
		result := measureUnmarshalTime(270000)
		results <- result
	}()

	go func() {
		for result := range results {
			fmt.Println(result)
		}
	}()

	wg.Wait()
}

func measureUnmarshalTime(numberOfSolutions int) (result string) {
	solutionLen := len(getSolution())
	totalBytes := solutionLen * numberOfSolutions
	size := bytefmt.ByteSize(uint64(totalBytes))

	startGettingData := time.Now()
	jsonData := getJSON(numberOfSolutions)
	elapsedGettingData := time.Since(startGettingData)
	var results auto.ResultList

	start := time.Now()
	if err := json.Unmarshal(jsonData, &results); err != nil {
		fmt.Printf("Unexpected Error: %v", err)
	}
	elapsed := time.Since(start)

	result = fmt.Sprintf("Solutions %d.\n\tSize of %s bytes.\n\tGetting data (time taken) %f seconds.\n\tUnmarshall (time taken) %f seconds",
		numberOfSolutions, size, elapsedGettingData.Seconds(), elapsed.Seconds())

	for range results {
	}

	return
}

func serializeExample(results auto.ResultList) {

	jsonBytes, err := json.Marshal(results)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(jsonBytes))

}

func parseExample() (results auto.ResultList, err error) {

	data := []byte(`
	[
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
		},
		{
		  "_id": "5c5b8c8c320a519df96efca5",
		  "index": 1,
		  "guid": "4b383e02-1609-4a59-890b-e80b0c33951f",
		  "isActive": true,
		  "balance": "$3,502.23",
		  "picture": "http://placehold.it/32x32",
		  "age": 23,
		  "eyeColor": "brown",
		  "name": "Harding Head",
		  "gender": "male",
		  "company": "ZAPPIX",
		  "email": "hardinghead@zappix.com",
		  "phone": "+1 (828) 453-4000",
		  "address": "805 Lexington Avenue, Saddlebrooke, Indiana, 154",
		  "about": "Lorem magna excepteur sit pariatur tempor dolore eiusmod aliquip eiusmod non aute deserunt. Ut cupidatat ad non non esse duis do ipsum in. Culpa duis irure non occaecat dolor laborum deserunt. Laboris pariatur enim nisi est consectetur commodo veniam ullamco incididunt. Excepteur velit esse ut voluptate officia id pariatur reprehenderit quis fugiat sit aliqua. Aute duis ullamco est voluptate anim proident. Proident nostrud sint adipisicing incididunt nostrud labore.\r\n",
		  "registered": "2016-08-08T09:53:27 +04:00",
		  "latitude": -37.721887,
		  "longitude": 67.115351,
		  "tags": [
			"in",
			"id",
			"excepteur",
			"deserunt",
			"et",
			"minim",
			"ea"
		  ],
		  "friends": [
			{
			  "id": 0,
			  "name": "Bradley Puckett",
			  "parents": [
				{
				  "phone": 0,
				  "name": "Keisha Blackburn"
				},
				{
				  "phone": 1,
				  "name": "Kirsten Savage"
				}
			  ]
			},
			{
			  "id": 1,
			  "name": "Jacquelyn Cleveland",
			  "parents": [
				{
				  "phone": 0,
				  "name": "Ines Hamilton"
				},
				{
				  "phone": 1,
				  "name": "Cobb Tyson"
				}
			  ]
			},
			{
			  "id": 2,
			  "name": "April Johnston",
			  "parents": [
				{
				  "phone": 0,
				  "name": "Patrice Wong"
				},
				{
				  "phone": 1,
				  "name": "Veronica Sherman"
				}
			  ]
			}
		  ],
		  "greeting": "Hello, Harding Head! You have 5 unread messages.",
		  "favoriteFruit": "strawberry"
		},
		{
		  "_id": "5c5b8c8cd53c71244ce48237",
		  "index": 2,
		  "guid": "94c35b58-9cf3-4bb4-89d6-d8ede83cce91",
		  "isActive": false,
		  "balance": "$3,303.95",
		  "picture": "http://placehold.it/32x32",
		  "age": 33,
		  "eyeColor": "blue",
		  "name": "Becker Garrett",
		  "gender": "male",
		  "company": "MARQET",
		  "email": "beckergarrett@marqet.com",
		  "phone": "+1 (874) 558-3466",
		  "address": "275 Junius Street, Lydia, Alabama, 2239",
		  "about": "Ullamco ullamco ad nulla laborum nostrud elit mollit nostrud enim ad. Dolor voluptate sunt commodo aliquip deserunt aliquip aliquip consectetur Lorem esse consectetur reprehenderit ullamco non. Et irure enim id occaecat esse exercitation eiusmod sunt irure exercitation deserunt est aliqua. Minim elit anim id proident esse amet consequat ipsum enim fugiat.\r\n",
		  "registered": "2014-06-18T07:47:46 +04:00",
		  "latitude": 82.295091,
		  "longitude": 1.816057,
		  "tags": [
			"sit",
			"voluptate",
			"nulla",
			"id",
			"minim",
			"ex",
			"quis"
		  ],
		  "friends": [
			{
			  "id": 0,
			  "name": "Marcia Daugherty",
			  "parents": [
				{
				  "phone": 0,
				  "name": "Claudette Ellison"
				},
				{
				  "phone": 1,
				  "name": "Erickson Riggs"
				}
			  ]
			},
			{
			  "id": 1,
			  "name": "Parks Faulkner",
			  "parents": [
				{
				  "phone": 0,
				  "name": "Iva Terry"
				},
				{
				  "phone": 1,
				  "name": "Francisca Potter"
				}
			  ]
			},
			{
			  "id": 2,
			  "name": "Sarah Moreno",
			  "parents": [
				{
				  "phone": 0,
				  "name": "Jeannie Mcmillan"
				},
				{
				  "phone": 1,
				  "name": "Allen Gordon"
				}
			  ]
			}
		  ],
		  "greeting": "Hello, Becker Garrett! You have 4 unread messages.",
		  "favoriteFruit": "apple"
		}
	  ]
	`)

	err = json.Unmarshal(data, &results)
	return
}
