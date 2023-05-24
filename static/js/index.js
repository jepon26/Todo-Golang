// index.js

// Get all the TODO items
let todoItems = document.querySelectorAll('.todos p');

// Add click event listener to each TODO item
todoItems.forEach(function(item) {
  item.addEventListener('click', function() {
    item.classList.toggle('finish');
  });
});
