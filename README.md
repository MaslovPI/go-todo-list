# Todo App

This is a projec inspired by: [goprojects](https://github.com/dreamsofcode-io/goprojects/blob/main/01-todo-list/README.md)

## Goal

Create an cli application for managing tasks in the terminal.

```bash
tasks
```

## Requirements

- [ ] CRUD operations

  - [ ] add

    ```
    tasks add <description>
    ```

    for example:

    ```
    tasks add "Tidy my desk"
    ```

  - [ ] list
        This method should return a list of all of the **uncompleted** tasks, with the option to return all tasks regardless of whether or not they are completed.

    for example:

    ```
    $ tasks list
    ID    Task                                                Created
    1     Tidy up my desk                                     a minute ago
    3     Change my keyboard mapping to use escape/control    a few seconds ago
    ```

    or for showing all tasks, using a flag (such as -a or --all)

    ```
    $ tasks list -a
    ID    Task                                                Created          Done
    1     Tidy up my desk                                     2 minutes ago    false
    2     Write up documentation for new project feature      a minute ago     true
    3     Change my keyboard mapping to use escape/control    a minute ago     false
    ```

  - [ ] complete
        To mark a task as done, add in the following method

    ```
    tasks complete <taskid>
    ```

  - [ ] delete
        The following method should be implemented to delete a task from the data store

    ```
    tasks delete <taskid>
    ```

- [ ] Use csv to store data

  ```
  ID,Description,CreatedAt,IsComplete
  1,My new task,2024-07-27T16:45:19-05:00,true
  2,Finish this video,2024-07-27T16:45:26-05:00,true
  3,Find a video editor,2024-07-27T16:45:31-05:00,false
  ```

- [ ] Write errors and diagnostics to stderr stream and write output to stdout

## Packages to use

- [x] `encoding/csv` for writing out as a csv file
- [ ] `strconv` for turning types into strings and visa versa
- [ ] `text/tabwriter` for writing out tab aligned output
- [ ] `os` for opening and reading files
- [x] `github.com/spf13/cobra` for the command line interface
- [ ] `github.com/mergestat/timediff` for displaying relative friendly time differences (1 hour ago, 10 minutes ago, etc)

## Extra tasks

- [ ] Change is complete to a timestamp
- [ ] Replace CSV with sqlite
- [ ] Add due date field
