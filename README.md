# actionlog

actionlog is a package created as an answer to a quick coding challenge. It allows a user to track activities, their timing, and see the average time of each activity.

# Installation

Installation is done with go get 

    go get -u github.com/cheinrichs/actionlog

# Features
These are the functions that are supported by the actionlog package

### AddAction

This function accepts a json serialized string of the form below and maintains an average time for each action

    actionlog.AddAction("{\"action\":\"jump\", \"time\":100}")

### GetStats

This function returns a serialized json array of the average time for each action that has been provided to the addAction function
    
    actionlog.GetStats()

# Example
To see an example use, navigate to the example folder in this directory and run 

    go build

and then call the executable with any valid json actions like so:

    ./example "{\"action\":\"dance\", \"time\":1600}" "{\"action\":\"fly\", \"time\":180}" "{\"action\":\"dance\", \"time\":250}"

# Testing

Unit tests are included with the actionlog package. Once you have the package installed, you can run them by navigating to the downloaded file and running:

    go test