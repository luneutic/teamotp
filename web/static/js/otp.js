const S_COPIED = "Copied!"
const S_LOWAST = "&lowast;&lowast;&lowast;&lowast;&lowast;&lowast;"
const COPIED_DURATION = 1500

function fallback_copy_text_to_clipboard(text) {
	const tmp_textarea = document.createElement("textarea")
	tmp_textarea.value = text

	// Avoid scrolling to bottom
	tmp_textarea.style.position = "fixed"
	tmp_textarea.style.top = "0"
	tmp_textarea.style.left = "0"

	document.body.appendChild(tmp_textarea)
	tmp_textarea.focus()
	tmp_textarea.select()

	try {
		const success = document.execCommand("copy")
		if (success) {
			console.log("Fallback copy successful. Please switch to https")
		} else {
			alert("Copy was unsuccessful")
		}
	} catch (err) {
		alert("Copy failed", err)
	}

	document.body.removeChild(tmp_textarea)
}

function copy_text_to_clipboard(text) {
	if (!navigator.clipboard) {
		fallback_copy_text_to_clipboard(text)
		return
	}

	navigator.clipboard.writeText(otp)
}

const getOtpValue = async (id) => {
	const resp = await fetch(`/otp/${id}`)

	if (!resp.ok) {
		throw new Error(`error: ${resp.status} ${resp.statusText}`)
	}

	const data = await resp.text()
	return data
}

const copyOtp = (otpCard) => {
	const id = otpCard.dataset.id

	getOtpValue(id).then((otp) => {
		copy_text_to_clipboard(otp)

		const otpValSpan = otpCard.querySelector(`.otp-value`)
		otpValSpan.textContent = S_COPIED
		setTimeout(() => { otpValSpan.innerHTML = S_LOWAST }, COPIED_DURATION)
	})
}

// === Add OTP Modal ===

const otpForm = document.querySelector(".modal#add-otp")

const showOtpForm = () => {
	otpForm.style.display = "block"
}

const hideOtpForm = () => {
	otpForm.style.display = "none"
}

const modalOutsideClick = (ev) => {
	if (ev.target === otpForm) {
		hideOtpForm()
	}
}

otpForm.addEventListener("click", modalOutsideClick)
