function $(id) {
	return document.getElementById(id);
}

function m() {
	const l = $('log'),
		title = $('title'),
		general = $('general'),
		dropzone = $('dropzone'),
		doc = $('doc'),
		files = $('files'),
		w = new Worker('worker.js');

	let extSupported = [];

	$('title-set').addEventListener('click', () => {
		doc.hidden = true;
		w.postMessage({
			type: 'title',
			title: title.value,
		});
	});

	$('gen').addEventListener('click', () => {
		w.postMessage({
			type: 'ask',
		});
	});

	$('reset').addEventListener('click', () => {
		w.postMessage({
			type: 'reset',
		});
	});

	// Input files
	files.addEventListener('change', () => {
		doc.hidden = true;
		Array.from(files.files).forEach(f => w.postMessage({
			type: 'blob',
			name: f.name,
			blob: f,
		}));
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
		doc.hidden = true;
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
			window.URL.revokeObjectURL(doc.href);
			doc.href = URL.createObjectURL(data.blob);
			doc.hidden = false;
			break;
		default:
			console.error('Unknown message type:', data);
		}
	};
}
document.readyState == 'loading' ? document.addEventListener('DOMContentLoaded', m, {
	once: true
}) : m();