var unirest = require("unirest");
exports.randomWord = () => {
 return new Promise((resolve, reject) => {
			var req = unirest("GET", "https://wordsapiv1.p.mashape.com/words/");
			req.query({
				"random": "true"
			}).headers({
			"X-Mashape-Key":"0reA09JpgnmshGI6Z2Sxl6usmjoWp1aEIV4jsn1ImdkLbThVb6"
			}).end(function (res) {
				if (res.error)  reject(res);
				    resolve(res);
			});


	});
}