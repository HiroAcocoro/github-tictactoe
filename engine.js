const fs = require("fs");
const { moves } = require("./moves");
const { mainMessage } = require("./readme-constants");

const emptyBoard =
	"||a|b|c|\n|---|---|---|---|\n|1|⬛|⬛|⬛|\n|2|⬛|⬛|⬛|\n|3|⬛|⬛|⬛|";

const letterToNum = (letter) => {
	const letterMap = {
		a: 1,
		b: 2,
		c: 3,
	};
	return letterMap[letter.toLowerCase()] || 0;
};

const renderMoves = () => {
	// 27 is the start of the board
	let renderBoard = emptyBoard;
	let isCircle = true;
	moves.forEach((move) => {
		// abc notation
		const xAxis = letterToNum(move[0]) * 2 + 1;

		// 123 notation
		const yAxis = +move[1] * 10;

		const boardArr = renderBoard.split("");
		boardArr[xAxis + yAxis + 17] = isCircle ? "⭕" : "❌";
		renderBoard = boardArr.join("");
		isCircle = !isCircle;
	});
	fs.writeFileSync("README.md", mainMessage + renderBoard, "utf8");
};

renderMoves();
