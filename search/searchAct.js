// getDoc
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

// execute the command
function searchAct() {
	document.getElementById("std").hidden = true ;
	document.getElementById("searchResult").hidden = false ;
	document.getElementById("help").hidden = !searchActList.help.length ;
	searchActLs("File");
	searchActLs("Lang");
	searchActLs("Type");
}

// Hide or not an article of listing
function searchActLs(Name) {
	const name = Name.toLowerCase() ;
	var present = searchActList.ls.includes("")
		|| searchActList.ls.includes(name) ;
	document.getElementById("ls"+Name+"Art").hidden = !present ;
	if (present) {
		syncList(document.getElementById("ls"+Name+"List"), arg[name])
	}
}

// Synchonise a DOM list and a Js array
function syncList(ul,pat) {
	for (let li of ul.getElementsByTagName("li")) {
		li.hidden = searchMatch(li,pat) < 1 ;
	}
}
