(function () {
	searchInput = document.getElementById("searchInput");
	searchInput.oninput = search ;
	window.onload = search ;
})();

function search() {
	arg = searchParse(searchInput.value);
	if (arg === null) {
		searchReset();
	} else if (arg.cmd.length > 0) {
		searchAct(arg)
	} else {
		searchElement(arg);
	}
}

// Parse the string from inout element and return an object
function searchParse(input) {
	if (input.length === 0) {
		return null
	} else {
		var arg = {
			file:[],
			lang:[],
			name:[],
			type:[],
			all:[],
			cmd:[],
		}
		for (let el of input.split(/\s+/)) {
			if (el === "") {
				continue
			}
			if (el[0] === "$") {
				let r = el.replace(/\$(\w*).*/, "$1")
				if (r) {
					arg.cmd.push(r)
				}
			} else {
				let cat = false ;
				let val ;
				if (/\w+:.*/.test(el)) {
					cat = el.replace(/(\w+):.*/, "$1");
					val = el.replace(/^\w+:(.*)/, "$1");
				} else {
					val = el;
				}
				switch (cat) {
					case "file":arg.file.push(val);break;
					case "lang":arg.lang.push(val);break;
					case "name":arg.name.push(val);break;
					case "type":arg.type.push(val);break;
					case false:
						arg.all.push(val);
						break;
				}
			}
		}
		return arg
	}
}

function searchAct() {
	console.log("searchAct");
}

// Reset the search
function searchReset() {
	document.getElementById("std").hidden = false ;
	document.getElementById("searchResult").hidden = true ;
	for (let item of document.querySelectorAll("#list>li")) {
		item.hidden = false ;
		searchResetElement(item.querySelector(".fileRef"));
		searchResetElement(item.querySelector(".lang"));
		searchResetElement(item.querySelector(".elementTitle"));
		searchResetElement(item.querySelector(".type"));
	}
}

// Reset an simple elem
function searchResetElement(el) {
	el.innerHTML = el.textContent ;
}

// Search in the list of element, hidden the not found item
function searchElement(arg) {
	document.getElementById("std").hidden = false ;
	document.getElementById("searchResult").hidden = true ;
	for (let item of document.querySelectorAll("#list>li")) {
		let notFound = 0 ;
		notFound += searchMatch(item.querySelector(".fileRef"),arg.file,arg.all);
		notFound += searchMatch(item.querySelector(".lang"),arg.lang,arg.all);
		notFound += searchMatch(item.querySelector(".elementTitle"),arg.name,arg.all);
		notFound += searchMatch(item.querySelector(".type"),arg.type,arg.all);
		item.hidden = notFound < 1 ;
	}
}

// Print the matched string (from the array of string) in the element.
// It return -10 if the tab if not empty and the pattern not found,
// return 1 if there are a match, and return 0 if tab is empty and not found
function searchMatch(el,tab, all) {
	var pattern = new RegExp("("+tab.concat(all).join("|")+")", "gui");
	el.innerHTML = el.textContent.replace(pattern,'<span class=find>$1</span>') ;
	if (pattern.test(el.textContent)) {
		return 1 ;
	} else if (tab.length !== 0) {
		return -10 ;
	} else {
		return 0 ;
	}
}
