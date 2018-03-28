# otame_swaggo

try to [Swaggo](https://github.com/swaggo)

## Usage
1. Download swag
```
go get -u github.com/swaggo/swag/cmd/swag
```
2. Run tht swag in project root folder which contains main.go file, The swag will parse your comments and generate required files(docs folder and docs/doc.go).
```
swag init
```
3. copy generated docs/swagger/swagger.yaml and paste it to [swagger online editor](https://editor.swagger.io/#)
