class Carro {
    constructor(modelo, chassi, portas) {
        this.modelo = modelo;
        this.chassi = chassi;
        this.portas = portas;
    }

    andar() {
        console.log('andar');
    }
}

class Sonix extends Carro {
    constructor(modelo, chassi, portas, cv) {
        super(modelo, chassi, portas);
        this.cv = cv;
    }

    abrirTetoSolar() {
        console.log('abrirTetoSolar');
    }
}

const carro = new Carro('Honda', 'abc', 4);

const sonix = new Sonix('Sonix', 'abcdef', 5, 190);

console.log(carro);
carro.andar();

console.log(sonix);
sonix.abrirTetoSolar();
