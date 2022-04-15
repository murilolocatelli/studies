(function() {
	'use strict';

	var teste = {
		prop: 1,
		prop2: 2
	};

	console.log(teste);
	console.log(Array.prototype);

	function MeuObjeto() {
		this.minhaProp = "minha prop"
		return "retorno";
	}

	MeuObjeto.prototype.outraProp = "Outra prop";

	MeuObjeto.prototype.outraFunc = function() {
		return "Outra func"; 
	}

	console.log(new MeuObjeto().minhaProp);
	console.log(new MeuObjeto().outraProp);
	console.log(new MeuObjeto().outraFunc());

})();