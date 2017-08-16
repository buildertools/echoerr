/*
Copyright 2017 Jeff Nickoloff "jeff@allingeek.com"

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
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Print a message to STDOUT and exit with the specified code.\n\nUSAGE:\n\techoerr [exitcode] <echo text...>\n")
		os.Exit(1)
	}
	if s, err := strconv.Atoi(os.Args[1]); err == nil {
		fmt.Println(strings.Join(os.Args[2:], ` `))
		os.Exit(s)
	}
	fmt.Println(strings.Join(os.Args[1:], ` `))
	os.Exit(1)
}
