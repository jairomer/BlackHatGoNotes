# Fundamentals

## Common Go Tool Commands

**go run**
- Compile and execute the _main package_.

**go build**
- Compile the main package and its dependencies into a self-contained binary.
- Note: Hello world generates a binary 1.9 MB in size.
- If we use flags `-ldflags "-w -s"`, we can remove debug information and reduce such binary to 1.3MB.
- Expect around 30% improvement.

**Cross-Compiling**
- Creating binaries for different architectures.
- We need to pass to the build command a _constraint_ with the operating system and the architecture.
- Set `GOOS` and `GOARCH`.

**go doc**
- Command to interrogate documentation about a package, function, method or variable.
- This documentation is embedded as comments through your code.
- Example: `go doc fmt.Println`

**go get**
- Integrated third party manager for a project.
- Two additional tools, dep and mod, to lock dependencies.
- Prevents backward compatibility issues.


**go fmt**
- Automatic formatting for Go sources.

**golint and go vet**
- golint reports style mistakes and specification level errors. Needs to be installed separately.
- go vet is a similar tool to detect bad practices on the source code that a compiler might miss.

## Go Syntax

### Data types

Two types, primitive and complex data types.

**Primitive Data Types**


```Go
bool, uintptr, byte, rune,
string,
int, int8, int16. int32, int64,
uint, uint8, uint16, uint32, uint64,
float32, float64,
complex64 and complex128

// Variables declared as:
var x = "Hello World"   // Implicitly typed
z := int(43)            // Explicitly typed

// Slices and maps
// Slices are like arrays with dynamic allocation and copy/move constructors.
// Maps are associative arrays, unordered lists of key/value pairs.

var slc = make([]string, 0)     // Declare the slice.
slc = append(slc, "some string")    // Extend and reassign.

var mp = make(map[string]string) // Map from one string to another.
mp["some key"] = "some value"    // Extend with a value for a given key.

// Pointers, Structs, and Interfaces

// A pointer is the same concept as in C.
var count = int(42)
ptr := &count       // Get memory address.
fmt.Println(*ptr)   // Print value stored at ptr memory address.
*ptr = 100          // Assign value '100' to value on memory location, can it be different type? NO!

// A struct defines new data types from other ones.

type Persons struct {   // Define a new struct.
    Name string
    Age int
}

// Define a new method that uses a reference to a person.
// If no reference or copy is included, then it is called a function.
func (p *Person) SayHello() {
    fmt.Println("Hello,", p.Name)
}

func main() {
    var guy = new(Person) // Initializes an instance of a 'Person'
    guy.Name = "Dave"
    guy.SayHello()
}

// An interface in Go is a blueprint or contract that defines the expected
// set of actions that any concrete implementation must fulfill in order
// to be considered a type of that interface.

// If it can SayHello, then it is a Friend.
type Friend interface {
    SayHello()
}

func Greet (f Friend) {
    f.SayHello()
}

```



