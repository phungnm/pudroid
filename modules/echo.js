'use strict';
// Imports the Google Cloud client library
  const {Translate} = require('@google-cloud/translate');
module.exports = (bot) => {
  bot.on('message', (payload, chat, data) => {
    const text = payload.message.text;
    if (data.captured) { return; }



 
  // Instantiates a client
  const translate = new Translate({projectId});
 
  // The text to translate
  const text = 'Hello, world!';
 
  // The target language
  const target = 'ru';
 
  // Translates some text into Russian
  const [translation] = await translate.translate(text, target);
  console.log(`Text: ${text}`);
  console.log(`Translation: ${translation}`);
}





    chat.say(`Echo: ${text}`);
  });
};
