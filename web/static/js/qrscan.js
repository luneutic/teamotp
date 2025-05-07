const result_div = document.querySelector("#qr-result")

function qrCodeSuccessCallback (decodedText, decodedResult) {
	html5_qrcode.stop()

	fetch('/otp/add', {
		method: 'POST',
		headers: {
			'Content-Type': 'text/plain',
		},
		body: decodedText,
	})
	.then(resp => {
		if (!resp.ok) {
			throw new Error(`error: ${resp.status} ${resp.statusText}`)
		}

		window.location.href = resp.url
	})
	.then(data => {})
	.catch(err => {})
}

let config = {
	verbose: false,
	fps: 10,
	qrbox: { width: 250, height: 250 },
	supportedScanTypes: [Html5QrcodeScanType.SCAN_TYPE_CAMERA],
	formatsToSupport: [Html5QrcodeSupportedFormats.QR_CODE],
}

let html5_qrcode = new Html5Qrcode(
	"qr-reader",
	config,
)
html5_qrcode.start({ facingMode: "environment" }, config, qrCodeSuccessCallback)
