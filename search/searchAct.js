function searchAct() {
	document.getElementById("std").hidden = true ;
	document.getElementById("searchResult").hidden = false ;
	document.getElementById("help").hidden = !arg.cmd.includes("help") ;
	if (arg.cmd.includes("ls")) {
		document.getElementById("lsFileArt").hidden = false ;
		syncList(document.getElementById("lsFileList"), arg.file)
	} else {
		document.getElementById("lsFileArt").hidden = true ;
	}
}

// Synchonise a DOM list and a Js array
function syncList(ul,pat) {
	for (let li of ul.querySelectorAll("li")) {
		li.hidden = searchMatch(li,pat) < 1 ;
	}
}

// Generate the list item for a Js array
function genList(ul, tab) {
	for (let item of tab) {
		let li = document.createElement("li");
		li.textContent = item ;
		ul.appendChild(li) ;
	}
}
