class List {
    constructor() {
        this.list = [];
    }

    add(data) {
        this.list.push(data);

        console.log(this.list);
    }
}

class TodoList extends List {
    constructor() {
        super();

        this.usuario = 'Murilo';
    }

    mostrarUsuario() {
        console.log(this.usuario);
    }
}

class Matematica {
    static somar(a, b) {
        return a + b;
    }
}

const minhaLista = new TodoList();

minhaLista.add('Novo Todo');

console.log(Matematica.somar(1, 6));

const arr = [1, 2, 3, 4];

// Map
const newArr = arr.map(function(item, index) {
    return item + index;
});

console.log(newArr);

// Reduce
const sum = arr.reduce(function(total, next) {
    return total + next;
});

console.log(sum);

// Reduce arrow function
const sum2 = arr.reduce((total, next) => total + next);

console.log(sum2);

// Filter
const filter = arr.filter(function(item) {
    return item % 2 === 0;
});

console.log(filter);

// Find
const find = arr.find(function(item) {
    return item === 4; 
});

console.log(find);

// Desestruturação

const usuario = {
    nome: 'Murilo',
    idade: 33,
    endereco: {
        cidade: 'Franca',
        estado: 'SP'
    }
}

const { nome, idade, endereco: { cidade }} = usuario;

console.log(nome);
console.log(idade);
console.log(cidade);

// Spread
const arr3 = [...arr, ...newArr];

console.log(arr3);

import somar, * as funcoes from './funcoes'

console.log('Somar: ' + somar(1, 3));
console.log('Subtrair: ' + funcoes.subtratir(1, 3));
console.log('Dividir: ' + funcoes.dividir(1, 3));
console.log('Multiplicar: ' + funcoes.multiplicar(1, 3));