# checkspaces
checkspaces checks if there is a space between // and directives.

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
./main.go:9:2: There is a space between slash and the directive: go:embed
```

# Config
You can add any directives that you want to check whether if there is a space between // and directives.
You can place config file as `checkspaces.yml`.
```yaml
directive:
  - any
  - something:else 
```

checkspaces searches for checkspaces.yml in directories up to the root from the file directory which analyzing currently. (not the working directory(command executed))

You can use the -checkspaces.configPath flag at runtime to indicate config by an absolute path.
