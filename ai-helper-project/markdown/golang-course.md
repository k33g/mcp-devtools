# Complete Go Programming Course

## Table of Contents

1. [Introduction to Go](#1-introduction-to-go)
2. [Getting Started](#2-getting-started)
3. [Basic Syntax and Concepts](#3-basic-syntax-and-concepts)
4. [Data Types and Variables](#4-data-types-and-variables)
5. [Control Structures](#5-control-structures)
6. [Functions](#6-functions)
7. [Arrays, Slices, and Maps](#7-arrays-slices-and-maps)
8. [Structs and Methods](#8-structs-and-methods)
9. [Interfaces](#9-interfaces)
10. [Error Handling](#10-error-handling)
11. [Concurrency and Goroutines](#11-concurrency-and-goroutines)
12. [Packages and Modules](#12-packages-and-modules)
13. [Testing](#13-testing)
14. [Advanced Topics](#14-advanced-topics)
15. [Best Practices](#15-best-practices)
16. [Exercises and Projects](#16-exercises-and-projects)

---

## 1. Introduction to Go

### What is Go?

Go (also known as Golang) is an open-source programming language developed by Google in 2007 and released in 2009. It was created by Robert Griesemer, Rob Pike, and Ken Thompson to address the challenges of modern software development.

### Key Features

- **Simple and readable syntax**
- **Fast compilation**
- **Built-in concurrency support**
- **Garbage collection**
- **Static typing**
- **Cross-platform compatibility**
- **Rich standard library**

### Why Choose Go?

- High performance similar to C/C++
- Simple syntax like Python
- Excellent for backend services, APIs, and microservices
- Strong support for concurrent programming
- Great tooling and testing framework
- Growing ecosystem and community

### Use Cases

- Web servers and APIs
- Microservices
- DevOps tools (Docker, Kubernetes)
- Command-line applications
- Network programming
- Cloud-native applications

---

## 2. Getting Started

### Installation

#### On macOS
```bash
# Using Homebrew
brew install go

# Or download from official website
# https://golang.org/dl/
```

#### On Linux
```bash
# Download and extract
wget https://golang.org/dl/go1.21.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz

# Add to PATH
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```

#### On Windows
Download the installer from https://golang.org/dl/ and run it.

### Verify Installation
```bash
go version
```

### Setting up Your Workspace

#### Environment Variables
```bash
# Add to your shell profile (.bashrc, .zshrc, etc.)
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

#### Create Your First Program
```bash
mkdir hello-world
cd hello-world
go mod init hello-world
```

Create `main.go`:
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

Run your program:
```bash
go run main.go
```

---

## 3. Basic Syntax and Concepts

### Package Declaration
Every Go file starts with a package declaration:
```go
package main
```

### Import Statements
```go
import "fmt"

// Multiple imports
import (
    "fmt"
    "os"
    "strings"
)
```

### The main Function
```go
func main() {
    // Entry point of the program
    fmt.Println("Hello, Go!")
}
```

### Comments
```go
// Single line comment

/*
Multi-line
comment
*/
```

### Basic Structure
```go
package main

import "fmt"

// Global variable
var globalVar = "I'm global"

func main() {
    // Local variable
    localVar := "I'm local"
    fmt.Println(globalVar, localVar)
}
```

---

## 4. Data Types and Variables

### Basic Types

#### Numeric Types
```go
// Integers
var i int = 42
var i8 int8 = 127
var i16 int16 = 32767
var i32 int32 = 2147483647
var i64 int64 = 9223372036854775807

// Unsigned integers
var ui uint = 42
var ui8 uint8 = 255
var ui16 uint16 = 65535
var ui32 uint32 = 4294967295
var ui64 uint64 = 18446744073709551615

// Floating point
var f32 float32 = 3.14
var f64 float64 = 3.141592653589793

// Complex numbers
var c64 complex64 = 1 + 2i
var c128 complex128 = 1 + 2i
```

#### String and Boolean
```go
var s string = "Hello, Go!"
var b bool = true
```

#### Byte and Rune
```go
var by byte = 'A'      // byte is alias for uint8
var r rune = '🐹'      // rune is alias for int32 (Unicode code point)
```

### Variable Declarations

#### Using var keyword
```go
var name string = "John"
var age int = 30
var isActive bool = true

// Multiple variables
var (
    name     string = "John"
    age      int    = 30
    isActive bool   = true
)

// Type inference
var name = "John"  // string inferred
var age = 30       // int inferred
```

#### Short declaration (inside functions only)
```go
func main() {
    name := "John"
    age := 30
    isActive := true
    
    // Multiple assignment
    x, y := 10, 20
}
```

#### Zero Values
```go
var i int     // 0
var f float64 // 0.0
var b bool    // false
var s string  // ""
```

### Constants
```go
const Pi = 3.14159
const (
    StatusOK = 200
    StatusNotFound = 404
    StatusInternalServerError = 500
)

// iota for auto-incrementing constants
const (
    Sunday = iota    // 0
    Monday          // 1
    Tuesday         // 2
    Wednesday       // 3
    Thursday        // 4
    Friday          // 5
    Saturday        // 6
)
```

### Type Conversion
```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)

// String conversions
import "strconv"

str := strconv.Itoa(42)        // int to string
num, err := strconv.Atoi("42") // string to int
```

---

## 5. Control Structures

### If Statements
```go
age := 20

if age >= 18 {
    fmt.Println("Adult")
}

// If-else
if age >= 18 {
    fmt.Println("Adult")
} else {
    fmt.Println("Minor")
}

// If-else if-else
if age < 13 {
    fmt.Println("Child")
} else if age < 18 {
    fmt.Println("Teenager")
} else {
    fmt.Println("Adult")
}

// If with initialization
if num := 42; num%2 == 0 {
    fmt.Println("Even")
} else {
    fmt.Println("Odd")
}
```

### Switch Statements
```go
day := "Monday"

switch day {
case "Monday":
    fmt.Println("Start of work week")
case "Friday":
    fmt.Println("TGIF!")
case "Saturday", "Sunday":
    fmt.Println("Weekend")
default:
    fmt.Println("Regular day")
}

// Switch without expression (like if-else chain)
age := 25
switch {
case age < 18:
    fmt.Println("Minor")
case age < 65:
    fmt.Println("Adult")
default:
    fmt.Println("Senior")
}

// Switch with initialization
switch num := 42; {
case num < 0:
    fmt.Println("Negative")
case num == 0:
    fmt.Println("Zero")
default:
    fmt.Println("Positive")
}
```

### Loops

#### For Loop (the only loop in Go)
```go
// Traditional for loop
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// While-like loop
i := 0
for i < 5 {
    fmt.Println(i)
    i++
}

// Infinite loop
for {
    fmt.Println("Forever")
    break // Use break to exit
}

// Range loop (for arrays, slices, maps, strings)
numbers := []int{1, 2, 3, 4, 5}
for index, value := range numbers {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}

// Range with only value
for _, value := range numbers {
    fmt.Println(value)
}

// Range with only index
for index := range numbers {
    fmt.Println(index)
}

// Range over string (runes)
for i, char := range "Hello" {
    fmt.Printf("Index: %d, Char: %c\n", i, char)
}
```

### Break and Continue
```go
for i := 0; i < 10; i++ {
    if i == 3 {
        continue // Skip iteration
    }
    if i == 7 {
        break // Exit loop
    }
    fmt.Println(i)
}

// Labeled break and continue
outer:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if i == 1 && j == 1 {
            break outer // Break out of outer loop
        }
        fmt.Printf("%d,%d ", i, j)
    }
}
```

---

## 6. Functions

### Basic Function Syntax
```go
func functionName(parameter1 type1, parameter2 type2) returnType {
    // function body
    return value
}
```

### Examples
```go
// Simple function
func greet(name string) string {
    return "Hello, " + name
}

// Function with multiple parameters
func add(a, b int) int {
    return a + b
}

// Function with multiple return values
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

// Named return values
func rectangle(length, width float64) (area, perimeter float64) {
    area = length * width
    perimeter = 2 * (length + width)
    return // naked return
}

// Variadic functions
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

func main() {
    fmt.Println(greet("Alice"))
    fmt.Println(add(5, 3))
    
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }
    
    a, p := rectangle(5, 3)
    fmt.Printf("Area: %.2f, Perimeter: %.2f\n", a, p)
    
    fmt.Println(sum(1, 2, 3, 4, 5))
}
```

### Function Types and Variables
```go
// Function as a type
type operation func(int, int) int

func add(a, b int) int {
    return a + b
}

func multiply(a, b int) int {
    return a * b
}

func main() {
    var op operation
    op = add
    fmt.Println(op(3, 4)) // 7
    
    op = multiply
    fmt.Println(op(3, 4)) // 12
}
```

### Anonymous Functions and Closures
```go
func main() {
    // Anonymous function
    result := func(a, b int) int {
        return a + b
    }(5, 3)
    fmt.Println(result) // 8
    
    // Closure
    multiplier := func(factor int) func(int) int {
        return func(num int) int {
            return num * factor
        }
    }
    
    double := multiplier(2)
    triple := multiplier(3)
    
    fmt.Println(double(5)) // 10
    fmt.Println(triple(5)) // 15
}
```

### Higher-Order Functions
```go
func applyOperation(numbers []int, op func(int) int) []int {
    result := make([]int, len(numbers))
    for i, num := range numbers {
        result[i] = op(num)
    }
    return result
}

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    
    // Double all numbers
    doubled := applyOperation(numbers, func(n int) int {
        return n * 2
    })
    fmt.Println(doubled) // [2 4 6 8 10]
}
```

---

## 7. Arrays, Slices, and Maps

### Arrays
```go
// Array declaration
var numbers [5]int
fmt.Println(numbers) // [0 0 0 0 0]

// Array initialization
var fruits = [3]string{"apple", "banana", "cherry"}
colors := [...]string{"red", "green", "blue"} // compiler counts

// Accessing elements
fmt.Println(fruits[0]) // apple
fruits[1] = "orange"

// Array length
fmt.Println(len(fruits)) // 3

// Iterating over array
for i, fruit := range fruits {
    fmt.Printf("Index %d: %s\n", i, fruit)
}
```

### Slices (Dynamic Arrays)
```go
// Slice declaration
var numbers []int
fmt.Println(numbers == nil) // true

// Creating slices
numbers = make([]int, 5)      // length 5, capacity 5
numbers = make([]int, 3, 5)   // length 3, capacity 5

// Slice literals
fruits := []string{"apple", "banana", "cherry"}
primes := []int{2, 3, 5, 7, 11}

// Appending to slices
fruits = append(fruits, "date")
fruits = append(fruits, "elderberry", "fig")

// Slicing
fmt.Println(fruits[1:3])  // [banana cherry]
fmt.Println(fruits[:2])   // [apple banana]
fmt.Println(fruits[2:])   // [cherry date elderberry fig]
fmt.Println(fruits[:])    // all elements

// Length and capacity
fmt.Println(len(fruits)) // length
fmt.Println(cap(fruits)) // capacity
```

### Slice Operations
```go
// Copy slices
original := []int{1, 2, 3, 4, 5}
copied := make([]int, len(original))
copy(copied, original)

// Remove element at index
index := 2
slice := append(slice[:index], slice[index+1:]...)

// Insert element at index
index = 1
value := 99
slice = append(slice[:index], append([]int{value}, slice[index:]...)...)

// 2D slices
matrix := [][]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}
```

### Maps
```go
// Map declaration
var ages map[string]int
ages = make(map[string]int)

// Map literal
ages := map[string]int{
    "Alice": 30,
    "Bob":   25,
    "Carol": 35,
}

// Adding/updating elements
ages["David"] = 40
ages["Alice"] = 31 // update

// Accessing elements
age, exists := ages["Alice"]
if exists {
    fmt.Println("Alice is", age, "years old")
}

// Deleting elements
delete(ages, "Bob")

// Iterating over map
for name, age := range ages {
    fmt.Printf("%s is %d years old\n", name, age)
}

// Check if key exists
if age, ok := ages["Eve"]; ok {
    fmt.Println("Eve is", age)
} else {
    fmt.Println("Eve not found")
}
```

### Advanced Map Usage
```go
// Map of slices
scores := map[string][]int{
    "Alice": {95, 87, 92},
    "Bob":   {78, 82, 85},
}

// Map of maps
students := map[string]map[string]interface{}{
    "Alice": {
        "age":    25,
        "grade":  "A",
        "active": true,
    },
}

// Set implementation using map
set := make(map[string]bool)
set["apple"] = true
set["banana"] = true

// Check membership
if set["apple"] {
    fmt.Println("apple is in the set")
}
```

---

## 8. Structs and Methods

### Defining Structs
```go
type Person struct {
    FirstName string
    LastName  string
    Age       int
    Email     string
}

// Anonymous struct
var config = struct {
    Host string
    Port int
}{
    Host: "localhost",
    Port: 8080,
}
```

### Creating and Using Structs
```go
// Different ways to create struct instances
var p1 Person
p1.FirstName = "John"
p1.LastName = "Doe"
p1.Age = 30

p2 := Person{
    FirstName: "Jane",
    LastName:  "Smith",
    Age:       25,
    Email:     "jane@example.com",
}

p3 := Person{"Bob", "Johnson", 35, "bob@example.com"}

// Accessing fields
fmt.Println(p2.FirstName, p2.LastName)

// Pointer to struct
p4 := &Person{"Alice", "Brown", 28, "alice@example.com"}
fmt.Println(p4.FirstName) // Go automatically dereferences
```

### Struct Methods
```go
type Rectangle struct {
    Width  float64
    Height float64
}

// Method with receiver
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

// Pointer receiver (can modify the struct)
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

func (r *Rectangle) SetDimensions(width, height float64) {
    r.Width = width
    r.Height = height
}

func main() {
    rect := Rectangle{Width: 5, Height: 3}
    
    fmt.Println("Area:", rect.Area())
    fmt.Println("Perimeter:", rect.Perimeter())
    
    rect.Scale(2)
    fmt.Printf("After scaling: %+v\n", rect)
}
```

### Struct Embedding (Composition)
```go
type Animal struct {
    Name string
    Age  int
}

func (a Animal) Speak() {
    fmt.Printf("%s makes a sound\n", a.Name)
}

type Dog struct {
    Animal // Embedded struct
    Breed  string
}

func (d Dog) Bark() {
    fmt.Printf("%s barks loudly!\n", d.Name)
}

// Method overriding
func (d Dog) Speak() {
    fmt.Printf("%s the %s barks\n", d.Name, d.Breed)
}

func main() {
    dog := Dog{
        Animal: Animal{Name: "Buddy", Age: 3},
        Breed:  "Golden Retriever",
    }
    
    dog.Speak() // Uses Dog's Speak method
    dog.Bark()
    
    // Can access embedded fields directly
    fmt.Println("Age:", dog.Age)
}
```

### Struct Tags
```go
import (
    "encoding/json"
    "fmt"
)

type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"-"` // Exclude from JSON
}

func main() {
    user := User{
        ID:       1,
        Username: "johndoe",
        Email:    "john@example.com",
        Password: "secret123",
    }
    
    jsonData, _ := json.Marshal(user)
    fmt.Println(string(jsonData))
    // Output: {"id":1,"username":"johndoe","email":"john@example.com"}
}
```

---

## 9. Interfaces

### Defining Interfaces
```go
type Writer interface {
    Write([]byte) (int, error)
}

type Reader interface {
    Read([]byte) (int, error)
}

// Interface composition
type ReadWriter interface {
    Reader
    Writer
}
```

### Implementing Interfaces
```go
import (
    "fmt"
    "strings"
)

type Shape interface {
    Area() float64
    Perimeter() float64
}

type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * 3.14159 * c.Radius
}

// Function that accepts any Shape
func PrintShapeInfo(s Shape) {
    fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
    rect := Rectangle{Width: 5, Height: 3}
    circle := Circle{Radius: 4}
    
    PrintShapeInfo(rect)
    PrintShapeInfo(circle)
}
```

### Empty Interface
```go
func PrintAnything(v interface{}) {
    fmt.Println(v)
}

func main() {
    PrintAnything(42)
    PrintAnything("hello")
    PrintAnything([]int{1, 2, 3})
}
```

### Type Assertion
```go
func main() {
    var i interface{} = "hello"
    
    // Type assertion
    s, ok := i.(string)
    if ok {
        fmt.Println("String value:", s)
    }
    
    // Type assertion without ok (panics if wrong type)
    s2 := i.(string)
    fmt.Println(s2)
}
```

### Type Switch
```go
func ProcessValue(v interface{}) {
    switch value := v.(type) {
    case string:
        fmt.Printf("String: %s (length: %d)\n", value, len(value))
    case int:
        fmt.Printf("Integer: %d (squared: %d)\n", value, value*value)
    case bool:
        fmt.Printf("Boolean: %t\n", value)
    case []int:
        fmt.Printf("Slice of ints: %v (length: %d)\n", value, len(value))
    default:
        fmt.Printf("Unknown type: %T\n", value)
    }
}

func main() {
    ProcessValue("hello")
    ProcessValue(42)
    ProcessValue(true)
    ProcessValue([]int{1, 2, 3})
    ProcessValue(3.14)
}
```

### Common Interfaces
```go
import (
    "fmt"
    "sort"
)

// Implementing sort.Interface
type Person struct {
    Name string
    Age  int
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
    people := []Person{
        {"Alice", 30},
        {"Bob", 25},
        {"Carol", 35},
    }
    
    sort.Sort(ByAge(people))
    fmt.Println(people)
}

// Implementing fmt.Stringer
func (p Person) String() string {
    return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}
```

---

## 10. Error Handling

### The Error Interface
```go
type error interface {
    Error() string
}
```

### Basic Error Handling
```go
import (
    "errors"
    "fmt"
    "strconv"
)

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Result:", result)
}
```

### Creating Custom Errors
```go
import (
    "fmt"
)

// Custom error type
type ValidationError struct {
    Field   string
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("validation error in field '%s': %s", e.Field, e.Message)
}

func validateAge(age int) error {
    if age < 0 {
        return ValidationError{
            Field:   "age",
            Message: "must be non-negative",
        }
    }
    if age > 150 {
        return ValidationError{
            Field:   "age", 
            Message: "must be less than 150",
        }
    }
    return nil
}

func main() {
    if err := validateAge(-5); err != nil {
        fmt.Println(err)
        
        // Type assertion to access custom error fields
        if validationErr, ok := err.(ValidationError); ok {
            fmt.Printf("Field: %s\n", validationErr.Field)
        }
    }
}
```

### Error Wrapping (Go 1.13+)
```go
import (
    "errors"
    "fmt"
)

func processFile(filename string) error {
    err := readFile(filename)
    if err != nil {
        return fmt.Errorf("failed to process file %s: %w", filename, err)
    }
    return nil
}

func readFile(filename string) error {
    return errors.New("file not found")
}

func main() {
    err := processFile("config.txt")
    if err != nil {
        fmt.Println("Error:", err)
        
        // Check if error is of specific type
        if errors.Is(err, errors.New("file not found")) {
            fmt.Println("File not found error detected")
        }
        
        // Unwrap error
        unwrapped := errors.Unwrap(err)
        if unwrapped != nil {
            fmt.Println("Unwrapped:", unwrapped)
        }
    }
}
```

### Panic and Recover
```go
func safeDivide(a, b float64) (result float64, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("panic recovered: %v", r)
        }
    }()
    
    if b == 0 {
        panic("division by zero")
    }
    result = a / b
    return
}

func main() {
    result, err := safeDivide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }
}
```

### Error Handling Best Practices
```go
// Don't ignore errors
result, err := someFunction()
if err != nil {
    // Handle the error appropriately
    return err // or log and return, or handle and continue
}

// Use errors.New for simple error messages
func validateInput(input string) error {
    if input == "" {
        return errors.New("input cannot be empty")
    }
    return nil
}

// Use fmt.Errorf for formatted error messages
func validateRange(value, min, max int) error {
    if value < min || value > max {
        return fmt.Errorf("value %d is outside range [%d, %d]", value, min, max)
    }
    return nil
}

// Define error variables for common errors
var (
    ErrNotFound     = errors.New("not found")
    ErrUnauthorized = errors.New("unauthorized")
    ErrInvalidInput = errors.New("invalid input")
)

func findUser(id int) (*User, error) {
    if id <= 0 {
        return nil, ErrInvalidInput
    }
    // ... search logic ...
    return nil, ErrNotFound
}
```

---

## 11. Concurrency and Goroutines

### Goroutines
```go
import (
    "fmt"
    "time"
)

func sayHello(name string) {
    for i := 0; i < 3; i++ {
        fmt.Printf("Hello, %s! (%d)\n", name, i+1)
        time.Sleep(time.Millisecond * 500)
    }
}

func main() {
    // Sequential execution
    sayHello("Alice")
    sayHello("Bob")
    
    // Concurrent execution with goroutines
    go sayHello("Alice")
    go sayHello("Bob")
    
    // Wait for goroutines to finish (not ideal)
    time.Sleep(time.Second * 2)
}
```

### Channels
```go
func main() {
    // Creating channels
    ch := make(chan string)
    
    // Sending to channel (in goroutine to avoid deadlock)
    go func() {
        ch <- "Hello"
        ch <- "World"
        close(ch)
    }()
    
    // Receiving from channel
    for message := range ch {
        fmt.Println(message)
    }
}
```

### Buffered Channels
```go
func main() {
    // Buffered channel
    ch := make(chan int, 3)
    
    // Can send without goroutine (up to buffer size)
    ch <- 1
    ch <- 2
    ch <- 3
    
    // Receiving
    fmt.Println(<-ch) // 1
    fmt.Println(<-ch) // 2
    fmt.Println(<-ch) // 3
}
```

### Channel Direction
```go
// Send-only channel
func sender(ch chan<- string) {
    ch <- "message"
    close(ch)
}

// Receive-only channel
func receiver(ch <-chan string) {
    for msg := range ch {
        fmt.Println("Received:", msg)
    }
}

func main() {
    ch := make(chan string)
    
    go sender(ch)
    receiver(ch)
}
```

### Select Statement
```go
import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)
    
    go func() {
        time.Sleep(2 * time.Second)
        ch1 <- "Channel 1"
    }()
    
    go func() {
        time.Sleep(1 * time.Second)
        ch2 <- "Channel 2"
    }()
    
    select {
    case msg1 := <-ch1:
        fmt.Println("Received from ch1:", msg1)
    case msg2 := <-ch2:
        fmt.Println("Received from ch2:", msg2)
    case <-time.After(3 * time.Second):
        fmt.Println("Timeout")
    }
}

// Non-blocking select
func nonBlockingReceive(ch <-chan string) {
    select {
    case msg := <-ch:
        fmt.Println("Received:", msg)
    default:
        fmt.Println("No message available")
    }
}
```

### Worker Pool Pattern
```go
import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job)
        time.Sleep(time.Millisecond * 500) // Simulate work
        results <- job * 2
    }
}

func main() {
    const numWorkers = 3
    const numJobs = 10
    
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)
    
    var wg sync.WaitGroup
    
    // Start workers
    for w := 1; w <= numWorkers; w++ {
        wg.Add(1)
        go worker(w, jobs, results, &wg)
    }
    
    // Send jobs
    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    close(jobs)
    
    // Wait for workers to finish
    go func() {
        wg.Wait()
        close(results)
    }()
    
    // Collect results
    for result := range results {
        fmt.Println("Result:", result)
    }
}
```

### Sync Package
```go
import (
    "fmt"
    "sync"
    "time"
)

// Mutex
type Counter struct {
    mu    sync.Mutex
    count int
}

func (c *Counter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.count++
}

func (c *Counter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.count
}

// WaitGroup
func main() {
    var wg sync.WaitGroup
    counter := &Counter{}
    
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for j := 0; j < 100; j++ {
                counter.Increment()
            }
        }()
    }
    
    wg.Wait()
    fmt.Println("Final count:", counter.Value())
}

// Once
var once sync.Once
var config *Config

func GetConfig() *Config {
    once.Do(func() {
        config = loadConfig()
    })
    return config
}
```

### Context Package
```go
import (
    "context"
    "fmt"
    "time"
)

func longRunningOperation(ctx context.Context) error {
    for i := 0; i < 10; i++ {
        select {
        case <-ctx.Done():
            return ctx.Err()
        case <-time.After(500 * time.Millisecond):
            fmt.Printf("Step %d completed\n", i+1)
        }
    }
    return nil
}

func main() {
    // Context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()
    
    err := longRunningOperation(ctx)
    if err != nil {
        fmt.Println("Operation cancelled:", err)
    } else {
        fmt.Println("Operation completed successfully")
    }
}
```

---

## 12. Packages and Modules

### Package Basics
```go
// Every Go file belongs to a package
package main // executable package

package utils // library package

// Package names should be:
// - lowercase
// - short and descriptive
// - single words if possible
```

### Creating a Package
Create directory `mathutils/mathutils.go`:
```go
package mathutils

import "math"

// Exported function (starts with capital letter)
func Add(a, b float64) float64 {
    return a + b
}

// Exported function
func Sqrt(x float64) float64 {
    return math.Sqrt(x)
}

// unexported function (starts with lowercase letter)
func internal() {
    // Only accessible within the package
}

// Exported constant
const Pi = 3.14159

// Exported variable
var Version = "1.0.0"
```

### Using Packages
```go
package main

import (
    "fmt"
    "myproject/mathutils"
)

func main() {
    result := mathutils.Add(5, 3)
    fmt.Println("5 + 3 =", result)
    
    sqrt := mathutils.Sqrt(16)
    fmt.Println("√16 =", sqrt)
    
    fmt.Println("Pi =", mathutils.Pi)
}
```

### Go Modules
```bash
# Initialize a new module
go mod init myproject

# Add dependencies
go get github.com/gorilla/mux

# Remove unused dependencies
go mod tidy

# Download dependencies
go mod download
```

### go.mod file
```go
module myproject

go 1.21

require (
    github.com/gorilla/mux v1.8.0
    github.com/lib/pq v1.10.7
)

require (
    github.com/gorilla/context v1.1.1 // indirect
)
```

### Package Initialization
```go
package mypackage

import "fmt"

var config Config

// init function runs automatically when package is imported
func init() {
    fmt.Println("Initializing mypackage")
    config = loadConfig()
}

func init() {
    // Multiple init functions are allowed
    fmt.Println("Second init function")
}

type Config struct {
    DatabaseURL string
    Port        int
}

func loadConfig() Config {
    return Config{
        DatabaseURL: "localhost:5432",
        Port:        8080,
    }
}
```

### Package Documentation
```go
// Package mathutils provides basic mathematical operations.
//
// This package includes functions for arithmetic operations,
// mathematical constants, and utility functions.
package mathutils

// Add returns the sum of two numbers.
//
// Example:
//   result := Add(5, 3) // result is 8
func Add(a, b float64) float64 {
    return a + b
}

// Pi represents the mathematical constant π.
const Pi = 3.14159265359
```

Generate documentation:
```bash
go doc mathutils
go doc mathutils.Add
```

### Internal Packages
```
myproject/
  internal/
    database/
      db.go  // Only accessible by myproject packages
  api/
    handler.go
  main.go
```

---

## 13. Testing

### Basic Testing
Create `math_test.go`:
```go
package main

import "testing"

func Add(a, b int) int {
    return a + b
}

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    
    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}

func TestAddNegative(t *testing.T) {
    result := Add(-2, -3)
    expected := -5
    
    if result != expected {
        t.Errorf("Add(-2, -3) = %d; want %d", result, expected)
    }
}
```

Run tests:
```bash
go test
go test -v  # verbose output
go test ./... # test all packages
```

### Table-Driven Tests
```go
func TestAddTable(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive numbers", 2, 3, 5},
        {"negative numbers", -2, -3, -5},
        {"zero", 0, 5, 5},
        {"mixed", -2, 5, 3},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
            }
        })
    }
}
```

### Benchmarks
```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(2, 3)
    }
}

func BenchmarkAddComplex(b *testing.B) {
    for i := 0; i < b.N; i++ {
        for j := 0; j < 1000; j++ {
            Add(i, j)
        }
    }
}
```

Run benchmarks:
```bash
go test -bench=.
go test -bench=BenchmarkAdd
```

### Test Coverage
```bash
go test -cover
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Mocking and Interfaces
```go
// Interface for testing
type UserService interface {
    GetUser(id int) (*User, error)
    CreateUser(user *User) error
}

// Implementation
type DatabaseUserService struct {
    db *sql.DB
}

func (s *DatabaseUserService) GetUser(id int) (*User, error) {
    // Database logic
    return nil, nil
}

// Mock for testing
type MockUserService struct {
    users map[int]*User
}

func (m *MockUserService) GetUser(id int) (*User, error) {
    user, exists := m.users[id]
    if !exists {
        return nil, errors.New("user not found")
    }
    return user, nil
}

// Test using mock
func TestUserHandler(t *testing.T) {
    mockService := &MockUserService{
        users: map[int]*User{
            1: {ID: 1, Name: "John"},
        },
    }
    
    handler := NewUserHandler(mockService)
    
    user, err := handler.GetUserByID(1)
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }
    
    if user.Name != "John" {
        t.Errorf("Expected name 'John', got '%s'", user.Name)
    }
}
```

### Test Helpers
```go
func TestMain(m *testing.M) {
    // Setup before tests
    setup()
    
    // Run tests
    code := m.Run()
    
    // Cleanup after tests
    teardown()
    
    os.Exit(code)
}

func setup() {
    fmt.Println("Setting up tests")
}

func teardown() {
    fmt.Println("Tearing down tests")
}

// Helper function
func assertEqual(t *testing.T, actual, expected interface{}) {
    t.Helper() // Marks this function as a test helper
    if actual != expected {
        t.Errorf("Expected %v, got %v", expected, actual)
    }
}

func TestWithHelper(t *testing.T) {
    result := Add(2, 3)
    assertEqual(t, result, 5)
}
```

---

## 14. Advanced Topics

### Reflection
```go
import (
    "fmt"
    "reflect"
)

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    p := Person{Name: "John", Age: 30}
    
    // Get type information
    t := reflect.TypeOf(p)
    fmt.Println("Type:", t.Name())
    
    // Get value information
    v := reflect.ValueOf(p)
    fmt.Println("Value:", v)
    
    // Iterate over fields
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        value := v.Field(i)
        tag := field.Tag.Get("json")
        
        fmt.Printf("Field: %s, Type: %s, Value: %v, Tag: %s\n",
            field.Name, field.Type, value.Interface(), tag)
    }
    
    // Modify values using reflection (requires pointer)
    pPtr := reflect.ValueOf(&p)
    pValue := pPtr.Elem()
    
    nameField := pValue.FieldByName("Name")
    if nameField.CanSet() {
        nameField.SetString("Jane")
    }
    
    fmt.Printf("Modified: %+v\n", p)
}
```

### Generics (Go 1.18+)
```go
// Generic function
func Max[T comparable](a, b T) T {
    if a > b {
        return a
    }
    return b
}

// Generic type
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
    if len(s.items) == 0 {
        var zero T
        return zero, false
    }
    
    index := len(s.items) - 1
    item := s.items[index]
    s.items = s.items[:index]
    return item, true
}

func main() {
    // Using generic function
    fmt.Println(Max(5, 10))        // 10
    fmt.Println(Max("a", "b"))     // "b"
    
    // Using generic type
    intStack := Stack[int]{}
    intStack.Push(1)
    intStack.Push(2)
    
    value, ok := intStack.Pop()
    if ok {
        fmt.Println("Popped:", value) // 2
    }
    
    stringStack := Stack[string]{}
    stringStack.Push("hello")
    stringStack.Push("world")
}

// Type constraints
type Numeric interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
    ~float32 | ~float64
}

func Sum[T Numeric](numbers []T) T {
    var total T
    for _, num := range numbers {
        total += num
    }
    return total
}
```

### Build Tags
```go
// +build linux darwin

package main

import "fmt"

func main() {
    fmt.Println("This runs on Linux and macOS")
}
```

```go
// +build windows

package main

import "fmt"

func main() {
    fmt.Println("This runs on Windows")
}
```

Build with tags:
```bash
go build -tags="production"
```

### CGO (C Integration)
```go
/*
#include <stdio.h>
#include <stdlib.h>

void hello() {
    printf("Hello from C!\n");
}

int add(int a, int b) {
    return a + b;
}
*/
import "C"

import "fmt"

func main() {
    C.hello()
    
    result := C.add(5, 3)
    fmt.Printf("5 + 3 = %d\n", int(result))
}
```

### Unsafe Package
```go
import (
    "fmt"
    "unsafe"
)

func main() {
    // Convert string to byte slice without copying
    s := "hello"
    b := *(*[]byte)(unsafe.Pointer(&s))
    fmt.Printf("%v\n", b)
    
    // Get size of types
    fmt.Printf("Size of int: %d\n", unsafe.Sizeof(int(0)))
    fmt.Printf("Size of string: %d\n", unsafe.Sizeof(""))
    
    // Pointer arithmetic (dangerous!)
    arr := [3]int{1, 2, 3}
    ptr := unsafe.Pointer(&arr[0])
    
    // Move to next element
    ptr = unsafe.Pointer(uintptr(ptr) + unsafe.Sizeof(arr[0]))
    fmt.Printf("Second element: %d\n", *(*int)(ptr))
}
```

---

## 15. Best Practices

### Code Organization
```go
// Good package structure
myproject/
  cmd/
    server/
      main.go
  pkg/
    database/
      db.go
    handlers/
      user.go
    models/
      user.go
  internal/
    config/
      config.go
  go.mod
  go.sum
  README.md
```

### Naming Conventions
```go
// Good naming
type UserService struct {}
func (s *UserService) GetUserByID(id int) (*User, error) {}

// Package names: short, lowercase, single word
package user
package httputil

// Variable names: camelCase
var userName string
var maxRetries int

// Constants: CamelCase or ALL_CAPS for exported
const MaxConnections = 100
const defaultTimeout = 30 * time.Second

// Interface names: usually end with -er
type Reader interface {}
type Writer interface {}
type UserRepository interface {}
```

### Error Handling Best Practices
```go
// Don't ignore errors
result, err := someFunction()
if err != nil {
    return fmt.Errorf("failed to process: %w", err)
}

// Use early returns
func processUser(id int) error {
    user, err := getUser(id)
    if err != nil {
        return err
    }
    
    if !user.Active {
        return errors.New("user is inactive")
    }
    
    // Continue with processing
    return nil
}

// Define custom error types for specific cases
type ValidationError struct {
    Field   string
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("validation failed for %s: %s", e.Field, e.Message)
}
```

### Performance Tips
```go
// Pre-allocate slices when size is known
func processItems(count int) []Item {
    items := make([]Item, 0, count) // capacity = count
    // ... populate items
    return items
}

// Use string builder for concatenation
func buildString(parts []string) string {
    var builder strings.Builder
    builder.Grow(len(parts) * 10) // Pre-allocate
    
    for _, part := range parts {
        builder.WriteString(part)
    }
    return builder.String()
}

// Use buffered channels for better performance
ch := make(chan int, 100) // buffered

// Pool expensive objects
var pool = sync.Pool{
    New: func() interface{} {
        return make([]byte, 1024)
    },
}

func processData() {
    buf := pool.Get().([]byte)
    defer pool.Put(buf)
    // Use buf
}
```

### Concurrency Best Practices
```go
// Use context for cancellation
func processWithTimeout(ctx context.Context) error {
    done := make(chan error, 1)
    
    go func() {
        // Long running operation
        done <- heavyProcessing()
    }()
    
    select {
    case err := <-done:
        return err
    case <-ctx.Done():
        return ctx.Err()
    }
}

// Limit goroutine creation
func processItems(items []Item) error {
    const maxWorkers = 10
    sem := make(chan struct{}, maxWorkers)
    
    var wg sync.WaitGroup
    for _, item := range items {
        wg.Add(1)
        go func(item Item) {
            defer wg.Done()
            sem <- struct{}{} // acquire
            defer func() { <-sem }() // release
            
            processItem(item)
        }(item)
    }
    
    wg.Wait()
    return nil
}
```

### Testing Best Practices
```go
// Use table-driven tests
func TestValidateEmail(t *testing.T) {
    tests := []struct {
        name    string
        email   string
        wantErr bool
    }{
        {"valid email", "user@example.com", false},
        {"invalid email", "invalid-email", true},
        {"empty email", "", true},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := ValidateEmail(tt.email)
            if (err != nil) != tt.wantErr {
                t.Errorf("ValidateEmail() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}

// Use interfaces for testing
type UserRepository interface {
    GetUser(id int) (*User, error)
}

type UserService struct {
    repo UserRepository
}

// Easy to mock in tests
type MockUserRepository struct {
    users map[int]*User
}

func (m *MockUserRepository) GetUser(id int) (*User, error) {
    user, exists := m.users[id]
    if !exists {
        return nil, errors.New("user not found")
    }
    return user, nil
}
```

---

## 16. Exercises and Projects

### Beginner Exercises

#### Exercise 1: Calculator
Create a command-line calculator that supports basic operations.

```go
// Solution skeleton
package main

import (
    "fmt"
    "strconv"
    "os"
)

func add(a, b float64) float64 {
    return a + b
}

func subtract(a, b float64) float64 {
    return a - b
}

func multiply(a, b float64) float64 {
    return a * b
}

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

func main() {
    if len(os.Args) != 4 {
        fmt.Println("Usage: calculator <num1> <operator> <num2>")
        os.Exit(1)
    }
    
    // Parse arguments
    num1, err := strconv.ParseFloat(os.Args[1], 64)
    if err != nil {
        fmt.Printf("Invalid number: %s\n", os.Args[1])
        os.Exit(1)
    }
    
    operator := os.Args[2]
    
    num2, err := strconv.ParseFloat(os.Args[3], 64)
    if err != nil {
        fmt.Printf("Invalid number: %s\n", os.Args[3])
        os.Exit(1)
    }
    
    var result float64
    switch operator {
    case "+":
        result = add(num1, num2)
    case "-":
        result = subtract(num1, num2)
    case "*":
        result = multiply(num1, num2)
    case "/":
        result, err = divide(num1, num2)
        if err != nil {
            fmt.Printf("Error: %v\n", err)
            os.Exit(1)
        }
    default:
        fmt.Printf("Unknown operator: %s\n", operator)
        os.Exit(1)
    }
    
    fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operator, num2, result)
}
```

#### Exercise 2: Word Counter
Create a program that counts words, lines, and characters in a text file.

#### Exercise 3: Number Guessing Game
Implement a number guessing game where the computer picks a random number and the user tries to guess it.

### Intermediate Projects

#### Project 1: TODO CLI Application
```go
// todo.go
package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "strconv"
    "time"
)

type Todo struct {
    ID        int       `json:"id"`
    Task      string    `json:"task"`
    Done      bool      `json:"done"`
    CreatedAt time.Time `json:"created_at"`
}

type TodoList struct {
    Todos []Todo `json:"todos"`
}

const filename = "todos.json"

func (tl *TodoList) Load() error {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        if os.IsNotExist(err) {
            return nil // File doesn't exist yet
        }
        return err
    }
    return json.Unmarshal(data, tl)
}

func (tl *TodoList) Save() error {
    data, err := json.MarshalIndent(tl, "", "  ")
    if err != nil {
        return err
    }
    return ioutil.WriteFile(filename, data, 0644)
}

func (tl *TodoList) Add(task string) {
    id := len(tl.Todos) + 1
    todo := Todo{
        ID:        id,
        Task:      task,
        Done:      false,
        CreatedAt: time.Now(),
    }
    tl.Todos = append(tl.Todos, todo)
}

func (tl *TodoList) Complete(id int) error {
    for i := range tl.Todos {
        if tl.Todos[i].ID == id {
            tl.Todos[i].Done = true
            return nil
        }
    }
    return fmt.Errorf("todo with ID %d not found", id)
}

func (tl *TodoList) List() {
    if len(tl.Todos) == 0 {
        fmt.Println("No todos found.")
        return
    }
    
    for _, todo := range tl.Todos {
        status := "[ ]"
        if todo.Done {
            status = "[x]"
        }
        fmt.Printf("%d. %s %s\n", todo.ID, status, todo.Task)
    }
}

func main() {
    todoList := &TodoList{}
    todoList.Load()
    
    if len(os.Args) < 2 {
        fmt.Println("Usage: todo <command> [arguments]")
        fmt.Println("Commands:")
        fmt.Println("  add <task>    Add a new todo")
        fmt.Println("  list          List all todos")
        fmt.Println("  done <id>     Mark todo as completed")
        os.Exit(1)
    }
    
    command := os.Args[1]
    
    switch command {
    case "add":
        if len(os.Args) < 3 {
            fmt.Println("Usage: todo add <task>")
            os.Exit(1)
        }
        task := os.Args[2]
        todoList.Add(task)
        todoList.Save()
        fmt.Printf("Added: %s\n", task)
        
    case "list":
        todoList.List()
        
    case "done":
        if len(os.Args) < 3 {
            fmt.Println("Usage: todo done <id>")
            os.Exit(1)
        }
        id, err := strconv.Atoi(os.Args[2])
        if err != nil {
            fmt.Printf("Invalid ID: %s\n", os.Args[2])
            os.Exit(1)
        }
        err = todoList.Complete(id)
        if err != nil {
            fmt.Printf("Error: %v\n", err)
            os.Exit(1)
        }
        todoList.Save()
        fmt.Printf("Completed todo %d\n", id)
        
    default:
        fmt.Printf("Unknown command: %s\n", command)
        os.Exit(1)
    }
}
```

#### Project 2: HTTP REST API
Create a simple REST API for managing users.

#### Project 3: Concurrent Web Scraper
Build a web scraper that fetches data from multiple URLs concurrently.

### Advanced Projects

#### Project 1: Chat Server
```go
// Simple chat server implementation
package main

import (
    "bufio"
    "fmt"
    "net"
    "strings"
)

type Client struct {
    conn     net.Conn
    nickname string
    ch       chan string
}

type Server struct {
    clients    map[net.Conn]*Client
    joining    chan *Client
    leaving    chan *Client
    messages   chan string
}

func NewServer() *Server {
    return &Server{
        clients:  make(map[net.Conn]*Client),
        joining:  make(chan *Client),
        leaving:  make(chan *Client),
        messages: make(chan string),
    }
}

func (s *Server) Start() {
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        panic(err)
    }
    defer listener.Close()
    
    fmt.Println("Chat server started on :8080")
    
    go s.run()
    
    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        go s.handleConnection(conn)
    }
}

func (s *Server) run() {
    for {
        select {
        case client := <-s.joining:
            s.clients[client.conn] = client
            fmt.Printf("%s joined the chat\n", client.nickname)
            s.broadcast(fmt.Sprintf("%s joined the chat", client.nickname), client)
            
        case client := <-s.leaving:
            delete(s.clients, client.conn)
            close(client.ch)
            fmt.Printf("%s left the chat\n", client.nickname)
            s.broadcast(fmt.Sprintf("%s left the chat", client.nickname), client)
            
        case message := <-s.messages:
            s.broadcast(message, nil)
        }
    }
}

func (s *Server) broadcast(message string, sender *Client) {
    for _, client := range s.clients {
        if client != sender {
            select {
            case client.ch <- message:
            default:
                close(client.ch)
                delete(s.clients, client.conn)
            }
        }
    }
}

func (s *Server) handleConnection(conn net.Conn) {
    defer conn.Close()
    
    // Get nickname
    conn.Write([]byte("Enter your nickname: "))
    scanner := bufio.NewScanner(conn)
    scanner.Scan()
    nickname := strings.TrimSpace(scanner.Text())
    
    client := &Client{
        conn:     conn,
        nickname: nickname,
        ch:       make(chan string),
    }
    
    s.joining <- client
    
    // Start goroutine to send messages to client
    go func() {
        for message := range client.ch {
            conn.Write([]byte(message + "\n"))
        }
    }()
    
    // Read messages from client
    for scanner.Scan() {
        message := strings.TrimSpace(scanner.Text())
        if message != "" {
            s.messages <- fmt.Sprintf("%s: %s", nickname, message)
        }
    }
    
    s.leaving <- client
}

func main() {
    server := NewServer()
    server.Start()
}
```

#### Project 2: Distributed Key-Value Store
Implement a simple distributed key-value store with replication.

#### Project 3: Microservice with gRPC
Build a microservice architecture using gRPC for communication.

### Practice Challenges

1. **Algorithm Implementation**: Implement sorting algorithms (quicksort, mergesort)
2. **Data Structures**: Implement a binary search tree, hash table, or graph
3. **System Programming**: Create a file backup utility
4. **Network Programming**: Build a proxy server or load balancer
5. **Concurrency**: Implement a rate limiter or worker pool
6. **Web Development**: Create a blog engine with authentication
7. **Database Integration**: Build an ORM or database migration tool
8. **DevOps Tools**: Create a deployment script or monitoring tool

---

## Conclusion

This comprehensive Go course covers everything from basic syntax to advanced topics like concurrency, testing, and best practices. Go is a powerful language that excels in:

- **Backend services and APIs**
- **Microservices architecture** 
- **Command-line tools**
- **System programming**
- **Cloud-native applications**

### Next Steps

1. **Practice regularly** with coding exercises
2. **Build real projects** to gain experience
3. **Read Go source code** to learn from experts
4. **Contribute to open source** Go projects
5. **Stay updated** with Go releases and community

### Additional Resources

- **Official Documentation**: https://golang.org/doc/
- **Go Blog**: https://blog.golang.org/
- **Effective Go**: https://golang.org/doc/effective_go.html
- **Go Code Review Comments**: https://github.com/golang/go/wiki/CodeReviewComments
- **Awesome Go**: https://awesome-go.com/

Remember: The best way to learn Go is by writing Go code. Start with simple programs and gradually work your way up to more complex projects. Good luck with your Go journey!