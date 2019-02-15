# TODO on bash

todo logger in cmd
logging in project's .todo directory

## commands
``` bash
# make Todo collection
> todo init

# Adding todo
# $message is required
# $start, $end are default null
> todo add $message $start $end

# Get todo list
# --all option given list with checked item
> todo list [--all]
# then return...
id  |  name  |  message           |  start       |  end
1   |  -     |  my todo text      |  Today       |   -
2   |  -     |  second todo text  |  2019-02-15  |  2019-02-18

# complete todo
> todo check "my todo text"
# or
> todo check $todoId
```
