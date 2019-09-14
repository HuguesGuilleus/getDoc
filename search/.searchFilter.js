// getDoc
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

(function () {
	for (let filter of document.querySelectorAll(".filter")) {
		filter.addEventListener("click",addFilter) ;
	}
})();

// Add the content of a filter to searchInput
function addFilter() {
	var text = this.textContent ;
	if (!searchInputArray.includes(text)) {
		this.classList.add("filterActive");
		searchInput.value += " "+text ;
		search();
	}
}

// Disable filter who are remove from inpustSearch
function disableFilter() {
	for (let filter of document.getElementsByClassName("filterActive")) {
		if (!searchInputArray.includes(filter.textContent)) {
			filter.classList.remove("filterActive");
		}
	}
}
