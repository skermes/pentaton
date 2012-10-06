Sites = new Meteor.Collection("sites");

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
