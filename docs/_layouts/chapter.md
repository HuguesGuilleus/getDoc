---
layout: default
---

{%- comment -%}
	2019 GUILLEUS Hugues <ghugues@netc.fr>
	BSD 3-Clause "New" or "Revised" License
{%- endcomment -%}

{% assign remote = "getDoc" %}

<style>
	h1:first-of-type{
		display: none;
	}
	h1{
		margin: 0px !important;
	}
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

{% if page.path != "index.md" %}
<h1>
	<a href="./">{{remote}}</a>&nbsp;/
	{{page.title}}
	<span id=betaTag>
		{% octicon beaker %} Beta Doc
	</span>
</h1>
{% include lang.liquid %}
{%- endif %}

{{ content }}

<footer>
	<hr>
	{%- assign pageLang = page.path | split: '/' | first -%}
	{%- capture remoteURL -%}
		https://github.com/HuguesGuilleus/{{ remote }}/
	{%- endcapture -%}
	{%- capture remoteLicense -%}
		https://github.com/HuguesGuilleus/{{ remote }}/blob/master/LICENSE
	{%- endcapture -%}
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
