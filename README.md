# checkspaces
checkspaces checks if there is a space before go:embed.

# Install

```
go get github.com/masibw/checkspaces/cmd/checkspaces
```

```
go install github.com/masibw/checkspaces/cmd/checkspaces@latest
```

# Usage
```
go vet -vettool=`which checkspaces` ./...
```

# Example
```
var (
    // go:embed testfile.txt
    fileInvalid []byte
)
```

Output
```
./main.go:9:2: There is a space between slash and go:embed
```
