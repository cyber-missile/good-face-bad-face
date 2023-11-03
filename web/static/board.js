const delay = ms => new Promise(res => setTimeout(res, ms));

const animationFrame = () => new Promise(res => requestAnimationFrame(res));

const domUpdate = async () => {
	// no idea why this works
	// source: https://forum.freecodecamp.org/t/how-to-make-js-wait-until-dom-is-updated/122067/3
	await animationFrame();
	await animationFrame();
};

const turnover = async (card, target) => {
	card.style.transform = "perspective(5cm) rotateY(90deg)";
	await delay(500);
	card.style.transition = "no"
	card.style.transform = "perspective(5cm) rotateY(-90deg)";
	await domUpdate();
	card.style.transition = "";
	card.classList.remove("facedown");
	card.classList.remove("good");
	card.classList.remove("bad");
	card.classList.add(target);
	card.style.transform = "perspective(5cm) rotateY(0deg)";
	await delay(500);
};

const revealTop = async target => {
	const stack = document.getElementsByClassName("current")[0];
	const cards = stack.getElementsByClassName("card");
	const card = cards[cards.length - 1];

	await Promise.all([
		turnover(card, target),
		(async () => {
			card.style.marginLeft = "5vw";
		})()
	]);
};
