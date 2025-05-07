const S_COPIED = "Copied!"
const S_LOWAST = "&lowast;&lowast;&lowast;&lowast;&lowast;&lowast;"
const COPIED_DURATION = 1500

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
		navigator.clipboard.writeText(otp)

		const otpValSpan = otpCard.querySelector(`.otp-value`)
		otpValSpan.textContent = S_COPIED
		setTimeout(() => { otpValSpan.innerHTML = S_LOWAST }, COPIED_DURATION)
	})
}
