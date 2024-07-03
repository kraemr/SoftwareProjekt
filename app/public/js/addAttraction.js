var posX = 0;
var posY = 0;
function addAttraction() {
    // Create overlay element
    const overlay = document.createElement('div');
    overlay.id = 'addAttractionOverlay';
    overlay.style.position = 'fixed';
    overlay.style.top = '0';
    overlay.style.left = '0';
    overlay.style.width = '100%';
    overlay.style.height = '100%';
    overlay.style.backgroundColor = 'rgba(0, 0, 0, 0.5)';
    overlay.style.display = 'flex';
    overlay.style.justifyContent = 'center';
    overlay.style.alignItems = 'center';
    overlay.style.zIndex = '1000';

    // Create form container
    const formContainer = document.createElement('div');
    formContainer.style.backgroundColor = 'white';
    formContainer.style.padding = '20px';
    formContainer.style.borderRadius = '10px';
    formContainer.style.boxShadow = '0 4px 8px rgba(0, 0, 0, 0.1)';
    formContainer.style.width = '300px';

    // Add City, Title, Type, Info, and Image attributes to the addAttraction function
    const cityLabel = document.createElement('label');
    cityLabel.innerText = 'City';
    cityLabel.style.display = 'block';
    cityLabel.style.marginTop = '10px';

    const cityInput = document.createElement('input');
    cityInput.type = 'text';
    cityInput.classList.add('add-attraction-input');

    const addressLabel = document.createElement('label');
    addressLabel.innerText = 'Address';
    addressLabel.style.display = 'block';
    addressLabel.style.marginTop = '10px';

    const addressInput = document.createElement('input');
    addressInput.type = 'text';
    addressInput.classList.add('add-attraction-input');

    const houseNumber = document.createElement('label');
    houseNumber.innerText = 'House Number';
    houseNumber.style.display = 'block';
    houseNumber.style.marginTop = '10px';

    const houseNumberInput = document.createElement('input');
    houseNumberInput.type = 'text';
    houseNumberInput.classList.add('add-attraction-input');

    const titleLabel = document.createElement('label');
    titleLabel.innerText = 'Title:';
    titleLabel.style.display = 'block';
    titleLabel.style.marginTop = '10px';

    const titleInput = document.createElement('input');
    titleInput.type = 'text';
    titleInput.classList.add('add-attraction-input');

    const typeLabel = document.createElement('label');
    typeLabel.innerText = 'Category:';
    typeLabel.style.display = 'block';
    typeLabel.style.marginTop = '10px';

    const typeSelect = document.createElement('select');
    typeSelect.style.width = '100%';
    typeSelect.style.padding = '5px';
    typeSelect.style.marginTop = '5px';

    // Fetch categories from API and populate dropdown using getCategories function
    getCategories()
        .then(categories => {
            categories.forEach(category => {
                const option = document.createElement('option');
                option.value = category;
                option.innerText = category;
                typeSelect.appendChild(option);
            });
        })
        .catch(error => {
            console.error("There was a problem fetching categories:", error);
            alert("Failed to load categories: " + error.message);
        });


    const infoLabel = document.createElement('label');
    infoLabel.innerText = 'Description:';
    infoLabel.style.display = 'block';
    infoLabel.style.marginTop = '10px';

    const infoInput = document.createElement('textarea');
    infoInput.style.width = '100%';
    infoInput.style.padding = '5px';
    infoInput.style.marginTop = '5px';

    const imageLabel = document.createElement('label');
    imageLabel.innerText = 'Image URL:';
    imageLabel.style.display = 'block';
    imageLabel.style.marginTop = '10px';

    const imageInput = document.createElement('input');
    imageInput.type = 'text';
    imageInput.classList.add('add-attraction-input');

    const addButton = document.createElement('button');
    addButton.innerText = 'Add';
    addButton.style.marginTop = '10px';
    addButton.style.padding = '10px';
    addButton.style.backgroundColor = '#50879c';
    addButton.style.color = 'white';
    addButton.style.border = 'none';
    addButton.style.borderRadius = '5px';
    addButton.style.cursor = 'pointer';

    const cancelButton = document.createElement('button');
    cancelButton.innerText = 'Cancel';
    cancelButton.style.marginTop = '10px';
    cancelButton.style.marginLeft = '10px';
    cancelButton.style.padding = '10px';
    cancelButton.style.backgroundColor = '#ccc';
    cancelButton.style.color = 'black';
    cancelButton.style.border = 'none';
    cancelButton.style.borderRadius = '5px';
    cancelButton.style.cursor = 'pointer';

    // Append form elements to form container
    formContainer.appendChild(titleLabel);
    formContainer.appendChild(titleInput);
    formContainer.appendChild(cityLabel);
    formContainer.appendChild(cityInput);
    formContainer.appendChild(addressLabel);
    formContainer.appendChild(addressInput);
    formContainer.appendChild(houseNumber);
    formContainer.appendChild(houseNumberInput);
    formContainer.appendChild(typeLabel);
    formContainer.appendChild(typeSelect);
    formContainer.appendChild(infoLabel);
    formContainer.appendChild(infoInput);
    formContainer.appendChild(imageLabel);
    formContainer.appendChild(imageInput);
    formContainer.appendChild(addButton);
    formContainer.appendChild(cancelButton);

    // Append form container to overlay
    overlay.appendChild(formContainer);

    // Append overlay to body
    document.body.appendChild(overlay);

    // Add event listeners
    cancelButton.addEventListener('click', () => {
        document.body.removeChild(overlay);
    });

    addButton.addEventListener('click', () => {
        console.log('Add button clicked');
        const city = cityInput.value;
        const address = addressInput.value;
        const houseNumber = houseNumberInput.value;
        const title = titleInput.value;
        const type = typeSelect.value;
        const info = infoInput.value;
        const image = imageInput.value;
        const fullAddress = city + " " + address + " " + houseNumber;
        var apiUrl = 'https://nominatim.openstreetmap.org/search.php?q=' + encodeURIComponent(fullAddress) + '&format=geojson&limit=1&countrycodes=de'
        console.log(title, city, address, fullAddress, type, info, image)
        console.log(apiUrl)
        
        fetch(apiUrl)
            .then(response => response.json())
            .then(data => {
                if (data.features && data.features.length > 0) {
                    const coordinates = data.features[0].geometry.coordinates;
                    const posX = coordinates[0]; // Longitude
                    const posY = coordinates[1]; // Latitude
                    console.log(posX, posY);
    
                    // Proceed with the second fetch call only after coordinates are fetched
                    return fetch('/api/attractions', {
                        method: 'POST',
                        body: JSON.stringify({title, city, address, houseNumber, type, info, image, posX, posY }),
                        headers: { 'Content-Type': 'application/json' }
                    });
                } else {
                    throw new Error('No coordinates found for the given address');
                }
            })
            .then(() => {
                document.body.removeChild(overlay);
                window.location.reload();
            })
            .catch((error) => {
                console.error("There was a problem with the request:", error);
                alert("Failed to add attraction: " + error.message);
            });
    });
    
}