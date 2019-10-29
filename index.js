'use strict';
const BootBot = require('bootbot');
const config = require('config');
const dictionaryModule = require('./modules/dictionary');

const bot = new BootBot({
  accessToken: config.get('access_token'),
  verifyToken: config.get('verify_token'),
  appSecret: config.get('app_secret')
});
bot.setPersistentMenu(
    [{
      type: 'postback',
      title: 'Random a word',
      payload: 'PERSISTENT_MENU_RANDOM'
    },
    {
      type: 'postback',
      title: 'Search',
      payload: 'PERSISTENT_MENU_SEARCH'
    }]
);
bot.module(dictionaryModule);



var port = process.env.PORT || '3001';
bot.start(port);


