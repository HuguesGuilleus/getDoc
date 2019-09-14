---
layout: default
---

{%- comment -%}
	2019 GUILLEUS Hugues <ghugues@netc.fr>
	BSD 3-Clause "New" or "Revised" License
{%- endcomment -%}

<style>
	svg.octicon{
		height: 1em !important;
		width: 1em !important;
	}
	#betaTag{
		background: #ff9400;
		color:#3e1300;
		padding: 0.3ex 1ex  ;
		font-size:60%;
		display: inline-block;
		border-radius: 0.5ex;
		margin-left: 2em;
	}
	#betaTag path {
		stroke:#3e1300;
		fill:#3e1300;
		stroke-width: 0.1;
	}
</style>

{% if page.path != "index.md" -%}
	<h1 style="margin-bottom:0px;">
		{%- if page.name != "index.md" -%}
			<a href="./">{% octicon home %}</a>
		{%- endif %}
		{{page.title}}
		<span id=betaTag>
			{% octicon beaker %} Beta Doc
		</span>
	</h1>
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
