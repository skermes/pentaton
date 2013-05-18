Sites = new Meteor.Collection("sites");
Sections = new Meteor.Collection("sections");

Sites.allow({
  insert: function() { return true; },
  update: function() { return true; },
  remove: function() { return true; }
})

Meteor.startup(function() {
  var default_sites = [
    {type: "music", name: "KEXP", url: "http://kexp.org", color: "FEAC31"},
    {type: "music", name: "last.fm", url: "http://last.fm/home", color: "C70E14"},
    {type: "music", name: "MOG", url: "http://mog.com", color: "FF263A"},
    {type: "music", name: "thesixtyone", url: "http://thesixtyone.com", color: "4C4C4C"},
    {type: "music", name: "Pandora", url: "http://pandora.com", color: "0F6799"},
    {type: "music", name: "mixest", url: "http://mixest.com", color: "80CC70"},
    {type: "music", name: "tunes.io", url: "http://tunes.io", color: "0000FF"},
    {type: "music", name: "Aurgasm", url: "http://aurgasm.us", color: "067292"},
    {type: "music", name: "Earbits", url: "http://www.earbits.com/play", color: "292929"},
    {type: "read", name: "Reader", url: "http://www.google.com/reader/view/#overview-page", color: "4A8CF6"},
    {type: "read", name: "Hacker News", url: "http://news.ycombinator.com", color: "FF6600"},
    {type: "read", name: "Stellar", url: "http://stellar.io/skermes/flow", color: "FF7358"},
    {type: "read", name: "forums", url: "http://forums.xkcd.com/", color: "0A90D2"},
    {type: "read", name: "Mark watches", url: "http://markwatches.net/reviews/", color: "000000"},
    {type: "read", name: "Rock Paper Shotgun", url: "http://www.rockpapershotgun.com/", color: "1D1D1D"},
    {type: "read", name: "kindle", url: "https://read.amazon.com", color: "F5A74E"},
    {type: "movies", name: "Netflix", url: "http://movies.netflix.com/Default", color: "B9090B"},
    {type: "movies", name: "Hulu", url: "http://www.hulu.com/", color: "9BC83A"},
    {type: "movies", name: "TableTop", url: "http://tabletop.geekandsundry.com/", color: "77DD00"},
    {type: "movies", name: "Day [9]", url: "http://day9.tv/archives/", color: "FEAC2C"},
    {type: "games", name: "Kongegrate", url: "http://www.kongregate.com", color: "990000"},
    {type: "games", name: "Armor Games", url: "http://armorgames.com/", color: "2E353E"},
    {type: "games", name: "Jayisgames", url: "http://jayisgames.com", color: "9AB821"}
  ];
  if (Sites.find().count() !== default_sites.length) {
    Sites.remove({});
    for (var i = 0; i < default_sites.length; i++) {
      Sites.insert(default_sites[i]);
    }
  }

  var default_sections = [
    {name: "music", icon: "headphones", description: "Select music sites"},
    {name: "read", icon: "book", description: "Select reading sites"},
    {name: "movies", icon: "television", description: "Select movie sites"},
    {name: "games", icon: "controller", description: "Select game sites"}
  ];
  if (Sections.find().count() !== default_sections.length) {
    for (var i = 0; i < default_sections.length; i++) {
      Sections.insert(default_sections[i]);
    }
  }
});
