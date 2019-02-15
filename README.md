# TODO on bash

todo logger in cmd
logging in project's .todo directory

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

### Check
``` shell
# complete todo
> todo check "my todo text"
# or
> todo check $todoId
```

### Get
* options: --all, --expired
``` bash
# Get todo list
# --all option given list with checked item
> todo list [--all, --expired]
# then return...
id  |  name  |  message           |  start       |  end
1   |  -     |  my todo text      |  Today       |   -
2   |  -     |  second todo text  |  2019-02-15  |  2019-02-18
```

### Get expired todo list
``` shell
# get non completed and expired todo
> todo list --expired
# or
> todo expired
id  |  name  |  message           |  start       |  end
1   |  -     |  second todo text  |  2019-02-02  |  2019-02-13
2   |  -     |  second todo text  |  2019-02-15  |  2019-02-18
```
