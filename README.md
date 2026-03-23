# CLI-functions
This is a simple todo list using golang , CLI and json(for storing tasks in a local file)
# 📝 Task Tracker CLI

Simple CLI tool to manage tasks.

---

## 🚀 Usage

```bash
go run . <command> [arguments]
```

---

## ➕ Add Task

```bash
go run . add "brush my teeth"
```

---

## 📋 List Tasks

```bash
go run . list
```

```
ID: 1 | brush my teeth | status: todo | created: 2026-03-23 22:10:00
```

---

## ✏️ Update Task

```bash
go run . update 1 "brush my teeth and wash face"
```

---

## ❌ Delete Task

```bash
go run . delete 1
```

---

## ✅ Mark as Done

```bash
go run . done 1
```

---

## 🔄 Mark as In Progress

```bash
go run . in-progress 1
```

---

## 📌 Filter Tasks

### Done

```bash
go run . done-list
```

### In Progress

```bash
go run . progress-list
```

### Todo

```bash
go run . todo-list
```

---

## 📁 Storage

```
tasks.json
```

(auto-created)

---

## ⚠️ Notes

```bash
# Use quotes for multiple words
go run . add "buy groceries and cook"

# ID must be a number
go run . delete 2
```

---

## 🧪 Example Flow

```bash
go run . add "learn Go"
go run . add "build CLI"

go run . list

go run . in-progress 1
go run . done 1

go run . done-list
```

