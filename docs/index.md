---
title: General Index
---

Choose a language /// Choisi une langue

- [Français/French](fr/)
- [English](en/)

<script>
	(function () {
		for (let lang of navigator.languages) {
			switch (lang) {
				case "fr":
				case "fr-FR":
					document.location.href="fr/";
					return ;
				case "en":
				case "en-US":
					document.location.href="en/";
					return ;
			}
		}
		document.location.href="en/";
	})();
</script>
