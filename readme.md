====================================================================================================

Simple:

Download go | https://go.dev/dl/

cd to the dir and run `go build`

import your token into `/imput/tokens`
import the users by there id into `/memberids`

then run the `.exe` file

====================================================================================================
Need to do etc: 

1. imput your tokens in `input/tokens.txt`
2. Input proxies in config.json Only HTTP(s) rotating proxies of the format user:pass@ip:port or ip:port. Use High Quality proxies for improved functionality.
3. Scrape the UIDs of a server for Multi DM mode. Make a file `users.txt` in the same directory for it to output.
4. Add UID's of discord Users who you want to message in `input/memberids.txt`
5. Decide the delay by setting your config file `config.json`
6. Add your message in message.json. This can be an Embed. Use this website for building the embed easily. Alternatively, you can use the `Get message` feature in the program.
7. Remove any fields you `don't wish` to send
8. Writing <user> anywhere in the message content would ping the user
9. Run the binary


====================================================================================================
porxys:
1. proxy in format `User:pass@IP:port`


====================================================================================================
Embeds and normal msg users:

https://discordjs.guide/popular-topics/embeds.html#embed-preview for embeds examples `#Using an embed object`

1. 
{
    "content" : "hi <user> \n To change line\n Made by !Glitched!"
}


embed:

2. {
  "content" : "Hi Fellow Discord User \n This is an example message! Use \\n to change lines and to ping people use <user> don't forget to marry me nimz!",
  "embeds": 
      [{
        "type": "rich",
        "title": "This can be a link",
        "description": "You can have embeds however you like them. As long as you send them in the correct format, they will be sent!",
        "color": 3348893,
        "fields": [
          {
            "name": "<-- You can add any colour but be sure it's in decimal",
            "value": "This is an embed field, you can add this too. You can delete anything from here to not have it show up.",
            "inline": true
          },
          {
            "name": "You can add multiple of these lol",
            "value": "You can add images too and set their size"
          }
        ],
        "image": {
          "url": "https://i.imgur.com/RCBBege.png",
          "height": 0,
          "width": 0
        },
        "author": {
          "name": "Use this website to make your Embed easily. ",
          "url": "https://autocode.com/tools/discord/embed-builder/",
          "icon_url": "https://i.imgur.com/RCBBege.png"
        }
      }]
    
}
====================================================================================================