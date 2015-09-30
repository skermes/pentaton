{{define "js"}}
(function() {
  var run = function() {
    var addLinkElem = document.getElementById('add-link');
    var trigger = addLinkElem.getElementsByClassName('activate')[0];
    var colorInput = document.getElementById('add-link-color');
    var urlInput = document.getElementById('add-link-url');
    var nameInput = document.getElementById('add-link-name');
    var menu = document.getElementById('menu');
    var menuTrigger = document.getElementById('menu-trigger');

    var changeLinkFormBackground = function() {
      var color = colorInput.value;
      addLinkElem.style.backgroundColor = '#' + color;
    };

    var showNewLinkForm = function() {
      addLinkElem.className = 'active';
      changeLinkFormBackground();
    };

    var selectAll = function(e) {
      e.target.select();
    };

    var closeMenu = function() {
      menu.className = '';
      document.body.removeEventListener('click', closeMenu);
    };

    var showMenu = function() {
      menu.className = 'active';
      document.body.addEventListener('click', closeMenu);
    };

    var stop = function(e) {
      e.stopPropagation();
    }

    trigger.addEventListener('click', showNewLinkForm);
    colorInput.addEventListener('input', changeLinkFormBackground)
    colorInput.addEventListener('click', selectAll);
    urlInput.addEventListener('click', selectAll);
    nameInput.addEventListener('click', selectAll);
    menuTrigger.addEventListener('click', showMenu);
    menu.addEventListener('click', stop);
  }

  document.addEventListener('DOMContentLoaded', run);
})()
{{end}}
