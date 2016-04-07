// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/mattn/go-drainclose"
)

type test_struct struct {
	Test string
}

func plurkPost(w http.ResponseWriter, req *http.Request) {
	defer drainclose.Close(req.Body)
	decoder := json.NewDecoder(req.Body)
	var t test_struct
	err := decoder.Decode(&t)
	if err != nil {
		log.Panic(err)
	}
	log.Println(t.Test)
}

func serveHttpAPI(port string, existC chan bool) {
	go func() {
		if err, ok := <-existC; ok {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/", plurkPost)
	http.ListenAndServe(":"+port, mux)
}
