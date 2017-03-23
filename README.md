# notification

Desktop notification using native api

## Usage
```go
package main

import "github.com/kasari/notification"

func main() {
	notification.Message("hello").WithTitle("I'm Title").Push()
}
```

## ToDo
- [x] Max OS X
- [ ] Linux
- [ ] Windows
- [ ] Set icon
