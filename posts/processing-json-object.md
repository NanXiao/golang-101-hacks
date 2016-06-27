# Processing JSON object
----
[JSON](http://www.json.org/) is a commonly used and powerful data-interchange format, and `Go` provides a built-in [json](https://golang.org/pkg/encoding/json/) package to handle it. Let' see the following example:  

    package main
    
    import (
    	"encoding/json"
    	"log"
    	"fmt"
    )
    
    type People struct {
    	Name string
    	age int
    	Career string `json:"career"`
    	Married bool `json:",omitempty"`
    }
    
    func main()  {
    	p := &People{
    		Name: "Nan",
    		age: 34,
    		Career: "Engineer",
    	}
    
    	data, err := json.Marshal(p)
    	if err != nil {
    		log.Fatalf("JSON marshaling failed: %s", err)
    	}
    	fmt.Printf("%s\n", data)
    }
And the execution result is shown as follows:  

	{"Name":"Nan","career":"Engineer"}
The [Marshal](https://golang.org/pkg/encoding/json/#Marshal) function is used to serialize an interface into a `JSON` object. In our example, it encodes a `People` struct:  

(1) The `Name` member is encoded as our expectation:  

	"Name":"Nan"

(2) Where is the `age` field? We can't find it in our result. The cause is only exported members of struct can be marshaled, so that means only the name whose first letter capitalized can be encoded into `JSON` object (In our example, you should use `Age` instead of `age`).  

(3) The name of `Career` field is `career`, not `Career`:  

	"career":"Engineer"
That's because the following tag: `json:"career"`, which tells the `Marshal` function to use `career` in the `JSON` object.  

(4) We also can't see `Married` in the result although it has been exported, the magic behind is the `json:",omitempty"` tag which tells `Marshal` function no need to encode this member if it uses the default value.  

There is another [Unmarshal](https://golang.org/pkg/encoding/json/#Unmarshal) function which is used to parse a `JSON` object. See the following example which extends from the above one:  

    package main
    
    import (
    	"encoding/json"
    	"log"
    	"fmt"
    )
    
    type People struct {
    	Name string
    	age int
    	Career string `json:"career"`
    	Married bool `json:",omitempty"`
    }
    
    func main()  {
    	var p People
    	data, err := json.Marshal(&People{Name: "Nan", age: 34, Career: "Engineer", Married: true})
    
    	if err != nil {
    		log.Fatalf("JSON marshaling failed: %s", err)
    	}
    
    	err = json.Unmarshal(data, &p)
    	if err != nil {
    		log.Fatalf("JSON unmarshaling failed: %s", err)
    	}
    
    	fmt.Println(p)
    }

The running result is like this:  

	{Nan 0 Engineer true}
We can see the `JSON` object is decoded successfully.  

Besides `Marshal` and `Unmarshal` functions, the `json` package also provides [Encoder](https://golang.org/pkg/encoding/json/#Encoder) and [Decoder](https://golang.org/pkg/encoding/json/#Decoder) structs which are used to process `JSON` object from stream. E.g., It is not uncommon to see code which handle `HTTP` likes this:  

	func postFunc(w http.ResponseWriter, r *http.Request) {
		......
	
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		......
	}

Because the mechanism of both methods are similar, it is not necessary to overtalk `Encoder` and `Decoder` here.  

References:  
[Package json](https://golang.org/pkg/encoding/json/);  
[The Go Programming Language](http://www.gopl.io/).