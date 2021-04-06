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

type Person struct {   // Define a new struct.
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

**Control structures**

_if-else_
```
// No wrapping on conditional check.
// All blocks require braces.
if x == 1 {
    fmt.Println("X is equal to 1")
} else {
    fmt.Println("X is not equal to 1")
}
```

If you need conditions with more than two choices. Use switch.

However, be careful when using switch with types that can be extended.
If this is the case, code reusability can be improved using interfaces and the double dispatch pattern.

Break statements are not required. Execution breaks automatically at the end of the matching case.

We can also get the type of a variable using `<interface_name>.(type)`.
However, using this tool is considered an antipattern is SWE circles. Use interface inheritance instead.

_switch_
```
switch x {
    case "foo":
        fmt.Println("Found foo")
    case "bar":
        fmt.Println("Found bar")
    default:
        fmt.Println("Default case")
}
```

The only control structure in Go to implement loops is the for-loop.

_for loop_
```
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

We will need a functionality to loop over collections.
We will use iterators for that.

_for loop with iterators_
```
collection := []int{2, 4, 8, 16}

// 'range()' will yield the index and the value of the item.
// This way, we can either read the value, or modify it using the
// index and the collection selection functionality.
for idx, val := range collection {
    fmt.Println(idx, val)
}

for idx, _ := range collection {
    fmt.Println(idx)
}


for _, val := range collection {
    fmt.Println(val)
}
```

**Concurrency**

- Go's concurrency model is based on something called _goroutines_.
- A goroutine is a function or method that can run simultaneously.
- A goroutine is also considered a _ligthweight thread_ because its creation cost is minimal.


_coroutine_
```

// Function to be executed concurrently.
func f() {
    fmt.Println("f function")
}

func main() {
    go f()  // Execute f

    // We need to keep the main thread alive until f() has finished
    // in order to see its results.
    time.Sleep(1 * time.Second)
    fmt.Println("main function")
}
```

**Error Handling**

- Go does not have _try-catchs_.
- Minimalistic approach, a function fails when it fails and must be handled at the point it occurs.

```
// Anything that implements this interface is an error.
// Rememeber that methods with Uppercase are packet-public.
type error interface {
    Error() string
}

// You can use any data type that implements such a method
// to implement custom errors.
type MyError string
func (e MyError) Error() string {
    return string(e)
}

func foo() error {
    return errors.New("Some Error Occurred")
}

func main() {
    // This is a common pattern to handling errors.
    // Notice this is similar to the way things are done in
    // POSIX C.
    if err := foo(); err != nill {
        // Handle the error
    }
}
```



**Handling Structured Data**
- Common encodings: JSON or XML.
- Go contains standard packages for data encoding: `encoding/json` and `encoding/xml`.

_Example_
```
type Foo struct {
    Bar string
    Baz string
}

// This example code is ignoring best practices for 
// the sake of simplicity.
func main() {
    // First define the struct.
    f := Foo{"Joe Junior", "Hello Shabado"}

    // Pass the struct instance to the marshall function.
    // This will return a byte slice (and an error).
    b, _ := json.Marshal(f)
    fmt.Println(string(b))

    json.Unmarshal(b, &f)
}
```
