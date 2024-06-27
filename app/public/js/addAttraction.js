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

    // Create form elements
    const formTitle = document.createElement('h2');
    formTitle.innerText = 'Add Attraction';

    const nameLabel = document.createElement('label');
    nameLabel.innerText = 'Name:';
    nameLabel.style.display = 'block';
    nameLabel.style.marginTop = '10px';

    const nameInput = document.createElement('input');
    nameInput.type = 'text';
    nameInput.style.width = '100%';
    nameInput.style.padding = '5px';
    nameInput.style.marginTop = '5px';

    const descriptionLabel = document.createElement('label');
    descriptionLabel.innerText = 'Description:';
    descriptionLabel.style.display = 'block';
    descriptionLabel.style.marginTop = '10px';

    const descriptionInput = document.createElement('textarea');
    descriptionInput.style.width = '100%';
    descriptionInput.style.padding = '5px';
    descriptionInput.style.marginTop = '5px';

    // Add City, Title, Type, Info, and Image attributes to the addAttraction function
    const cityLabel = document.createElement('label');
    cityLabel.innerText = 'City:';
    cityLabel.style.display = 'block';
    cityLabel.style.marginTop = '10px';

    const cityInput = document.createElement('input');
    cityInput.type = 'text';
    cityInput.style.width = '100%';
    cityInput.style.padding = '5px';
    cityInput.style.marginTop = '5px';

    const titleLabel = document.createElement('label');
    titleLabel.innerText = 'Title:';
    titleLabel.style.display = 'block';
    titleLabel.style.marginTop = '10px';

    const titleInput = document.createElement('input');
    titleInput.type = 'text';
    titleInput.style.width = '100%';
    titleInput.style.padding = '5px';
    titleInput.style.marginTop = '5px';

    const typeLabel = document.createElement('label');
    typeLabel.innerText = 'Category:';
    typeLabel.style.display = 'block';
    typeLabel.style.marginTop = '10px';

    const typeInput = document.createElement('input');
    typeInput.type = 'text';
    typeInput.style.width = '100%';
    typeInput.style.padding = '5px';
    typeInput.style.marginTop = '5px';

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
    imageInput.style.width = '100%';
    imageInput.style.padding = '5px';
    imageInput.style.marginTop = '5px';

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
    formContainer.appendChild(cityLabel);
    formContainer.appendChild(cityInput);
    formContainer.appendChild(titleLabel);
    formContainer.appendChild(titleInput);
    formContainer.appendChild(typeLabel);
    formContainer.appendChild(typeInput);
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
    const apiUrl = `https://nominatim.openstreetmap.org/search.php?q=${encodeURIComponent(address)}&format=json&limit=1`;


    addButton.addEventListener('click', () => {
        const name = nameInput.value;
        const description = descriptionInput.value;
        const city = cityInput.value;
        const title = titleInput.value;
        const type = typeInput.value;
        const info = infoInput.value;
        const image = imageInput.value;
        fetch(apiUrl)
            .then(response => response.json())
            .then(data => {
                if (data.length > 0) {
                    return {
                        posX: data[0].lat,
                        posY: data[0].lon
                    };
                } else {
                    throw new Error("No coordinates found for the given address");
                }
            }).catch(error => {
                console.error("There was a problem with the request:", error);
                alert("Failed to get coordinates: " + error.message);
            });

        fetch('/api/attractions', {
            method: 'POST',
            body: JSON.stringify({ name, description, city, title, type, info, image, posX, posY }),
            headers: { 'Content-Type': 'application/json' }
        }).then(() => {
            document.body.removeChild(overlay);
            window.location.reload();
        }).catch((error) => {
            console.error("There was a problem with the request:", error);
            alert("Failed to add attraction: ", error);
        });
    });

}