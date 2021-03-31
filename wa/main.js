function $(id) {
	return document.getElementById(id);
}

// All informations about one output format.
function output(ext) {
	this.ext = ext;
	this.gen = $('gen-' + ext);
	this.doc = $('doc-' + ext);
	this.reset = () => {
		this.gen.hidden = false;
		this.doc.hidden = true;
		window.URL.revokeObjectURL(this.url);
		this.url = '';
	};
	this.setURL = u => {
		this.doc.hidden = false;
		this.gen.hidden = true;
		window.URL.revokeObjectURL(this.url);
		this.url;
		this.doc.href = u;
	};
}

function m() {
	const l = $('log'),
		title = $('title'),
		general = $('general'),
		dropzone = $('dropzone'),

		outs = new Map([
			['text/html', new output('html')],
			['application/json', new output('json')],
			['application/xml', new output('xml')],
		]),

		files = $('files'),
		w = new Worker('worker.js');

	let extSupported = [];

	function setTitle() {
		resetOutput();
		w.postMessage({
			type: 'title',
			title: title.value,
		});
	}
	$('title-set').addEventListener('click', setTitle);
	title.addEventListener('keydown', e => {
		switch (e.key) {
		case 'Tab':
		case 'Enter':
			setTitle();
		}
	});

	function resetOutput() {
		outs.forEach(o => o.reset());
	}

	outs.forEach((o, k) => o.gen.addEventListener('click',
		() => w.postMessage({
			type: 'ask',
			format: o.ext,
		})
	));

	$('reset').addEventListener('click', () => {
		resetOutput();
		w.postMessage({
			type: 'reset',
		});
	});

	// Input files
	files.addEventListener('change', () => {
		resetOutput();
		Array.from(files.files).forEach(f => w.postMessage({
			type: 'blob',
			name: f.name,
			blob: f,
		}));
	});

	$('files-label').addEventListener('keydown', e => {
		switch (e.key) {
		case ' ':
		case 'Enter':
			files.click();
		}
	});

	// Drop files
	document.addEventListener('dragenter', e => {
		general.hidden = true;
		dropzone.hidden = false;
	}, false);
	document.addEventListener('dragover', e => {
		e.stopPropagation();
		e.preventDefault();
	}, false);
	['drop', 'dragleave', 'dragexit', 'dragend'].forEach(n => document.addEventListener(n, e => {
		e.stopPropagation();
		e.preventDefault();
		general.hidden = false;
		dropzone.hidden = true;
	}));
	document.addEventListener('drop', async e => {
		resetOutput();
		Array.from(e.dataTransfer.items)
			.map(e => e.webkitGetAsEntry())
			.map(async function readFileOrDir(f) {
				if (f.isDirectory) {
					f.createReader().readEntries(l => l.forEach(readFileOrDir));
				} else if (extSupported.some(ext => ext.test(f.fullPath))) {
					w.postMessage({
						type: 'blob',
						name: f.fullPath,
						blob: await new Promise((r, e) => f.file(r, e)),
					});
				}
			});
	});

	// Message from the worker
	w.onmessage = ({
		data
	}) => {
		switch (data.type) {
		case 'logReset':
			l.innerText = data.text || '';
			break;
		case 'logLine':
			l.innerText += data.line;
			break;
		case 'ext':
			extSupported = data.ext.map(e => new RegExp('[./]' + e + '$'));
			break;
		case 'doc':
			const o = outs.get(data.blob.type.replace(/;.*/, ''));
			if (o) {
				o.setURL(URL.createObjectURL(data.blob))
			} else {
				console.error('Unknwon Blob type:', data.blob.type);
			}
			break;
		default:
			console.error('Unknown message type:', data);
		}
	};
}
document.readyState == 'loading' ? document.addEventListener('DOMContentLoaded', m, {
	once: true
}) : m();