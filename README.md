# Csv2Json
This project is a tool for taking multiple csv files and converting them into a single json for ingestion elsewhere.

This project was completely generated using **ChatGPT** , please verify source code before including in other works.

## License
This project comes with the Unlicense. Please feel free to fork, open a PR, or clone and change at will :)


## Customize
You will need to clone and build/run from your projects root dir

```
cd /home/$USER/projects/csv2json
go run .
```
or
`go build .`
### Modifications
Change this line to fit your needs

`outputLocation := "./bunchofstuff.json"`

and these one too

```
filepath := "~/Desktop"
for _, file := range []string{filepath + "Bunchofstuff.csv", "Bunchofstuff2.csv", "Bunchofstuff3.csv"}
```

