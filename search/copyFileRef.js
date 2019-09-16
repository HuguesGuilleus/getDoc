// getDoc
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

document.addEventListener("DOMContentLoaded",()=>{
	for (let item of document.getElementsByClassName("fileRef")) {
		item.addEventListener("click",copy);
	}
},{once:true,});

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
	this.classList.add("copied");
	setTimeout(()=>{
		this.classList.remove("copied");
	},500);
}
