# go-sandbox
A playground to learn golang basics.

## Project 1: Library Managment System

This Go project teaches you structs and methods through a practical library management system. Here's what you'll learn:
Key Concepts Demonstrated:
1. Struct Definition

Book struct with various field types (int, string, bool, time.Time)
Library struct that contains slices of other structs
LibraryStats struct for data organization

2. Method Types

Value receivers (func (b Book) String()) - for methods that only read data
Pointer receivers (func (b *Book) Borrow()) - for methods that modify the struct

3. Constructor Pattern

NewLibrary() function that returns a pointer to an initialized struct

4. Interface Implementation

String() method implements the Stringer interface automatically

5. Error Handling

Methods return errors when operations can't be completed

Running the Project:

Save the code as library.go
Run with: go run library.go

Exercises to Try:

Add a Member struct with methods for tracking borrowing history
Add validation methods (e.g., IsValidISBN())
Create methods to sort books by different criteria
Add a Category field and methods to filter by category
Implement a Reservation system with its own struct and methods

The output will show you how structs maintain state and how methods can both read and modify that state, which are fundamental concepts in Go programming!