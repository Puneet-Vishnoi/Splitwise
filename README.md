# Splitwise Clone (Go)

A simple expense-sharing application written in Go, inspired by Splitwise. Supports multiple users, groups, and expenses with different split strategies.

## Features

- Add users and groups
- Record expenses with different split types:
  - Equal split
  - Exact split
  - Percentage split
  - Share-based split
- View group-wise balances

## Project Structure
/Splitwise
├── main.go
├── manager/
│ └── expense_manager.go
├── models/
│ ├── expense.go
│ ├── user.go
│ ├── group.go
│ ├── split.go

pgsql
Copy
Edit

## Split Types

- `Equal`: Divides the expense equally among participants.
- `Exact`: Allows exact amounts per user.
- `Percentage`: Splits based on given percentage per user.
- `Share`: Splits based on proportional share units.
