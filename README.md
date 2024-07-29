# Go Interfaces Example

This project provides a basic example of how to use interfaces in Go. It demonstrates defining an interface, implementing it with different types, and using it to achieve polymorphism.

## Overview

The program illustrates the following concepts:
1. Defining an interface.
2. Implementing the interface with different types.
3. Using the interface to call methods on different types.

## Usage

To run the program, ensure you have Go installed on your system. Then, follow these steps:

1. Save the code to a file, for example, `main.go`.
2. Open a terminal and navigate to the directory containing `main.go`.
3. Run the following command to execute the program:

```bash
go run main.go
```

## Code Explanation

The main components of the code are:

1. **Defining Types:**
   ```go
   type englishBot struct{}
   type spanishBot struct{}
   ```

   Two types, `englishBot` and `spanishBot`, are defined. These types do not have any fields or methods at this point.

2. **Defining the Interface:**
   ```go
   type bot interface {
       getGreeting() string
   }
   ```

   An interface named `bot` is defined. It declares a method `getGreeting` that returns a string. Any type that implements this method satisfies the `bot` interface.

3. **Implementing the Interface:**
   ```go
   func (englishBot) getGreeting() string {
       return "hi"
   }

   func (spanishBot) getGreeting() string {
       return "hola"
   }
   ```

   Both `englishBot` and `spanishBot` implement the `getGreeting` method, fulfilling the `bot` interface.

4. **Using the Interface:**
   ```go
   func printGreeting(b bot) {
       fmt.Println(b.getGreeting())
   }
   ```

   The `printGreeting` function takes a `bot` interface as a parameter and calls its `getGreeting` method. This function works with any type that implements the `bot` interface.

5. **Main Function:**
   ```go
   func main() {
       eb := englishBot{}
       sb := spanishBot{}

       printGreeting(sb)
       printGreeting(eb)
   }
   ```

   In the `main` function, instances of `englishBot` and `spanishBot` are created and passed to `printGreeting`, demonstrating polymorphism.

## Notes

- Interfaces in Go provide a way to specify methods that a type must implement without specifying how these methods are implemented.
- This example shows how different types can share common behavior through an interface, allowing for flexible and reusable code.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue if you have any suggestions or improvements.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
