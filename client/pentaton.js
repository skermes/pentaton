Sites = new Meteor.Collection("sites");

Template.section_selectors.sections = function() {
  return [
    {_id: "music", name: "music", icon: "headphones", description: "Select music sites"},
    {_id: "read", name: "read", icon: "book", description: "Select reading sites"},
    {_id: "movies", name: "movies", icon: "television", description: "Select movie sites"},
    {_id: "games", name: "games", icon: "controller", description: "Select game sites"}
  ];
}

Template.site_list.sites_by_3 = function() {
  var groups = [];
  var i = 0;
  Sites.find().forEach(function(site) {
    if (i % 3 === 0) {
      groups.push([]);
    }

    groups[groups.length - 1].push(site);
    i++;
  });
  return groups;
}

Template.site_row.sites = function() {
  return this;
}
