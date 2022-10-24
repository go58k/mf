# go-mongo-key-escape

[Escape MongoDB keys](http://docs.mongodb.org/manual/faq/developers/#faq-dollar-sign-escaping) for characters . , $ , & , + ,  * , ? , ] , [ , ) , ( , | , . , } and { .

## Example

```go
package main

import "github.com/go58k/mf"

func main() {
    mongo.Escape("event.thing")
    // event\uFF0Ething
    mongo.Unescape("event\uFF0Ething")
    // event.thing

    mongo.Escape("event$thing")
    // event\uFF04thing
    mongo.Unescape("event\uFF04thing")
    // event$thing
}
```

### Thanks 

https://github.com/segment-boneyard/go-mongo-key-escape

## License

MIT
