# Go Programming Lab – Structs, Interfaces, and Control Flow

## Overview

This lab is designed to reinforce several core Go concepts you have learned so far:

- Variables
- Functions
- Structs
- Conditional logic (`if` and `switch`)
- Interfaces

You will design a small program that models **different types of users in a system**. Each user type will have its own data and behavior, but the program will interact with them through a **shared interface**.

The goal of this lab is to help you practice struct design, implementing interfaces, and applying conditional logic in Go.

---

# Learning Objectives

By completing this lab, you should be able to:

- Design structs to represent different entities
- Implement methods on structs
- Define and use interfaces
- Use `if` and `switch` statements for decision making
- Write functions that work with interface types
- Organize program logic using functions

---

# Scenario

You are building a very simple **account management system**.

The system supports multiple types of users. Each user type has slightly different information and permissions, but they all share some common behavior.

The system needs to:

- Store information about users
- Identify their role
- Determine their access level
- Display information about them

To accomplish this, you will design structs representing different users and connect them through a common interface.

---

# Lab Requirements

## 1. Create a User Interface

Define an interface representing a generic user in the system.

The interface should represent the behaviors that **every user must support**, such as retrieving their name, identifying their role, and determining their level of access.

All user types must satisfy this interface.

---

## 2. Create User Types

Create two different user types using structs.

### Admin

Represents a system administrator.  
Administrators have the highest level of permissions in the system.

The struct should store information such as the administrator's name and department.

### Regular User

Represents a standard user of the system.  
Regular users have limited permissions.

The struct should store information such as the user's name and membership level.

---

## 3. Implement the Interface

Both user types must implement the interface created earlier.

Each type should provide its own behavior for:

- Returning the user's name
- Returning the user's role
- Returning the user's access level

Although the interface is shared, the actual values returned should differ between administrators and regular users.

---

## 4. Create a Function That Uses the Interface

Write a function responsible for displaying information about a user.

This function should accept a value that satisfies the user interface. The function should display details such as the user's name, role, and access level.

This function should work with **any type that implements the interface**, not just one specific struct.

---

## 5. Add Conditional Logic

Inside the user description function, include conditional logic to display different messages depending on the user's role and permissions.

You should use:

- An `if` statement to evaluate the user's access level and determine the type of system access they have.
- A `switch` statement to determine the message that should be shown for different roles.

---

## 6. Create Multiple Users

In your program's main function, create multiple users using both user types.

The system should include at least **three users**. Some should be administrators and some should be regular users.

---

## 7. Store Users in a Slice of Interfaces

Create a slice that stores multiple users through the interface type.

This demonstrates how interfaces allow different types to be handled uniformly.

Loop through the slice and process each user using the function you created earlier.

---

# Program Behavior

When the program runs, it should iterate through all users and display their information along with messages describing their permissions.

The program should clearly show how different user types share common behavior through the interface while still having their own characteristics.

---

# Constraints

Your solution must include:

- At least two different structs
- One interface
- Struct methods that satisfy the interface
- At least one function that accepts the interface type
- Use of both `if` and `switch` statements
- A slice that stores interface values
- A loop that processes multiple users

---

# Bonus Challenges (Optional)

## Add a Guest User

Add a third user type representing a guest. Guests have the lowest level of access in the system.

## Access Control Function

Create a function that determines whether a user is allowed to access an administrator-only feature.

## Greeting Function

Create a function that prints a welcome message tailored to the user's role.

---

# Deliverables

Your submission should include:

- The Go source file containing the program
- This README file describing the lab

The program must compile and run successfully.

---

# Estimated Time

Approximately **45–60 minutes** to complete.

---

# Expected output
User: Alice

Role: Admin

Access Level: 10

Administrative privileges granted

Full system access


User: Charlie

Role: User

Access Level: 1

Standard user permissions

Limited access
