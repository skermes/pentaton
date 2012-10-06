Sites = new Meteor.Collection("sites");

if (Sites.find().count() === 0) {
  var default_sites = [
    {name: "KEXP", url: "http://kexp.org"},
    {name: "last.fm", url: "http://last.fm/home"},
    {name: "MOG", url: "http://mog.com"},
    {name: "thesixtyone", url: "http://thesixtyone.com"},
    {name: "Pandora", url: "http://pandora.com"},
    {name: "mixest", url: "http://mixest.com"},
    {name: "tunes.io", url: "http://tunes.io"},
    {name: "Aurgasm", url: "http://aurgasm.us"}
  ];
  for (var i = 0; i < default_sites.length; i++) {
    Sites.insert(default_sites[i]);
  }
}
