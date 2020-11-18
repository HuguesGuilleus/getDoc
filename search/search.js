// getDoc
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

// The patern for parsing input element of action
const searchActPattern = /^\$(\w+)(?:\:(.*))?$/;
// The patern for parsing input element of simple search
const searchElPattern = /^(?:(\w*):)?(.*)$/;

// The parsed list of action from the searchInput
var searchActList = {};

// The parsed elements of search
var arg = null;

// The value of searchInput splited
var searchInputArray = [];

// The ref to the element input#searchInput
var searchInput = null;

document.addEventListener("DOMContentLoaded", () => {
	searchInput = document.getElementById("searchInput");
	searchInput.value = "";
	searchInput.oninput = search;
}, {
	once: true,
});

// Launch the search
function search() {
	if (searchInput.value) {
		arg = searchParse(searchInput.value);
		if (arg.cmd) {
			searchAct();
		} else {
			searchElement();
		}
	} else {
		searchReset();
	}
}

// Parse the string from inpout element and return an object.
function searchParse(input) {
	searchInputArray = input.split(/\s+/);
	searchActList = {
		ls: [],
		help: [],
	};
	if (input.length === 0) {
		return null
	} else {
		var cmd = false;
		var list = {
			file: [],
			lang: [],
			name: [],
			type: [],
			all: [],
		};
		for (let el of searchInputArray) {
			if (!el) continue;
			if (el[0] === "$") {
				let cat = el.replace(searchActPattern, "$1");
				if (searchActList[cat]) {
					let val = el.replace(searchActPattern, "$2");
					searchActList[cat].push(val);
					cmd = true;
				};
			} else {
				let val = el.replace(searchElPattern, "$2");
				switch (el.replace(searchElPattern, "$1")) {
				case "file":
					list.file.push(val);
					break;
				case "lang":
					list.lang.push(val);
					break;
				case "name":
					list.name.push(val);
					break;
				case "type":
					list.type.push(val);
					break;
				case "":
					list.all.push(val);
					break;
				}
			}
		}
		return {
			cmd: cmd,
			file: {
				pat: new RegExp("(" + list.file.concat(list.all).join("|") + ")", "gui"),
				notFound: list.file.length,
			},
			lang: {
				pat: new RegExp("(" + list.lang.concat(list.all).join("|") + ")", "gui"),
				notFound: list.lang.length,
			},
			name: {
				pat: new RegExp("(" + list.name.concat(list.all).join("|") + ")", "gui"),
				notFound: list.name.length,
			},
			type: {
				pat: new RegExp("(" + list.type.concat(list.all).join("|") + ")", "gui"),
				notFound: list.type.length,
			},
		};
	}
}

// Reset the search
function searchReset() {
	document.getElementById("searchResult").hidden = true;
	var sdt = document.getElementById("std");
	std.hidden = false;
	for (let c of ["fileRef", "lang", "elementTitle", "type"]) {
		for (let item of std.getElementsByClassName(c)) {
			searchResetElement(item)
		}
	}
	for (let li of document.getElementsByTagName("li")) {
		li.hidden = false;
	}
}

// Reset an simple elem
function searchResetElement(el) {
	// var list = el.getElementsByClassName("find");
	// if (list.length) {
	// 	const  text = el.textContent;
	// 	for (let find of list) {
	// 		find.remove();
	// 	}
	// 	el.textContent = text;
	// }
}

// Search in the list of element, hidden the not found item
function searchElement() {
	document.getElementById("std").hidden = false;
	document.getElementById("searchResult").hidden = true;
	for (let item of document.querySelectorAll("#list>li")) {
		item.hidden = (
			searchMatch(item.querySelector(".fileRef"), arg.file) +
			searchMatch(item.querySelector(".lang"), arg.lang) +
			searchMatch(item.querySelector(".elementTitle"), arg.name) +
			searchMatch(item.querySelector(".type"), arg.type)
		) < 1;
	}
}

// Print the matched string (from the array of string) in the element.
// It return -10 if the tab if not empty and the pattern not found,
// return 1 if there are a match, and return 0 if tab is empty and not found
function searchMatch(el, finder) {
	if (finder.pat.test(el.innerText)) {
		// el.innerHTML = el.textContent.replace(finder.pat, '<span class=find>$1</span>');
		return 1;
	} else {
		// searchResetElement(el);
		return -10 * finder.notFound;
	}
}
