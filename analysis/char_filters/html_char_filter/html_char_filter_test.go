//  Copyright (c) 2014 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.
package html_char_filter

import (
	"reflect"
	"testing"
)

func TestHtmlCharFilter(t *testing.T) {
	tests := []struct {
		input  []byte
		output []byte
	}{
		{
			input: []byte(`<!DOCTYPE html>
<html>
<body>

<h1>My First Heading</h1>

<p>My first paragraph.</p>

</body>
</html>`),
			output: []byte(`               
      
      

    My First Heading     

   My first paragraph.    

       
       `),
		},
	}

	for _, test := range tests {
		filter := NewHtmlCharFilter()
		output := filter.Filter(test.input)
		if !reflect.DeepEqual(output, test.output) {
			t.Errorf("Expected:\n`%s`\ngot:\n`%s`\nfor:\n`%s`\n", string(test.output), string(output), string(test.input))
		}
	}
}
