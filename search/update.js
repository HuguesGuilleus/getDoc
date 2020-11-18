// getDoc
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

document.addEventListener("DOMContentLoaded", () => {
	document.getElementById("updateNotifClose").onclick = () => {
		document.getElementById("updateNotif").remove();
	};
	document.getElementById("updateNotifUpdate").onclick = () => {
		document.location.reload(true);
	};
	setTimeout(() => {
		document.getElementById("updateNotif").hidden = false;
	}, 3 * 60 * 1000);
}, {
	once: true,
});
