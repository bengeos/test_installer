## Compile for Linux
env GOOS=linux go build

## Compile for Mac
env GOOS=darmin go build

## Compile for Windows
env GOOS=windows GOARCH=amd64 go build


## Compile for Linux on Specific File
env GOOS=linux go build -o installer-linux

## Compile for Mac
env GOOS=darmin go build -o installer-mac