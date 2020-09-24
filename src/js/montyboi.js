// Updates the location value.
const updateLocation = (loc) => {
    fetch('/update/location', {
        method: 'POST',
        headers: {
            'Content-Type': 'plain/text',
        },
        body: loc,
    })
    .then(response => location.reload());
}

// Updates the status value.
const updateStatus = () => {
    const selectedValue = document.getElementById('status-select').value;
    console.log(selectedValue);
    fetch('/update/status', {
        method: 'POST',
        headers: {
            'Content-Type': 'plain/text',
        },
        body: selectedValue,
    })
    .then(response => location.reload());
}