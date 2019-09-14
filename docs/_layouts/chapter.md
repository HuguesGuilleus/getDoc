---
layout: default
---

{%- comment -%}
	2019 GUILLEUS Hugues <ghugues@netc.fr>
	BSD 3-Clause "New" or "Revised" License
{%- endcomment -%}

<meta name="viewport" content="width=device-width, initial-scale=1">

{%- if page.path != "index.md" -%}
	<h1 style="margin-bottom:0px;">{{page.title}}</h1>
	{% include lang.liquid %}
{%- endif -%}

{{ content }}

<footer>
	<hr>
	{%- assign pageLang = page.path | split: '/' | first -%}
	{%- assign remoteURL = "https://github.com/HuguesGuilleus/getDoc/" -%}
	{%- assign remoteLicense = "https://github.com/HuguesGuilleus/getDoc/blob/master/LICENSE" -%}
	{%- case pageLang -%}
		{%- when "fr" -%}
		<a href="{{remoteLicense}}" title="License">
			{% octicon law %} BSD 3-Clause "New" or "Revised" License (License BSD trois clauses «Nouvelles» ou «Révisé»)
		</a><br>
		<a href="{{remoteURL}}" title="Dépôt GitHub">{% octicon mark-github %} GitHub</a>
		{%- else -%}
			<a href="{{remoteLicense}}" title="License">
				{% octicon law %} BSD 3-Clause "New" or "Revised" License
			</a><br>
			<a href="{{remoteURL}}" title="GitHub Repository">{% octicon mark-github %} GitHub</a>
	{%- endcase -%}
</footer>
