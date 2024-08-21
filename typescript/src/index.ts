// Basic Types
let id: number = 5
let company: string = 'Traversy Media'
let isPublished:boolean = true
let x: any = 'Hello'

let ids: number[] = [1, 2, 3, 4, 5]
let arr: any[] = [1, true, 'Hello']

// Tuple
let person: [number, string, boolean] = [1, 'Brad', true]

// Tuple Array
let employee: [number, string][]

employee = [
  [1, 'Brad'],
  [2, 'John'],
  [3, 'Jill']
]

// Union
let pid : string | number = 1
pid = '1'

// Enum
enum Direction1 {
  Up = 1,
  Down,
  Left,
  Right
}

console.log(Direction1.Left)

enum Direction2 {
  Up = 'Up',
  Down = 'Down',
  Left = 'Left',
  Right = 'Right'
}

console.log(Direction2.Left)

// Objects
  // First way
const user: {
  id: number,
  name: string
} = {
  id: 1,
  name: 'John'
}
  // Second way
type User = {
  id: number,
  name: string
}

const user2: User = {
  id: 1,
  name: 'John'
}

// Type Assertion
let cid: any = 1
  // First way
let customerId = <number>cid
  // Second way
let customerId2 = cid as number

// Functions
function addNum(x: number, y: number): number {
  return x + y
}

// Void
function log(message: string | number): void {
  console.log(message)
}

// Interface
interface UserInterface {
  readonly id: number, // cannot be changed
  name: string,
  age?: number // ? means optional
}

const user3: UserInterface = {
  id: 1,
  name: 'John'
}

interface MathFunc {
  (x: number, y: number): number
}

const add: MathFunc = (x: number, y: number): number => x + y
const sub: MathFunc = (x: number, y: number): number => x - y

interface PersonInterface {
  id: number,
  name: string,
  register(): string
}

//Classes
class Person implements PersonInterface {
  id: number
  name: string

  constructor(id: number, name: string) {
    this.id = id
    this.name = name
  }

  register() {
    return `${this.name} is now registered.`
  }
}

const brad = new Person(1, 'Brad')

// Subclasses
class Employee extends Person {
  position: string

  constructor(id: number, name: string, position: string) {
    super(id, name)
    this.position = position
  }
}

const emp = new Employee(3, 'Shawn', 'Developer')

console.log(emp.register())

// Generics
function getArray<T>(items: T[]): T[] {
  return new Array().concat(items)
}

let numArray = getArray<number>([1, 2, 3])
let stringArray = getArray<string>(['1', '2', '3'])

numArray.push(4)
stringArray.push('4')