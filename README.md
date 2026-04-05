

````markdown
# Task CLI

A simple Go command-line application to manage tasks. You can add, update, delete, and mark tasks as done or in progress, and list tasks by status.

---

## Features

- Add tasks with a description
- List all tasks
- Update task descriptions
- Delete tasks
- Mark tasks as done or in-progress
- List tasks by status (done, in-progress, todo)
- Per-command help with usage

---

## Installation

1. **Clone the repository:**

```bash
git clone <https://github.com/joelyakubu269/CLI-functions>
cd cliApp
````

2. **Build the binary (if not built yet or after code changes):**

```bash
go build -o taskcli
``

3. **Optional: Move the binary to your PATH:**

```bash
mv taskcli ~/go/bin/
export PATH=$PATH:~/go/bin
```

Now you can run `taskcli` from anywhere.

---

## Usage

Run commands as:

```bash
./taskcli <command> [flags]
```

Use `./taskcli help` to see all available commands.

---

## Commands

### 1. Add a Task

```bash
./taskcli add -description "Buy groceries"
```

* Adds a new task with a description.
* `-description` is **required**.

---

### 2. List All Tasks

```bash
./taskcli list
```

* Lists all tasks with their IDs, description, and status.

---

### 3. Delete a Task

```bash
./taskcli delete -id 4
```

* Deletes the task with the given ID.
* `-id` is **required**.

---

### 4. Update a Task

```bash
./taskcli update -id 2 -description "Complete homework"
```

* Updates the description of a task with the given ID.
* Both `-id` and `-description` are **required**.

---

### 5. Mark Task Done

```bash
./taskcli done -id 3
```

* Marks the task with the given ID as done.
* `-id` is **required**.

---

### 6. Mark Task In Progress

```bash
./taskcli taskInProgress -id 2
```

* Marks the task with the given ID as in-progress.
* `-id` is **required**.

---

### 7. List Tasks by Status

```bash
./taskcli list-done      # Completed tasks
./taskcli list-progress  # Tasks in progress
./taskcli list-todo      # Pending tasks
```

* These commands **don’t require flags**.

---

### 8. Help

```bash
./taskcli help
```

* Shows all commands and their usage.

````

---


