<p align="center">
<img alt="Logo Banner" src="https://repository-images.githubusercontent.com/156623891/051894fb-7427-4546-994e-17aed8df500d"/>
</p>

# JSON Scrubber
This library is for parsing a struct before returning it in JSON format and scrubbing it dynamically of certain fields. It supports unlimited fields and can filter them down.

## Installation

```
go get github.com/Fyb3roptik/go-json-scrubber
```

## Usage

```go
jsonscrubber.AddOnly(*interface{}, ...string)
jsonscrubber.RemoveOnly(*interface{}, ...string)
```

## Example

```go
type User struct {
	FirstName         string      `json:"first_name"`
	LastName          string      `json:"last_name"`
	Address           *Address    `json:"address"`
}

type Address struct {
	Address            string    `json:"address"`
	City               string    `json:"city"`
	State              string    `json:"state"`
	Zip                string    `json:"zip"`
}

address := &Address{City: "New York", State: "NY"}
u := &User{FirstName: "Foo", LastName: "Bar", address}

// Initial model fields go here
user := jsonscrubber.AddOnly(u, "first_name", "address").(map[string]interface{})

// You can also do sub structs fields too, this time specifying fields to remove
user["address"] = jsonscrubber.RemoveOnly(u.Address, "address", "state", "zip")

// Return the JSON
b, err := json.MarshalIndent(user)
if err != nil {
		panic(err.Error())
	}
fmt.Print(string(b))

// Return JSON in Revel Framework
return c.RenderJSON(user)
```

## Output
```json
{
	"address": {
		"city": "New York"
	},
	"first_name": "Foo"
}
```
