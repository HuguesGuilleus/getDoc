// getDoc
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

window.addEventListener("keydown", event=>{
	if (document.activeElement != searchInput && event.keyCode > 30) {
		searchInput.value += event.key ;
		searchInput.focus();
	}
});
