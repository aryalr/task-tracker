# Task-Tracker CLI Specification

## Command Structure:

Example:

### Add task:

```bash
mytask add "create user input" p1
```

output:

`successfully added task "create user input" id => 1 | priority => high`

### Update task:

```bash
mytask update 1 "add update user input" p2
```

output:

`successfully update task "add update user input" id => 1 | priority => medium`

### Delete task:

```bash
mytask delete 1
```

output:

`successfully delete task "add update user input" id => 1`

### List task:

```bash
mytask list
```

output:

```bash
task list:

id | Name                  | Priority

1  | add update user input | Medium
2  | create database port  | High
```

### Mark task:

```bash
mytask done 1
```

output:

`task marked done id => 1`

---

### description:


| Input    | Description   | value                                            |
| ---------- | --------------- | :------------------------------------------------- |
| mytask   | app name      |                                                  |
| add      | input type    | add, list, done, update, delete                  |
| update   | input type    | id                                               |
| done     | input type    | id                                               |
| delete   | input type    | id                                               |
| update   | input type    | id(required), name(optional), priority(optional) |
| priority | task priority | p1 = high, p2 = medium, p3 = low                 |

---

This app use/create json files in current working directory. filename: mytask.json
