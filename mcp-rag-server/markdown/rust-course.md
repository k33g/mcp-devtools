# Complete Rust Programming Course

## Table of Contents

1. [Introduction to Rust](#introduction-to-rust)
2. [Getting Started](#getting-started)
3. [Basic Syntax](#basic-syntax)
4. [Variables and Data Types](#variables-and-data-types)
5. [Control Flow](#control-flow)
6. [Functions](#functions)
7. [Ownership and Borrowing](#ownership-and-borrowing)
8. [Structs and Enums](#structs-and-enums)
9. [Pattern Matching](#pattern-matching)
10. [Error Handling](#error-handling)
11. [Collections](#collections)
12. [Iterators and Closures](#iterators-and-closures)
13. [Modules and Packages](#modules-and-packages)
14. [Lifetimes](#lifetimes)
15. [Traits](#traits)
16. [Generics](#generics)
17. [Smart Pointers](#smart-pointers)
18. [Concurrency](#concurrency)
19. [Async Programming](#async-programming)
20. [Testing](#testing)
21. [Project Structure](#project-structure)
22. [Advanced Topics](#advanced-topics)

---

## Introduction to Rust

Rust is a systems programming language that focuses on safety, speed, and concurrency. It prevents common programming errors like null pointer dereferences, buffer overflows, and memory leaks without requiring a garbage collector.

### Key Features:
- **Memory safety** without garbage collection
- **Zero-cost abstractions**
- **Thread safety**
- **Pattern matching**
- **Type inference**
- **Minimal runtime**

### Use Cases:
- System programming
- Web backends
- Command-line tools
- WebAssembly applications
- Blockchain and cryptocurrency
- Game engines
- Operating systems

---

## Getting Started

### Installation

Install Rust using rustup:
```bash
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
```

Update Rust:
```bash
rustup update
```

### Creating Your First Project

```bash
cargo new hello_world
cd hello_world
cargo run
```

### Cargo Commands
- `cargo new project_name` - Create new project
- `cargo build` - Build project
- `cargo run` - Build and run project
- `cargo check` - Check code without building
- `cargo test` - Run tests
- `cargo doc` - Generate documentation

---

## Basic Syntax

### Hello World

```rust
fn main() {
    println!("Hello, world!");
}
```

### Comments

```rust
// Single line comment

/*
Multi-line
comment
*/

/// Documentation comment for the next item
fn documented_function() {}

//! Documentation comment for the containing item
```

### Printing

```rust
fn main() {
    println!("Hello, world!");
    println!("The number is {}", 42);
    println!("Multiple values: {} and {}", "first", "second");
    println!("Named parameters: {name} is {age} years old", name="Alice", age=30);
    
    // Debug printing
    let x = vec![1, 2, 3];
    println!("{:?}", x);  // Debug format
    println!("{:#?}", x); // Pretty debug format
}
```

---

## Variables and Data Types

### Variables

```rust
fn main() {
    // Immutable by default
    let x = 5;
    // x = 6; // This would cause an error
    
    // Mutable variables
    let mut y = 5;
    y = 6; // This is fine
    
    // Constants
    const MAX_POINTS: u32 = 100_000;
    
    // Shadowing
    let x = x + 1; // New variable shadows the previous one
    let x = x * 2; // x is now 12
}
```

### Scalar Types

```rust
fn main() {
    // Integers
    let a: i8 = -128;      // 8-bit signed
    let b: u8 = 255;       // 8-bit unsigned
    let c: i32 = -2000;    // 32-bit signed (default)
    let d: u32 = 3000;     // 32-bit unsigned
    let e: i64 = -40000;   // 64-bit signed
    let f: u64 = 50000;    // 64-bit unsigned
    let g: isize = -100;   // pointer-sized signed
    let h: usize = 200;    // pointer-sized unsigned
    
    // Floating point
    let x = 2.0;      // f64 (default)
    let y: f32 = 3.0; // f32
    
    // Boolean
    let t = true;
    let f: bool = false;
    
    // Character
    let c = 'z';
    let emoji = '😎';
}
```

### Compound Types

```rust
fn main() {
    // Tuples
    let tup: (i32, f64, u8) = (500, 6.4, 1);
    let (x, y, z) = tup; // Destructuring
    let first = tup.0;   // Access by index
    
    // Arrays
    let a = [1, 2, 3, 4, 5];
    let a: [i32; 5] = [1, 2, 3, 4, 5];
    let a = [3; 5]; // [3, 3, 3, 3, 3]
    let first = a[0];
    
    // Slices
    let slice = &a[1..3]; // [2, 3]
}
```

### Strings

```rust
fn main() {
    // String literals (str)
    let s1 = "Hello, world!";
    
    // String type (growable)
    let mut s2 = String::new();
    let s3 = String::from("Hello");
    let s4 = "Hello".to_string();
    
    // String operations
    s2.push_str("Hello");
    s2.push('!');
    
    let combined = format!("{} {}", s3, s4);
}
```

---

## Control Flow

### if Expressions

```rust
fn main() {
    let number = 6;
    
    if number % 4 == 0 {
        println!("number is divisible by 4");
    } else if number % 3 == 0 {
        println!("number is divisible by 3");
    } else {
        println!("number is not divisible by 4 or 3");
    }
    
    // if as expression
    let condition = true;
    let number = if condition { 5 } else { 6 };
}
```

### Loops

```rust
fn main() {
    // loop
    let mut counter = 0;
    let result = loop {
        counter += 1;
        if counter == 10 {
            break counter * 2; // Return value from loop
        }
    };
    
    // while loop
    let mut number = 3;
    while number != 0 {
        println!("{}!", number);
        number -= 1;
    }
    
    // for loop
    let a = [10, 20, 30, 40, 50];
    for element in a {
        println!("the value is: {}", element);
    }
    
    // Range
    for number in (1..4).rev() {
        println!("{}!", number);
    }
}
```

---

## Functions

### Basic Functions

```rust
fn main() {
    another_function();
    function_with_parameter(5);
    let result = function_with_return(3, 4);
}

fn another_function() {
    println!("Another function.");
}

fn function_with_parameter(x: i32) {
    println!("The value of x is: {}", x);
}

fn function_with_return(x: i32, y: i32) -> i32 {
    x + y // No semicolon for expression
}

// Early return
fn early_return_example(x: i32) -> i32 {
    if x < 0 {
        return 0; // Early return with return keyword
    }
    x * 2
}
```

### Function Pointers

```rust
fn add(x: i32, y: i32) -> i32 {
    x + y
}

fn main() {
    let operation: fn(i32, i32) -> i32 = add;
    let result = operation(3, 4);
}
```

---

## Ownership and Borrowing

### Ownership Rules

1. Each value in Rust has a variable that's called its owner
2. There can only be one owner at a time
3. When the owner goes out of scope, the value will be dropped

```rust
fn main() {
    let s1 = String::from("hello");
    let s2 = s1; // s1 is moved to s2, s1 is no longer valid
    // println!("{}", s1); // This would cause an error
    println!("{}", s2); // This is fine
    
    // Clone for deep copy
    let s3 = s2.clone();
    println!("{} {}", s2, s3); // Both are valid
}
```

### References and Borrowing

```rust
fn main() {
    let s1 = String::from("hello");
    
    let len = calculate_length(&s1); // Borrow s1
    println!("The length of '{}' is {}.", s1, len); // s1 is still valid
    
    let mut s2 = String::from("hello");
    change(&mut s2); // Mutable borrow
    
    // Only one mutable reference at a time
    let r1 = &mut s2;
    // let r2 = &mut s2; // This would cause an error
    
    // Multiple immutable references are OK
    let s3 = String::from("hello");
    let r1 = &s3;
    let r2 = &s3;
    println!("{} and {}", r1, r2);
}

fn calculate_length(s: &String) -> usize {
    s.len()
}

fn change(s: &mut String) {
    s.push_str(", world");
}
```

### Slice References

```rust
fn main() {
    let s = String::from("hello world");
    let word = first_word(&s);
    println!("First word: {}", word);
    
    let a = [1, 2, 3, 4, 5];
    let slice = &a[1..3];
    assert_eq!(slice, &[2, 3]);
}

fn first_word(s: &str) -> &str {
    let bytes = s.as_bytes();
    
    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return &s[0..i];
        }
    }
    
    &s[..]
}
```

---

## Structs and Enums

### Structs

```rust
// Classic struct
struct User {
    username: String,
    email: String,
    sign_in_count: u64,
    active: bool,
}

// Tuple struct
struct Color(i32, i32, i32);
struct Point(i32, i32, i32);

// Unit struct
struct Unit;

impl User {
    // Associated function (constructor)
    fn new(username: String, email: String) -> User {
        User {
            username,
            email,
            active: true,
            sign_in_count: 1,
        }
    }
    
    // Method
    fn deactivate(&mut self) {
        self.active = false;
    }
    
    // Method that takes ownership
    fn into_email(self) -> String {
        self.email
    }
}

fn main() {
    let mut user1 = User::new(
        String::from("someusername123"),
        String::from("someone@example.com"),
    );
    
    user1.deactivate();
    
    // Struct update syntax
    let user2 = User {
        email: String::from("another@example.com"),
        username: String::from("anotherusername567"),
        ..user1 // Use remaining fields from user1
    };
    
    // Tuple structs
    let black = Color(0, 0, 0);
    let origin = Point(0, 0, 0);
}
```

### Enums

```rust
enum IpAddrKind {
    V4,
    V6,
}

enum IpAddr {
    V4(u8, u8, u8, u8),
    V6(String),
}

enum Message {
    Quit,
    Move { x: i32, y: i32 },
    Write(String),
    ChangeColor(i32, i32, i32),
}

impl Message {
    fn call(&self) {
        match self {
            Message::Quit => println!("Quit"),
            Message::Move { x, y } => println!("Move to ({}, {})", x, y),
            Message::Write(text) => println!("Write: {}", text),
            Message::ChangeColor(r, g, b) => println!("Change color to ({}, {}, {})", r, g, b),
        }
    }
}

fn main() {
    let home = IpAddr::V4(127, 0, 0, 1);
    let loopback = IpAddr::V6(String::from("::1"));
    
    let msg = Message::Write(String::from("hello"));
    msg.call();
}
```

### Option Enum

```rust
fn main() {
    let some_number = Some(5);
    let some_string = Some("a string");
    let absent_number: Option<i32> = None;
    
    // Using Option
    let x: i8 = 5;
    let y: Option<i8> = Some(5);
    
    // let sum = x + y; // This won't work
    let sum = x + y.unwrap_or(0); // This works
    
    // Safe handling
    match y {
        None => println!("No value"),
        Some(i) => println!("Value: {}", i),
    }
    
    // Using if let
    if let Some(value) = y {
        println!("Got a value: {}", value);
    }
}
```

---

## Pattern Matching

### match Expression

```rust
enum Coin {
    Penny,
    Nickel,
    Dime,
    Quarter(UsState),
}

enum UsState {
    Alabama,
    Alaska,
    // ... etc
}

fn value_in_cents(coin: Coin) -> u8 {
    match coin {
        Coin::Penny => {
            println!("Lucky penny!");
            1
        }
        Coin::Nickel => 5,
        Coin::Dime => 10,
        Coin::Quarter(state) => {
            println!("State quarter from {:?}!", state);
            25
        }
    }
}

fn main() {
    // Matching with Option
    let x: i8 = 5;
    let y: Option<i8> = Some(5);
    
    match y {
        None => None,
        Some(i) => Some(i + 1),
    };
    
    // Catch-all patterns
    let dice_roll = 9;
    match dice_roll {
        3 => add_fancy_hat(),
        7 => remove_fancy_hat(),
        _ => (), // Do nothing for other values
    }
    
    // Using variables in patterns
    match dice_roll {
        3 => add_fancy_hat(),
        7 => remove_fancy_hat(),
        other => move_player(other),
    }
}

fn add_fancy_hat() {}
fn remove_fancy_hat() {}
fn move_player(num_spaces: u8) {}
```

### if let

```rust
fn main() {
    let some_u8_value = Some(3u8);
    
    // Instead of this:
    match some_u8_value {
        Some(3) => println!("three"),
        _ => (),
    }
    
    // You can write this:
    if let Some(3) = some_u8_value {
        println!("three");
    }
    
    // With else
    if let Some(x) = some_u8_value {
        println!("Value: {}", x);
    } else {
        println!("No value");
    }
}
```

### Advanced Patterns

```rust
fn main() {
    // Destructuring structs
    struct Point {
        x: i32,
        y: i32,
    }
    
    let p = Point { x: 0, y: 7 };
    
    match p {
        Point { x, y: 0 } => println!("On the x axis at {}", x),
        Point { x: 0, y } => println!("On the y axis at {}", y),
        Point { x, y } => println!("On neither axis: ({}, {})", x, y),
    }
    
    // Range patterns
    let x = 5;
    match x {
        1..=5 => println!("one through five"),
        _ => println!("something else"),
    }
    
    // Guards
    let num = Some(4);
    match num {
        Some(x) if x < 5 => println!("less than five: {}", x),
        Some(x) => println!("{}", x),
        None => (),
    }
    
    // @ bindings
    enum Message {
        Hello { id: i32 },
    }
    
    let msg = Message::Hello { id: 5 };
    match msg {
        Message::Hello { id: id_variable @ 3..=7 } => {
            println!("Found an id in range: {}", id_variable)
        }
        Message::Hello { id: 10..=12 } => {
            println!("Found an id in another range")
        }
        Message::Hello { id } => {
            println!("Found some other id: {}", id)
        }
    }
}
```

---

## Error Handling

### panic! Macro

```rust
fn main() {
    // Explicit panic
    panic!("crash and burn");
    
    // Panic from invalid array access
    let v = vec![1, 2, 3];
    v[99]; // This will panic
}
```

### Result Type

```rust
use std::fs::File;
use std::io::ErrorKind;

fn main() {
    // Basic Result handling
    let f = File::open("hello.txt");
    
    let f = match f {
        Ok(file) => file,
        Err(error) => {
            panic!("Problem opening the file: {:?}", error)
        },
    };
    
    // Matching on different errors
    let f = File::open("hello.txt");
    
    let f = match f {
        Ok(file) => file,
        Err(error) => match error.kind() {
            ErrorKind::NotFound => match File::create("hello.txt") {
                Ok(fc) => fc,
                Err(e) => panic!("Problem creating the file: {:?}", e),
            },
            other_error => {
                panic!("Problem opening the file: {:?}", other_error)
            }
        },
    };
}
```

### Shortcuts: unwrap and expect

```rust
use std::fs::File;

fn main() {
    // unwrap: panic on error
    let f = File::open("hello.txt").unwrap();
    
    // expect: panic with custom message
    let f = File::open("hello.txt")
        .expect("Failed to open hello.txt");
}
```

### Propagating Errors

```rust
use std::fs::File;
use std::io::{self, Read};

fn read_username_from_file() -> Result<String, io::Error> {
    let f = File::open("hello.txt");
    
    let mut f = match f {
        Ok(file) => file,
        Err(e) => return Err(e),
    };
    
    let mut s = String::new();
    
    match f.read_to_string(&mut s) {
        Ok(_) => Ok(s),
        Err(e) => Err(e),
    }
}

// Shortcut with ? operator
fn read_username_from_file_short() -> Result<String, io::Error> {
    let mut f = File::open("hello.txt")?;
    let mut s = String::new();
    f.read_to_string(&mut s)?;
    Ok(s)
}

// Even shorter
fn read_username_from_file_shortest() -> Result<String, io::Error> {
    let mut s = String::new();
    File::open("hello.txt")?.read_to_string(&mut s)?;
    Ok(s)
}

fn main() {
    match read_username_from_file() {
        Ok(username) => println!("Username: {}", username),
        Err(e) => println!("Error: {}", e),
    }
}
```

### Custom Error Types

```rust
use std::fmt;

#[derive(Debug)]
enum MyError {
    Io(std::io::Error),
    Parse(std::num::ParseIntError),
    Custom(String),
}

impl fmt::Display for MyError {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        match self {
            MyError::Io(err) => write!(f, "IO error: {}", err),
            MyError::Parse(err) => write!(f, "Parse error: {}", err),
            MyError::Custom(msg) => write!(f, "Custom error: {}", msg),
        }
    }
}

impl std::error::Error for MyError {}

impl From<std::io::Error> for MyError {
    fn from(error: std::io::Error) -> Self {
        MyError::Io(error)
    }
}

impl From<std::num::ParseIntError> for MyError {
    fn from(error: std::num::ParseIntError) -> Self {
        MyError::Parse(error)
    }
}

fn do_something() -> Result<i32, MyError> {
    let contents = std::fs::read_to_string("number.txt")?;
    let number: i32 = contents.trim().parse()?;
    Ok(number * 2)
}
```

---

## Collections

### Vectors

```rust
fn main() {
    // Creating vectors
    let mut v: Vec<i32> = Vec::new();
    let v = vec![1, 2, 3]; // vec! macro
    
    // Adding elements
    let mut v = Vec::new();
    v.push(5);
    v.push(6);
    v.push(7);
    v.push(8);
    
    // Reading elements
    let third: &i32 = &v[2];
    println!("The third element is {}", third);
    
    match v.get(2) {
        Some(third) => println!("The third element is {}", third),
        None => println!("There is no third element."),
    }
    
    // Iterating
    let v = vec![100, 32, 57];
    for i in &v {
        println!("{}", i);
    }
    
    // Mutable iteration
    let mut v = vec![100, 32, 57];
    for i in &mut v {
        *i += 50;
    }
    
    // Storing different types with enums
    enum SpreadsheetCell {
        Int(i32),
        Float(f64),
        Text(String),
    }
    
    let row = vec![
        SpreadsheetCell::Int(3),
        SpreadsheetCell::Text(String::from("blue")),
        SpreadsheetCell::Float(10.12),
    ];
}
```

### HashMaps

```rust
use std::collections::HashMap;

fn main() {
    // Creating hash maps
    let mut scores = HashMap::new();
    scores.insert(String::from("Blue"), 10);
    scores.insert(String::from("Yellow"), 50);
    
    // From vectors
    let teams = vec![String::from("Blue"), String::from("Yellow")];
    let initial_scores = vec![10, 50];
    let scores: HashMap<_, _> = teams.into_iter().zip(initial_scores.into_iter()).collect();
    
    // Accessing values
    let team_name = String::from("Blue");
    let score = scores.get(&team_name);
    
    match score {
        Some(s) => println!("Score: {}", s),
        None => println!("Team not found"),
    }
    
    // Iterating
    for (key, value) in &scores {
        println!("{}: {}", key, value);
    }
    
    // Updating
    let mut scores = HashMap::new();
    scores.insert(String::from("Blue"), 10);
    
    // Overwriting
    scores.insert(String::from("Blue"), 25);
    
    // Only insert if key doesn't exist
    scores.entry(String::from("Yellow")).or_insert(50);
    scores.entry(String::from("Blue")).or_insert(50); // Won't overwrite
    
    // Update based on old value
    let text = "hello world wonderful world";
    let mut map = HashMap::new();
    
    for word in text.split_whitespace() {
        let count = map.entry(word).or_insert(0);
        *count += 1;
    }
    
    println!("{:?}", map);
}
```

### Other Collections

```rust
use std::collections::{VecDeque, HashSet, BTreeMap, BTreeSet};

fn main() {
    // VecDeque (double-ended queue)
    let mut deque = VecDeque::new();
    deque.push_back(1);
    deque.push_front(2);
    
    // HashSet
    let mut books = HashSet::new();
    books.insert("A Dance With Dragons".to_string());
    books.insert("To Kill a Mockingbird".to_string());
    
    if !books.contains("The Winds of Winter") {
        println!("We have {} books, but The Winds of Winter ain't one.", books.len());
    }
    
    // BTreeMap (sorted map)
    let mut movie_reviews = BTreeMap::new();
    movie_reviews.insert("Office Space", 4);
    movie_reviews.insert("Pulp Fiction", 5);
    
    // BTreeSet (sorted set)
    let mut set = BTreeSet::new();
    set.insert(3);
    set.insert(1);
    set.insert(2);
    // Iteration will be in sorted order
}
```

---

## Iterators and Closures

### Closures

```rust
fn main() {
    // Basic closure
    let expensive_closure = |num| {
        println!("calculating slowly...");
        thread::sleep(Duration::from_secs(2));
        num
    };
    
    // Type annotations (optional)
    let add_one = |x: i32| -> i32 { x + 1 };
    let add_one = |x| x + 1; // Type inferred
    
    // Capturing environment
    let x = vec![1, 2, 3];
    let equal_to_x = move |z| z == x; // move keyword transfers ownership
    // println!("can't use x here: {:?}", x); // x is moved
    
    let y = vec![1, 2, 3];
    assert!(equal_to_x(y));
    
    // Closure traits
    // FnOnce - takes ownership of captured variables
    // FnMut - can change the captured variables
    // Fn - borrows values immutably
}

use std::thread;
use std::time::Duration;

// Using closures as parameters
fn simulated_expensive_calculation(intensity: u32) -> u32 {
    println!("calculating slowly...");
    thread::sleep(Duration::from_secs(2));
    intensity
}

struct Cacher<T>
where
    T: Fn(u32) -> u32,
{
    calculation: T,
    value: Option<u32>,
}

impl<T> Cacher<T>
where
    T: Fn(u32) -> u32,
{
    fn new(calculation: T) -> Cacher<T> {
        Cacher {
            calculation,
            value: None,
        }
    }
    
    fn value(&mut self, arg: u32) -> u32 {
        match self.value {
            Some(v) => v,
            None => {
                let v = (self.calculation)(arg);
                self.value = Some(v);
                v
            },
        }
    }
}
```

### Iterators

```rust
fn main() {
    // Creating iterators
    let v1 = vec![1, 2, 3];
    let v1_iter = v1.iter(); // Iterator over references
    let v1_iter = v1.into_iter(); // Iterator that takes ownership
    let mut v1 = vec![1, 2, 3];
    let v1_iter = v1.iter_mut(); // Iterator over mutable references
    
    // Using iterators
    let v1 = vec![1, 2, 3];
    for val in v1.iter() {
        println!("Got: {}", val);
    }
    
    // Iterator adaptors
    let v1: Vec<i32> = vec![1, 2, 3];
    let v2: Vec<_> = v1.iter().map(|x| x + 1).collect();
    
    // Consuming adaptors
    let v1 = vec![1, 2, 3];
    let total: i32 = v1.iter().sum();
    
    // More iterator methods
    let v1: Vec<i32> = vec![1, 2, 3];
    let v2: Vec<i32> = v1
        .iter()
        .filter(|&x| *x > 1)
        .map(|x| x * 2)
        .collect();
    
    // find
    let v = vec![1, 2, 3, 4, 5];
    let found = v.iter().find(|&&x| x > 3);
    
    // any and all
    let v = vec![2, 4, 6, 8];
    let all_even = v.iter().all(|&x| x % 2 == 0);
    let any_greater_than_5 = v.iter().any(|&x| x > 5);
    
    // enumerate
    for (index, value) in v.iter().enumerate() {
        println!("{}: {}", index, value);
    }
    
    // zip
    let names = vec!["Alice", "Bob", "Charlie"];
    let ages = vec![30, 25, 35];
    for (name, age) in names.iter().zip(ages.iter()) {
        println!("{} is {} years old", name, age);
    }
}

// Custom iterator
struct Counter {
    current: usize,
    max: usize,
}

impl Counter {
    fn new(max: usize) -> Counter {
        Counter { current: 0, max }
    }
}

impl Iterator for Counter {
    type Item = usize;
    
    fn next(&mut self) -> Option<Self::Item> {
        if self.current < self.max {
            let current = self.current;
            self.current += 1;
            Some(current)
        } else {
            None
        }
    }
}

fn using_custom_iterator() {
    let counter = Counter::new(3);
    for n in counter {
        println!("{}", n);
    }
}
```

---

## Modules and Packages

### Module System

```rust
// src/lib.rs
mod front_of_house {
    pub mod hosting {
        pub fn add_to_waitlist() {}
        
        fn seat_at_table() {} // private
    }
    
    mod serving {
        fn take_order() {}
        fn serve_order() {}
        fn take_payment() {}
    }
}

pub fn eat_at_restaurant() {
    // Absolute path
    crate::front_of_house::hosting::add_to_waitlist();
    
    // Relative path
    front_of_house::hosting::add_to_waitlist();
}

// Using super
fn serve_order() {}

mod back_of_house {
    fn fix_incorrect_order() {
        cook_order();
        super::serve_order(); // Call function in parent module
    }
    
    fn cook_order() {}
    
    // Public struct with private fields
    pub struct Breakfast {
        pub toast: String,
        seasonal_fruit: String, // private
    }
    
    impl Breakfast {
        pub fn summer(toast: &str) -> Breakfast {
            Breakfast {
                toast: String::from(toast),
                seasonal_fruit: String::from("peaches"),
            }
        }
    }
    
    // Public enum (all variants are public)
    pub enum Appetizer {
        Soup,
        Salad,
    }
}

// Using use
use crate::front_of_house::hosting;

pub fn eat_at_restaurant() {
    hosting::add_to_waitlist();
}

// Re-exporting
pub use crate::front_of_house::hosting;

// Using external packages
use std::collections::HashMap;
use std::fmt::Result;
use std::io::Result as IoResult; // Alias

// Nested paths
use std::{cmp::Ordering, io};
use std::io::{self, Write}; // Import io and io::Write

// Glob operator
use std::collections::*;
```

### Separating Modules into Files

```rust
// src/lib.rs
mod front_of_house;

pub use crate::front_of_house::hosting;

pub fn eat_at_restaurant() {
    hosting::add_to_waitlist();
}

// src/front_of_house.rs
pub mod hosting;

// src/front_of_house/hosting.rs
pub fn add_to_waitlist() {}
```

### Package Structure

```
restaurant/
├── Cargo.toml
└── src/
    ├── lib.rs
    ├── main.rs
    └── front_of_house/
        ├── mod.rs
        └── hosting.rs
```

### Cargo.toml

```toml
[package]
name = "restaurant"
version = "0.1.0"
edition = "2021"

[dependencies]
rand = "0.8"
serde = { version = "1.0", features = ["derive"] }

[dev-dependencies]
tokio-test = "0.4"

[build-dependencies]
cc = "1.0"
```

---

## Lifetimes

### Lifetime Basics

```rust
fn main() {
    let r;
    
    {
        let x = 5;
        r = &x; // Error: x doesn't live long enough
    }
    
    println!("r: {}", r);
}

// Lifetime annotations in function signatures
fn longest<'a>(x: &'a str, y: &'a str) -> &'a str {
    if x.len() > y.len() {
        x
    } else {
        y
    }
}

fn main() {
    let string1 = String::from("abcd");
    let string2 = "xyz";
    
    let result = longest(string1.as_str(), string2);
    println!("The longest string is {}", result);
}
```

### Lifetime Annotations in Structs

```rust
struct ImportantExcerpt<'a> {
    part: &'a str,
}

impl<'a> ImportantExcerpt<'a> {
    fn level(&self) -> i32 {
        3
    }
    
    fn announce_and_return_part(&self, announcement: &str) -> &str {
        println!("Attention please: {}", announcement);
        self.part
    }
}

fn main() {
    let novel = String::from("Call me Ishmael. Some years ago...");
    let first_sentence = novel.split('.').next().expect("Could not find a '.'");
    let i = ImportantExcerpt {
        part: first_sentence,
    };
}
```

### Lifetime Elision Rules

1. Each parameter that is a reference gets its own lifetime parameter
2. If there is exactly one input lifetime parameter, that lifetime is assigned to all output lifetime parameters
3. If there are multiple input lifetime parameters, but one of them is `&self` or `&mut self`, the lifetime of `self` is assigned to all output lifetime parameters

```rust
// These function signatures are equivalent due to lifetime elision:
fn first_word(s: &str) -> &str {
    // Compiler infers: fn first_word<'a>(s: &'a str) -> &'a str
    let bytes = s.as_bytes();
    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return &s[0..i];
        }
    }
    &s[..]
}
```

### Static Lifetime

```rust
fn main() {
    let s: &'static str = "I have a static lifetime.";
    // String literals always have 'static lifetime
}

// Generic lifetimes with traits
use std::fmt::Display;

fn longest_with_an_announcement<'a, T>(
    x: &'a str,
    y: &'a str,
    ann: T,
) -> &'a str
where
    T: Display,
{
    println!("Announcement! {}", ann);
    if x.len() > y.len() {
        x
    } else {
        y
    }
}
```

---

## Traits

### Defining and Implementing Traits

```rust
pub trait Summary {
    fn summarize(&self) -> String;
    
    // Default implementation
    fn summarize_default(&self) -> String {
        String::from("(Read more...)")
    }
    
    // Default implementation that calls other methods
    fn summarize_author(&self) -> String;
    
    fn summarize_with_author(&self) -> String {
        format!("(Read more from {}...)", self.summarize_author())
    }
}

pub struct NewsArticle {
    pub headline: String,
    pub location: String,
    pub author: String,
    pub content: String,
}

impl Summary for NewsArticle {
    fn summarize(&self) -> String {
        format!("{}, by {} ({})", self.headline, self.author, self.location)
    }
    
    fn summarize_author(&self) -> String {
        format!("{}", self.author)
    }
}

pub struct Tweet {
    pub username: String,
    pub content: String,
    pub reply: bool,
    pub retweet: bool,
}

impl Summary for Tweet {
    fn summarize(&self) -> String {
        format!("{}: {}", self.username, self.content)
    }
    
    fn summarize_author(&self) -> String {
        format!("@{}", self.username)
    }
}

fn main() {
    let tweet = Tweet {
        username: String::from("horse_ebooks"),
        content: String::from("of course, as you probably already know, people"),
        reply: false,
        retweet: false,
    };
    
    println!("1 new tweet: {}", tweet.summarize());
}
```

### Traits as Parameters

```rust
// Trait bound syntax
pub fn notify(item: &impl Summary) {
    println!("Breaking news! {}", item.summarize());
}

// Generic with trait bound
pub fn notify<T: Summary>(item: &T) {
    println!("Breaking news! {}", item.summarize());
}

// Multiple trait bounds
pub fn notify(item: &(impl Summary + Display)) {
    // ...
}

pub fn notify<T: Summary + Display>(item: &T) {
    // ...
}

// Where clauses for readability
fn some_function<T, U>(t: &T, u: &U) -> i32
where
    T: Display + Clone,
    U: Clone + Debug,
{
    // ...
}
```

### Returning Types that Implement Traits

```rust
fn returns_summarizable() -> impl Summary {
    Tweet {
        username: String::from("horse_ebooks"),
        content: String::from("of course, as you probably already know, people"),
        reply: false,
        retweet: false,
    }
}

// This won't work because we might return different types:
// fn returns_summarizable(switch: bool) -> impl Summary {
//     if switch {
//         NewsArticle { ... }
//     } else {
//         Tweet { ... }  // Error!
//     }
// }
```

### Conditional Implementation

```rust
use std::fmt::Display;

struct Pair<T> {
    x: T,
    y: T,
}

impl<T> Pair<T> {
    fn new(x: T, y: T) -> Self {
        Self { x, y }
    }
}

impl<T: Display + PartialOrd> Pair<T> {
    fn cmp_display(&self) {
        if self.x >= self.y {
            println!("The largest member is x = {}", self.x);
        } else {
            println!("The largest member is y = {}", self.y);
        }
    }
}

// Blanket implementations
impl<T: Display> ToString for T {
    // This implements ToString for any type that implements Display
}
```

### Important Standard Traits

```rust
// Clone trait
#[derive(Clone)]
struct Point {
    x: i32,
    y: i32,
}

// Copy trait (marker trait)
#[derive(Copy, Clone)]
struct Point {
    x: i32,
    y: i32,
}

// Debug trait
#[derive(Debug)]
struct Rectangle {
    width: u32,
    height: u32,
}

// PartialEq and Eq
#[derive(PartialEq)]
struct Person {
    name: String,
    age: u32,
}

// PartialOrd and Ord
#[derive(PartialOrd, PartialEq)]
struct Height(u32);

// Hash
use std::collections::HashMap;

#[derive(Hash, Eq, PartialEq)]
struct Person {
    name: String,
    age: u32,
}

fn main() {
    let mut map = HashMap::new();
    let person = Person {
        name: "Alice".to_string(),
        age: 30,
    };
    map.insert(person, "Engineer");
}
```

---

## Generics

### Generic Functions

```rust
fn largest<T: PartialOrd>(list: &[T]) -> &T {
    let mut largest = &list[0];
    
    for item in list {
        if item > largest {
            largest = item;
        }
    }
    
    largest
}

fn main() {
    let number_list = vec![34, 50, 25, 100, 65];
    let result = largest(&number_list);
    println!("The largest number is {}", result);
    
    let char_list = vec!['y', 'm', 'a', 'q'];
    let result = largest(&char_list);
    println!("The largest char is {}", result);
}
```

### Generic Structs

```rust
struct Point<T> {
    x: T,
    y: T,
}

impl<T> Point<T> {
    fn x(&self) -> &T {
        &self.x
    }
}

// Implementation for specific type
impl Point<f32> {
    fn distance_from_origin(&self) -> f32 {
        (self.x.powi(2) + self.y.powi(2)).sqrt()
    }
}

// Multiple generic parameters
struct Point<T, U> {
    x: T,
    y: U,
}

impl<T, U> Point<T, U> {
    fn mixup<V, W>(self, other: Point<V, W>) -> Point<T, W> {
        Point {
            x: self.x,
            y: other.y,
        }
    }
}

fn main() {
    let integer = Point { x: 5, y: 10 };
    let float = Point { x: 1.0, y: 4.0 };
    let mixed = Point { x: 5, y: 4.0 };
    
    let p1 = Point { x: 5, y: 10.4 };
    let p2 = Point { x: "Hello", y: 'c' };
    let p3 = p1.mixup(p2);
    
    println!("p3.x = {}, p3.y = {}", p3.x, p3.y);
}
```

### Generic Enums

```rust
enum Option<T> {
    Some(T),
    None,
}

enum Result<T, E> {
    Ok(T),
    Err(E),
}

// Custom generic enum
enum MyResult<T, E> {
    Success(T),
    Failure(E),
}
```

### Advanced Generic Concepts

```rust
// Associated types
pub trait Iterator {
    type Item; // Associated type
    
    fn next(&mut self) -> Option<Self::Item>;
}

struct Counter {
    current: usize,
    max: usize,
}

impl Iterator for Counter {
    type Item = usize;
    
    fn next(&mut self) -> Option<Self::Item> {
        if self.current < self.max {
            let current = self.current;
            self.current += 1;
            Some(current)
        } else {
            None
        }
    }
}

// Generic associated types (GATs)
trait IntoIterator {
    type Item;
    type IntoIter: Iterator<Item = Self::Item>;
    
    fn into_iter(self) -> Self::IntoIter;
}

// Higher-rank trait bounds (HRTB)
fn example<F>(f: F) 
where
    F: for<'a> Fn(&'a str) -> &'a str,
{
    // F works for any lifetime 'a
}

// Type aliases with generics
type Result<T> = std::result::Result<T, std::io::Error>;

fn read_file() -> Result<String> {
    std::fs::read_to_string("file.txt")
}
```

---

## Smart Pointers

### Box<T>

```rust
fn main() {
    // Basic Box usage
    let b = Box::new(5);
    println!("b = {}", b);
    
    // Box for recursive data structures
    enum List {
        Cons(i32, Box<List>),
        Nil,
    }
    
    use List::{Cons, Nil};
    
    let list = Cons(1, Box::new(Cons(2, Box::new(Cons(3, Box::new(Nil))))));
}

// Implementing Deref trait
use std::ops::Deref;

struct MyBox<T>(T);

impl<T> MyBox<T> {
    fn new(x: T) -> MyBox<T> {
        MyBox(x)
    }
}

impl<T> Deref for MyBox<T> {
    type Target = T;
    
    fn deref(&self) -> &Self::Target {
        &self.0
    }
}

fn hello(name: &str) {
    println!("Hello, {}!", name);
}

fn main() {
    let m = MyBox::new(String::from("Rust"));
    hello(&m); // Deref coercion: &MyBox<String> -> &String -> &str
}
```

### Rc<T> (Reference Counted)

```rust
use std::rc::Rc;

enum List {
    Cons(i32, Rc<List>),
    Nil,
}

use List::{Cons, Nil};

fn main() {
    let a = Rc::new(Cons(5, Rc::new(Cons(10, Rc::new(Nil)))));
    println!("count after creating a = {}", Rc::strong_count(&a));
    
    let b = Cons(3, Rc::clone(&a));
    println!("count after creating b = {}", Rc::strong_count(&a));
    
    {
        let c = Cons(4, Rc::clone(&a));
        println!("count after creating c = {}", Rc::strong_count(&a));
    }
    
    println!("count after c goes out of scope = {}", Rc::strong_count(&a));
}
```

### RefCell<T> and Interior Mutability

```rust
use std::cell::RefCell;
use std::rc::Rc;

#[derive(Debug)]
enum List {
    Cons(Rc<RefCell<i32>>, Rc<List>),
    Nil,
}

use List::{Cons, Nil};

fn main() {
    let value = Rc::new(RefCell::new(5));
    
    let a = Rc::new(Cons(Rc::clone(&value), Rc::new(Nil)));
    let b = Cons(Rc::new(RefCell::new(3)), Rc::clone(&a));
    let c = Cons(Rc::new(RefCell::new(4)), Rc::clone(&a));
    
    *value.borrow_mut() += 10;
    
    println!("a after = {:?}", a);
    println!("b after = {:?}", b);
    println!("c after = {:?}", c);
}

// Mock object example
pub trait Messenger {
    fn send(&self, msg: &str);
}

pub struct LimitTracker<'a, T: Messenger> {
    messenger: &'a T,
    value: usize,
    max: usize,
}

impl<'a, T> LimitTracker<'a, T>
where
    T: Messenger,
{
    pub fn new(messenger: &T, max: usize) -> LimitTracker<T> {
        LimitTracker {
            messenger,
            value: 0,
            max,
        }
    }
    
    pub fn set_value(&mut self, value: usize) {
        self.value = value;
        
        let percentage_of_max = self.value as f64 / self.max as f64;
        
        if percentage_of_max >= 1.0 {
            self.messenger.send("Error: You are over your quota!");
        } else if percentage_of_max >= 0.9 {
            self.messenger.send("Urgent warning: You've used up over 90% of your quota!");
        } else if percentage_of_max >= 0.75 {
            self.messenger.send("Warning: You've used up over 75% of your quota");
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use std::cell::RefCell;
    
    struct MockMessenger {
        sent_messages: RefCell<Vec<String>>,
    }
    
    impl MockMessenger {
        fn new() -> MockMessenger {
            MockMessenger {
                sent_messages: RefCell::new(vec![]),
            }
        }
    }
    
    impl Messenger for MockMessenger {
        fn send(&self, message: &str) {
            self.sent_messages.borrow_mut().push(String::from(message));
        }
    }
    
    #[test]
    fn it_sends_an_over_75_percent_warning_message() {
        let mock_messenger = MockMessenger::new();
        let mut limit_tracker = LimitTracker::new(&mock_messenger, 100);
        
        limit_tracker.set_value(80);
        
        assert_eq!(mock_messenger.sent_messages.borrow().len(), 1);
    }
}
```

### Arc<T> and Mutex<T>

```rust
use std::sync::{Arc, Mutex};
use std::thread;

fn main() {
    let counter = Arc::new(Mutex::new(0));
    let mut handles = vec![];
    
    for _ in 0..10 {
        let counter = Arc::clone(&counter);
        let handle = thread::spawn(move || {
            let mut num = counter.lock().unwrap();
            *num += 1;
        });
        handles.push(handle);
    }
    
    for handle in handles {
        handle.join().unwrap();
    }
    
    println!("Result: {}", *counter.lock().unwrap());
}
```

### Weak References

```rust
use std::rc::{Rc, Weak};
use std::cell::RefCell;

#[derive(Debug)]
struct Node {
    value: i32,
    parent: RefCell<Weak<Node>>,
    children: RefCell<Vec<Rc<Node>>>,
}

fn main() {
    let leaf = Rc::new(Node {
        value: 3,
        parent: RefCell::new(Weak::new()),
        children: RefCell::new(vec![]),
    });
    
    println!("leaf parent = {:?}", leaf.parent.borrow().upgrade());
    
    let branch = Rc::new(Node {
        value: 5,
        parent: RefCell::new(Weak::new()),
        children: RefCell::new(vec![Rc::clone(&leaf)]),
    });
    
    *leaf.parent.borrow_mut() = Rc::downgrade(&branch);
    
    println!("leaf parent = {:?}", leaf.parent.borrow().upgrade());
}
```

---

## Concurrency

### Threads

```rust
use std::thread;
use std::time::Duration;

fn main() {
    // Basic thread creation
    let handle = thread::spawn(|| {
        for i in 1..10 {
            println!("hi number {} from the spawned thread!", i);
            thread::sleep(Duration::from_millis(1));
        }
    });
    
    for i in 1..5 {
        println!("hi number {} from the main thread!", i);
        thread::sleep(Duration::from_millis(1));
    }
    
    handle.join().unwrap(); // Wait for thread to finish
    
    // Moving data into threads
    let v = vec![1, 2, 3];
    
    let handle = thread::spawn(move || {
        println!("Here's a vector: {:?}", v);
    });
    
    handle.join().unwrap();
}
```

### Message Passing

```rust
use std::sync::mpsc;
use std::thread;
use std::time::Duration;

fn main() {
    // Single producer, single consumer
    let (tx, rx) = mpsc::channel();
    
    thread::spawn(move || {
        let val = String::from("hi");
        tx.send(val).unwrap();
        // println!("val is {}", val); // Error: val was moved
    });
    
    let received = rx.recv().unwrap();
    println!("Got: {}", received);
    
    // Multiple messages
    let (tx, rx) = mpsc::channel();
    
    thread::spawn(move || {
        let vals = vec![
            String::from("hi"),
            String::from("from"),
            String::from("the"),
            String::from("thread"),
        ];
        
        for val in vals {
            tx.send(val).unwrap();
            thread::sleep(Duration::from_secs(1));
        }
    });
    
    for received in rx {
        println!("Got: {}", received);
    }
    
    // Multiple producers
    let (tx, rx) = mpsc::channel();
    
    let tx1 = tx.clone();
    thread::spawn(move || {
        let vals = vec![
            String::from("hi"),
            String::from("from"),
            String::from("the"),
            String::from("thread"),
        ];
        
        for val in vals {
            tx1.send(val).unwrap();
            thread::sleep(Duration::from_secs(1));
        }
    });
    
    thread::spawn(move || {
        let vals = vec![
            String::from("more"),
            String::from("messages"),
            String::from("for"),
            String::from("you"),
        ];
        
        for val in vals {
            tx.send(val).unwrap();
            thread::sleep(Duration::from_secs(1));
        }
    });
    
    for received in rx {
        println!("Got: {}", received);
    }
}
```

### Shared State Concurrency

```rust
use std::sync::{Arc, Mutex};
use std::thread;

fn main() {
    let counter = Arc::new(Mutex::new(0));
    let mut handles = vec![];
    
    for _ in 0..10 {
        let counter = Arc::clone(&counter);
        let handle = thread::spawn(move || {
            let mut num = counter.lock().unwrap();
            *num += 1;
        });
        handles.push(handle);
    }
    
    for handle in handles {
        handle.join().unwrap();
    }
    
    println!("Result: {}", *counter.lock().unwrap());
}

// RwLock for reader-writer scenarios
use std::sync::RwLock;

fn rwlock_example() {
    let data = Arc::new(RwLock::new(vec![1, 2, 3]));
    let mut handles = vec![];
    
    // Multiple readers
    for i in 0..5 {
        let data = Arc::clone(&data);
        let handle = thread::spawn(move || {
            let reader = data.read().unwrap();
            println!("Reader {}: {:?}", i, *reader);
        });
        handles.push(handle);
    }
    
    // One writer
    let data = Arc::clone(&data);
    let handle = thread::spawn(move || {
        let mut writer = data.write().unwrap();
        writer.push(4);
        println!("Writer: {:?}", *writer);
    });
    handles.push(handle);
    
    for handle in handles {
        handle.join().unwrap();
    }
}
```

### Atomic Types

```rust
use std::sync::atomic::{AtomicUsize, Ordering};
use std::sync::Arc;
use std::thread;

fn main() {
    let counter = Arc::new(AtomicUsize::new(0));
    let mut handles = vec![];
    
    for _ in 0..10 {
        let counter = Arc::clone(&counter);
        let handle = thread::spawn(move || {
            for _ in 0..1000 {
                counter.fetch_add(1, Ordering::SeqCst);
            }
        });
        handles.push(handle);
    }
    
    for handle in handles {
        handle.join().unwrap();
    }
    
    println!("Result: {}", counter.load(Ordering::SeqCst));
}
```

### Thread Pools

```rust
use std::sync::{mpsc, Arc, Mutex};
use std::thread;

pub struct ThreadPool {
    workers: Vec<Worker>,
    sender: mpsc::Sender<Job>,
}

type Job = Box<dyn FnOnce() + Send + 'static>;

impl ThreadPool {
    pub fn new(size: usize) -> ThreadPool {
        assert!(size > 0);
        
        let (sender, receiver) = mpsc::channel();
        let receiver = Arc::new(Mutex::new(receiver));
        
        let mut workers = Vec::with_capacity(size);
        
        for id in 0..size {
            workers.push(Worker::new(id, Arc::clone(&receiver)));
        }
        
        ThreadPool { workers, sender }
    }
    
    pub fn execute<F>(&self, f: F)
    where
        F: FnOnce() + Send + 'static,
    {
        let job = Box::new(f);
        self.sender.send(job).unwrap();
    }
}

struct Worker {
    id: usize,
    thread: thread::JoinHandle<()>,
}

impl Worker {
    fn new(id: usize, receiver: Arc<Mutex<mpsc::Receiver<Job>>>) -> Worker {
        let thread = thread::spawn(move || loop {
            let job = receiver.lock().unwrap().recv().unwrap();
            println!("Worker {} got a job; executing.", id);
            job();
        });
        
        Worker { id, thread }
    }
}
```

---

## Async Programming

### Basic Async/Await

```rust
// Add to Cargo.toml:
// [dependencies]
// tokio = { version = "1", features = ["full"] }

use tokio::time::{sleep, Duration};

async fn hello_world() {
    println!("Hello, world!");
}

async fn say_hello() {
    sleep(Duration::from_secs(1)).await;
    println!("Hello!");
}

#[tokio::main]
async fn main() {
    hello_world().await;
    say_hello().await;
    
    // Running multiple futures concurrently
    let future1 = say_hello();
    let future2 = say_hello();
    let future3 = say_hello();
    
    // Wait for all futures to complete
    tokio::join!(future1, future2, future3);
}
```

### Async Functions and Return Types

```rust
use std::future::Future;

// These are equivalent:
async fn foo() -> u8 {
    5
}

fn foo() -> impl Future<Output = u8> {
    async {
        5
    }
}

// Using Box for dynamic dispatch
fn returns_async() -> Box<dyn Future<Output = u8> + Send> {
    Box::new(async {
        5
    })
}
```

### Working with Streams

```rust
use tokio_stream::{self as stream, StreamExt};

#[tokio::main]
async fn main() {
    let mut stream = stream::iter(vec![1, 2, 3, 4, 5]);
    
    while let Some(item) = stream.next().await {
        println!("Got: {}", item);
    }
    
    // Transform streams
    let doubled: Vec<_> = stream::iter(vec![1, 2, 3, 4, 5])
        .map(|x| x * 2)
        .collect()
        .await;
    
    println!("Doubled: {:?}", doubled);
}
```

### Async I/O

```rust
use tokio::fs::File;
use tokio::io::{self, AsyncReadExt, AsyncWriteExt};

#[tokio::main]
async fn main() -> io::Result<()> {
    // Read a file
    let mut file = File::open("example.txt").await?;
    let mut contents = String::new();
    file.read_to_string(&mut contents).await?;
    println!("File contents: {}", contents);
    
    // Write to a file
    let mut file = File::create("output.txt").await?;
    file.write_all(b"Hello, async world!").await?;
    
    Ok(())
}
```

### Async Channels

```rust
use tokio::sync::{mpsc, oneshot};
use tokio::time::{sleep, Duration};

#[tokio::main]
async fn main() {
    // Multiple producer, single consumer
    let (tx, mut rx) = mpsc::channel(32);
    
    let tx2 = tx.clone();
    tokio::spawn(async move {
        tx.send("Hello").await.unwrap();
        tx.send("World").await.unwrap();
    });
    
    tokio::spawn(async move {
        tx2.send("from").await.unwrap();
        tx2.send("tokio").await.unwrap();
    });
    
    while let Some(message) = rx.recv().await {
        println!("Got: {}", message);
    }
    
    // Oneshot channel
    let (tx, rx) = oneshot::channel();
    
    tokio::spawn(async move {
        sleep(Duration::from_secs(1)).await;
        tx.send("Hello from oneshot").unwrap();
    });
    
    match rx.await {
        Ok(message) => println!("Received: {}", message),
        Err(_) => println!("Sender dropped"),
    }
}
```

### Async Mutex and RwLock

```rust
use tokio::sync::{Mutex, RwLock};
use std::sync::Arc;

#[tokio::main]
async fn main() {
    // Async Mutex
    let data = Arc::new(Mutex::new(0));
    let mut handles = vec![];
    
    for _ in 0..10 {
        let data = Arc::clone(&data);
        let handle = tokio::spawn(async move {
            let mut lock = data.lock().await;
            *lock += 1;
        });
        handles.push(handle);
    }
    
    for handle in handles {
        handle.await.unwrap();
    }
    
    println!("Result: {}", *data.lock().await);
    
    // Async RwLock
    let data = Arc::new(RwLock::new(vec![1, 2, 3]));
    let mut handles = vec![];
    
    // Multiple readers
    for i in 0..5 {
        let data = Arc::clone(&data);
        let handle = tokio::spawn(async move {
            let reader = data.read().await;
            println!("Reader {}: {:?}", i, *reader);
        });
        handles.push(handle);
    }
    
    // One writer
    let data = Arc::clone(&data);
    let handle = tokio::spawn(async move {
        let mut writer = data.write().await;
        writer.push(4);
        println!("Writer: {:?}", *writer);
    });
    handles.push(handle);
    
    for handle in handles {
        handle.await.unwrap();
    }
}
```

### Select and Cancellation

```rust
use tokio::time::{sleep, Duration, timeout};

#[tokio::main]
async fn main() {
    // Using select!
    tokio::select! {
        _ = sleep(Duration::from_secs(1)) => {
            println!("First future completed");
        }
        _ = sleep(Duration::from_secs(2)) => {
            println!("Second future completed");
        }
    }
    
    // Timeout
    let result = timeout(Duration::from_secs(1), sleep(Duration::from_secs(2))).await;
    match result {
        Ok(_) => println!("Operation completed in time"),
        Err(_) => println!("Operation timed out"),
    }
    
    // Cancellation with select!
    let (tx, mut rx) = tokio::sync::mpsc::channel(1);
    
    tokio::spawn(async move {
        sleep(Duration::from_secs(2)).await;
        tx.send("Done").await.unwrap();
    });
    
    tokio::select! {
        msg = rx.recv() => {
            println!("Received: {:?}", msg);
        }
        _ = sleep(Duration::from_secs(1)) => {
            println!("Timed out waiting for message");
        }
    }
}
```

---

## Testing

### Unit Tests

```rust
#[cfg(test)]
mod tests {
    use super::*;
    
    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }
    
    #[test]
    fn it_adds_two() {
        assert_eq!(add_two(2), 4);
    }
    
    #[test]
    #[should_panic]
    fn greater_than_100() {
        Guess::new(200);
    }
    
    #[test]
    #[should_panic(expected = "Guess value must be less than or equal to 100")]
    fn greater_than_100_with_message() {
        Guess::new(200);
    }
    
    #[test]
    fn it_works_with_result() -> Result<(), String> {
        if 2 + 2 == 4 {
            Ok(())
        } else {
            Err(String::from("two plus two does not equal four"))
        }
    }
    
    #[test]
    #[ignore]
    fn expensive_test() {
        // Test that takes a long time
    }
}

pub fn add_two(a: i32) -> i32 {
    a + 2
}

pub struct Guess {
    value: i32,
}

impl Guess {
    pub fn new(value: i32) -> Guess {
        if value < 1 || value > 100 {
            panic!("Guess value must be between 1 and 100, got {}.", value);
        }
        
        Guess { value }
    }
}

// Testing with custom assertions
fn prints_and_returns_10(a: i32) -> i32 {
    println!("I got the value {}", a);
    10
}

#[cfg(test)]
mod tests {
    use super::*;
    
    #[test]
    fn this_test_will_pass() {
        let value = prints_and_returns_10(4);
        assert_eq!(10, value);
    }
    
    #[test]
    fn this_test_will_fail() {
        let value = prints_and_returns_10(8);
        assert_eq!(5, value);
    }
}
```

### Integration Tests

```rust
// tests/integration_test.rs
use my_crate;

#[test]
fn it_adds_two() {
    assert_eq!(4, my_crate::add_two(2));
}

// tests/common/mod.rs
pub fn setup() {
    // Setup code common to multiple integration tests
}

// tests/another_test.rs
use my_crate;

mod common;

#[test]
fn it_works() {
    common::setup();
    assert_eq!(4, my_crate::add_two(2));
}
```

### Property-Based Testing

```rust
// Add to Cargo.toml:
// [dev-dependencies]
// quickcheck = "1"
// quickcheck_macros = "1"

use quickcheck_macros::quickcheck;

#[quickcheck]
fn reverse_reverse_is_identity(xs: Vec<i32>) -> bool {
    let mut ys = xs.clone();
    ys.reverse();
    ys.reverse();
    xs == ys
}

#[quickcheck]
fn sort_is_idempotent(mut xs: Vec<i32>) -> bool {
    xs.sort();
    let mut ys = xs.clone();
    ys.sort();
    xs == ys
}
```

### Benchmark Tests

```rust
// Add to Cargo.toml:
// [dev-dependencies]
// criterion = "0.3"

// benches/my_benchmark.rs
use criterion::{black_box, criterion_group, criterion_main, Criterion};

fn fibonacci(n: u64) -> u64 {
    match n {
        0 => 1,
        1 => 1,
        n => fibonacci(n-1) + fibonacci(n-2),
    }
}

fn criterion_benchmark(c: &mut Criterion) {
    c.bench_function("fib 20", |b| b.iter(|| fibonacci(black_box(20))));
}

criterion_group!(benches, criterion_benchmark);
criterion_main!(benches);
```

---

## Project Structure

### Cargo.toml Configuration

```toml
[package]
name = "my_project"
version = "0.1.0"
edition = "2021"
authors = ["Your Name <you@example.com>"]
description = "A sample Rust project"
license = "MIT OR Apache-2.0"
repository = "https://github.com/username/my_project"
homepage = "https://github.com/username/my_project"
documentation = "https://docs.rs/my_project"
readme = "README.md"
keywords = ["cli", "example"]
categories = ["command-line-utilities"]

[dependencies]
serde = { version = "1.0", features = ["derive"] }
tokio = { version = "1", features = ["full"] }
clap = "3.0"

[dev-dependencies]
tokio-test = "0.4"

[build-dependencies]
cc = "1.0"

# Binary targets
[[bin]]
name = "main"
path = "src/main.rs"

[[bin]]
name = "tool"
path = "src/bin/tool.rs"

# Library target
[lib]
name = "my_project"
path = "src/lib.rs"

# Example targets
[[example]]
name = "example1"
path = "examples/example1.rs"

# Workspace configuration
[workspace]
members = [
    "crate1",
    "crate2",
]

# Features
[features]
default = ["std"]
std = []
serde_support = ["serde"]
```

### Typical Project Structure

```
my_project/
├── Cargo.toml
├── Cargo.lock
├── README.md
├── LICENSE
├── .gitignore
├── src/
│   ├── main.rs          # Binary crate entry point
│   ├── lib.rs           # Library crate entry point
│   ├── bin/
│   │   └── tool.rs      # Additional binary
│   └── modules/
│       ├── mod.rs
│       ├── parser.rs
│       └── utils.rs
├── tests/
│   ├── integration_test.rs
│   └── common/
│       └── mod.rs
├── examples/
│   └── example1.rs
├── benches/
│   └── benchmark.rs
├── docs/
│   └── api.md
└── target/              # Build artifacts (generated)
    └── debug/
    └── release/
```

### Workspace Configuration

```toml
# Workspace Cargo.toml
[workspace]
members = [
    "app",
    "utils",
    "models"
]

# Shared dependencies
[workspace.dependencies]
serde = "1.0"
tokio = "1.0"

# Individual crate Cargo.toml
[package]
name = "app"
version = "0.1.0"
edition = "2021"

[dependencies]
utils = { path = "../utils" }
models = { path = "../models" }
serde = { workspace = true }
```

---

## Advanced Topics

### Macros

#### Declarative Macros (macro_rules!)

```rust
// Simple macro
macro_rules! say_hello {
    () => {
        println!("Hello!");
    };
}

// Macro with parameters
macro_rules! create_function {
    ($func_name:ident) => {
        fn $func_name() {
            println!("You called {:?}()", stringify!($func_name));
        }
    };
}

create_function!(foo);
create_function!(bar);

// More complex macro with repetitions
macro_rules! vec {
    ( $( $x:expr ),* ) => {
        {
            let mut temp_vec = Vec::new();
            $(
                temp_vec.push($x);
            )*
            temp_vec
        }
    };
}

fn main() {
    say_hello!();
    foo();
    bar();
    
    let v = vec![1, 2, 3];
    println!("{:?}", v);
}

// Pattern matching in macros
macro_rules! calculate {
    (eval $e:expr) => {{
        {
            let val: usize = $e; // Force types to be integers
            println!("{} = {}", stringify!($e), val);
        }
    }};
}

// Usage
calculate! {
    eval 1 + 2 // hehehe `eval` is _not_ a Rust keyword!
}
```

#### Procedural Macros

```rust
// Cargo.toml
// [lib]
// proc-macro = true
// [dependencies]
// syn = "1.0"
// quote = "1.0"
// proc-macro2 = "1.0"

use proc_macro::TokenStream;
use quote::quote;
use syn;

// Function-like procedural macro
#[proc_macro]
pub fn make_answer(_item: TokenStream) -> TokenStream {
    "fn answer() -> u32 { 42 }".parse().unwrap()
}

// Derive macro
#[proc_macro_derive(HelloMacro)]
pub fn hello_macro_derive(input: TokenStream) -> TokenStream {
    let ast = syn::parse(input).unwrap();
    impl_hello_macro(&ast)
}

fn impl_hello_macro(ast: &syn::DeriveInput) -> TokenStream {
    let name = &ast.ident;
    let gen = quote! {
        impl HelloMacro for #name {
            fn hello_macro() {
                println!("Hello, Macro! My name is {}!", stringify!(#name));
            }
        }
    };
    gen.into()
}

// Attribute macro
#[proc_macro_attribute]
pub fn route(args: TokenStream, input: TokenStream) -> TokenStream {
    // Implementation would parse the args and modify the input
    input
}
```

### Unsafe Rust

```rust
fn main() {
    // Raw pointers
    let mut num = 5;
    
    let r1 = &num as *const i32;
    let r2 = &mut num as *mut i32;
    
    // Arbitrary memory address (dangerous!)
    let address = 0x012345usize;
    let r = address as *const i32;
    
    // Dereferencing raw pointers (requires unsafe)
    unsafe {
        println!("r1 is: {}", *r1);
        println!("r2 is: {}", *r2);
    }
    
    // Calling unsafe functions
    unsafe {
        dangerous();
    }
    
    // Safe abstraction over unsafe code
    use std::slice;
    
    let mut v = vec![1, 2, 3, 4, 5, 6];
    let r = &mut v[..];
    let (a, b) = split_at_mut(r, 3);
    
    assert_eq!(a, &mut [1, 2, 3]);
    assert_eq!(b, &mut [4, 5, 6]);
}

unsafe fn dangerous() {
    // Unsafe operations
}

fn split_at_mut(slice: &mut [i32], mid: usize) -> (&mut [i32], &mut [i32]) {
    let len = slice.len();
    let ptr = slice.as_mut_ptr();
    
    assert!(mid <= len);
    
    unsafe {
        (
            std::slice::from_raw_parts_mut(ptr, mid),
            std::slice::from_raw_parts_mut(ptr.add(mid), len - mid),
        )
    }
}

// Accessing or modifying mutable static variables
static HELLO_WORLD: &str = "Hello, world!";
static mut COUNTER: usize = 0;

fn add_to_count(inc: usize) {
    unsafe {
        COUNTER += inc;
    }
}

fn main() {
    add_to_count(3);
    
    unsafe {
        println!("COUNTER: {}", COUNTER);
    }
}
```

### Foreign Function Interface (FFI)

```rust
// Calling C functions from Rust
extern "C" {
    fn abs(input: i32) -> i32;
}

fn main() {
    unsafe {
        println!("Absolute value of -3 according to C: {}", abs(-3));
    }
}

// Calling Rust functions from other languages
#[no_mangle]
pub extern "C" fn call_from_c() {
    println!("Just called a Rust function from C!");
}

// More complex example with strings
use std::os::raw::{c_char, c_int};
use std::ffi::{CStr, CString};

#[no_mangle]
pub extern "C" fn hello_from_rust(name: *const c_char) -> *mut c_char {
    let c_str = unsafe {
        assert!(!name.is_null());
        CStr::from_ptr(name)
    };
    
    let recipient = match c_str.to_str() {
        Err(_) => "there",
        Ok(string) => string,
    };
    
    let response = format!("Hello {}!", recipient);
    let c_string = CString::new(response).expect("CString::new failed");
    c_string.into_raw()
}

#[no_mangle]
pub extern "C" fn free_string(s: *mut c_char) {
    unsafe {
        if s.is_null() { return }
        CString::from_raw(s)
    };
}
```

### Advanced Trait Usage

```rust
// Associated types vs generics
trait Iterator {
    type Item;
    fn next(&mut self) -> Option<Self::Item>;
}

// Operator overloading
use std::ops::Add;

#[derive(Debug, Copy, Clone, PartialEq)]
struct Point {
    x: i32,
    y: i32,
}

impl Add for Point {
    type Output = Point;
    
    fn add(self, other: Point) -> Point {
        Point {
            x: self.x + other.x,
            y: self.y + other.y,
        }
    }
}

// Default generic type parameters
impl<Rhs = Self> Add<Rhs> for Point
where
    Rhs: Into<Point>,
{
    type Output = Point;
    
    fn add(self, rhs: Rhs) -> Point {
        let rhs = rhs.into();
        Point {
            x: self.x + rhs.x,
            y: self.y + rhs.y,
        }
    }
}

// Fully Qualified Syntax
trait Pilot {
    fn fly(&self);
}

trait Wizard {
    fn fly(&self);
}

struct Human;

impl Pilot for Human {
    fn fly(&self) {
        println!("This is your captain speaking.");
    }
}

impl Wizard for Human {
    fn fly(&self) {
        println!("Up!");
    }
}

impl Human {
    fn fly(&self) {
        println!("*waving arms furiously*");
    }
}

fn main() {
    let person = Human;
    Pilot::fly(&person);
    Wizard::fly(&person);
    person.fly();
}

// Supertraits
use std::fmt;

trait OutlinePrint: fmt::Display {
    fn outline_print(&self) {
        let output = self.to_string();
        let len = output.len();
        println!("{}", "*".repeat(len + 4));
        println!("*{}*", " ".repeat(len + 2));
        println!("* {} *", output);
        println!("*{}*", " ".repeat(len + 2));
        println!("{}", "*".repeat(len + 4));
    }
}

// Newtype pattern
struct Wrapper(Vec<String>);

impl fmt::Display for Wrapper {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "[{}]", self.0.join(", "))
    }
}

// Type aliases
type Kilometers = i32;
type Thunk = Box<dyn Fn() + Send + 'static>;

// Never type
fn bar() -> ! {
    panic!()
}

// Function pointers
fn add_one(x: i32) -> i32 {
    x + 1
}

fn do_twice(f: fn(i32) -> i32, arg: i32) -> i32 {
    f(arg) + f(arg)
}

fn main() {
    let answer = do_twice(add_one, 5);
    println!("The answer is: {}", answer);
    
    // Closures can be coerced to function pointers
    let list_of_numbers = vec![1, 2, 3];
    let list_of_strings: Vec<String> =
        list_of_numbers.iter().map(|i| i.to_string()).collect();
        
    // Or using function pointers
    let list_of_strings: Vec<String> =
        list_of_numbers.iter().map(ToString::to_string).collect();
}
```

---

## Resources and Next Steps

### Essential Tools
- **rustc**: The Rust compiler
- **cargo**: Package manager and build system
- **rustfmt**: Code formatter
- **clippy**: Linter for common mistakes
- **rust-analyzer**: Language server for IDEs

### Useful Cargo Commands
```bash
cargo new project_name      # Create new project
cargo build                 # Build project
cargo run                   # Build and run
cargo test                  # Run tests
cargo check                 # Check without building
cargo clippy               # Run linter
cargo fmt                  # Format code
cargo doc --open           # Generate and open documentation
cargo install package      # Install package globally
cargo update               # Update dependencies
```

### Learning Resources
- [The Rust Programming Language Book](https://doc.rust-lang.org/book/)
- [Rust by Example](https://doc.rust-lang.org/rust-by-example/)
- [The Rustonomicon](https://doc.rust-lang.org/nomicon/) (Advanced topics)
- [Rust Reference](https://doc.rust-lang.org/reference/)
- [Cargo Book](https://doc.rust-lang.org/cargo/)

### Community and Ecosystem
- [crates.io](https://crates.io/) - Package registry
- [docs.rs](https://docs.rs/) - Documentation hosting
- [Rust Forums](https://users.rust-lang.org/)
- [Rust Discord](https://discord.gg/rust-lang)
- [r/rust subreddit](https://reddit.com/r/rust)

### Popular Crates by Category

#### Web Development
- **axum** / **warp** / **actix-web**: Web frameworks
- **hyper**: HTTP library
- **reqwest**: HTTP client
- **serde**: Serialization framework

#### CLI Applications
- **clap**: Command line argument parser
- **structopt**: Derive-based CLI parsing
- **colored**: Terminal colors
- **indicatif**: Progress bars

#### Async Runtime
- **tokio**: Async runtime
- **async-std**: Alternative async runtime
- **futures**: Async utilities

#### Database
- **sqlx**: Async SQL toolkit
- **diesel**: ORM and query builder
- **rusqlite**: SQLite bindings

#### Error Handling
- **anyhow**: Flexible error handling
- **thiserror**: Derive macros for error types
- **color-eyre**: Error reporting

#### Data Structures
- **serde**: Serialization
- **chrono**: Date and time
- **uuid**: UUID generation
- **regex**: Regular expressions

This course provides a comprehensive foundation in Rust programming. Practice these concepts by building projects and exploring the vast Rust ecosystem!