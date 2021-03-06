/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/logitick/todo-api/pkg/adding"
	"github.com/logitick/todo-api/pkg/http/jsonapi"
	"github.com/logitick/todo-api/pkg/listing"
	"github.com/logitick/todo-api/pkg/storage/memory"
	"github.com/logitick/todo-api/pkg/updating"
)

func main() {
	store := new(memory.Storage)
	ls := listing.NewService(store)
	as := adding.NewService(store)
	us := updating.NewService(store)
	router := jsonapi.Handler(ls, as, us)
	fmt.Println("serving: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
