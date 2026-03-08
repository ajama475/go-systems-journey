# Go Programming Lab – Structs, Interfaces, and Control Flow

## Overview

This lab reinforces the Go concepts you have learned so far:

- Variables
- Functions
- Structs
- Conditional logic (`if` and `switch`)
- Interfaces

You will build a small program that models **different types of users in a system**. Each user type will have its own information and behavior, but the program will interact with them through a **shared interface**.

The goal is to practice struct design, implementing interfaces, and applying conditional logic in Go.

---

# Learning Objectives

By completing this lab, you should be able to:

- Design structs to represent real-world entities
- Implement methods on structs
- Define and use interfaces
- Apply conditional logic using `if` and `switch`
- Write functions that work with interface types

---

# Scenario

You are building a simple **account management system**.

The system supports multiple types of users. Each type of user has slightly different information and permissions, but they all share common behaviors.

The system must be able to:

- Identify who the user is
- Determine the user's role
- Determine the user's level of access
- Display information about the user

To accomplish this, you will design structs representing different users and connect them through a shared interface.

---

# Lab Requirements

## 1. Create a User Interface

Define an interface that represents a generic user in the system.

The interface should represent behaviors that **every user must support**, such as retrieving their name, identifying their role, and determining their access level.

All user types must satisfy this interface.

---

## 2. Create User Types

Create two user types using structs.

### Admin

Represents a system administrator. Administrators have the highest level of permissions.

The struct should contain information such as the administrator's name and department.

### Regular User

Represents a standard user of the system. Regular users have limited permissions.

The struct should contain information such as the user's name and membership level.

---

## 3. Implement the Interface

Both structs must implement the interface created earlier.

Each type should provide its own behavior for:

- Returning the user's name
- Returning the user's role
- Returning the user's access level

Although the interface is shared, the values returned by each type should reflect their different roles in the system.

---

## 4. Create a Function That Uses the Interface

Write a function responsible for displaying information about a user.

This function should accept a parameter whose type is the user interface.

Inside the function, display details such as the user's name, role, and access level.

The important requirement is that this function must work with **any type that satisfies the interface**, not just one specific struct.

---

## 5. Add Conditional Logic

Inside the user description function, add decision-making logic.

Use an `if` statement to determine whether the user has high or limited system access based on their access level.

Use a `switch` statement to determine the message displayed for different user roles.

This part of the lab is meant to reinforce how Go handles conditional logic.

---

## 6. Create Multiple Users

In the `main` function, create multiple users using the structs you defined.

You should create **at least three users**, including both administrators and regular users.

---

## 7. Process Each User

Call the function you created earlier to display the information for each user.

Do this by manually calling the function for each user you created in the main program.

This step ensures the program demonstrates how the interface works with multiple types.

---

# Program Behavior

When the program runs, it should display information about each user and print messages that reflect their permissions and role in the system.

The program should clearly demonstrate how different structs can be treated uniformly through an interface while still behaving differently.

---

# Constraints

Your solution must include:

- At least two different structs
- One interface
- Struct methods that satisfy the interface
- A function that accepts the interface type
- At least one `if` statement
- At least one `switch` statement
- Multiple user objects created in the program

---

# Bonus Challenges (Optional)

## Add a Guest User

Create a third user type representing a guest with very limited access.

## Access Control Function

Create a function that determines whether a user is allowed to access administrator-only features.

## Greeting Function

Create a function that prints a greeting message based on the user's role.

---

# Deliverables

Your submission should include:

- The Go source file containing the program
- This README file describes the lab

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
