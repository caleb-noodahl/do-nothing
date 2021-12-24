# com/caleb-noodahl/do-nothing/examples/getting-started
## step 1

git pull
pull the code back from github

```bash
git clone https://github.com/caleb-noodahl/do-nothing.git
```

## step 2

dependency restore
download the dependencies specified in the go.mod file

```bash
go mod download
```

## step 3

build
make a build creating a do-nothing.exe file

```bash
go build .
```

