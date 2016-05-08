# mongo-progess
An application for displaying the progress of currently running index creation process for mongo.

Can be downloaded from [here](https://github.com/plasma147/mongo-progess/raw/0.0.1-release/mongo-progress) 

###Demo: 

[![asciicast](https://asciinema.org/a/byaxgodr2ho4q0ebnga1uorbx.png)](https://asciinema.org/a/byaxgodr2ho4q0ebnga1uorbx)

###Building/Running

Build using `go install`. Assuming `$GOPATH/bin` is on the path, run by executing `mongo-progress`. 

Once the application is running it will either connect to an existing index creation process or wait until it detects a new index being applied. 
It will then display the progress of process. The application will run indefinitely so requires Ctrl-C to terminate it. 

