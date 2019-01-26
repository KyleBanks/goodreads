# goodreads

An unofficial [Goodreads API](https://www.goodreads.com/api/index) client written in Go. 

## Usage

The first thing you'll need to do is register for a Goodreads API Key: [https://www.goodreads.com/api/keys](https://www.goodreads.com/api/keys)

Once you have your key, you can initialize a Goodreads client like so:

```
package main

import (
    "os"
    "github.com/KyleBanks/goodreads"
)

func main() {
    key := os.GetEnv("API_KEY")	
    c := goodreads.NewClient(key)
}
```

With a client initialized, simply call the API methods as needed:

```
u, err := c.UserShow("38763538")
if err != nil {
    panic(err)
}
fmt.Printf("Loaded user details of %s:\n", u.Name)
```

The client function names match those of the Goodreads API documentation. For example, `user.show` is `UserShow` above. To see the full list of supported methods:

```
$ go doc github.com/KyleBanks/goodreads Client 
```

## Examples

Example code is available in the [examples/](./examples) directory.

After you've obtained a Goodreads API Key, you can run the examples like so:

```
$ API_KEY="api key" go run example/example.go
```
