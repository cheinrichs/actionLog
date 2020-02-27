# actionlog

actionlog is a package created to demonstrate my abilities. It allows a user to track activities, their timing, and see the average time of each activity tracked.

# Installation

Installation is done with go get 

    go get -u github.com/cheinrichs/actionlog

# Features
These are the functions that are supported by the actionlog package

### AddAction

This function accepts a json serialized string of the form below and maintains an average time for each action

### GetStats

This function returns a serialized json array of the average time for each action that has been provided to the addAction function

# Testing

Unit tests are included with the actionlog package. You can run them by navigating to the downloaded file and running

    go test