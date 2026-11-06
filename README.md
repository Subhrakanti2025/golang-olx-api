

<!---------------------------------------------- START ------------------------------------------------------->
<!-- # Go Package Documentation – Generic Cheat Sheet

## 1. Find a package
Search for the package on:
https://pkg.go.dev/

## 2. Install a package
```bash
go get <package-import-path>
```

Example:
```bash
go get golang.org/x/crypto/bcrypt
```

## 3. Read package documentation
```bash
go doc <package-import-path>
```

Example:
```bash
go doc golang.org/x/crypto/bcrypt
```

## 4. Read a specific function
```bash
go doc <package-import-path>.<FunctionName>
```

Example:
```bash
go doc golang.org/x/crypto/bcrypt.GenerateFromPassword
```

## 5. Check package documentation online
https://pkg.go.dev/

## 6. Understand function signatures

Example:
```go
func FunctionName(param1 string, param2 int) (string, error)
```

- `FunctionName` → Function name
- `param1 string` → First input parameter
- `param2 int` → Second input parameter
- `string` → First return value
- `error` → Second return value

## 7. Useful Go dependency commands

```bash
go get <package>       # Add or update a dependency
go mod tidy            # Clean up dependencies
go list -m all         # List all dependencies
go doc <package>       # Read package documentation
go help doc            # Read go doc command help
```

## My workflow

1. Find the package.
2. Read its documentation.
3. Install the package.
4. Understand the required functions.
5. Check parameters and return values.
6. Implement it in my project.
7. Handle errors and test the implementation. -->

<!------------------------------------------------ END -------------------------------------------------->