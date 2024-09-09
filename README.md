# GOtWebo2, a discord bot created from [BotWebo2](https://github.com/gulis1/BotWebo2) in Go.

## Commands
* [music commands](#music)
* [danbooru commands](#danbooru)
* [sauce commands](#sauce)
* [anilist commands](#anilist)
---
## Features
* [anime/manga news](#news)
* [emotes](#emotes)
---


<h3 id="music"> Music commands </h3>
Use command /play with the link of YouTube videos and playlist to add songs.

<h4>Other arguments of the command to manipulate the list</h4>

* [**list**]                                                   _Shows current playlist._
* [**loop**]                                                   _Changes current loop settings._
* [**skip**]                                                   _If no number is provided, skips to the next song._
* [**song**]                                                   _Displays info about the current song._
* [**empty**]                                                  _Removes all songs from the playlist._


TODO
* [**rload** <username>]                                       _Loads completed anime list from anilist username._
* [**rplay**]                                                  _Plays random songs from the anime list loaded._
* [**rstop**]                                                  _Stop playing random theme._
* [**ruser**]                                                  _Shows current list owner._
* [**shuffle**]                                                _Shuffles the playlist._


<h3 id="danbooru">Danbooru commands</h3>

* [**/danbooru** tags]                                          _Sends a random image from danbooru with the specified tag._

<h3 id="sauce"> Saucenao commands</h3>

* [**/sauce** url]                                              _Finds the source for the image in the url._


<h3 id="anilist"> Anilist commands</h3>

* [**/anime** name]                                             _Gets the remaining time until the next episode of the specified anime._

  
---
<h3 id="news">Anime news</h3>

First you must have 2 discord channels called "anime-webonews" and "manga-webonews".

Then, every 20 mins the bot will send the most recent anime news.

Created using the [animenewsnetwork.com](https://www.animenewsnetwork.com) newsfeed.

[File here](https://www.animenewsnetwork.com/news/atom.xml)

<h3 id="emotes">Emotes</h3>
The bot will send an image related to the text, try it yourself with /image [text].
  
* yes
* no
* pray
* please
* smug
* trembling
* pekora
* haacham
