// getDoc
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

// Function to test the speed of searchElement()
function benchSearchElement(nb) {
	nb = Number(nb);
	if (isNaN(nb)) {
		throw new TypeError("nb is not a number")
	}
	var begin = (new Date()).getTime();
	for (var i = 0; i < nb; i++) {
		arg = searchParse(searchInput.value);
		searchElement();
	}
	return ( (new Date()).getTime() - begin )/nb ;
}

function testSearchPattern() {
	var input = [
		{
			inp: "aaa:bbb",
			cat: "aaa",
			val: "bbb",
		},
		{
			inp: "aaa:",
			cat: "aaa",
			val: "",
		},
		{
			inp: "bbb",
			cat: "",
			val: "bbb",
		},
	];
	let i = 0;
	let allValidate = true;
	for (let t of input) {
		if (t.inp.replace(searchPattern,"$1") !== t.cat) {
			allValidate = false ;
			console.log(i,"cat:", t.cat, t.inp.replace(searchPattern,"$1"))
		}
		if (t.inp.replace(searchPattern,"$2") !== t.val) {
			allValidate = false ;
			console.log(i,"val:", t.val, t.inp.replace(searchPattern,"$2"))
		}
		i++;
	}
	if (allValidate) {
		console.log("Test réussi");
	} else {
		console.log("test échoué");
	}
}
