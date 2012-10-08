Sections = new Meteor.Collection("sections");
Sites = new Meteor.Collection("sites");

Template.section_selectors.sections = function() {
  return Sections.find();
};

Template.section_selector.active_class = function() {
  var active = Session.get("active_section");
  if (active === undefined) { return this.name === "read" ? "active" : ""; }
  return active === this._id ? "active" : "";
};

Template.section_selector.events({
  'click': function(e) {
    Session.set("active_section", this._id);
  }
});

Template.footer.events({
  'click .toggle': function(e) {
    $(".footer .content").slideToggle(500);
    $("html, body").animate({
      scrollTop: $(document).height()
    }, 500);
  }
});

Template.site_list.sites_by_3 = function() {
  var active = Session.get("active_section");
  var active_section = "read";
  if (active !== undefined) { active_section = Sections.findOne({_id: active}).name; }
  var groups = [];
  var i = 0;
  Sites.find({type: active_section}).forEach(function(site) {
    if (i % 3 === 0) {
      groups.push([]);
    }

    groups[groups.length - 1].push(site);
    i++;
  });
  return groups;
};

Template.site_row.sites = function() {
  return this;
};
