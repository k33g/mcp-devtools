## Hello World
Basic program structure and main function
```swift
print("Hello, World!")
```

----------

## Variable Declaration
Different ways to declare and initialize variables
```swift
let name = "John"              // Immutable constant
var age = 30                   // Mutable variable
let city: String = "New York"  // With type annotation
var height: Double = 5.9       // Explicit type

age += 1 // Modify mutable variable

print("Name: \(name), Age: \(age), City: \(city)")
```

----------

## Constants and Variables
Working with constants and variables
```swift
let pi = 3.14159
let maxUsers = 100
var currentUsers = 0

// Multiple variable declaration
let x = 10, y = 20, z = 30

// Optional variables
var optionalName: String? = "Alice"
var implicitOptional: String? = nil

print("Pi: \(pi), Max Users: \(maxUsers)")
print("Optional name: \(optionalName ?? "No name")")
```

----------

## Data Types
Working with basic data types
```swift
let integer: Int = 42
let float: Float = 3.14
let double: Double = 3.14159
let boolean: Bool = true
let character: Character = "🍎"
let string: String = "Swift"

// Type inference
let inferredInt = 42
let inferredDouble = 3.14
let inferredString = "Hello"

print("Int: \(integer), Double: \(double), Bool: \(boolean)")
print("Character: \(character), String: \(string)")
```

----------

## Arrays
Working with arrays
```swift
var fruits = ["apple", "banana", "cherry"]
let numbers: [Int] = [1, 2, 3, 4, 5]

// Array operations
fruits.append("date")
fruits.insert("orange", at: 1)
fruits.remove(at: 0)

print("Fruits: \(fruits)")
print("Count: \(fruits.count)")

// Iterating over arrays
for (index, fruit) in fruits.enumerated() {
    print("\(index): \(fruit)")
}
```

----------

## Dictionaries
Creating and manipulating dictionaries
```swift
var scores = ["Alice": 95, "Bob": 87, "Charlie": 92]
let emptyDict: [String: Int] = [:]

// Dictionary operations
scores["David"] = 88
scores["Alice"] = 98 // Update value

if let aliceScore = scores["Alice"] {
    print("Alice's score: \(aliceScore)")
}

for (name, score) in scores {
    print("\(name): \(score)")
}

scores.removeValue(forKey: "Bob")
print("Updated scores: \(scores)")
```

----------

## Optionals
Working with optional values
```swift
var optionalString: String? = "Hello"
var nilString: String? = nil

// Optional binding
if let unwrapped = optionalString {
    print("String value: \(unwrapped)")
} else {
    print("String is nil")
}

// Nil coalescing
let defaultValue = optionalString ?? "Default"
print("Value: \(defaultValue)")

// Force unwrapping (use carefully)
// let forced = optionalString!

// Optional chaining
let length = optionalString?.count
print("Length: \(length ?? 0)")
```

----------

## Control Flow - If/Else
Conditional statements
```swift
let age = 25

if age >= 18 {
    print("Adult")
} else {
    print("Minor")
}

let score = 85
let grade: String

if score >= 90 {
    grade = "A"
} else if score >= 80 {
    grade = "B"
} else if score >= 70 {
    grade = "C"
} else {
    grade = "F"
}

print("Grade: \(grade)")

// Ternary operator
let status = age >= 18 ? "adult" : "minor"
print("Status: \(status)")
```

----------

## Loops
Different types of loops
```swift
// For-in loop with range
for i in 1...5 {
    print("Number: \(i)")
}

// For-in loop with array
let colors = ["red", "green", "blue"]
for color in colors {
    print("Color: \(color)")
}

// While loop
var count = 3
while count > 0 {
    print("Countdown: \(count)")
    count -= 1
}

// Repeat-while loop
var number = 1
repeat {
    print("Repeat: \(number)")
    number += 1
} while number <= 3
```

----------

## Switch Statements
Pattern matching with switch
```swift
let day = 3

switch day {
case 1:
    print("Monday")
case 2:
    print("Tuesday")
case 3:
    print("Wednesday")
case 4...7:
    print("Rest of week")
default:
    print("Invalid day")
}

let character = "a"
switch character {
case "a", "e", "i", "o", "u":
    print("Vowel")
case "b"..."z":
    print("Consonant")
default:
    print("Not a letter")
}
```

----------

## Functions
Function declaration and parameters
```swift
func greet(name: String) -> String {
    return "Hello, \(name)!"
}

func add(_ a: Int, _ b: Int) -> Int {
    return a + b
}

func divide(a: Double, by b: Double) -> Double? {
    guard b != 0 else { return nil }
    return a / b
}

// Function with multiple return values
func minMax(array: [Int]) -> (min: Int, max: Int)? {
    guard !array.isEmpty else { return nil }
    return (array.min()!, array.max()!)
}

print(greet(name: "Swift"))
print(add(5, 3))

if let result = divide(a: 10, by: 2) {
    print("Division result: \(result)")
}
```

----------

## Closures
Using closures and higher-order functions
```swift
let numbers = [1, 2, 3, 4, 5]

// Map
let doubled = numbers.map { $0 * 2 }
print("Doubled: \(doubled)")

// Filter
let evens = numbers.filter { $0 % 2 == 0 }
print("Evens: \(evens)")

// Reduce
let sum = numbers.reduce(0) { $0 + $1 }
print("Sum: \(sum)")

// Closure as parameter
func performOperation(_ numbers: [Int], operation: (Int) -> Int) -> [Int] {
    return numbers.map(operation)
}

let squared = performOperation(numbers) { $0 * $0 }
print("Squared: \(squared)")
```

----------

## Structs
Defining and using structs
```swift
struct Person {
    var name: String
    var age: Int
    
    init(name: String, age: Int) {
        self.name = name
        self.age = age
    }
    
    func greet() -> String {
        return "Hello, I'm \(name) and I'm \(age) years old"
    }
    
    mutating func haveBirthday() {
        age += 1
    }
}

var person = Person(name: "Alice", age: 30)
print(person.greet())

person.haveBirthday()
print("After birthday: \(person.age)")
```

----------

## Classes
Defining and using classes
```swift
class Vehicle {
    var brand: String
    var year: Int
    
    init(brand: String, year: Int) {
        self.brand = brand
        self.year = year
    }
    
    func start() {
        print("\(brand) is starting...")
    }
}

class Car: Vehicle {
    var doors: Int
    
    init(brand: String, year: Int, doors: Int) {
        self.doors = doors
        super.init(brand: brand, year: year)
    }
    
    override func start() {
        print("Car \(brand) with \(doors) doors is starting...")
    }
}

let car = Car(brand: "Toyota", year: 2022, doors: 4)
car.start()
```

----------

## Enums
Defining and using enums
```swift
enum Direction {
    case north, south, east, west
}

enum Planet: Int {
    case mercury = 1, venus, earth, mars
}

enum Result {
    case success(String)
    case failure(Error)
}

let direction = Direction.north
print("Going \(direction)")

switch direction {
case .north:
    print("Heading north")
case .south:
    print("Heading south")
case .east, .west:
    print("Heading east or west")
}

// Enum with associated values
let result = Result.success("Operation completed")
switch result {
case .success(let message):
    print("Success: \(message)")
case .failure(let error):
    print("Error: \(error)")
}
```

----------

## Protocols
Defining and implementing protocols
```swift
protocol Drawable {
    func draw()
}

protocol Colorable {
    var color: String { get set }
}

struct Circle: Drawable, Colorable {
    var radius: Double
    var color: String
    
    func draw() {
        print("Drawing a \(color) circle with radius \(radius)")
    }
}

struct Rectangle: Drawable {
    var width: Double
    var height: Double
    
    func draw() {
        print("Drawing rectangle \(width)x\(height)")
    }
}

let circle = Circle(radius: 5.0, color: "red")
let rectangle = Rectangle(width: 10, height: 20)

circle.draw()
rectangle.draw()
```

----------

## Extensions
Extending existing types
```swift
extension String {
    func reversed() -> String {
        return String(self.reversed())
    }
    
    var wordCount: Int {
        return self.components(separatedBy: .whitespacesAndNewlines)
            .filter { !$0.isEmpty }.count
    }
}

extension Int {
    func squared() -> Int {
        return self * self
    }
    
    var isEven: Bool {
        return self % 2 == 0
    }
}

let text = "Hello Swift World"
print("Reversed: \(text.reversed())")
print("Word count: \(text.wordCount)")

let number = 5
print("\(number) squared: \(number.squared())")
print("\(number) is even: \(number.isEven)")
```

----------

## Error Handling
Working with errors and try-catch
```swift
enum ValidationError: Error {
    case tooShort
    case tooLong
    case invalidCharacters
}

func validatePassword(_ password: String) throws -> Bool {
    if password.count < 6 {
        throw ValidationError.tooShort
    }
    if password.count > 20 {
        throw ValidationError.tooLong
    }
    return true
}

// Using do-catch
do {
    try validatePassword("abc")
    print("Password is valid")
} catch ValidationError.tooShort {
    print("Password is too short")
} catch ValidationError.tooLong {
    print("Password is too long")
} catch {
    print("Unknown error: \(error)")
}

// Using try?
if let isValid = try? validatePassword("validpassword") {
    print("Password validation result: \(isValid)")
}
```

----------

## Generics
Using generic types and functions
```swift
func swapValues<T>(_ a: inout T, _ b: inout T) {
    let temp = a
    a = b
    b = temp
}

struct Stack<Element> {
    private var items: [Element] = []
    
    mutating func push(_ item: Element) {
        items.append(item)
    }
    
    mutating func pop() -> Element? {
        return items.popLast()
    }
    
    var isEmpty: Bool {
        return items.isEmpty
    }
}

var a = 5, b = 10
swapValues(&a, &b)
print("a: \(a), b: \(b)")

var stringStack = Stack<String>()
stringStack.push("first")
stringStack.push("second")
print("Popped: \(stringStack.pop() ?? "empty")")
```

----------

## Property Observers
Observing property changes
```swift
class TemperatureMonitor {
    var temperature: Double = 0.0 {
        willSet {
            print("Temperature will change from \(temperature) to \(newValue)")
        }
        didSet {
            if temperature > oldValue {
                print("Temperature increased by \(temperature - oldValue)")
            } else if temperature < oldValue {
                print("Temperature decreased by \(oldValue - temperature)")
            }
        }
    }
}

let monitor = TemperatureMonitor()
monitor.temperature = 25.0
monitor.temperature = 30.0
monitor.temperature = 20.0
```

----------

## Computed Properties
Properties with custom getters and setters
```swift
struct Rectangle {
    var width: Double
    var height: Double
    
    var area: Double {
        return width * height
    }
    
    var perimeter: Double {
        get {
            return 2 * (width + height)
        }
    }
    
    var center: (x: Double, y: Double) {
        get {
            return (width / 2, height / 2)
        }
        set {
            width = newValue.x * 2
            height = newValue.y * 2
        }
    }
}

var rect = Rectangle(width: 10, height: 20)
print("Area: \(rect.area)")
print("Perimeter: \(rect.perimeter)")
print("Center: \(rect.center)")

rect.center = (x: 15, y: 25)
print("New dimensions: \(rect.width) x \(rect.height)")
```

----------

## Lazy Properties
Properties that are computed only when needed
```swift
class DataManager {
    lazy var expensiveResource: String = {
        print("Computing expensive resource...")
        return "Computed value"
    }()
    
    lazy var configuration: [String: Any] = {
        return [
            "apiURL": "https://api.example.com",
            "timeout": 30,
            "retryCount": 3
        ]
    }()
}

let manager = DataManager()
print("Manager created")

// Resource is computed only when first accessed
print("Resource: \(manager.expensiveResource)")
print("Config: \(manager.configuration)")
```

----------

## String Manipulation
Common string operations
```swift
let text = "  Hello, Swift World!  "

print("Original: '\(text)'")
print("Trimmed: '\(text.trimmingCharacters(in: .whitespaces))'")
print("Uppercase: \(text.uppercased())")
print("Lowercase: \(text.lowercased())")
print("Contains 'Swift': \(text.contains("Swift"))")

// String interpolation
let name = "Alice"
let age = 30
let greeting = "Hello, \(name)! You are \(age) years old."
print(greeting)

// String components
let words = text.trimmingCharacters(in: .whitespaces)
    .components(separatedBy: " ")
print("Words: \(words)")

// String replacement
let replaced = text.replacingOccurrences(of: "Swift", with: "iOS")
print("Replaced: \(replaced)")
```

----------

## Date and Time
Working with dates and times
```swift
import Foundation

let now = Date()
print("Current date: \(now)")

let formatter = DateFormatter()
formatter.dateStyle = .medium
formatter.timeStyle = .short
print("Formatted: \(formatter.string(from: now))")

// Custom date format
formatter.dateFormat = "yyyy-MM-dd HH:mm:ss"
print("Custom format: \(formatter.string(from: now))")

// Date arithmetic
let calendar = Calendar.current
let tomorrow = calendar.date(byAdding: .day, value: 1, to: now)!
let oneHourLater = calendar.date(byAdding: .hour, value: 1, to: now)!

print("Tomorrow: \(formatter.string(from: tomorrow))")
print("One hour later: \(formatter.string(from: oneHourLater))")
```

----------

## File I/O
Reading and writing files
```swift
import Foundation

let content = "Hello, Swift!"
let filename = "test.txt"

// Write to file
do {
    try content.write(toFile: filename, atomically: true, encoding: .utf8)
    print("File written successfully")
} catch {
    print("Write error: \(error)")
}

// Read from file
do {
    let readContent = try String(contentsOfFile: filename, encoding: .utf8)
    print("File content: \(readContent)")
} catch {
    print("Read error: \(error)")
}

// Working with URLs
if let documentsPath = FileManager.default.urls(for: .documentDirectory, 
                                                in: .userDomainMask).first {
    let fileURL = documentsPath.appendingPathComponent("data.txt")
    print("File URL: \(fileURL)")
}
```

----------

## JSON Handling
Encoding and decoding JSON
```swift
import Foundation

struct Person: Codable {
    let name: String
    let age: Int
    let email: String?
}

let person = Person(name: "Alice", age: 30, email: "alice@example.com")

// Encode to JSON
do {
    let jsonData = try JSONEncoder().encode(person)
    let jsonString = String(data: jsonData, encoding: .utf8)!
    print("JSON: \(jsonString)")
    
    // Decode from JSON
    let decodedPerson = try JSONDecoder().decode(Person.self, from: jsonData)
    print("Decoded: \(decodedPerson.name), \(decodedPerson.age)")
} catch {
    print("JSON error: \(error)")
}

// Working with dictionaries
let dict: [String: Any] = ["name": "Bob", "age": 25]
if let jsonData = try? JSONSerialization.data(withJSONObject: dict) {
    print("Dict JSON: \(String(data: jsonData, encoding: .utf8)!)")
}
```

----------

## Networking
Making HTTP requests
```swift
import Foundation

func fetchData(from urlString: String, completion: @escaping (Data?, Error?) -> Void) {
    guard let url = URL(string: urlString) else {
        completion(nil, NSError(domain: "Invalid URL", code: 0))
        return
    }
    
    let task = URLSession.shared.dataTask(with: url) { data, response, error in
        completion(data, error)
    }
    task.resume()
}

// GET request
fetchData(from: "https://httpbin.org/json") { data, error in
    if let error = error {
        print("Error: \(error)")
        return
    }
    
    if let data = data, let string = String(data: data, encoding: .utf8) {
        print("Response: \(string)")
    }
}

// POST request
var request = URLRequest(url: URL(string: "https://httpbin.org/post")!)
request.httpMethod = "POST"
request.setValue("application/json", forHTTPHeaderField: "Content-Type")

let postData = ["name": "John", "age": 30]
request.httpBody = try? JSONSerialization.data(withJSONObject: postData)
```

----------

## Grand Central Dispatch
Concurrency with GCD
```swift
import Foundation

// Background queue
DispatchQueue.global(qos: .background).async {
    print("Background task started")
    Thread.sleep(forTimeInterval: 2)
    
    // Update UI on main queue
    DispatchQueue.main.async {
        print("UI update on main queue")
    }
}

// Concurrent execution
let concurrentQueue = DispatchQueue(label: "concurrent", attributes: .concurrent)

for i in 1...5 {
    concurrentQueue.async {
        print("Task \(i) on concurrent queue")
    }
}

// Serial queue
let serialQueue = DispatchQueue(label: "serial")
for i in 1...3 {
    serialQueue.async {
        print("Serial task \(i)")
    }
}

// Dispatch group
let group = DispatchGroup()
group.enter()
DispatchQueue.global().async {
    Thread.sleep(forTimeInterval: 1)
    print("Group task 1 completed")
    group.leave()
}

group.notify(queue: .main) {
    print("All group tasks completed")
}
```

----------

## Testing
Unit testing with XCTest
```swift
import XCTest

class Calculator {
    func add(_ a: Int, _ b: Int) -> Int {
        return a + b
    }
    
    func divide(_ a: Double, _ b: Double) -> Double? {
        guard b != 0 else { return nil }
        return a / b
    }
}

class CalculatorTests: XCTestCase {
    var calculator: Calculator!
    
    override func setUp() {
        super.setUp()
        calculator = Calculator()
    }
    
    override func tearDown() {
        calculator = nil
        super.tearDown()
    }
    
    func testAddition() {
        let result = calculator.add(2, 3)
        XCTAssertEqual(result, 5)
    }
    
    func testDivision() {
        let result = calculator.divide(10, 2)
        XCTAssertEqual(result, 5.0)
    }
    
    func testDivisionByZero() {
        let result = calculator.divide(10, 0)
        XCTAssertNil(result)
    }
}
```

----------

## Memory Management
Working with weak and strong references
```swift
class Parent {
    var name: String
    var children: [Child] = []
    
    init(name: String) {
        self.name = name
    }
    
    deinit {
        print("Parent \(name) is being deallocated")
    }
}

class Child {
    var name: String
    weak var parent: Parent? // Weak reference to avoid retain cycle
    
    init(name: String, parent: Parent) {
        self.name = name
        self.parent = parent
    }
    
    deinit {
        print("Child \(name) is being deallocated")
    }
}

// Usage
var parent: Parent? = Parent(name: "John")
let child = Child(name: "Alice", parent: parent!)
parent?.children.append(child)

parent = nil // Parent will be deallocated due to weak reference
```

----------

## Property Wrappers
Custom property wrappers
```swift
@propertyWrapper
struct Clamped {
    private var value: Int
    private let range: ClosedRange<Int>
    
    init(wrappedValue: Int, _ range: ClosedRange<Int>) {
        self.range = range
        self.value = min(max(wrappedValue, range.lowerBound), range.upperBound)
    }
    
    var wrappedValue: Int {
        get { value }
        set { value = min(max(newValue, range.lowerBound), range.upperBound) }
    }
}

struct GameStats {
    @Clamped(0...100) var health: Int = 100
    @Clamped(0...10) var level: Int = 1
}

var stats = GameStats()
stats.health = 150 // Clamped to 100
stats.level = -5   // Clamped to 0

print("Health: \(stats.health), Level: \(stats.level)")
```