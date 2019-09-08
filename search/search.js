(function () {
	h1 = document.querySelector("h1"); ///////// for dev
	searchInput = document.getElementById("searchInput");
	searchInput.oninput = ()=>{
		arg = searchParse(searchInput.value);
		if (arg === null) {
			searchReset();
		} else if (arg.cmd.length > 0) {
			searchAct(arg)
		} else {
			searchElement(arg);
		}
	}
})();

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
			if (el[0] === "$") {
				let r = el.replace(/\$(\w*).*/, "$1")
				if (r) {
					arg.cmd.push(r)
				}
			} else {
				let cat = false ;
				if (/\w+:.*/.test(el)) {
					cat = el.replace(/(\w+):.*/, "$1");
				}
				let val = el.replace(/:?(.*)/, "$1");
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
		let name = item.querySelector(".elementTitle");
		name.innerHTML = name.textContent ;
		item.hidden = false ;
	}
}

// Search in the list of element, hidden the not found item
function searchElement(arg) {
	console.log("searchElement");
	document.getElementById("std").hidden = false ;
	document.getElementById("searchResult").hidden = true ;
	for (let item of document.querySelectorAll("#list>li")) {
		let notFound = true
		let name = item.querySelector(".elementTitle");
		notFound = notFound && searchMatch(name, arg.all.concat(arg.name)) ;
		item.hidden = notFound ;
	}
}

// Print the matched string (from the array of string) in the element.
// It return true if no matched and false if match.
function searchMatch(el,tab) {
	if (tab.length === 0) {
		el.innerHTML = el.textContent ;
		return true ;
	} else {
		var pattern = new RegExp("("+tab.join("|")+")", "g");
		el.innerHTML = el.textContent.replace(pattern,'<span class=find>$1</span>') ;
		return !pattern.test(el.textContent) ;
	}
}
