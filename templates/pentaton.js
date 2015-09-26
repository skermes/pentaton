{{define "js"}}
(function() {
  var run = function() {
    var addLinkElem = document.getElementById('add-link');
    var trigger = addLinkElem.getElementsByClassName('activate')[0];

    var showNewLinkForm = function() {
      addLinkElem.className = 'active';
    }

    trigger.addEventListener('click', showNewLinkForm);
  }

  document.addEventListener('DOMContentLoaded', run);
})()
{{end}}
