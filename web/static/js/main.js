function get(url) {
	fetch(url, {
		method: 'GET',
		redirect: 'follow',
	})
	.then((response) => {
		if (!response.ok) {
			throw new Error(`error: ${resp.status} ${resp.statusText}`)
		}
		window.location.href = resp.url
	})
	.catch(err => {})
}
