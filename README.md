# 1practice
No-SQL DBMS development
Usage: ./dbms --query
All commands:
SADD <value>: Add a value to the set
SREM <value>: Remove a value from the set
SISMEMBER <value>: Performs a check whether an element is part of a set
SPUSH <value>: Add a value to the stack
SPOP: Remove the top value from the stack
QPUSH <value>: Add the top value to the queue
QPOP: Remove the first value from the queue
HSET <key> <value>: Add the value, key to the hash-table
HDEL <key>: Delete the key to the hash-table
HGET <key>: Reads the value by key in the hash-table
EXIT: Exit the program
Program example: ./dbms --QPUSH mk
