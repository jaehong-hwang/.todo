# TODO on bash

todo logger working in bash
they are logging in .todo directory

## commands

### Init
``` shell
# make Todo collection
> todo init
```

### Add
* \$message is required
* \$start, \$end are default null
``` shell
# Adding todo
> todo add $message $start $end
```

### Change Status
Change the state of work
A todo can have the following statuses:
* wait
* work
* done
``` shell
# let's work
> todo state work $todoId
```

### Get
The Get command gets a list of tasks that have not completed their status by default.
If you want to get a specific list, you can use the following options.
* --all: Gets the all list
* --expired: Gets the list past the end date.
* --today: It has not been completed, and I will get what I need to do today.
* --state=$state: Gets the tasks that fit the state
``` bash
# Get todo list
# --all option given list with checked item
> todo list [--all, --expired]
# then return...
id  |  name  |  message           | status   |  start       |  end
1   |  -     |  my todo text      | work     |  Today       |   -
2   |  -     |  second todo text  | waiting  |  2019-02-15  |  2019-02-18

# complete todo
> todo check "my todo text"
# or
> todo check $todoId
```
