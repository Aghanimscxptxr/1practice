package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

const maxSize = 100

type HashSet struct {
	data [maxSize]string
}

type Node struct {
	data string
	next *Node
}

type Stack struct {
	head       *Node
	lastPopped string
}

type Queue struct {
	front        *Node
	rear         *Node
	lastDequeued string
}

type HashTable struct {
	data [maxSize][]KeyValue
}

type KeyValue struct {
	Key   string
	Value string
}

// ---------------------------------------------STACK---------------------------------------------------
func (stack *Stack) Push(val string) {
	newNode := &Node{data: val, next: stack.head}
	stack.head = newNode
	//fmt.Printf("%s\n", val)
}

func (stack *Stack) Pop() error {
	if stack.head == nil {
		return errors.New("Stack is empty")
	}
	stack.lastPopped = stack.head.data
	stack.head = stack.head.next
	return nil
}

// ------------------------------------------------SET--------------------------------------------
func thash(s string) int {
	hashValue := 0
	for _, char := range s {
		hashValue += int(char)
	}
	return hashValue % maxSize
}

// SADD
func (set *HashSet) Insert(key string) {
	if set.Contains(key) {
		fmt.Printf("Value %s already exists in the set. Not added.\n", key)
	} else {
		hashValue := thash(key)
		for i := hashValue; ; i = (i + 1) % maxSize {
			if set.data[i] == "" {
				set.data[i] = key
				//fmt.Printf("%s\n", key)
				return
			}
		}
	}
}

// SREM
func (set *HashSet) Delete(key string) {
	hashValue := hash(key)
	i := hashValue
	for {
		if set.data[i] == key {
			set.data[i] = ""
			//fmt.Printf("%s\n", key)
			return
		}
		i = (i + 1) % maxSize
		if i == hashValue {
			fmt.Printf("%s does not exist in the set.\n", key)
			return
		}
	}
}

// SISMEMBER
func (set *HashSet) Contains(key string) bool {
	hashValue := hash(key)
	i := hashValue
	for {
		if set.data[i] == key {
			//fmt.Printf("%s exists in the set. Returning true.\n", key)
			return true
		}
		i = (i + 1) % maxSize
		if i == hashValue || set.data[i] == "" {
			//fmt.Printf("%s does not exist in the set. Returning false.\n", key)
			return false
		}
	}
}

// -----------------------------------QUEUE--------------------------------------
// QPUSH
func (queue *Queue) Enqueue(val string) {

	newNode := &Node{data: val, next: nil}
	if queue.front == nil {
		queue.front = newNode
		queue.rear = newNode
	} else {
		queue.rear.next = newNode
		queue.rear = newNode
	}
	//fmt.Printf("%s\n", val)
}

// QPOP
func (queue *Queue) Dequeue() error {
	if queue.front == nil {
		return errors.New("Queue is empty")
	}
	queue.lastDequeued = queue.front.data
	queue.front = queue.front.next
	if queue.front == nil {
		queue.rear = nil
	}
	return nil
}

// ---------------------------HASH-TABLE-------------------------------
func hash(s string) int {
	hashValue := 0
	for _, char := range s {
		hashValue += int(char)
	}
	return hashValue % maxSize
}

func (ht *HashTable) HSET(key string, value string) {
	hashValue := hash(key)

	// Проверка на уникальность ключа
	for i := 0; i < len(ht.data[hashValue]); i++ {
		if ht.data[hashValue][i].Key == key {
			fmt.Printf("%s is already occupied, select another key.\n", key)
			return
		}
	}

	for i := 0; ; i = (i + 1) % maxSize {
		// Инициализация среза, если он пустой
		if ht.data[hashValue] == nil {
			ht.data[hashValue] = make([]KeyValue, maxSize)
		}

		// Если найдена пустая ячейка или достигнут предел, вставляем новую пару ключ-значение
		if ht.data[hashValue][i].Key == "" {
			ht.data[hashValue][i] = KeyValue{Key: key, Value: value}
			//fmt.Printf("HSET: key=%s, value=%s\n", key, value)
			return
		}
	}
}

func (ht *HashTable) HDEL(key string) {
	hashValue := hash(key)
	for i, kv := range ht.data[hashValue] {
		if kv.Key == key {
			ht.data[hashValue] = append(ht.data[hashValue][:i], ht.data[hashValue][i+1:]...)
			//fmt.Printf("%s\n",key)
			return
		}
	}
	fmt.Printf("%s does not exist in the hash table.\n", key)
}

func (ht *HashTable) HGET(key string) (string, error) {
	hashValue := hash(key)
	for i := 0; ; i = (i + 1) % maxSize {
		if ht.data[hashValue] == nil {
			break
		}
		if ht.data[hashValue][i].Key == key {
			fmt.Printf("%s\n", ht.data[hashValue][i].Value)
			return ht.data[hashValue][i].Value, nil
		}
		if ht.data[hashValue][i].Key == "" {
			break
		}
	}

	message := fmt.Sprintf("Key not found: %s", key)
	fmt.Printf("%s not found in the hash table.\n", key)
	return "", errors.New(message)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	mySet := HashSet{}
	stack := Stack{}
	queue := Queue{}
	hashTable := HashTable{}
	for {
		scanner.Scan()
		commandLine := scanner.Text()

		if !strings.HasPrefix(commandLine, "./dbms --") {
			fmt.Println("Please start with ./dbms --query")
			continue
		}
		parts := strings.Fields(commandLine[len("./dbms --"):])

		if len(parts) < 1 {
			fmt.Println("Invalid command format")
			continue
		}

		command := parts[0]
		arguments := parts[1:]

		switch command {
		case "EXIT":
			fmt.Println("Exiting the program")
			return
		case "HELP":
			fmt.Println("Available Commands:")
			fmt.Println("SADD <value>: Add a value to the set")
			fmt.Println("SREM <value>: Remove a value from the set")
			fmt.Println("SISMEMBER <value>: Performs a check whether an element is part of a set")
			fmt.Println("SPUSH <value>: Add a value to the stack")
			fmt.Println("SPOP: Remove the top value from the stack")
			fmt.Println("QPUSH <value>: Add the top value to the queue")
			fmt.Println("QPOP: Remove the first value from the queue")
			fmt.Println("HSET <key> <value>: Add the value, key to the hash-table")
			fmt.Println("HDEL <key>: Delete the key to the hash-table")
			fmt.Println("HGET <key>: Reads the value by key in the hash-table")
			fmt.Println("EXIT: Exit the program")
			continue
		case "SADD":
			if len(arguments) != 1 {
				fmt.Println("Invalid format for SADD. Usage: SADD <value>")
			} else {
				mySet.Insert(arguments[0])
			}
		case "SREM":
			if len(arguments) != 1 {
				fmt.Println("Invalid format for SREM. Usage: SREM <value>")
			} else {
				mySet.Delete(arguments[0])
			}
		case "SISMEMBER":
			if len(arguments) != 1 {
				fmt.Println("Invalid format for SISMEMBER. Usage: SISMEMBER <value>")
			} else {
				if !(mySet.Contains(arguments[0])) {
					println("FALSE")
				} else {
					println("TRUE")
				}
			}
		case "SPUSH":
			if len(arguments) != 1 {
				fmt.Println("Invalid format for SPUSH. Usage: SPUSH <value>")
			} else {
				stack.Push(arguments[0])
			}
		case "SPOP":
			if len(arguments) != 0 {
				fmt.Println("Invalid format for SPOP. Usage: SPOP")
			} else {
				serr := stack.Pop()
				if serr != nil {
					fmt.Println(serr)
				} else {
					fmt.Printf("%s\n", stack.lastPopped)
				}
			}
		case "QPUSH":
			if len(arguments) != 1 {
				fmt.Println("Invalid format for QPUSH. Usage: QPUSH <value>")
			} else {
				queue.Enqueue(arguments[0])
			}
		case "QPOP":
			if len(arguments) != 0 {
				fmt.Println("Invalid format for QPOP. Usage: QPOP")
			} else {
				qerr := queue.Dequeue()
				if qerr != nil {
					fmt.Println(qerr)
				} else {
					fmt.Printf("%s\n", queue.lastDequeued)
				}
			}
		case "HSET":
			if len(arguments) != 2 {
				fmt.Println("Invalid format for HSET. Usage: HSET <key><value>")
			} else {
				hashTable.HSET(arguments[0], arguments[1])
			}
		case "HDEL":
			if len(arguments) != 1 {
				fmt.Println("Invalid format for HDEL. Usage: HDEL <key>")
			} else {
				hashTable.HDEL(arguments[0])
			}
		case "HGET":
			if len(arguments) != 1 {
				fmt.Println("Invalid format for HGET. Usage: HGET <key>")
			} else {
				hashTable.HGET(arguments[0])
			}
		default:
			fmt.Println("Unrecognized command")
		}
	}
}
