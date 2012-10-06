Sites = new Meteor.Collection("sites");

if (Sites.find().count() === 0) {
  var default_sites = [
    {name: "KEXP", url: "http://kexp.org", color: "FEAC31"},
    {name: "last.fm", url: "http://last.fm/home", color: "C70E14"},
    {name: "MOG", url: "http://mog.com", color: "FF263A"},
    {name: "thesixtyone", url: "http://thesixtyone.com", color: "4C4C4C"},
    {name: "Pandora", url: "http://pandora.com", color: "0F6799"},
    {name: "mixest", url: "http://mixest.com", color: "80CC70"},
    {name: "tunes.io", url: "http://tunes.io", color: "0000FF"},
    {name: "Aurgasm", url: "http://aurgasm.us", color: "067292"},
    {name: "Earbits", url: "http://www.earbits.com/play", color: "292929"}
  ];
  for (var i = 0; i < default_sites.length; i++) {
    Sites.insert(default_sites[i]);
  }
}
