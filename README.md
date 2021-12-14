# Auto reload gin
```
go get github.com/codegangsta/gin
gin --appPort 5000 --port 3000
```

# Log 
```
go get "github.com/Sirupsen/logrus"
```
### Log level
```
log.Trace("Something very low level.")

log.Debug("Useful debugging information.")

log.Info("Something noteworthy happened!")

log.Warn("You should probably take a look at this.")

log.Error("Something failed but I'm not quitting.")

// Calls os.Exit(1) after logging
log.Fatal("Bye.")

// Calls panic() after logging
log.Panic("I'm bailing.")
```