# GoFusion
GoFusion is a utility library for Go that provides a collection of useful functions to simplify and enhance your code. It is inspired by the popular JavaScript library Lodash, but built specifically for Go.

[![Go Reference](https://pkg.go.dev/badge/github.com/your-username/your-package-name.svg)](https://pkg.go.dev/github.com/vatsalpatel/gofusion)

## Installation

Use the following `go get` command to install the package:

```
go get github.com/vatsalpatel/gofusion
```

## Usage
Import the package into your Go code:
```
import (
    ut github.com/vatsalpatel/gofusion
)

func main() {
	mapp := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	ut.Keys[string](mapp)
}

```

## Testing
Run the following command:
```
go test
```


## Contributing
Contributions are welcome! If you encounter any issues or have suggestions for improvements, please open an issue or submit a pull request on GitHub.

## License
This project is licensed under the MIT License - see the LICENSE file for details.