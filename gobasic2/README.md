# Go Programming Lab – User Directory Manager

## Overview

This lab is designed to reinforce the Go concepts you have learned so far:

- Variables
- Functions
- Structs
- Methods
- Conditional logic (`if` and `switch`)
- Interfaces
- Arrays
- Slices
- Maps

You will build a small program that manages a simple **user directory system**. The program will store users, organize their information, and allow basic lookups and reports.

The goal is to combine your newer data structure knowledge with the earlier Go concepts you already learned.

---

# Learning Objectives

By completing this lab, you should be able to:

- Use **structs** to model data
- Use **methods** to attach behavior to types
- Use an **interface** to describe shared behavior
- Use **arrays** for fixed-size data
- Use **slices** for flexible collections
- Use **maps** to connect keys to values
- Use **functions** to organize logic
- Use **if** and **switch** to control program behavior

---

# Scenario

You are building a small **user directory manager** for a school club.

Each user in the system has:

- A name
- A role
- An age
- A list of interests

The system should be able to:

- Store several users
- Group them by ID
- Show their information
- Classify them based on role
- Display some fixed information using an array
- Work with users through a shared interface

---

# Lab Requirements

## 1. Create an Interface

Create an interface that represents a user in the system.

This interface should describe the basic behaviors every user must have, such as:

- returning their name
- returning their role
- returning their age

You may choose the exact method names, but all user types must satisfy this interface.

---

## 2. Create Structs

Create at least **two user types** using structs.

Examples:
- `Admin`
- `Member`

Each struct should contain user-related fields.

At minimum, each user should have:

- name
- age
- role-related information

You should also include a field that stores the user's interests using a **slice**.

---

## 3. Implement the Interface

Both structs must implement the interface.

Even though the two structs are different, the program should still be able to treat them both as users through the shared interface.

---

## 4. Use an Array

Create one **array** that stores fixed system information.

This array should have exactly 3 items.

Possible examples:
- three department names
- three supported membership types
- three fixed club rules
- three office locations

The purpose of this step is to practice arrays as fixed-size collections.

---

## 5. Use Slices

Each user should have a slice storing multiple interests, skills, or hobbies.

Examples:
- badminton
- coding
- reading
- design

Your program should display these interests as part of the user's information.

This step is meant to help you practice storing flexible-length data.

---

## 6. Use a Map

Create a **map** that connects a user ID to a user.

Example idea:

- key: integer ID
- value: user struct

This map will act like a simple directory or registry.

The program should use the map to store at least **three users**.

---

## 7. Create Functions

Write functions that help manage the program.

Your program should include functions for tasks such as:

- displaying user information
- checking whether a user is an admin or member
- looking up a user by ID
- printing system information

The exact function names are up to you, but your solution should be clearly organized into separate functions.

---

## 8. Add Conditional Logic

Use both:

- `if`
- `switch`

Ideas:
- use `if` to check whether a user is old enough for a certain permission
- use `if` to check whether a map lookup succeeded
- use `switch` to display different messages based on the user's role

This part is meant to combine logic with your data structures.

---

## 9. Work With Multiple Users

Create multiple users in `main()`.

Your program should work with at least:

- one admin
- two regular members

Store them in your map and process them one by one.

You may manually call your functions for each user if you do not want to rely heavily on loops yet.

---

# Program Behavior

When the program runs, it should:

- show the fixed system information from the array
- create several users
- store them in a map
- display each user's information
- show their interests
- classify them based on their role
- allow a lookup by user ID
- print a helpful message if a user ID does not exist

---

# Constraints

Your solution must include:

- at least **one interface**
- at least **two structs**
- at least **one array**
- at least **one slice**
- at least **one map**
- at least **three users**
- at least **two functions**
- at least one `if` statement
- at least one `switch` statement

---

# Bonus Challenges (Optional)

## Bonus 1: Add a Guest Type

Create a third struct for a guest user with limited access.

## Bonus 2: Add a Function for Interest Count

Write a function that tells how many interests a user has.

## Bonus 3: Add a Role Summary

Write a function that prints a different message depending on the user's role.

## Bonus 4: Add a Search Result Message

When searching the map by ID, print one message if the user exists and another if the user does not.

---

# Deliverables

Your submission should include:

- `main.go`
- `README.md`

Your program must compile and run successfully.

---

# Suggested Project Structure

go-user-directory/
- main.go
- README.md

---

# Estimated Time

Approximately 60–90 minutes.

---

# Success Checklist

Before submitting, make sure your program:

- uses arrays correctly for fixed-size data
- uses slices correctly for flexible data
- uses maps correctly for key-value lookup
- uses structs and interfaces together
- uses functions to keep the code organized
- uses `if` and `switch`
- runs without compile errors