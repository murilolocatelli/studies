(function() {
	'use strict';

	var scope = 'global scope';

	function checkScope() {
		var scope = 'local scope';
		function func() {
			return scope;
		}
		return func;
	}

	console.log(checkScope()());
})();