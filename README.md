# mongo-progess
Display progress of currently running index creation process for mongo

###Demo: 

[![asciicast](https://asciinema.org/a/8miunghmesh7v4d1vwbyp95cc.png)](https://asciinema.org/a/8miunghmesh7v4d1vwbyp95cc)

###Building/Running

Build using `go install`. Assuming `$GOPATH/bin` is on the path, run by executing `mongo-progress`. 

Once the application is running it will wait until it detects an index being applied and will then display the progress of process.

