// The parsed elements of search
var arg = null ;

// The value of searchInput splited
var searchInputArray = [];

// The ref to the element input#searchInput
var searchInput = null ;

(function () {
	searchInput = document.getElementById("searchInput");
	searchInput.value = "";
	searchInput.oninput = search ;
	window.onload = search ;
})();

// Launch the search
function search() {
	arg = searchParse(searchInput.value);
	if (arg === null) {
		searchReset();
	} else if (arg.cmd.length > 0) {
		searchAct(arg)
	} else {
		searchElement();
	}
}

// The patern for parsing input element
const searchPattern = /^(?:(\w*):)?(.*)$/ ;

// Parse the string from inout element and return an object
function searchParse(input) {
	searchInputArray = input.split(/\s+/) ;
	if (input.length === 0) {
		return null
	} else {
		var list = {
			file:[],
			lang:[],
			name:[],
			type:[],
			all:[],
			cmd:[],
		}
		for (let el of searchInputArray) {
			var cat = el.replace(searchPattern, "$1")
			var val = el.replace(searchPattern, "$2")
			if (val === "") continue ;
			if (/\$.+/.test(val)) {
				list.cmd.push(val.substr(1))
			} else {
				switch (cat) {
					case "file":list.file.push(val);break;
					case "lang":list.lang.push(val);break;
					case "name":list.name.push(val);break;
					case "type":list.type.push(val);break;
					case "":    list.all.push(val);break;
				}
			}
		}
		return {
			cmd: list.cmd,
			file: {
				pat: new RegExp("("+list.file.concat(list.all).join("|")+")", "gui"),
				notFound: list.file.length ,
			},
			lang: {
				pat: new RegExp("("+list.lang.concat(list.all).join("|")+")", "gui"),
				notFound: list.lang.length ,
			},
			name: {
				pat: new RegExp("("+list.name.concat(list.all).join("|")+")", "gui"),
				notFound: list.name.length ,
			},
			type: {
				pat: new RegExp("("+list.type.concat(list.all).join("|")+")", "gui"),
				notFound: list.type.length ,
			},
		}
	}
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
	if (el.querySelectorAll(".find").length > 0) {
		el.innerHTML = el.textContent ;
	}
}

// Search in the list of element, hidden the not found item
function searchElement() {
	document.getElementById("std").hidden = false ;
	document.getElementById("searchResult").hidden = true ;
	for (let item of document.querySelectorAll("#list>li")) {
		item.hidden = (
			searchMatch(item.querySelector(".fileRef"),arg.file)
			+ searchMatch(item.querySelector(".lang"),arg.lang)
			+ searchMatch(item.querySelector(".elementTitle"),arg.name)
			+ searchMatch(item.querySelector(".type"),arg.type)
		) < 1 ;
	}
}

// Print the matched string (from the array of string) in the element.
// It return -10 if the tab if not empty and the pattern not found,
// return 1 if there are a match, and return 0 if tab is empty and not found
function searchMatch(el, finder ) {
	if (finder.pat.test(el.textContent)) {
		el.innerHTML = el.textContent.replace(finder.pat,'<span class=find>$1</span>');
		return 1
	} else {
		searchResetElement(el)
		return -10 * finder.notFound ;
	}
}
