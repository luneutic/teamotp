const timer_bars = document.querySelectorAll(".timer-bar")

const initialize_timers = () => {
	const seconds = new Date().getSeconds()

	for (timer of timer_bars) {
		const period = timer.dataset.period
		const delay = period - (seconds % period)

		timer.style.animationDuration = `${period}s`
		timer.style.animationDelay = `-${period - delay}s`
	}
}

document.addEventListener('DOMContentLoaded', () => {
	initialize_timers()
})
