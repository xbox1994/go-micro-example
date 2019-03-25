# qson

Convert URL query strings into JSON so that you can more easily parse them into structs/maps in Go.

I wrote this to help someone in the Gopher Slack, so it isn't really 100% complete but it should be stable enough to use in most production environments and work well as a starting point if you need something more custom.

If you end up using the package, feel free to submit any bugs and feature requests and I'll try to get to those updated time permitting.

## Usage

You can either turn a URL query param into a JSON byte array, or unmarshal that directly into a Go object.

Transforming the URL query param into a JSON byte array:

```go
import "github.com/joncalhoun/qson"

func main() {
  b, err := qson.ToJSON("bar%5Bone%5D%5Btwo%5D=2&bar[one][red]=112")
  if err != nil {
    panic(err)
  }
  fmt.Println(string(b))
  // Should output: {"bar":{"one":{"red":112,"two":2}}}
}
```

Or unmarshalling directly into a Go object using JSON struct tags:

```go
import "github.com/joncalhoun/qson"

type unmarshalT struct {
	A string     `json:"a"`
	B unmarshalB `json:"b"`
}
type unmarshalB struct {
	C int `json:"c"`
}

func main() {
  var out unmarshalT
  query := "a=xyz&b[c]=456"
  err := Unmarshal(&out, query)
  if err != nil {
  	t.Error(err)
  }
  // out should equal
  //   unmarshalT{
	// 	  A: "xyz",
	// 	  B: unmarshalB{
	// 	  	C: 456,
	// 	  },
	//   }
}
```

To get a query string like in the two previous examples you can use the `RawQuery` field on the [net/url.URL](https://golang.org/pkg/net/url/#URL) type.
