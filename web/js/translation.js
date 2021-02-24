// getDoc
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

document.addEventListener("DOMContentLoaded", () => {
	var tem = document.getElementById("translate").content;
	for (let lang of navigator.languages) {
		let trad = tem.querySelector("#lang_" + lang);
		if (trad) {
			translate(trad, lang);
			return;
		}
	}
	translate(tem.querySelector("#lang_en"), "en")
}, {
	once: true,
});

// translate apply the translation for a specific language
function translate(trad, lang) {
	document.documentElement.lang = lang;
	// inputSearch
	document.getElementById("searchInput").placeholder =
		trad.querySelector("#placeholder").textContent;
	// footer
	document.body.appendChild(trad.querySelector("footer"));
	// Help
	var help = trad.querySelector("#help");
	document.getElementById("searchResult").appendChild(help);
	// remove traduction elements
	document.getElementById("translate").remove();
}
