
require('./setup');
require('./editor');

function resizeInput() {
    $(this).attr('size', $(this).val().length);
}

let share_url = $('#share-url');
share_url.keyup(resizeInput)
    // resize on page load
    .each(resizeInput);
$('#share').click(() => {
    share_url.addClass('active');
    share_url.focus();
});


import React from 'react';
import ReactDOM from 'react-dom';

ReactDOM.render(
    <h1>Hello, world!</h1>,
    document.getElementById('root')
);