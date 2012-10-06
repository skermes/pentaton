Sites = new Meteor.Collection("sites");

Template.site_list.sites = function() {
  return Sites.find();
}
