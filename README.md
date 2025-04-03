# Go Utils Library

![GitHub Tag](https://img.shields.io/github/v/tag/go-universal/utils?sort=semver&label=version)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-universal/utils.svg)](https://pkg.go.dev/github.com/go-universal/utils)
[![License](https://img.shields.io/badge/license-ISC-blue.svg)](https://github.com/go-universal/utils/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-universal/utils)](https://goreportcard.com/report/github.com/go-universal/utils)
![Contributors](https://img.shields.io/github/contributors/go-universal/utils)
![Issues](https://img.shields.io/github/issues/go-universal/utils)

This library provides a collection of utility functions for Go, covering various domains such as pointers, numbers, strings, web, images, and file operations.

## Installation

To install the library, use the following command:

```sh
go get github.com/go-universal/utils
```

## API Documentation

### Pointer Utilities

#### `PointerOf`

Returns the pointer of a given value.

```go
package main

import (
    "fmt"
    "github.com/go-universal/utils"
)

func main() {
    value := 42
    ptr := utils.PointerOf(value)
    fmt.Println(*ptr) // Output: 42
}
```

#### `SafeValue`

Returns the value of a pointer or an empty value if the pointer is nil.

```go
package main

import (
    "fmt"
    "github.com/go-universal/utils"
)

func main() {
    var ptr *int
    value := utils.SafeValue(ptr)
    fmt.Println(value) // Output: 0
}
```

#### `ValueOf`

Returns the value of a pointer or a fallback value if the pointer is nil.

```go
package main

import (
    "fmt"
    "github.com/go-universal/utils"
)

func main() {
    var ptr *int
    value := utils.ValueOf(ptr, 10)
    fmt.Println(value) // Output: 10
}
```

#### `Alter`

Returns the value of a pointer or a fallback value if the pointer is nil or zero.

```go
package main

import (
    "fmt"
    "github.com/go-universal/utils"
)

func main() {
    var ptr int
    value := utils.Alter(&ptr, 10)
    fmt.Println(value) // Output: 10
}
```

#### `NullableOf`

Returns nil if the value is zero.

```go
package main

import (
    "fmt"
    "github.com/go-universal/utils"
)

func main() {
    value := 0
    ptr := utils.NullableOf(&value)
    fmt.Println(ptr) // Output: <nil>
}
```

#### `IsEmpty`

Checks if a pointer is nil or zero.

```go
package main

import (
    "fmt"
    "github.com/go-universal/utils"
)

func main() {
    var ptr *int
    isEmpty := utils.IsEmpty(ptr)
    fmt.Println(isEmpty) // Output: true
}
```

#### `IsSame`

Checks if two pointer values are equal.

```go
package main

import (
    "fmt"
    "github.com/go-universal/utils"
)

func main() {
    a := 42
    b := 42
    isSame := utils.IsSame(&a, &b)
    fmt.Println(isSame) // Output: true
}
```

### Number Utilities

#### `Abs`

Returns the absolute value of a number.

```go
package main

import (
    "fmt"
    "github.com/go-universal/utils"
)

func main() {
    value := -42
    absValue := utils.Abs(value)
    fmt.Println(absValue) // Output: 42
}
```

#### `Round`

Returns the nearest integer, rounding half away from zero.

```go
package main

import (
    "fmt"
    "github.com/go-universal/utils"
)

func main() {
    value := 42.5
    roundedValue := utils.Round[int](value)
    fmt.Println(roundedValue) // Output: 43
}
```

#### `RoundUp`

Returns the nearest larger integer (ceil).

```go
package main

import (
    "fmt"
    "github.com/go-universal/utils"
)

func main() {
    value := 42.3
    roundedValue := utils.RoundUp[int](value)
    fmt.Println(roundedValue) // Output: 43
}
```

#### `RoundDown`

Returns the nearest smaller integer.

```go
package main

import (
    "fmt"
    "github.com/go-universal/utils"
)

func main() {
    value := 42.9
    roundedValue := utils.RoundDown[int](value)
    fmt.Println(roundedValue) // Output: 42
}
```

#### `Min`

Returns the smallest value among the given numbers or zero.

```go
package main

import (
    "fmt"
    "github.com/go-universal/utils"
)

func main() {
    values := []int{42, 10, 56}
    minValue := utils.Min(values...)
    fmt.Println(minValue) // Output: 10
}
```

#### `Max`

Returns the largest value among the given numbers or zero.

```go
package main

import (
    "fmt"
    "github.com/go-universal/utils"
)

func main() {
    values := []int{42, 10, 56}
    maxValue := utils.Max(values...)
    fmt.Println(maxValue) // Output: 56
}
```

### String Utilities

#### `ExtractNumbers`

Extracts numbers from a string.

```go
package main

import (
    "fmt"
    "github.com/go-universal/utils"
)

func main() {
    str := "abc123def"
    numbers := utils.ExtractNumbers(str)
    fmt.Println(numbers) // Output: 123
}
```

#### `ExtractAlphaNum`

Extracts alphanumeric characters from a string.

```go
package main

import (
    "fmt"
    "github.com/go-universal/utils"
)

func main() {
    str := "abc123!@#"
    alphaNum := utils.ExtractAlphaNum(str)
    fmt.Println(alphaNum) // Output: abc123
}
```

#### `ExtractAlphaNumPersian`

Extracts English and Persian alphanumeric characters from a string.

```go
package main

import (
    "fmt"
    "github.com/go-universal/utils"
)

func main() {
    str := "abc123!@#سلام"
    alphaNum := utils.ExtractAlphaNumPersian(str)
    fmt.Println(alphaNum) // Output: abc123سلام
}
```

#### `RandomString`

Generates a random string from a given set of characters.

```go
package main

import (
    "fmt"
    "github.com/go-universal/utils"
)

func main() {
    randomStr := utils.RandomString(10, "abcdef123456")
    fmt.Println(randomStr) // Output: Random 10-character string from the given set
}
```

#### `RandomNumeric`

Returns a random numeric string.

```go
package main

import (
    "fmt"
    "github.com/go-universal/utils"
)

func main() {
    randomNum := utils.RandomNumeric(10)
    fmt.Println(randomNum) // Output: Random 10 digit number
}
```

#### `RandomAlphaNum`

Generates a random alphanumeric string with uppercase characters.

```go
package main

import (
    "fmt"
    "github.com/go-universal/utils"
)

func main() {
    randomAlphaNum := utils.RandomAlphaNum(10)
    fmt.Println(randomAlphaNum) // Output: Random 10-character alphanumeric string
}
```

#### `Slugify`

Makes a URL-friendly slug from strings.

```go
package main

import (
    "fmt"
    "github.com/go-universal/utils"
)

func main() {
    slug := utils.Slugify("Hello World!")
    fmt.Println(slug) // Output: hello-world
}
```

#### `SlugifyPersian`

Creates a URL-friendly slug from strings, supporting English and Persian characters.

```go
package main

import (
    "fmt"
    "github.com/go-universal/utils"
)

func main() {
    slug := utils.SlugifyPersian("سلام دنیا!")
    fmt.Println(slug) // Output: سلام-دنیا
}
```

#### `Concat`

Concatenates non-empty strings with a separator.

```go
package main

import (
    "fmt"
    "github.com/go-universal/utils"
)

func main() {
    result := utils.Concat("-", "hello", "", "world")
    fmt.Println(result) // Output: hello-world
}
```

#### `FormatNumber`

Formats a number with a comma separator.

```go
package main

import (
    "fmt"
    "github.com/go-universal/utils"
)

func main() {
    formatted := utils.FormatNumber("%d Dollars", 100000)
    fmt.Println(formatted) // Output: 100,000 Dollars
}
```

#### `FormatRx`

Formats a string using a regex pattern with match groups.

```go
package main

import (
    "fmt"
    "github.com/go-universal/utils"
)

func main() {
    formatted, err := utils.FormatRx("123456", `(\d{3})(\d{2})(\d{1})`, "($1) $2-$3")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(formatted) // Output: (123) 45-6
}
```

### Web Utilities

#### `RelativeURL`

Returns the relative URL path of a file with respect to the root directory.

```go
package main

import (
    "fmt"
    "github.com/go-universal/utils"
)

func main() {
relativeURL := utils.RelativeURL("/root", "/root/path/to/file")
fmt.Println(relativeURL) // Output: path/to/file
}
```

#### `AbsoluteURL`

Returns the absolute URL path of a file with respect to the root directory.

```go
package main

import (
    "fmt"
    "github.com/go-universal/utils"
)

func main() {
    absoluteURL := utils.AbsoluteURL("/root", "/root/path/to/file")
    fmt.Println(absoluteURL) // Output: /path/to/file
}
```

#### `SanitizeRaw`

Sanitizes input to raw text.

```go
package main

import (
    "fmt"
    "github.com/go-universal/utils"
)

func main() {
    sanitized := utils.SanitizeRaw("<script>alert('xss')</script> and more", true)
    fmt.Println(sanitized) // Output: and more
}
```

#### `SanitizeCommon`

Sanitizes input to HTML with common allowed tags.

```go
package main

import (
    "fmt"
    "github.com/go-universal/utils"
)

func main() {
    sanitized := utils.SanitizeCommon("<b>bold</b> <script>alert('xss')</script>", true)
    fmt.Println(sanitized) // Output: <b>bold</b>
}
```

### File Utilities

#### `NormalizePath`

Joins and normalizes a file path using slash separator.

```go
package main

import (
    "fmt"
    "github.com/go-universal/utils"
)

func main() {
    path := utils.NormalizePath("/root", "path//to\\some", "file")
    fmt.Println(path) // Output: /root/path/to/some/file
}
```

#### `CreateDirectory`

Creates a nested directory.

```go
package main

import (
    "fmt"
    "github.com/go-universal/utils"
)

func main() {
    err := utils.CreateDirectory("/path/to/directory")
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("Directory created successfully")
    }
}
```

#### `IsDirectory`

Checks if a path is a directory.

```go
package main

import (
    "fmt"
    "github.com/go-universal/utils"
)

func main() {
    isDir, err := utils.IsDirectory("/path/to/directory")
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(isDir) // Output: true or false
    }
}
```

#### `GetSubDirectory`

Returns a list of subdirectories in a directory.

```go
package main

import (
    "fmt"
    "github.com/go-universal/utils"
)

func main() {
    subDirs, err := utils.GetSubDirectory("/path/to/directory")
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(subDirs) // Output: [subdir1 subdir2 ...]
    }
}
```

#### `ClearDirectory`

Deletes all files and subdirectories in a directory.

```go
package main

import (
    "fmt"
    "github.com/go-universal/utils"
)

func main() {
    err := utils.ClearDirectory("/path/to/directory")
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("Directory cleared successfully")
    }
}
```

#### `FileExists`

Checks if a file exists.

```go
package main

import (
    "fmt"
    "github.com/go-universal/utils"
)

func main() {
    exists, err := utils.FileExists("/path/to/file")
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(exists) // Output: true or false
    }
}
```

#### `FindFile`

Searches a directory for a file matching a pattern and returns the first match.

```go
package main

import (
    "fmt"
    "github.com/go-universal/utils"
)

func main() {
    file := utils.FindFile("/path/to/directory", ".*\\.txt$")
    if file != nil {
        fmt.Println(*file) // Output: /path/to/directory/file.txt
    } else {
        fmt.Println("File not found")
    }
}
```

#### `FindFiles`

Searches a directory for files matching a pattern.

```go
package main

import (
    "fmt"
    "github.com/go-universal/utils"
)

func main() {
    files := utils.FindFiles("/path/to/directory", ".*\\.txt$")
    fmt.Println(files) // Output: [/path/to/directory/file1.txt /path/to/directory/file2.txt ...]
}
```

#### `GetMime`

Returns file MIME info from content.

```go
package main

import (
    "fmt"
    "github.com/go-universal/utils"
    "os"
)

func main() {
    data, _ := os.ReadFile("/path/to/file.txt")
    mime := utils.GetMime(data)
    fmt.Println(mime.String()) // Output: text/plain; charset=utf-8
}
```

#### `GetExtension`

Returns the file extension.

```go
package main

import (
    "fmt"
    "github.com/go-universal/utils"
)

func main() {
    ext := utils.GetExtension("file.txt")
    fmt.Println(ext) // Output: txt
}
```

#### `GetFilename`

Returns the file name without the extension.

```go
package main

import (
    "fmt"
    "github.com/go-universal/utils"
)

func main() {
    name := utils.GetFilename("file.txt")
    fmt.Println(name) // Output: file
}
```

#### `TimestampedFile`

Returns a file name with a timestamp prefix.

```go
package main

import (
    "fmt"
    "github.com/go-universal/utils"
)

func main() {
    timestamped := utils.TimestampedFile("file.txt")
    fmt.Println(timestamped) // Output: file-1633024800000.txt
}
```

#### `NumberedFile`

Generates a unique numbered file name.

```go
package main

import (
    "fmt"
    "github.com/go-universal/utils"
)

func main() {
    fileName, err := utils.NumberedFile("/path/to/directory", "file.txt")
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(fileName) // Output: file-1.txt (if file.txt already exists)
    }
}
```

### Image Utilities

#### `CreateThumbnail`

Generates a thumbnail image from the provided image data.

```go
package main

import (
    "fmt"
    "github.com/go-universal/utils"
    "os"
)

func main() {
    imgData, _ := os.ReadFile("image.jpg")
    thumbnail, format, err := utils.CreateThumbnail(imgData, 150)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("Thumbnail created in format: %s\n", format)
}
```
