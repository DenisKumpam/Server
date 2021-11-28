let button = document.querySelector('#validate-form-button');
button.addEventListener('click', function (){
    let nameInput = document.querySelector('#name')
    let mailInput = document.querySelector('#mail')
    let passInput = document.querySelector('#password')
    let form = {
        name: nameInput.value,
        mail: mailInput.value,
        password: passInput.value,
    }

    fetch( window.location.origin + '/validate', {
        method: 'POST', // *GET, POST, PUT, DELETE, etc.
        mode: 'cors', // no-cors, *cors, same-origin
        cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
        credentials: 'same-origin', // include, *same-origin, omit
        headers: {
            'Content-Type': 'application/json'
            // 'Content-Type': 'application/x-www-form-urlencoded',
        },
        redirect: 'follow', // manual, *follow, error
        referrerPolicy: 'no-referrer', // no-referrer, *client
        body: JSON.stringify(form) // body data type must match "Content-Type" header
    }).then(response => response.json())
        .then(data => {
            console.log('Success:', data);
        })
        .catch((error) => {
        console.error('Error:', error);
    })
})

