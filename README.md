sconfig
===============

### setup
```
go get -u github.com/marshalys/sconfig
```

### usage

```
import "github.com/marshalys/sconfig"

...

// init
config := sconfig.New()

// reader config file
err := config.LoadConfig(configPath)

// get config item
val, ok = config.Get("context.timeoutMilliseconds")

// get config item by struct
type cacheSetting struct {
	DefaultExpirationSeconds uint32
	CleanupIntervalSeconds   uint32
}
cache := &cacheSetting{}
err := config.UnmarshalKey("cache.memory", cache)
```

other usage, please see test code.