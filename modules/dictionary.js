'use strict';

var dictionary = require("../functions/dictionary");







module.exports = (bot) => {

bot.setGreetingText('Hey there! Welcome to PuDroid!');
bot.setGetStartedButton((payload, chat) => {
  chat.say('Welcome to Pudroid. What are you looking for?');
});


bot.on('postback:PERSISTENT_MENU_RANDOM', (payload, chat) => {
	
	dictionary.randomWord().then(result=>{
		console.log(result);
		 chat.say(`OK`); 
	})

});

};
