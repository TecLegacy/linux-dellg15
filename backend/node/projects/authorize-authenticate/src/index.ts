const firstname: string = 'John'
const lastname = 'Doe'
const age: number = 30

const sayHello = (name: string) => {
    console.log(`Hello ${name}`)
}
sayHello('John Doe')

console.log(firstname, lastname, age)

const obj = { name: 'John Doe', age: 30 }
console.log(obj)
const a = 3
function b() {
    const a = 10
    console.log(a)
}
b()

console.log(a)
