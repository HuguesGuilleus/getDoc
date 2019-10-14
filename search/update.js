// getDoc
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

document.addEventListener("DOMContentLoaded", ()=>{
	document.getElementById("updateNotifClose").onclick = ()=>{
		document.getElementById("updateNotif").remove();
	};
	setTimeout(()=>{
		document.getElementById("updateNotif").hidden = false ;
	}, 0* 2*60*1000);
},{once:true,});
