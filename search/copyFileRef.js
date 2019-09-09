(function () {
	for (let item of document.querySelectorAll(".element .fileRef")) {
		item.addEventListener("click",copy)
	}
})();

function copy() {
	var text = this.textContent ;
	if(!text) return ;
	var input = document.createElement("input");
	input.type = "text" ;
	input.value = text ;
	document.body.appendChild(input);
	input.select();
	document.execCommand("copy");
	input.remove();
	this.classList.add("copied")
	setTimeout(()=>{
		this.classList.remove("copied")
	},500);
}
